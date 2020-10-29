/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nef_models

type AfAckInfo struct {
	AfTransId string `json:"afTransId,omitempty" bson:"afTransId"`

	AckResult *AfResultInfo `json:"ackResult" bson:"ackResult"`

	Gpsi string `json:"gpsi,omitempty" bson:"gpsi"`
}
