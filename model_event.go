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
// Event struct for Event
type Event struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	CreatedOn time.Time `json:"createdOn,omitempty"`
	State string `json:"state,omitempty"`
	Stage string `json:"stage,omitempty"`
	Resource string `json:"resource,omitempty"`
}
