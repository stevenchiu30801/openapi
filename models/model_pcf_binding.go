/*
 * Nbsf_Management
 *
 * Binding Support Management Service API. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type PcfBinding struct {
	Supi string `json:"supi,omitempty" bson:"supi"`

	Gpsi string `json:"gpsi,omitempty" bson:"gpsi"`

	Ipv4Addr string `json:"ipv4Addr,omitempty" bson:"ipv4Addr"`

	Ipv6Prefix string `json:"ipv6Prefix,omitempty" bson:"ipv6Prefix"`

	// The additional IPv6 Address Prefixes of the served UE.
	AddIpv6Prefixes []string `json:"addIpv6Prefixes,omitempty" bson:"addIpv6Prefixes"`

	IpDomain string `json:"ipDomain,omitempty" bson:"ipDomain"`

	MacAddr48 string `json:"macAddr48,omitempty" bson:"macAddr48"`

	// The additional MAC Addresses of the served UE.
	AddMacAddrs []string `json:"addMacAddrs,omitempty" bson:"addMacAddrs"`

	Dnn string `json:"dnn" bson:"dnn"`

	// Fully Qualified Domain Name
	PcfFqdn string `json:"pcfFqdn,omitempty" bson:"pcfFqdn"`

	// IP end points of the PCF hosting the Npcf_PolicyAuthorization service.
	PcfIpEndPoints []IpEndPoint `json:"pcfIpEndPoints,omitempty" bson:"pcfIpEndPoints"`

	PcfDiamHost string `json:"pcfDiamHost,omitempty" bson:"pcfDiamHost"`

	PcfDiamRealm string `json:"pcfDiamRealm,omitempty" bson:"pcfDiamRealm"`

	// Fully Qualified Domain Name
	PcfSmFqdn string `json:"pcfSmFqdn,omitempty" bson:"pcfSmFqdn"`

	// IP end points of the PCF hosting the Npcf_SMPolicyControl service.
	PcfSmIpEndPoints []IpEndPoint `json:"pcfSmIpEndPoints,omitempty" bson:"pcfSmIpEndPoints"`

	Snssai *Snssai `json:"snssai" bson:"snssai"`

	SuppFeat string `json:"suppFeat,omitempty" bson:"suppFeat"`

	PcfId string `json:"pcfId,omitempty" bson:"pcfId"`

	PcfSetId string `json:"pcfSetId,omitempty" bson:"pcfSetId"`

	RecoveryTime *time.Time `json:"recoveryTime,omitempty" bson:"recoveryTime"`

	ParaCom *ParameterCombination `json:"paraCom,omitempty" bson:"paraCom"`

	BindLevel BindingLevel `json:"bindLevel,omitempty" bson:"bindLevel"`

	Ipv4FrameRouteList []string `json:"ipv4FrameRouteList,omitempty" bson:"ipv4FrameRouteList"`

	Ipv6FrameRouteList []string `json:"ipv6FrameRouteList,omitempty" bson:"ipv6FrameRouteList"`
}
