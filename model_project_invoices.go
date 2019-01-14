/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ProjectInvoices struct {
	Id string `json:"_id,omitempty"`
	InvoiceNo string `json:"invoiceNo,omitempty"`
	IssueDate string `json:"issueDate,omitempty"`
	Summary string `json:"summary,omitempty"`
	Project string `json:"project,omitempty"`
	Organisation string `json:"organisation,omitempty"`
	Duplicate ProjectInvoicesDuplicate `json:"duplicate,omitempty"`
}
