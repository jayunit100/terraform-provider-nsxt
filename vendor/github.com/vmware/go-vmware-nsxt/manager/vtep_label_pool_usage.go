/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

// Pool usage statistics in a pool
type VtepLabelPoolUsage struct {

	// Total number of allocated IDs in a pool
	AllocatedIds int64 `json:"allocated_ids,omitempty"`

	// Total number of free IDs in a pool
	FreeIds int64 `json:"free_ids,omitempty"`

	// Total number of IDs in a pool
	TotalIds int64 `json:"total_ids,omitempty"`
}
