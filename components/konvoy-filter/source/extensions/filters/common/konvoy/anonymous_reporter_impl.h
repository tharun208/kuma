#pragma once

#include "extensions/filters/common/konvoy/anonymous_reporter.h"

#include <cstdlib>

#include "envoy/event/dispatcher.h"
#include "envoy/network/address.h"
#include "envoy/network/dns.h"
#include "envoy/network/io_handle.h"
#include "envoy/runtime/runtime.h"
#include "envoy/singleton/instance.h"

namespace Envoy {
namespace Extensions {
namespace Filters {
namespace Common {
namespace Konvoy {

namespace {
// Environment varibles

// E.g., KONVOY_ANONYMOUS_REPORTS=off
const char KONVOY_ANONYMOUS_REPORTS[] = "KONVOY_ANONYMOUS_REPORTS";
// E.g., KONVOY_ANONYMOUS_REPORTS_ADDRESS=kong-hf.konghq.com
const char KONVOY_ANONYMOUS_REPORTS_HOST[] = "KONVOY_ANONYMOUS_REPORTS_HOST";
// E.g., KONVOY_ANONYMOUS_REPORTS_PORT=61829
const char KONVOY_ANONYMOUS_REPORTS_PORT[] = "KONVOY_ANONYMOUS_REPORTS_PORT";
// E.g., KONVOY_ANONYMOUS_REPORTS_INTERVAL=1200000
const char KONVOY_ANONYMOUS_REPORTS_INTERVAL[] = "KONVOY_ANONYMOUS_REPORTS_INTERVAL";

// Values

const char KONVOY_ANONYMOUS_REPORTS_DISABLED[] = "off";
const char KONVOY_ANONYMOUS_REPORTS_DEFAULT_HOST[] = "kong-hf.konghq.com";
const char KONVOY_ANONYMOUS_REPORTS_DEFAULT_PORT[] = "61829";
const uint32_t KONVOY_ANONYMOUS_REPORTS_DEFAULT_INTERVAL_MS = 1200000;
} // namespace

/**
 * Abstracts access to environment variables.
 */
using EnvVarSource = std::function<std::string(const std::string&)>;

/**
 * Returns the source backed by the actual environment variables.
 */
EnvVarSource std_env();

/**
 * Sends report messages.
 */
class Sender {
public:
  virtual ~Sender() {}

  virtual void send(const std::string& message) PURE;
};

using SenderPtr = std::unique_ptr<Sender>;

/**
 * Sends report messages over UDP.
 */
class UdpSender : public Sender {
public:
  explicit UdpSender(Network::Address::InstanceConstSharedPtr address);
  ~UdpSender();

  void send(const std::string& message) override;

private:
  Network::IoHandlePtr io_handle_;
};

/**
 * No-op version of the Sender.
 */
class NoOpSender : public Sender {
public:
  void send(const std::string&) override {}
};

/**
 * Abstracts cancelable ongoing operation.
 */
class Cancellable {
public:
  virtual ~Cancellable() {}

  /**
   * Cancel the ongoing operation.
   */
  virtual void cancel() PURE;

  /**
   * Mark operation as complete to prevent accidental cancelling in the future.
   */
  virtual void disarm() PURE;
};

using CancellablePtr = std::unique_ptr<Cancellable>;

/**
 * Wraps another object of a type that has cancel() method but does not derive from Cancellable.
 */
template <class C> class CancellableWrapper : public Cancellable {
public:
  explicit CancellableWrapper(C* cancellable) : cancellable_(cancellable) {}
  ~CancellableWrapper() { cancel(); }

  void cancel() override {
    if (cancellable_) {
      cancellable_->cancel();
      cancellable_ = nullptr;
    }
  }

  void disarm() override { cancellable_ = nullptr; }

private:
  C* cancellable_{};
};

/**
 * Abstracts construction of a Sender.
 *
 * Might be an asynchronous operation because of DNS resolution.
 */
class SenderFactory {
public:
  virtual ~SenderFactory() {}

  /**
   * Callback that delivers a Sender instance.
   */
  using Cb = std::function<void(SenderPtr)>;

  virtual CancellablePtr createSender(EnvVarSource env, Cb cb) PURE;
};

using SenderFactorySharedPtr = std::shared_ptr<SenderFactory>;

/**
 * Constructs a UdpSender.
 */
class UdpSenderFactory : public SenderFactory {
public:
  explicit UdpSenderFactory(Network::DnsResolverSharedPtr dns_resolver);

  CancellablePtr createSender(EnvVarSource env, SenderFactory::Cb cb) override;

private:
  Network::DnsResolverSharedPtr dns_resolver_;
};

/**
 * AnonymousReporter that sends report messages over UDP.
 */
class SimpleAnonymousReporter : public Singleton::Instance,
                                public Filters::Common::Konvoy::AnonymousReporter {
public:
  SimpleAnonymousReporter(EnvVarSource env, Event::Dispatcher& dispatcher,
                          Runtime::RandomGenerator& random);
  SimpleAnonymousReporter(EnvVarSource env, Event::Dispatcher& dispatcher,
                          Runtime::RandomGenerator& random, SenderFactorySharedPtr sender_factory);

  void observeUsageOfHttpFilter() override;
  void observeUsageOfNetworkFilter() override;

protected:
  static std::string createUniqueId(Runtime::RandomGenerator& random);
  static std::string createPingMessage(const std::string& unique_id);

private:
  void onTimerTick();
  void sendPing();

  const std::string unique_id_;
  const std::string ping_message_;
  SenderPtr sender_;
  const SenderFactorySharedPtr sender_factory_;
  CancellablePtr cancellable_;
  const Event::TimerPtr timer_;
  std::chrono::milliseconds interval_;
};

/**
 * No-op version of the AnonymousReporter.
 */
class NoOpAnonymousReporter : public Singleton::Instance,
                              public Filters::Common::Konvoy::AnonymousReporter {
public:
  void observeUsageOfHttpFilter() override {}
  void observeUsageOfNetworkFilter() override {}
};

} // namespace Konvoy
} // namespace Common
} // namespace Filters
} // namespace Extensions
} // namespace Envoy
