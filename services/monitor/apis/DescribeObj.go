// Copyright 2018 JDCLOUD.COM
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
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type DescribeObjRequest struct {

    core.JDCloudRequest

    /* region  */
    RegionId string `json:"regionId"`

    /* namespaceUID  */
    NamespaceUID string `json:"namespaceUID"`

    /* objUID  */
    ObjUID string `json:"objUID"`
}

/*
 * param regionId: region (Required)
 * param namespaceUID: namespaceUID (Required)
 * param objUID: objUID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeObjRequest(
    regionId string,
    namespaceUID string,
    objUID string,
) *DescribeObjRequest {

	return &DescribeObjRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cm/namespaces/{namespaceUID}/objs/{objUID}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        NamespaceUID: namespaceUID,
        ObjUID: objUID,
	}
}

/*
 * param regionId: region (Required)
 * param namespaceUID: namespaceUID (Required)
 * param objUID: objUID (Required)
 */
func NewDescribeObjRequestWithAllParams(
    regionId string,
    namespaceUID string,
    objUID string,
) *DescribeObjRequest {

    return &DescribeObjRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cm/namespaces/{namespaceUID}/objs/{objUID}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        NamespaceUID: namespaceUID,
        ObjUID: objUID,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeObjRequestWithoutParam() *DescribeObjRequest {

    return &DescribeObjRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cm/namespaces/{namespaceUID}/objs/{objUID}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: region(Required) */
func (r *DescribeObjRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param namespaceUID: namespaceUID(Required) */
func (r *DescribeObjRequest) SetNamespaceUID(namespaceUID string) {
    r.NamespaceUID = namespaceUID
}

/* param objUID: objUID(Required) */
func (r *DescribeObjRequest) SetObjUID(objUID string) {
    r.ObjUID = objUID
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeObjRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeObjResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeObjResult `json:"result"`
}

type DescribeObjResult struct {
    NamespaceName string `json:"namespaceName"`
    ObjName string `json:"objName"`
}