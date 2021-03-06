//
//Copyright [2016] [SnapRoute Inc]
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//	 Unless required by applicable law or agreed to in writing, software
//	 distributed under the License is distributed on an "AS IS" BASIS,
//	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	 See the License for the specific language governing permissions and
//	 limitations under the License.
//
// _______  __       __________   ___      _______.____    __    ____  __  .___________.  ______  __    __
// |   ____||  |     |   ____\  \ /  /     /       |\   \  /  \  /   / |  | |           | /      ||  |  |  |
// |  |__   |  |     |  |__   \  V  /     |   (----` \   \/    \/   /  |  | `---|  |----`|  ,----'|  |__|  |
// |   __|  |  |     |   __|   >   <       \   \      \            /   |  |     |  |     |  |     |   __   |
// |  |     |  `----.|  |____ /  .  \  .----)   |      \    /\    /    |  |     |  |     |  `----.|  |  |  |
// |__|     |_______||_______/__/ \__\ |_______/        \__/  \__/     |__|     |__|      \______||__|  |__|
//

package models

type Vlan struct {
	baseObj
	VlanId        int32  `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY: "*", MIN:"1", MAX: "4094", DESCRIPTION: "802.1Q tag/Vlan ID for vlan being provisioned"`
	IntfList      string `DESCRIPTION: "List of interface names or ifindex values to  be added as tagged members of the vlan", DEFAULT:""`
	UntagIntfList string `DESCRIPTION: "List of interface names or ifindex values to  be added as untagged members of the vlan", DEFAULT:""`
}

type VlanState struct {
	baseObj
	VlanId    int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY: "*", DESCRIPTION: "802.1Q tag/Vlan ID for vlan being provisioned"`
	VlanName  string `DESCRIPTION: "System assigned vlan name"`
	OperState string `DESCRIPTION: "Operational state of vlan interface"`
	IfIndex   int32  `DESCRIPTION: "System assigned interface id for this vlan interface"`
}

type IPv4Intf struct {
	baseObj
	IntfRef string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION: "Interface name or ifindex of port/lag or vlan on which this IPv4 object is configured"`
	IpAddr  string `DESCRIPTION: "Interface IP/Net mask in CIDR format to provision on switch interface", STRLEN:"18"`
}

type IPv4IntfState struct {
	baseObj
	IntfRef           string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "System assigned interface id of L2 interface (port/lag/vlan) to which this IPv4 object is linked"`
	IfIndex           int32  `DESCRIPTION: "System assigned interface id for this IPv4 interface"`
	IpAddr            string `DESCRIPTION: "Interface IP/Net mask in CIDR format to provision on switch interface"`
	OperState         string `DESCRIPTION: "Operational state of this IP interface"`
	NumUpEvents       int32  `DESCRIPTION: "Number of times the operational state transitioned from DOWN to UP"`
	LastUpEventTime   string `DESCRIPTION: "Timestamp corresponding to the last DOWN to UP operational state change event"`
	NumDownEvents     int32  `DESCRIPTION: "Number of times the operational state transitioned from UP to DOWN"`
	LastDownEventTime string `DESCRIPTION: "Timestamp corresponding to the last UP to DOWN operational state change event"`
	L2IntfType        string `DESCRIPTION: "Type of L2 interface on which IP has been configured (Port/Lag/Vlan)"`
	L2IntfId          int32  `DESCRIPTION: "Id of the L2 interface. Port number/lag id/vlan id."`
}

type Port struct {
	baseObj
	IntfRef      string `SNAPROUTE: "KEY", ACCESS:"rw", DESCRIPTION: "Front panel port name or system assigned interface id"`
	IfIndex      int32  `DESCRIPTION: "System assigned interface id for this port. Read only attribute"`
	Description  string `DESCRIPTION: "User provided string description", DEFAULT:"FP Port", STRLEN:"64"`
	PhyIntfType  string `DESCRIPTION: "Type of internal phy interface", STRLEN:"16" SELECTION: GMII/SGMII/QSMII/SFI/XFI/XAUI/XLAUI/RXAUI/CR/CR2/CR4/KR/KR2/KR4/SR/SR2/SR4/SR10/LR/LR4`
	AdminState   string `DESCRIPTION: "Administrative state of this port", STRLEN:"4" SELECTION: ON/OFF`
	MacAddr      string `DESCRIPTION: "Mac address associated with this port", STRLEN:"17"`
	Speed        int32  `DESCRIPTION: "Port speed in Mbps", MIN: 10, MAX: "100000"`
	Duplex       string `DESCRIPTION: "Duplex setting for this port", STRLEN:"16" SELECTION: Half Duplex/Full Duplex`
	Autoneg      string `DESCRIPTION: "Autonegotiation setting for this port", STRLEN:"4" SELECTION: ON/OFF`
	MediaType    string `DESCRIPTION: "Type of media inserted into this port", STRLEN:"16"`
	Mtu          int32  `DESCRIPTION: "Maximum transmission unit size for this port"`
	BreakOutMode string `DESCRIPTION: "Break out mode for the port. Only applicable on ports that support breakout. Valid modes - 1x40, 4x10", STRLEN:"6" SELECTION: 1x40(1)/4x10(2)`
}

type PortState struct {
	baseObj
	IntfRef           string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "Front panel port name or system assigned interface id"`
	IfIndex           int32  `DESCRIPTION: "System assigned interface id for this port"`
	Name              string `DESCRIPTION: "System assigned vlan name"`
	OperState         string `DESCRIPTION: "Operational state of front panel port"`
	NumUpEvents       int32  `DESCRIPTION: "Number of times the operational state transitioned from DOWN to UP"`
	LastUpEventTime   string `DESCRIPTION: "Timestamp corresponding to the last DOWN to UP operational state change event"`
	NumDownEvents     int32  `DESCRIPTION: "Number of times the operational state transitioned from UP to DOWN"`
	LastDownEventTime string `DESCRIPTION: "Timestamp corresponding to the last UP to DOWN operational state change event"`
	Pvid              int32  `DESCRIPTION: "The vlanid assigned to untagged traffic ingressing this port"`
	IfInOctets        int64  `DESCRIPTION: "RFC2233 Total number of octets received on this port"`
	IfInUcastPkts     int64  `DESCRIPTION: "RFC2233 Total number of unicast packets received on this port"`
	IfInDiscards      int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that were discarded"`
	IfInErrors        int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that contained an error"`
	IfInUnknownProtos int64  `DESCRIPTION: "RFC2233 Total number of inbound packets discarded due to unknown protocol"`
	IfOutOctets       int64  `DESCRIPTION: "RFC2233 Total number of octets transmitted on this port"`
	IfOutUcastPkts    int64  `DESCRIPTION: "RFC2233 Total number of unicast packets transmitted on this port"`
	IfOutDiscards     int64  `DESCRIPTION: "RFC2233 Total number of error free packets discarded and not transmitted"`
	IfOutErrors       int64  `DESCRIPTION: "RFC2233 Total number of packets discarded and not transmitted due to packet errors"`
	ErrDisableReason  string `DESCRIPTION: "Reason explaining why port has been disabled by protocol code"`
	PresentInHW       string `DESCRIPTION: "Indication of whether this port object maps to a physical port. Set to 'No' for ports that are not broken out."`
}

