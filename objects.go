package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type ConfigObj interface {
	UnmarshalObject(data []byte) (ConfigObj, error)
	CreateDBTable(dbHdl *sql.DB) error
	StoreObjectInDb(dbHdl *sql.DB) (int64, error)
	DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error
	GetKey() (string, error)
	GetSqlKeyStr(string) (string, error)
	GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error)
	CompareObjectsAndDiff(updateKeys map[string]bool, dbObj ConfigObj) ([]bool, error)
	MergeDbAndConfigObj(dbObj ConfigObj, attrSet []bool) (ConfigObj, error)
	UpdateObjectInDb(dbV4Route ConfigObj, attrSet []bool, dbHdl *sql.DB) error
}

//
// This file is handcoded for now. Eventually this would be generated by yang compiler
//
type IPV4Route struct {
	BaseObj
	DestinationNw     string `SNAPROUTE: "KEY"`
	NetworkMask       string `SNAPROUTE: "KEY"`
	Cost              uint32
	NextHopIp         string
	OutgoingIntfType  string
	OutgoingInterface string
	Protocol          string
}

func (obj IPV4Route) UnmarshalObject(body []byte) (ConfigObj, error) {
	var v4Route IPV4Route
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &v4Route); err != nil {
			fmt.Println("### Trouble in unmarshaling IPV4Route from Json", body)
		}
	}
	return v4Route, err
}

type BGPGlobalConfig struct {
	BaseObj
	ASNum    uint32
	RouterId string `SNAPROUTE: "KEY"`
}

func (obj BGPGlobalConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf BGPGlobalConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling BPGlobalConfig from Json", body)
		}
	}
	return gConf, err
}

type BGPGlobalState struct {
	BaseObj
	AS            uint32
	RouterId      string
	TotalPaths    uint32
	TotalPrefixes uint32
}

func (obj BGPGlobalState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gState BGPGlobalState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gState); err != nil {
			fmt.Println("### Trouble in unmarshalling BPGlobalConfig from Json", body)
		}
	}
	return gState, err
}

type PeerType int

const (
	PeerTypeInternal PeerType = iota
	PeerTypeExternal
)

type BgpCounters struct {
	Update       uint64
	Notification uint64
}

type BGPMessages struct {
	Sent     BgpCounters
	Received BgpCounters
}

type BGPQueues struct {
	Input  uint32
	Output uint32
}

type BGPNeighborConfig struct {
	BaseObj
	PeerAS                  uint32
	LocalAS                 uint32
	AuthPassword            string
	Description             string
	NeighborAddress         string `SNAPROUTE: "KEY"`
	RouteReflectorClusterId uint32
	RouteReflectorClient    bool
	MultiHopEnable          bool
	MultiHopTTL             uint8
}

func (obj BGPNeighborConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var nConf BGPNeighborConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &nConf); err != nil {
			fmt.Println("### Trouble in unmarshaling BGPNeighborConfig from Json", body)
		}
	}
	return nConf, err
}

type BGPNeighborState struct {
	BaseObj
	PeerAS                  uint32
	LocalAS                 uint32
	PeerType                PeerType
	AuthPassword            string
	Description             string
	NeighborAddress         string
	SessionState            uint32
	Messages                BGPMessages
	Queues                  BGPQueues
	RouteReflectorClusterId uint32
	RouteReflectorClient    bool
	MultiHopEnable          bool
	MultiHopTTL             uint8
}

func (obj BGPNeighborState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var nState BGPNeighborState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &nState); err != nil {
			fmt.Println("### Trouble in unmarshaling BGPNeighborConfig from Json", body)
		}
	}
	return nState, err
}

type VlanConfig struct {
	BaseObj
	VlanId           int32 `SNAPROUTE: "KEY"`
	IfIndexList      string
	UntagIfIndexList string
}

type VlanState struct {
	BaseObj
	VlanId    int32
	IfIndex   int32
	VlanName  string
	OperState string
}

type IPv4Intf struct {
	BaseObj
	IpAddr  string `SNAPROUTE: "KEY"`
	IfIndex int32
}

func (obj VlanConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var vlanObj VlanConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &vlanObj); err != nil {
			fmt.Println("### Trouble in unmarshaling VlanConfig from Json", body)
		}
	}
	return vlanObj, err
}

