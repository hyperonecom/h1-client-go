/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package h1
// IamProjectSaCreate struct for IamProjectSaCreate
type IamProjectSaCreate struct {
	Name string `json:"name"`
	Service string `json:"service,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
}