/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package errorresolver

// Metadata related to a given error_id
type ErrorResolverInfo struct {

	// The error id for which metadata information is needed
	ErrorId int64 `json:"error_id"`

	// Indicates whether there is a resolver associated with the error or not
	ResolverPresent bool `json:"resolver_present"`

	// User supplied metadata that might be required by the resolver
	UserMetadata *ErrorResolverUserMetadata `json:"user_metadata,omitempty"`
}
