/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ContainerProcess struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	User string `json:"user,omitempty"`
	Parent string `json:"parent,omitempty"`
}