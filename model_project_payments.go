/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ProjectPayments struct {
	Id string `json:"_id,omitempty"`
	Name string `json:"name,omitempty"`
	Services string `json:"services,omitempty"`
	Flavour string `json:"flavour,omitempty"`
	ModifiedOn string `json:"modifiedOn,omitempty"`
	ModifiedBy string `json:"modifiedBy,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	CreatedOn string `json:"createdOn,omitempty"`
	AccessRights string `json:"accessRights,omitempty"`
	Processing string `json:"processing,omitempty"`
	Created string `json:"created,omitempty"`
	Queue []Event `json:"queue,omitempty"`
	State string `json:"state,omitempty"`
	Tag string `json:"tag,omitempty"`
	Project string `json:"project,omitempty"`
	CreditsFree float32 `json:"creditsFree,omitempty"`
	Credits float32 `json:"credits,omitempty"`
	Channel string `json:"channel,omitempty"`
	Amount string `json:"amount,omitempty"`
	Type string `json:"type,omitempty"`
}
