/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package h1
import (
	"time"
)
// Ip struct for Ip
type Ip struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Flavour string `json:"flavour,omitempty"`
	ModifiedOn time.Time `json:"modifiedOn,omitempty"`
	ModifiedBy string `json:"modifiedBy,omitempty"`
	CreatedOn time.Time `json:"createdOn,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	State string `json:"state,omitempty"`
	Project string `json:"project,omitempty"`
	Uri string `json:"uri,omitempty"`
	Address string `json:"address,omitempty"`
	Fqdn string `json:"fqdn,omitempty"`
	Network string `json:"network,omitempty"`
	PtrRecord string `json:"ptrRecord,omitempty"`
	Persistent bool `json:"persistent,omitempty"`
	Associated IpAssociated `json:"associated,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
}
