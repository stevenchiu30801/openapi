/*
 * 3gpp-pfd-management
 *
 * API for PFD management. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nef_models

type PfdManagement struct {

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self string `json:"self,omitempty" bson:"self"`

	SupportedFeatures string `json:"supportedFeatures,omitempty" bson:"supportedFeatures"`

	// Each element uniquely identifies the PFDs for an external application identifier. Each element is identified in the map via an external application identifier as key. The response shall include successfully provisioned PFD data of application(s).
	PfdDatas map[string]PfdData `json:"pfdDatas" bson:"pfdDatas"`

	// Supplied by the SCEF and contains the external application identifiers for which PFD(s) are not added or modified successfully. The failure reason is also included. Each element provides the related information for one or more external application identifier(s) and is identified in the map via the failure identifier as key.
	PfdReports map[string]PfdReport `json:"pfdReports,omitempty" bson:"pfdReports"`

	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination string `json:"notificationDestination,omitempty" bson:"notificationDestination"`

	// Set to true by the SCS/AS to request the SCEF to send a test notification as defined in subclause 5.2.5.3. Set to false or omitted otherwise.
	RequestTestNotification bool `json:"requestTestNotification,omitempty" bson:"requestTestNotification"`

	WebsockNotifConfig *WebsockNotifConfig `json:"websockNotifConfig,omitempty" bson:"websockNotifConfig"`
}
