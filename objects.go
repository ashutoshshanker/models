package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type ConfigObj interface {
	UnmarshalObject(data []byte) (ConfigObj, error)
	CreateDBTable(dbHdl *sql.DB) error
	StoreObjectInDb(dbHdl *sql.DB) (int64, error)
	DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error
	GetKey() (string, error)
	GetSqlKeyStr (string) (string, error)
	GetObjectFromDb(objKey string, dbHdl *sql.DB) (ConfigObj, error)
	CompareObjectsAndDiff(dbObj ConfigObj) ([]byte, error)
	UpdateObjectInDb(dbV4Route ConfigObj, attrSet []byte, dbHdl *sql.DB) error
}

type BaseObj struct{}

func (b BaseObj) CreateDBTable(dbHdl *sql.DB) error {
	return nil
}

func (b BaseObj) StoreObjectInDb(dbHdl *sql.DB) (objId int64, err error) {
	return objId, err
}

func (b BaseObj) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	return nil
}

func (b BaseObj) GetKey() (string, error) {
	s := ""
	return s, nil
}
func (b BaseObj) GetSqlKeyStr(objKey string) (string, error) {
	s := ""
	return s, nil
}

func (b BaseObj) CompareObjectsAndDiff(dbObj ConfigObj) ([]byte, error) {
	var arr []byte
	return arr, nil
}

func (b BaseObj) UpdateObjectInDb(dbV4Route ConfigObj, attrSet []byte, dbHdl *sql.DB) error {
	return nil
}

//
// This file is handcoded for now. Eventually this would be generated by yang compiler
//
type IPV4Route struct {
	BaseObj
	DestinationNw     string
	NetworkMask       string
	Cost              int
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
	PeerAS          uint32
	LocalAS         uint32
	AuthPassword    string
	Description     string
	NeighborAddress string `SNAPROUTE: "KEY"`
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
	PeerAS          uint32
	LocalAS         uint32
	PeerType        PeerType
	AuthPassword    string
	Description     string
	NeighborAddress string
	SessionState    uint32
	Messages        BGPMessages
	Queues          BGPQueues
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

/* Start - Asicd objects */
type Vlan struct {
	BaseObj
	VlanId      int32
	Ports       string
	PortTagType string
}

/* FIXME : RouterIf needs to be replaced by generic
 * layer 2 object name e.x Port-21 or Vlan-5 etc.
 * Internally this l2 object name can be translated
 * into appropriate key.
 */
type IPv4Intf struct {
	BaseObj
	IpAddr   string
	RouterIf int32
	IfType   int32
}

type IPv4Neighbor struct {
	BaseObj
	IpAddr   string
	MacAddr  string
	VlanId   int32
	RouterIf int32
}

func (obj Vlan) UnmarshalObject(body []byte) (ConfigObj, error) {
	var vlanObj Vlan
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &vlanObj); err != nil {
			fmt.Println("### Trouble in unmarshaling Vlan from Json", body)
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

func (obj IPv4Neighbor) UnmarshalObject(body []byte) (ConfigObj, error) {
	var v4Nbr IPv4Neighbor
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &v4Nbr); err != nil {
			fmt.Println("### Trouble in unmarshaling IPv4Neighbor from Json", body)
		}
	}
	return v4Nbr, err
}
