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
// RegistryCredential struct for RegistryCredential
type RegistryCredential struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	CreatedBy string `json:"createdBy,omitempty"`
	CreatedOn time.Time `json:"createdOn,omitempty"`
	Type string `json:"type"`
	Value string `json:"value"`
	Fingerprint string `json:"fingerprint,omitempty"`
	Token string `json:"token,omitempty"`
}
