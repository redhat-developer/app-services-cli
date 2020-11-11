/*
 * Managed Service API
 *
 * Managed Service API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package msclient

import (
	"time"
)

// KafkaRequestAllOf struct for KafkaRequestAllOf
type KafkaRequestAllOf struct {
	Status              string    `json:"status,omitempty"`
	CloudProvider       string    `json:"cloud_provider,omitempty"`
	MultiAz             bool      `json:"multi_az,omitempty"`
	Region              string    `json:"region,omitempty"`
	Owner               string    `json:"owner,omitempty"`
	Name                string    `json:"name,omitempty"`
	BootstrapServerHost string    `json:"bootstrapServerHost,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
}
