#include "extensions/filters/common/konvoy/anonymous_reporter_impl.h"

#include <list>

#include "common/common/utility.h"
#include "common/network/utility.h"

#include "absl/strings/str_join.h"
#include "absl/strings/str_replace.h"

namespace Envoy {
namespace Extensions {
namespace Filters {
namespace Common {
namespace Konvoy {

EnvVarSource std_env() {
  return [](const std::string& env_var) {
    auto value = std::getenv(env_var.c_str());
    return value != nullptr ? value : "";
  };
}

UdpSender::UdpSender(Network::Address::InstanceConstSharedPtr address)
    : io_handle_{address->socket(Network::Address::SocketType::Datagram)} {
  if (io_handle_->fd() == -1) {
    throw EnvoyException(fmt::format("failed to open socket to: {}", address->asString()));
  }

  const Api::SysCallIntResult result = address->connect(io_handle_->fd());

  if (result.rc_ == -1) {
    throw EnvoyException(
        fmt::format("failed to connect to: {}/{}", address->asString(), io_handle_->fd()));
  }
}

UdpSender::~UdpSender() {
  if (io_handle_->isOpen()) {
    RELEASE_ASSERT(io_handle_->close().err_ == nullptr, "");
  }
}

void UdpSender::send(const std::string& message) {
  ::send(io_handle_->fd(), message.c_str(), message.size(), MSG_DONTWAIT);
}

UdpSenderFactory::UdpSenderFactory(Network::DnsResolverSharedPtr dns_resolver)
    : dns_resolver_{dns_resolver} {}

CancellablePtr UdpSenderFactory::createSender(EnvVarSource env, SenderFactory::Cb cb) {
  std::string host{env(KONVOY_ANONYMOUS_REPORTS_HOST)};
  if (host == "") {
    host = KONVOY_ANONYMOUS_REPORTS_DEFAULT_HOST;
  }
  std::string port{env(KONVOY_ANONYMOUS_REPORTS_PORT)};
  if (port == "") {
    port = KONVOY_ANONYMOUS_REPORTS_DEFAULT_PORT;
  }
  uint64_t port64 = 0;
  if (!StringUtil::atoull(port.c_str(), port64, 10) || port64 > 65535) {
    ENVOY_LOG_MISC(trace, "invalid port value: {}", port);
    cb({});
    return {};
  }
  try {
    // if host value is already an IP address, we don't need DNS resolution
    auto address = Network::Utility::parseInternetAddress(host, port64);
    cb(std::make_unique<UdpSender>(std::move(address)));
    return {};
  } catch (const EnvoyException&) {
    ENVOY_LOG_MISC(trace, "need to resolve DNS name");

    // get DNS resolution involved
    auto* active_query = dns_resolver_->resolve(
        host, Network::DnsLookupFamily::Auto,
        [cb, host,
         port64](const std::list<Network::Address::InstanceConstSharedPtr>&& address_list) {
          ENVOY_LOG_MISC(trace, "DNS resolution is complete");

          if (address_list.empty()) {
            ENVOY_LOG_MISC(trace, "unable to resolve DNS name: {}", host);
            return;
          }
          auto& address = address_list.front();
          Network::Address::InstanceConstSharedPtr address_with_port =
              Network::Utility::getAddressWithPort(*address, port64);

          try {
            cb(std::make_unique<UdpSender>(std::move(address_with_port)));
          } catch (const EnvoyException& e) {
            ENVOY_LOG_MISC(trace, "failed to create UDP sender: {}", e.what());
            cb({});
          }
        });

    return active_query == nullptr
               ? CancellablePtr{}
               : std::make_unique<CancellableWrapper<Network::ActiveDnsQuery>>(active_query);
  }
}

std::string SimpleAnonymousReporter::createUniqueId(Runtime::RandomGenerator& random) {
  return absl::StrReplaceAll(random.uuid(), {{"-", ""}});
}

std::string SimpleAnonymousReporter::createPingMessage(const std::string& unique_id) {
  std::initializer_list<std::pair<std::string, std::string>> args = {
      {"version", "0.1"}, {"product", "konvoy"}, {"signal", "ping"}, {"unique_id", unique_id}};
  return absl::StrJoin(args, ";", absl::PairFormatter("="));
}

SimpleAnonymousReporter::SimpleAnonymousReporter(EnvVarSource env, Event::Dispatcher& dispatcher,
                                                 Runtime::RandomGenerator& random)
    : SimpleAnonymousReporter(
          env, dispatcher, random,
          std::make_shared<UdpSenderFactory>(dispatcher.createDnsResolver({}))) {}

SimpleAnonymousReporter::SimpleAnonymousReporter(EnvVarSource env, Event::Dispatcher& dispatcher,
                                                 Runtime::RandomGenerator& random,
                                                 SenderFactorySharedPtr sender_factory)
    : unique_id_{createUniqueId(random)},
      ping_message_{createPingMessage(unique_id_)}, sender_{std::make_unique<NoOpSender>()},
      sender_factory_{std::move(sender_factory)}, timer_{dispatcher.createTimer(
                                                      [this]() { onTimerTick(); })} {
  // initializagion of the actual sender might take some time
  cancellable_ = sender_factory_->createSender(env, [this](SenderPtr sender) {
    // avoid calling cancel() in destructor
    if (cancellable_) {
      cancellable_->disarm();
    }
    // sender might not have been constructed, e.g. due to failed DNS resolution
    if (sender) {
      ENVOY_LOG_MISC(trace, "using UDP sender");

      sender_.swap(sender);
    } else {
      ENVOY_LOG_MISC(trace, "falling back to no-op mode");
    }
  });
  uint64_t duration_ms = 0;
  if (!StringUtil::atoull(env(KONVOY_ANONYMOUS_REPORTS_INTERVAL).c_str(), duration_ms, 10) ||
      duration_ms <= 0) {
    duration_ms = KONVOY_ANONYMOUS_REPORTS_DEFAULT_INTERVAL_MS;
  }
  ENVOY_LOG_MISC(trace, "interval set to: {} millis", duration_ms);
  interval_ = std::chrono::milliseconds(duration_ms);
  timer_->enableTimer(interval_);
}

void SimpleAnonymousReporter::observeUsageOfHttpFilter() {}

void SimpleAnonymousReporter::observeUsageOfNetworkFilter() {}

void SimpleAnonymousReporter::sendPing() { sender_->send(ping_message_); }

void SimpleAnonymousReporter::onTimerTick() {
  sendPing();
  timer_->enableTimer(interval_);
}

} // namespace Konvoy
} // namespace Common
} // namespace Filters
} // namespace Extensions
} // namespace Envoy
