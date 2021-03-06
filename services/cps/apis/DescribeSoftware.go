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
    cps "github.com/jdcloud-api/jdcloud-sdk-go/services/cps/models"
)

type DescribeSoftwareRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 操作系统系统类型ID，调用接口（describeOS）获取云物理服务器支持的操作系统  */
    OsTypeId string `json:"osTypeId"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param osTypeId: 操作系统系统类型ID，调用接口（describeOS）获取云物理服务器支持的操作系统 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSoftwareRequest(
    regionId string,
    osTypeId string,
) *DescribeSoftwareRequest {

	return &DescribeSoftwareRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/os/{osTypeId}/softwares",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        OsTypeId: osTypeId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param osTypeId: 操作系统系统类型ID，调用接口（describeOS）获取云物理服务器支持的操作系统 (Required)
 */
func NewDescribeSoftwareRequestWithAllParams(
    regionId string,
    osTypeId string,
) *DescribeSoftwareRequest {

    return &DescribeSoftwareRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/os/{osTypeId}/softwares",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        OsTypeId: osTypeId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSoftwareRequestWithoutParam() *DescribeSoftwareRequest {

    return &DescribeSoftwareRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/os/{osTypeId}/softwares",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *DescribeSoftwareRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param osTypeId: 操作系统系统类型ID，调用接口（describeOS）获取云物理服务器支持的操作系统(Required) */
func (r *DescribeSoftwareRequest) SetOsTypeId(osTypeId string) {
    r.OsTypeId = osTypeId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSoftwareRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSoftwareResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSoftwareResult `json:"result"`
}

type DescribeSoftwareResult struct {
    Softwares []cps.Software `json:"softwares"`
}