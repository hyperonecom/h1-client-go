/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package h1
// NetworkingProjectNetgwCreate struct for NetworkingProjectNetgwCreate
type NetworkingProjectNetgwCreate struct {
	Name string `json:"name,omitempty"`
	Public NetgwPublic `json:"public,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
}