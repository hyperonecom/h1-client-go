/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject21 struct {
	Service string `json:"service,omitempty"`
	Network string `json:"network,omitempty"`
	Ip string `json:"ip,omitempty"`
	Tag map[string]interface{} `json:"tag,omitempty"`
	VlanIds string `json:"vlanIds,omitempty"`
}
