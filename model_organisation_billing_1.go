/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package h1
// OrganisationBilling1 struct for OrganisationBilling1
type OrganisationBilling1 struct {
	Email string `json:"email,omitempty"`
	Company string `json:"company,omitempty"`
	Address BillingAddress1 `json:"address,omitempty"`
}