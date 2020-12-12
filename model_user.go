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
// User struct for User
type User struct {
	Id string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	FamilyName string `json:"familyName,omitempty"`
	Name string `json:"name,omitempty"`
	CreatedOn time.Time `json:"createdOn,omitempty"`
	ModifiedOn time.Time `json:"modifiedOn,omitempty"`
	Lang string `json:"lang,omitempty"`
	Phone string `json:"phone,omitempty"`
	Limit UserLimit `json:"limit,omitempty"`
	Uri string `json:"uri,omitempty"`
}
