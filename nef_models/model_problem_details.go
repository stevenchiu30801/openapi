/*
 * 3gpp-traffic-influence
 *
 * API for AF traffic influence © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nef_models

type ProblemDetails struct {

	// string providing an URI formatted according to IETF RFC 3986.
	Type string `json:"type,omitempty" bson:"type"`

	// A short, human-readable summary of the problem type. It should not change from occurrence to occurrence of the problem.
	Title string `json:"title,omitempty" bson:"title"`

	// The HTTP status code for this occurrence of the problem.
	Status int32 `json:"status,omitempty" bson:"status"`

	// A human-readable explanation specific to this occurrence of the problem.
	Detail string `json:"detail,omitempty" bson:"detail"`

	// string providing an URI formatted according to IETF RFC 3986.
	Instance string `json:"instance,omitempty" bson:"instance"`

	// A machine-readable application error cause specific to this occurrence of the problem. This IE should be present and provide application-related error information, if available.
	Cause string `json:"cause,omitempty" bson:"cause"`

	// Description of invalid parameters, for a request rejected due to invalid parameters.
	InvalidParams []InvalidParam `json:"invalidParams,omitempty" bson:"invalidParams"`
}
