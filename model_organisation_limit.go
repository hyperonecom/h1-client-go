/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type OrganisationLimit struct {
	Organisation OrganisationLimitOrganisation `json:"organisation,omitempty"`
	Project ProjectLimit `json:"project,omitempty"`
}