func (obj VlanState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var vlanObj VlanState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &vlanObj); err != nil {
			fmt.Println("### Trouble in unmarshaling VlanState from Json", body)
		}
	}
	return vlanObj, err
}

func (obj IPv4Intf) UnmarshalObject(body []byte) (ConfigObj, error) {
	var v4Intf IPv4Intf
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &v4Intf); err != nil {
			fmt.Println("### Trouble in unmarshaling IPv4Intf from Json", body)
		}
	}
	return v4Intf, err
}

/* ARP */
type ArpConfig struct {
	BaseObj
	// placeholder to create a key
	ArpConfigKey string `SNAPROUTE: "KEY"`
	Timeout      int32
}

func (obj ArpConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var arpConfigObj ArpConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &arpConfigObj); err != nil {
			fmt.Println("### Trouble in unmarshaling ArpConfig from Json", body)
		}
	}

	return arpConfigObj, err
}

type ArpEntry struct {
	BaseObj
	IpAddr         string `SNAPROUTE: "KEY"`
	MacAddr        string
	Vlan           uint32
	Intf           string
	ExpiryTimeLeft string
}

func (obj ArpEntry) UnmarshalObject(body []byte) (ConfigObj, error) {
	var arpEntryObj ArpEntry
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &arpEntryObj); err != nil {
			fmt.Println("### Trouble in unmarshaling ArpConfig from Json", body)
		}
	}
	return arpEntryObj, err
}

/* PortInterface */
type PortConfig struct {
	BaseObj
	IfIndex     int32 `SNAPROUTE: "KEY"`
	Name        string
	Description string
	Type        string
	AdminState  string
	OperState   string
	MacAddr     string
	Speed       int32
	Duplex      string
	Autoneg     string
	MediaType   string
	Mtu         int32
}

func (obj PortConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var portConfigObj PortConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &portConfigObj); err != nil {
			fmt.Println("### Trouble in unmarshaling PortIntfConfig from Json", body)
		}
	}

	return portConfigObj, err
}

type PortState struct {
	BaseObj
	IfIndex   int32
	PortStats []int64
}

func (obj PortState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var portStateObj PortState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &portStateObj); err != nil {
			fmt.Println("### Trouble in unmarshaling PortIntfState from Json", body)
		}
	}

	return portStateObj, err
}

type UserConfig struct {
	BaseObj
	UserName    string `SNAPROUTE: "KEY"`
	Password    string
	Description string
	Previledge  string
}

func (obj UserConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var userConfigObj UserConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &userConfigObj); err != nil {
			fmt.Println("### Trouble in unmarshaling UserConfig from Json", body)
		}
	}

	return userConfigObj, err
}

type UserState struct {
	BaseObj
	UserName      string
	LastLoginTime time.Time
	LastLoginIp   string
	NumAPICalled  uint32
}

func (obj UserState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var userStateObj UserState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &userStateObj); err != nil {
			fmt.Println("### Trouble in unmarshaling UserState from Json", body)
		}
	}

	return userStateObj, err
}

type OspfAreaEntryConfig struct {
	BaseObj
	AreaIdKey                           string `SNAPROUTE: "KEY"`
	AuthType                            int32
	ImportAsExtern                      int32
	AreaSummary                         int32
	AreaNssaTranslatorRole              int32
	AreaNssaTranslatorStabilityInterval int32
}

func (obj OspfAreaEntryConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfAreaEntryConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfAreaEntryConfig from Json", body)
		}
	}
	return gConf, err
}

type OspfAreaEntryState struct {
	BaseObj
	AreaIdKey                string `SNAPROUTE: "KEY"`
	SpfRuns                  uint32
	AreaBdrRtrCount          uint32
	AsBdrRtrCount            uint32
	AreaLsaCount             uint32
	AreaLsaCksumSum          int32
	AreaNssaTranslatorState  int32
	AreaNssaTranslatorEvents uint32
}

func (obj OspfAreaEntryState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfAreaEntryState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfAreaEntryState from Json", body)
		}
	}
	return gConf, err
}

type OspfStubAreaEntryConfig struct {
	BaseObj
	StubTOSKey     int32  `SNAPROUTE: "KEY"`
	StubAreaIdKey  string `SNAPROUTE: "KEY"`
	StubMetric     int32
	StubMetricType int32
}

