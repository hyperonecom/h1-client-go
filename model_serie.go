/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package h1
import (
	"time"
)
// Serie struct for Serie
type Serie struct {
	Time time.Time `json:"time"`
	Value float32 `json:"value"`
}
