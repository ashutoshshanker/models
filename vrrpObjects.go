package models

/*
 * This DS will be used while Created/Deleting Vrrp Intf Config
 */
type VrrpIntf struct {
	ConfigObj
	IfIndex  int32 `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: ""Interface index for which VRRP Config needs to be done"`
	VRID     int32 `SNAPROUTE: "KEY", DESCRIPTION: "Virtual Router's Unique Identifier"`
	Priority int32 `DESCRIPTION: "Sending VRRP router's priority for
	   the virtual router", DEFAULT: "100"`
	VirtualIPv4Addr       string `DESCRIPTION: "Virtual Router Identifier"`
	AdvertisementInterval int32  `DESCRIPTION: "Time interval between ADVERTISEMENTS", DEFAULT:"1"`
	PreemptMode           bool   `DESCRIPTION: "Controls whether a (starting or restarting) higher-priority Backup router preempts a lower-priority Master router", DEFAULT: "true"`
	AcceptMode            bool   `DESCRIPTION: "Controls whether a virtual router in Master state will accept packets addressed to the address owner's IPvX address as its own if it is not the IPvX address owner.", DEFAULT:"false"`
}

type VrrpIntfState struct {
	ConfigObj
	IfIndex                 int32  `SNAPROUTE: "KEY", ACCESS:"r",  MULTIPLICITY:"*", DESCRIPTION: "Interface index for which VRRP state is requested"`
	VRID                    int32  `SNAPROUTE: "KEY", DESCRIPTION: "Virtual Router's Unique Identifier"`
	IntfIpAddr              string `DESCRIPTION: "Ip Address of Interface where VRRP is configured"`
	Priority                int32  `DESCRIPTION: "Virtual router's Priority"`
	VirtualIPv4Addr         string `DESCRIPTION: "Ip Address of Virtual Router"`
	AdvertisementInterval   int32  `DESCRIPTION: "Time interval between Advertisements"`
	PreemptMode             bool   `DESCRIPTION: "States Whether Preempt is Supported or not"`
	VirtualRouterMACAddress string `DESCRIPTION: "VRRP router's Mac Address"`
	SkewTime                int32  `DESCRIPTION: "Time to skew Master Down Interval"`
	MasterDownTimer         int32  `DESCRIPTION: "Time interval for Backup to declare Master down"`
	VrrpState               string `DESCRIPTION: "Current vrrp state i.e. backup or master"`
}

type VrrpVridState struct {
	ConfigObj
	IfIndex          int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION:"Interface index for which VRRP state is requested"`
	VRID             int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION:"Virtual Router's Unique Identifier""`
	AdverRx          int32  `DESCRIPTION:"Total number of advertisement packets received"`
	AdverTx          int32  `DESCRIPTION:"Total number of advertisement packets send"`
	LastAdverRx      string `DESCRIPTION:"Time when last advertisement packet was received"`
	LastAdverTx      string `DESCRIPTION:"Time when last advertisement packet was send out"`
	MasterIp         string `DESCRIPTION:"Ip Address of the Master VRRP"`
	CurrentState     string `DESCRIPTION:"Current State of Local VRRP"`
	PreviousState    string `DESCRIPTION:"Previous State of Local VRRP"`
	TransitionReason string `DESCRIPTION:"Reason why we moved from previous state -> current state"`
}
