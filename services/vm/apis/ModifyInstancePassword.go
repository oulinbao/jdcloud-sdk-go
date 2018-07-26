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

type ModifyInstancePasswordRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 云主机ID  */
    InstanceId string `json:"instanceId"`

    /* 密码，<a href="https://www.jdcloud.com/help/detail/3870/isCatalog/1">参考公共参数规范</a>。  */
    Password string `json:"password"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 * param password: 密码，<a href="https://www.jdcloud.com/help/detail/3870/isCatalog/1">参考公共参数规范</a>。 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyInstancePasswordRequest(
    regionId string,
    instanceId string,
    password string,
) *ModifyInstancePasswordRequest {

	return &ModifyInstancePasswordRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstancePassword",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        Password: password,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: 云主机ID (Required)
 * param password: 密码，<a href="https://www.jdcloud.com/help/detail/3870/isCatalog/1">参考公共参数规范</a>。 (Required)
 */
func NewModifyInstancePasswordRequestWithAllParams(
    regionId string,
    instanceId string,
    password string,
) *ModifyInstancePasswordRequest {

    return &ModifyInstancePasswordRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstancePassword",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Password: password,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyInstancePasswordRequestWithoutParam() *ModifyInstancePasswordRequest {

    return &ModifyInstancePasswordRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:modifyInstancePassword",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ModifyInstancePasswordRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: 云主机ID(Required) */
func (r *ModifyInstancePasswordRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param password: 密码，<a href="https://www.jdcloud.com/help/detail/3870/isCatalog/1">参考公共参数规范</a>。(Required) */
func (r *ModifyInstancePasswordRequest) SetPassword(password string) {
    r.Password = password
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyInstancePasswordRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyInstancePasswordResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyInstancePasswordResult `json:"result"`
}

type ModifyInstancePasswordResult struct {
}