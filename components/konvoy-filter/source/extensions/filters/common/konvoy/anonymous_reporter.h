#pragma once

#include <memory>

#include "envoy/common/pure.h"

namespace Envoy {
namespace Extensions {
namespace Filters {
namespace Common {
namespace Konvoy {

/**
 * Anonymous reporter.
 */
class AnonymousReporter {
public:
  virtual ~AnonymousReporter() {}

  /**
   * Observe usage of Konvoy Http Filter.
   */
  virtual void observeUsageOfHttpFilter() PURE;

  /**
   * Observe usage of Konvoy Network Filter.
   */
  virtual void observeUsageOfNetworkFilter() PURE;
};

using AnonymousReporterSharedPtr = std::shared_ptr<AnonymousReporter>;

} // namespace Konvoy
} // namespace Common
} // namespace Filters
} // namespace Extensions
} // namespace Envoy
