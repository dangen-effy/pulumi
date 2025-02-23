// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apitype

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ResourceSearchResponse struct {
	Total        *int64                    `json:"total,omitempty"`
	Resources    []ResourceResult          `json:"resources,omitempty"`
	Aggregations map[string]Aggregation    `json:"aggregations,omitempty"`
	Pagination   *ResourceSearchPagination `json:"pagination,omitempty"`
}

// ResourceResult is the user-facing type for our indexed resources.
type ResourceResult struct {
	Created      *string          `json:"created,omitempty"`
	Custom       *bool            `json:"custom,omitempty"`
	Delete       *bool            `json:"delete,omitempty"`
	Dependencies []string         `json:"dependencies,omitempty"`
	ID           *string          `json:"id,omitempty"`
	Modified     *string          `json:"modified,omitempty"`
	Module       *string          `json:"module"`
	Name         *string          `json:"name,omitempty"`
	Package      *string          `json:"package"`
	ParentURN    *string          `json:"parent.urn"`
	Pending      *string          `json:"pending,omitempty"`
	Program      *string          `json:"project,omitempty"`
	Protected    *bool            `json:"protected,omitempty"`
	ProviderURN  *string          `json:"provider.urn"`
	Stack        *string          `json:"stack,omitempty"`
	Type         *string          `json:"type,omitempty"`
	URN          *string          `json:"urn,omitempty"`
	Teams        []string         `json:"teams,omitempty"`
	Properties   *json.RawMessage `json:"properties,omitempty"`
}

// Aggregation collects the top 5 aggregated values for the requested dimension.
type Aggregation struct {
	Others  *int64              `json:"others,omitempty"`
	Results []AggregationBucket `json:"results,omitempty"`
}

// AggregationBucket is an example value for the requested aggregation, with a
// count of how many resources share that value.
type AggregationBucket struct {
	Name  *string `json:"name,omitempty"`
	Count *int64  `json:"count,omitempty"`
}

// ResourceSearchPagination provides links for paging through results.
type ResourceSearchPagination struct {
	Previous *string `json:"previous,omitempty"`
	Next     *string `json:"next,omitempty"`
	Cursor   *string `json:"cursor,omitempty"`
}

type PulumiQueryResponse struct {
	Query string `json:"query"`
}

type PulumiQueryRequest struct {
	Query string `url:"query"`
}

// ParseQueryParams takes a list of parameters passed into the CLI
// Search commands (either in the form of `key=value` or a bare Pulumi query)
// and returns a PulumiQueryRequest struct that can be used to make a request
// to the Pulumi API.
//
// See https://www.pulumi.com/docs/pulumi-cloud/insights/search/#query-syntax for reference
func ParseQueryParams(rawParams []string) *PulumiQueryRequest {
	var queryString strings.Builder
	for _, param := range rawParams {
		if key, value, ok := strings.Cut(param, "="); ok {
			fmt.Fprintf(&queryString, "%s:%s", key, value)
		} else {
			queryString.WriteString(param)
		}
	}
	return &PulumiQueryRequest{Query: queryString.String()}
}
