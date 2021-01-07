package v3_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	mesh_core "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"

	. "github.com/kumahq/kuma/pkg/xds/envoy/listeners"

	util_proto "github.com/kumahq/kuma/pkg/util/proto"
	envoy_common "github.com/kumahq/kuma/pkg/xds/envoy"
)

var _ = Describe("UDPProxyConfigurer", func() {

	type testCase struct {
		listenerName     string
		listenerAddress  string
		listenerPort     uint32
		listenerProtocol mesh_core.Protocol
		statsName        string
		cluster          envoy_common.ClusterSubset
		expected         string
	}

	DescribeTable("should generate proper Envoy config",
		func(given testCase) {
			// when
			listener, err := NewListenerBuilder(envoy_common.APIV3).
				Configure(InboundListener(given.listenerName, given.listenerAddress, given.listenerPort, given.listenerProtocol)).
				Configure(UDPProxy(given.statsName, given.cluster)).
				Build()
			// then
			Expect(err).ToNot(HaveOccurred())

			// when
			actual, err := util_proto.ToYAML(listener)
			Expect(err).ToNot(HaveOccurred())
			// and
			Expect(actual).To(MatchYAML(given.expected))
		},
		Entry("basic udp_proxy with a single destination cluster", testCase{
			listenerName:     "inbound:192.168.0.1:8080",
			listenerAddress:  "192.168.0.1",
			listenerPort:     8080,
			listenerProtocol: mesh_core.ProtocolUDP,
			statsName:        "localhost:8080",
			cluster: envoy_common.ClusterSubset{
				ClusterName: "localhost:8080",
			},
			expected: `
        name: inbound:192.168.0.1:8080
        reusePort: true
        trafficDirection: INBOUND
        address:
          socketAddress:
            address: 192.168.0.1
            portValue: 8080
            protocol: UDP
        listenerFilters:
          - name: envoy.filters.udp_listener.udp_proxy
            typedConfig:
              '@type': type.googleapis.com/envoy.extensions.filters.udp.udp_proxy.v3.UdpProxyConfig
              cluster: localhost:8080
              statPrefix: localhost_8080
`,
		}),
	)
})
