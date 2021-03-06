/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

// Object used internally as an interface to handle multple model versions.
type EdgelbPoolContainer struct {
	ApiVersion EdgelbApiVersion `json:"apiVersion,omitempty"`
	Name       string           `json:"name,omitempty"`
	Namespace  string           `json:"namespace,omitempty"`
	V2         EdgelbV2Pool     `json:"v2,omitempty"`
}
