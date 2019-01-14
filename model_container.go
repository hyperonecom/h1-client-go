/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

type Container struct {
	Id string `json:"_id,omitempty"`
	Name string `json:"name,omitempty"`
	Services []ProjectServices `json:"services,omitempty"`
	Flavour string `json:"flavour,omitempty"`
	ModifiedOn time.Time `json:"modifiedOn,omitempty"`
	ModifiedBy string `json:"modifiedBy,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	CreatedOn time.Time `json:"createdOn,omitempty"`
	AccessRights []string `json:"accessRights,omitempty"`
	Processing bool `json:"processing,omitempty"`
	Created bool `json:"created,omitempty"`
	Queue []Event `json:"queue,omitempty"`
	State string `json:"state,omitempty"`
	Tag map[string]interface{} `json:"tag,omitempty"`
	Project string `json:"project,omitempty"`
	Url string `json:"url,omitempty"`
}
