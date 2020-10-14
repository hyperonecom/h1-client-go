/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package h1
// DnsRecordset struct for DnsRecordset
type DnsRecordset struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name"`
	Type string `json:"type"`
	Ttl float32 `json:"ttl,omitempty"`
	Record []DnsRecord `json:"record,omitempty"`
}