func (obj OspfStubAreaEntryConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfStubAreaEntryConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfStubAreaEntryConfig from Json", body)
		}
	}
	return gConf, err
}

type OspfLsdbEntryState struct {
	BaseObj
	LsdbRouterIdKey   string `SNAPROUTE: "KEY"`
	LsdbAreaIdKey     string `SNAPROUTE: "KEY"`
	LsdbLsidKey       string `SNAPROUTE: "KEY"`
	LsdbTypeKey       int32  `SNAPROUTE: "KEY"`
	LsdbSequence      int32
	LsdbAge           int32
	LsdbChecksum      int32
	LsdbAdvertisement string
}

func (obj OspfLsdbEntryState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfLsdbEntryState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfLsdbEntryState from Json", body)
		}
	}
	return gConf, err
}

type OspfAreaRangeEntryConfig struct {
	BaseObj
	AreaRangeAreaIdKey string `SNAPROUTE: "KEY"`
	AreaRangeNetKey    string `SNAPROUTE: "KEY"`
	AreaRangeMask      string `SNAPROUTE: "KEY"`
	AreaRangeEffect    int32
}

func (obj OspfAreaRangeEntryConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfAreaRangeEntryConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfAreaRangeEntryConfig from Json", body)
		}
	}
	return gConf, err
}

type OspfHostEntryConfig struct {
	BaseObj
	HostIpAddressKey string `SNAPROUTE: "KEY"`
	HostTOSKey       int32  `SNAPROUTE: "KEY"`
	HostMetric       int32
	HostCfgAreaID    string
}

func (obj OspfHostEntryConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfHostEntryConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfHostEntryConfig from Json", body)
		}
	}
	return gConf, err
}

type OspfHostEntryState struct {
	BaseObj
	HostIpAddressKey string `SNAPROUTE: "KEY"`
	HostTOSKey       int32  `SNAPROUTE: "KEY"`
	HostAreaID       string
}

func (obj OspfHostEntryState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfHostEntryState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfHostEntryState from Json", body)
		}
	}
	return gConf, err
}

type OspfIfEntryConfig struct {
	BaseObj
	IfIpAddressKey        string `SNAPROUTE: "KEY"`
	AddressLessIfKey      int32  `SNAPROUTE: "KEY"`
	IfAreaId              string
	IfType                int32
	IfAdminStat           int32
	IfRtrPriority         int32
	IfTransitDelay        int32
	IfRetransInterval     int32
	IfHelloInterval       int32
	IfRtrDeadInterval     int32
	IfPollInterval        int32
	IfAuthKey             string
	IfMulticastForwarding int32
	IfDemand              bool
	IfAuthType            int32
}

func (obj OspfIfEntryConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfIfEntryConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfIfEntryConfig from Json", body)
		}
	}
	return gConf, err
}

type OspfIfEntryState struct {
	BaseObj
	IfIpAddressKey             string `SNAPROUTE: "KEY"`
	AddressLessIfKey           int32  `SNAPROUTE: "KEY"`
	IfState                    int32
	IfDesignatedRouter         string
	IfBackupDesignatedRouter   string
	IfEvents                   uint32
	IfLsaCount                 uint32
	IfLsaCksumSum              uint32
	IfDesignatedRouterId       string
	IfBackupDesignatedRouterId string
}

func (obj OspfIfEntryState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfIfEntryState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfIfEntryState from Json", body)
		}
	}
	return gConf, err
}

type OspfIfMetricEntryConfig struct {
	BaseObj
	IfMetricAddressLessIfKey int32  `SNAPROUTE: "KEY"`
	IfMetricTOSKey           int32  `SNAPROUTE: "KEY"`
	IfMetricIpAddressKey     string `SNAPROUTE: "KEY"`
	IfMetricValue            int32
}

func (obj OspfIfMetricEntryConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfIfMetricEntryConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfIfMetricEntryConfig from Json", body)
		}
	}
	return gConf, err
}

