package mesh

import (
	"net"

	mesh_proto "github.com/Kong/kuma/api/mesh/v1alpha1"
)

var ipv4loopback = net.IPv4(127, 0, 0, 1)

func (d *DataplaneResource) UsesInterface(address net.IP, port uint32) bool {
	return d.UsesInboundInterface(address, port) || d.UsesOutboundInterface(address, port)
}

func (d *DataplaneResource) UsesInboundInterface(address net.IP, port uint32) bool {
	if d == nil {
		return false
	}
	for _, inbound := range d.Spec.Networking.GetInbound() {
		iface, err := mesh_proto.ParseInboundInterface(inbound.Interface)
		if err != nil {
			continue
		}
		// compare against port and IP address of the dataplane
		if port == iface.DataplanePort && overlap(address, net.ParseIP(iface.DataplaneIP)) {
			return true
		}
		// compare against port and IP address of the application
		if port == iface.WorkloadPort && overlap(address, ipv4loopback) {
			return true
		}
	}
	return false
}

func (d *DataplaneResource) UsesOutboundInterface(address net.IP, port uint32) bool {
	if d == nil {
		return false
	}
	for _, outbound := range d.Spec.Networking.GetOutbound() {
		oface, err := mesh_proto.ParseOutboundInterface(outbound.Interface)
		if err != nil {
			continue
		}
		// compare against port and IP address of the dataplane
		if port == oface.DataplanePort && overlap(address, net.ParseIP(oface.DataplaneIP)) {
			return true
		}
	}
	return false
}

func overlap(address1 net.IP, address2 net.IP) bool {
	if address1.IsUnspecified() || address2.IsUnspecified() {
		// wildcard match (either IPv4 address "0.0.0.0" or the IPv6 address "::")
		return true
	}
	// exact match
	return address1.Equal(address2)
}
