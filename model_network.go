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

type Network struct {
	Id           string                 `json:"_id,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Services     []ContainerServices    `json:"services,omitempty"`
	Flavour      string                 `json:"flavour,omitempty"`
	ModifiedOn   time.Time              `json:"modifiedOn,omitempty"`
	ModifiedBy   string                 `json:"modifiedBy,omitempty"`
	CreatedBy    string                 `json:"createdBy,omitempty"`
	CreatedOn    time.Time              `json:"createdOn,omitempty"`
	AccessRights []string               `json:"accessRights,omitempty"`
	Processing   string                 `json:"processing,omitempty"`
	Created      string                 `json:"created,omitempty"`
	Queue        []Event                `json:"queue,omitempty"`
	State        string                 `json:"state,omitempty"`
	Tag          map[string]interface{} `json:"tag,omitempty"`
	Project      string                 `json:"project,omitempty"`
	Address      string                 `json:"address,omitempty"`
	Broadcast    string                 `json:"broadcast,omitempty"`
	Gateway      string                 `json:"gateway,omitempty"`
	Type         string                 `json:"type,omitempty"`
	Dns          NetworkDns             `json:"dns,omitempty"`
	Netgw        string                 `json:"netgw,omitempty"`
}