type OspfVirtIfEntryConfig struct {
	BaseObj
	VirtIfNeighborKey     string `SNAPROUTE: "KEY"`
	VirtIfAreaIdKey       string `SNAPROUTE: "KEY"`
	VirtIfTransitDelay    int32
	VirtIfRetransInterval int32
	VirtIfHelloInterval   int32
	VirtIfRtrDeadInterval int32
	VirtIfAuthKey         string
	VirtIfAuthType        int32
}

func (obj OspfVirtIfEntryConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfVirtIfEntryConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfVirtIfEntryConfig from Json", body)
		}
	}
	return gConf, err
}

type OspfNbrEntryConfig struct {
	BaseObj
	NbrAddressLessIndexKey int32  `SNAPROUTE: "KEY"`
	NbrIpAddrKey           string `SNAPROUTE: "KEY"`
	NbrPriority            int32
}

func (obj OspfNbrEntryConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfNbrEntryConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfNbrEntryConfig from Json", body)
		}
	}
	return gConf, err
}

type OspfNbrEntryState struct {
	BaseObj
	NbrAddressLessIndexKey     int32  `SNAPROUTE: "KEY"`
	NbrIpAddrKey               string `SNAPROUTE: "KEY"`
	NbrRtrId                   string
	NbrOptions                 int32
	NbrState                   int32
	NbrEvents                  uint32
	NbrLsRetransQLen           uint32
	NbmaNbrPermanence          int32
	NbrHelloSuppressed         bool
	NbrRestartHelperStatus     int32
	NbrRestartHelperAge        uint32
	NbrRestartHelperExitReason int32
}

func (obj OspfNbrEntryState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfNbrEntryState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfNbrEntryState from Json", body)
		}
	}
	return gConf, err
}

type OspfVirtNbrEntryState struct {
	BaseObj
	VirtNbrRtrIdKey                string `SNAPROUTE: "KEY"`
	VirtNbrAreaKey                 string `SNAPROUTE: "KEY"`
	VirtNbrIpAddr                  string
	VirtNbrOptions                 int32
	VirtNbrState                   int32
	VirtNbrEvents                  uint32
	VirtNbrLsRetransQLen           uint32
	VirtNbrHelloSuppressed         bool
	VirtNbrRestartHelperStatus     int32
	VirtNbrRestartHelperAge        uint32
	VirtNbrRestartHelperExitReason int32
}

func (obj OspfVirtNbrEntryState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfVirtNbrEntryState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfVirtNbrEntryState from Json", body)
		}
	}
	return gConf, err
}

type OspfExtLsdbEntryState struct {
	BaseObj
	ExtLsdbTypeKey       int32  `SNAPROUTE: "KEY"`
	ExtLsdbRouterIdKey   string `SNAPROUTE: "KEY"`
	ExtLsdbLsidKey       string `SNAPROUTE: "KEY"`
	ExtLsdbSequence      int32
	ExtLsdbAge           int32
	ExtLsdbChecksum      int32
	ExtLsdbAdvertisement string
}

func (obj OspfExtLsdbEntryState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfExtLsdbEntryState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfExtLsdbEntryState from Json", body)
		}
	}
	return gConf, err
}

type OspfAreaAggregateEntryConfig struct {
	BaseObj
	AreaAggregateMaskKey     string `SNAPROUTE: "KEY"`
	AreaAggregateLsdbTypeKey int32  `SNAPROUTE: "KEY"`
	AreaAggregateNetKey      string `SNAPROUTE: "KEY"`
	AreaAggregateAreaIDKey   string `SNAPROUTE: "KEY"`
	AreaAggregateEffect      int32
	AreaAggregateExtRouteTag uint32
}

func (obj OspfAreaAggregateEntryConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfAreaAggregateEntryConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfAreaAggregateEntryConfig from Json", body)
		}
	}
	return gConf, err
}

type OspfLocalLsdbEntryState struct {
	BaseObj
	LocalLsdbTypeKey          int32  `SNAPROUTE: "KEY"`
	LocalLsdbAddressLessIfKey int32  `SNAPROUTE: "KEY"`
	LocalLsdbRouterIdKey      string `SNAPROUTE: "KEY"`
	LocalLsdbLsidKey          string `SNAPROUTE: "KEY"`
	LocalLsdbIpAddressKey     string
	LocalLsdbLsid             string
	LocalLsdbRouterId         string
	LocalLsdbSequence         int32
	LocalLsdbAge              int32
	LocalLsdbChecksum         int32
	LocalLsdbAdvertisement    string
}

