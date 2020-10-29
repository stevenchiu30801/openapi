/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nef_models

type AfResultStatus string

// List of AfResultStatus
const (
	AfResultStatus_SUCCESS              AfResultStatus = "SUCCESS"
	AfResultStatus_TEMPORARY_CONGESTION AfResultStatus = "TEMPORARY_CONGESTION"
	AfResultStatus_RELOC_NO_ALLOWED     AfResultStatus = "RELOC_NO_ALLOWED"
	AfResultStatus_OTHER                AfResultStatus = "OTHER"
)