type MacTableEntryState struct {
	baseObj
	MacAddr string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "MAC Address", USESTATEDB:"true"`
	VlanId  int32  `DESCRIPTION: "Vlan id corresponding to which mac was learned", DEFAULT:0`
	Port    int32  `DESCRIPTION: "Port number on which mac was learned", DEFAULT:0`
}

type IPv4RouteHwState struct {
	baseObj
	DestinationNw    string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "IP address of the route in CIDR format"`
	NextHopIps       string `DESCRIPTION: "next hop ip list for the route"`
	RouteCreatedTime string `DESCRIPTION :"Time when the route was added"`
	RouteUpdatedTime string `DESCRIPTION :"Time when the route was last updated"`
}

type ArpEntryHwState struct {
	baseObj
	IpAddr  string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", QPARAM: "optional" ,DESCRIPTION: "Neighbor's IP Address"`
	MacAddr string `DESCRIPTION: "MAC address of the neighbor machine with corresponding IP Address", QPARAM: "optional" `
	Vlan    string `DESCRIPTION: "Vlan ID of the Router Interface to which neighbor is attached to", QPARAM: "optional" `
	Port    string `DESCRIPTION: "Router Interface to which neighbor is attached to", QPARAM: "optional" `
}

type LogicalIntf struct {
	baseObj
	Name string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION: "Name of logical interface"`
	Type string `DESCRIPTION: "Type of logical interface (e.x. loopback)", DEFAULT:"Loopback", STRLEN:"16"`
}

type LogicalIntfState struct {
	baseObj
	Name              string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: "Name of logical interface"`
	IfIndex           int32  `DESCRIPTION: "System assigned interface id for this logical interface"`
	SrcMac            string `DESCRIPTION: "Source Mac assigned to the interface"`
	OperState         string `DESCRIPTION: "Operational state of logical interface"`
	IfInOctets        int64  `DESCRIPTION: "RFC2233 Total number of octets received on this port"`
	IfInUcastPkts     int64  `DESCRIPTION: "RFC2233 Total number of unicast packets received on this port"`
	IfInDiscards      int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that were discarded"`
	IfInErrors        int64  `DESCRIPTION: "RFC2233 Total number of inbound packets that contained an error"`
	IfInUnknownProtos int64  `DESCRIPTION: "RFC2233 Total number of inbound packets discarded due to unknown protocol"`
	IfOutOctets       int64  `DESCRIPTION: "RFC2233 Total number of octets transmitted on this port"`
	IfOutUcastPkts    int64  `DESCRIPTION: "RFC2233 Total number of unicast packets transmitted on this port"`
	IfOutDiscards     int64  `DESCRIPTION: "RFC2233 Total number of error free packets discarded and not transmitted"`
	IfOutErrors       int64  `DESCRIPTION: "RFC2233 Total number of packets discarded and not transmitted due to packet errors"`
}

type SubIPv4Intf struct {
	baseObj
	IpAddr  string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION:"Ip Address for the interface"`
	IntfRef string `SNAPROUTE: "KEY", ACCESS:"w", DESCRIPTION:"Intf name of system generated id (ifindex) of the ipv4Intf where sub interface is to be configured"`
	Type    string `DESCRIPTION:"Type of interface, e.g. Secondary or Virtual", STRLEN:"16"`
	MacAddr string `DESCRIPTION:"Mac address to be used for the sub interface. If none specified IPv4Intf mac address will be used", STRLEN:"17"`
	Enable  bool   `DESCRIPTION:"Enable or disable this interface", DEFAULT:false`
}
