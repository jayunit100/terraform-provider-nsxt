/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package licensing

import (
	"github.com/vmware/go-vmware-nsxt/common"
)

type CapacityUsage struct {

	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`

	Schema string `json:"_schema,omitempty"`

	Self *common.SelfResourceLink `json:"_self,omitempty"`

	// type of the capacity field
	CapacityType string `json:"capacity_type,omitempty"`

	// count of number of items of capacity_type
	UsageCount int64 `json:"usage_count,omitempty"`
}