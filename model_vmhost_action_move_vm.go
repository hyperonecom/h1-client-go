/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VmhostActionMoveVm struct {
	Vm string `json:"vm"`
	Vmhost string `json:"vmhost,omitempty"`
}