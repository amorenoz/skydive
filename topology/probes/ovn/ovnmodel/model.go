// Code generated by ovnmetagen
// DO NOT EDIT.

package ovnmodel

import (
	"github.com/ovn-org/libovsdb/model"

	"github.com/skydive-project/skydive/graffiti/graph"
)

// FullDatabaseModel returns the DatabaseModel object to be used in libovsdb
func FullDatabaseModel() (*model.DBModel, error) {
	return model.NewDBModel("OVN_Northbound", map[string]model.Model{
		"ACL":                         &ACL{},
		"Address_Set":                 &AddressSet{},
		"Connection":                  &Connection{},
		"DHCP_Options":                &DHCPOptions{},
		"DNS":                         &DNS{},
		"Forwarding_Group":            &ForwardingGroup{},
		"Gateway_Chassis":             &GatewayChassis{},
		"HA_Chassis":                  &HAChassis{},
		"HA_Chassis_Group":            &HAChassisGroup{},
		"Load_Balancer":               &LoadBalancer{},
		"Load_Balancer_Health_Check":  &LoadBalancerHealthCheck{},
		"Logical_Router":              &LogicalRouter{},
		"Logical_Router_Policy":       &LogicalRouterPolicy{},
		"Logical_Router_Port":         &LogicalRouterPort{},
		"Logical_Router_Static_Route": &LogicalRouterStaticRoute{},
		"Logical_Switch":              &LogicalSwitch{},
		"Logical_Switch_Port":         &LogicalSwitchPort{},
		"Meter":                       &Meter{},
		"Meter_Band":                  &MeterBand{},
		"NAT":                         &NAT{},
		"NB_Global":                   &NBGlobal{},
		"Port_Group":                  &PortGroup{},
		"QoS":                         &QoS{},
		"SSL":                         &SSL{},
	})
}

// Decoders returns all the decoder functions indexed by Name
func Decoders() map[string]graph.MetadataDecoder {
	return map[string]graph.MetadataDecoder{

		"ACL": ACLDecoder,

		"AddressSet": AddressSetDecoder,

		"Connection": ConnectionDecoder,

		"DHCPOptions": DHCPOptionsDecoder,

		"DNS": DNSDecoder,

		"ForwardingGroup": ForwardingGroupDecoder,

		"GatewayChassis": GatewayChassisDecoder,

		"HAChassis": HAChassisDecoder,

		"HAChassisGroup": HAChassisGroupDecoder,

		"LoadBalancer": LoadBalancerDecoder,

		"LoadBalancerHealthCheck": LoadBalancerHealthCheckDecoder,

		"LogicalRouter": LogicalRouterDecoder,

		"LogicalRouterPolicy": LogicalRouterPolicyDecoder,

		"LogicalRouterPort": LogicalRouterPortDecoder,

		"LogicalRouterStaticRoute": LogicalRouterStaticRouteDecoder,

		"LogicalSwitch": LogicalSwitchDecoder,

		"LogicalSwitchPort": LogicalSwitchPortDecoder,

		"Meter": MeterDecoder,

		"MeterBand": MeterBandDecoder,

		"NAT": NATDecoder,

		"NBGlobal": NBGlobalDecoder,

		"PortGroup": PortGroupDecoder,

		"QoS": QoSDecoder,

		"SSL": SSLDecoder,
	}
}