func (obj OspfLocalLsdbEntryState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfLocalLsdbEntryState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfLocalLsdbEntryState from Json", body)
		}
	}
	return gConf, err
}

type OspfVirtLocalLsdbEntryState struct {
	BaseObj
	VirtLocalLsdbTypeKey        int32  `SNAPROUTE: "KEY"`
	VirtLocalLsdbLsidKey        string `SNAPROUTE: "KEY"`
	VirtLocalLsdbTransitAreaKey string `SNAPROUTE: "KEY"`
	VirtLocalLsdbNeighborKey    string `SNAPROUTE: "KEY"`
	VirtLocalLsdbRouterIdKey    string `SNAPROUTE: "KEY"`
	VirtLocalLsdbNeighbor       string
	VirtLocalLsdbSequence       int32
	VirtLocalLsdbAge            int32
	VirtLocalLsdbChecksum       int32
	VirtLocalLsdbAdvertisement  string
}

func (obj OspfVirtLocalLsdbEntryState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfVirtLocalLsdbEntryState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfVirtLocalLsdbEntryState from Json", body)
		}
	}
	return gConf, err
}

type OspfAsLsdbEntryState struct {
	BaseObj
	AsLsdbRouterIdKey   string `SNAPROUTE: "KEY"`
	AsLsdbLsidKey       string `SNAPROUTE: "KEY"`
	AsLsdbTypeKey       int32  `SNAPROUTE: "KEY"`
	AsLsdbSequence      int32
	AsLsdbAge           int32
	AsLsdbChecksum      int32
	AsLsdbAdvertisement string
}

func (obj OspfAsLsdbEntryState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfAsLsdbEntryState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfAsLsdbEntryState from Json", body)
		}
	}
	return gConf, err
}

type OspfAreaLsaCountEntryState struct {
	BaseObj
	AreaLsaCountLsaTypeKey int32  `SNAPROUTE: "KEY"`
	AreaLsaCountAreaIdKey  string `SNAPROUTE: "KEY"`
	AreaLsaCountNumber     uint32
}

func (obj OspfAreaLsaCountEntryState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfAreaLsaCountEntryState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfAreaLsaCountEntryState from Json", body)
		}
	}
	return gConf, err
}

type OspfGlobalConfig struct {
	BaseObj
	RouterIdKey              string `SNAPROUTE: "KEY"`
	AdminStat                int32
	VersionNumber            int32
	AreaBdrRtrStatus         bool
	ASBdrRtrStatus           bool
	ExternLsaCount           uint32
	ExternLsaCksumSum        int32
	TOSSupport               bool
	OriginateNewLsas         uint32
	RxNewLsas                uint32
	ExtLsdbLimit             int32
	MulticastExtensions      int32
	ExitOverflowInterval     int32
	DemandExtensions         bool
	RFC1583Compatibility     bool
	OpaqueLsaSupport         bool
	ReferenceBandwidth       uint32
	RestartSupport           int32
	RestartInterval          int32
	RestartStrictLsaChecking bool
	RestartStatus            int32
	RestartAge               uint32
	RestartExitReason        int32
	AsLsaCount               uint32
	AsLsaCksumSum            uint32
	StubRouterSupport        bool
	StubRouterAdvertisement  int32
	DiscontinuityTime        uint32
}

func (obj OspfGlobalConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var gConf OspfGlobalConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &gConf); err != nil {
			fmt.Println("### Trouble in unmarshalling OspfGeneralGroup from Json", body)
		}
	}
	return gConf, err
}

type Login struct {
	BaseObj
	UserName string
	Password string
}

func (obj Login) UnmarshalObject(body []byte) (ConfigObj, error) {
	var loginObj Login
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &loginObj); err != nil {
			fmt.Println("### Trouble in unmarshaling Login from Json", body)
		}
	}

	return loginObj, err
}

type Logout struct {
	BaseObj
	UserName  string
	SessionId uint32
}

func (obj Logout) UnmarshalObject(body []byte) (ConfigObj, error) {
	var logoutObj Logout
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &logoutObj); err != nil {
			fmt.Println("### Trouble in unmarshaling Logout from Json", body)
		}
	}

	return logoutObj, err
}
