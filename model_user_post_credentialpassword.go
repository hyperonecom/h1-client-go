/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UserPostCredentialpassword struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Value string `json:"value"`
	Token string `json:"token,omitempty"`
}
