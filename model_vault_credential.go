/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VaultCredential struct {
	Password VaultCredentialPassword `json:"password,omitempty"`
	Certificate VaultCredentialPassword `json:"certificate,omitempty"`
}