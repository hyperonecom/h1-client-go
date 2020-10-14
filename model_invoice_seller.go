/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package h1
// InvoiceSeller struct for InvoiceSeller
type InvoiceSeller struct {
	Company string `json:"company"`
	Address InvoiceSellerAddress `json:"address,omitempty"`
	Nip string `json:"nip"`
}
