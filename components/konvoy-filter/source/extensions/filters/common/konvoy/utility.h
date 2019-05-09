#include "extensions/filters/common/konvoy/anonymous_reporter_impl.h"

#include "envoy/server/filter_config.h"

namespace Envoy {
namespace Extensions {
namespace Filters {
namespace Common {
namespace Konvoy {

class Utility {
public:
  /**
   * Return an anonymous reporter.
   */
  static AnonymousReporterSharedPtr
  anonymousReporter(Server::Configuration::FactoryContext& factory_context,
                    EnvVarSource env = std_env());
};

} // namespace Konvoy
} // namespace Common
} // namespace Filters
} // namespace Extensions
} // namespace Envoy
