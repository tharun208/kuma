#include "extensions/filters/common/konvoy/utility.h"

#include "common/network/utility.h"

#include "absl/strings/match.h"

namespace Envoy {
namespace Extensions {
namespace Filters {
namespace Common {
namespace Konvoy {

// Singleton registration via macro defined in envoy/singleton/manager.h
SINGLETON_MANAGER_REGISTRATION(anonymous_reporter);

AnonymousReporterSharedPtr
Utility::anonymousReporter(Server::Configuration::FactoryContext& factory_context,
                           EnvVarSource env) {
  return factory_context.singletonManager().getTyped<AnonymousReporter>(
      SINGLETON_MANAGER_REGISTERED_NAME(anonymous_reporter),
      [&factory_context, &env]() -> Singleton::InstanceSharedPtr {
        bool disabled = absl::EqualsIgnoreCase(env(KONVOY_ANONYMOUS_REPORTS),
                                               KONVOY_ANONYMOUS_REPORTS_DISABLED);
        if (disabled) {
          ENVOY_LOG_MISC(trace, "anonymous reports have been disabled");
          return std::make_shared<NoOpAnonymousReporter>();
        } else {
          return std::make_shared<SimpleAnonymousReporter>(env, factory_context.dispatcher(),
                                                           factory_context.random());
        }
      });
}

} // namespace Konvoy
} // namespace Common
} // namespace Filters
} // namespace Extensions
} // namespace Envoy
