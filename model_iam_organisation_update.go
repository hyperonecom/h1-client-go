/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package h1
// IamOrganisationUpdate struct for IamOrganisationUpdate
type IamOrganisationUpdate struct {
	Name string `json:"name,omitempty"`
	Billing OrganisationBilling1 `json:"billing,omitempty"`
}