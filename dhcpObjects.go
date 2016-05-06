package models

/*
 * This DS will be used while Created/Deleting DHCP Config
 */
type DhcpGlobalConfig struct {
	ConfigObj
	// placeholder to create a key
	DhcpConfigKey    string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"1", DESCRIPTION: "DHCP global config"`
	Enable           bool   `DESCRIPTION: "DHCP Server enable/disable control DEFAULT: false"`
	DefaultLeaseTime uint32 `DESCRIPTION: "Default Lease Time in seconds DEFAULT: 600"`
	MaxLeaseTime     uint32 `DESCRIPTION: "Max Lease Time in seconds DEFAULT: 7200"`
}

type DhcpIntfConfig struct {
	ConfigObj
	// placeholder to create a key
	IntfRef       string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Interface name or ifindex of L3 interface object on which Dhcp Server need to be configured"`
	Subnet        string `DESCRIPTION: "Subnet"`
	SubnetMask    string `DESCRIPTION: "Subnet Mask"`
	IPAddrRange   string `DESCRIPTION: "Range of IP Addresses DEFAULT: All the IP Addresses in the given subnet"`
	BroadcastAddr string `DESCRIPTION: "Broadcast Address DEFAULT: Last IP Addresses in the given subnet"`
	RouterAddr    string `DESCRIPTION: "Router Address DEFAULT: First IP Addresses in the given subnet"`
	DNSServerAddr string `DESCRIPTION: "Comma seperated List of DNS Server Address DEFAULT: none"`
	DomainName    string `DESCRIPTION: "Domain Name Address DEFAULT: none"`
	Enable        bool   `DESCRIPTION: "Enable and Disable Control DEFAULT: false"`
}
