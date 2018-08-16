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

type ModifyAuditRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Instance ID  */
    InstanceId string `json:"instanceId"`

    /* 增加审计项，多个审计项之间用英文逗号，分号或空格分隔，例如DATABASE_OBJECT_ACCESS_GROUP,ACKUP_RESTORE_GROUP (Optional) */
    Add *string `json:""`

    /* 删除审计项，多个审计项之间用英文逗号，分号或空格分隔，例如DATABASE_OBJECT_ACCESS_GROUP,ACKUP_RESTORE_GROUP如删除了所有审计项，则审计自动关闭 (Optional) */
    Drop *string `json:""`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyAuditRequest(
    regionId string,
    instanceId string,
) *ModifyAuditRequest {

	return &ModifyAuditRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/audit:modifyAudit",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 * param add: 增加审计项，多个审计项之间用英文逗号，分号或空格分隔，例如DATABASE_OBJECT_ACCESS_GROUP,ACKUP_RESTORE_GROUP (Optional)
 * param drop: 删除审计项，多个审计项之间用英文逗号，分号或空格分隔，例如DATABASE_OBJECT_ACCESS_GROUP,ACKUP_RESTORE_GROUP如删除了所有审计项，则审计自动关闭 (Optional)
 */
func NewModifyAuditRequestWithAllParams(
    regionId string,
    instanceId string,
    add *string,
    drop *string,
) *ModifyAuditRequest {

    return &ModifyAuditRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/audit:modifyAudit",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        Add: add,
        Drop: drop,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyAuditRequestWithoutParam() *ModifyAuditRequest {

    return &ModifyAuditRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/audit:modifyAudit",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *ModifyAuditRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: Instance ID(Required) */
func (r *ModifyAuditRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param add: 增加审计项，多个审计项之间用英文逗号，分号或空格分隔，例如DATABASE_OBJECT_ACCESS_GROUP,ACKUP_RESTORE_GROUP(Optional) */
func (r *ModifyAuditRequest) SetAdd(add string) {
    r.Add = &add
}

/* param drop: 删除审计项，多个审计项之间用英文逗号，分号或空格分隔，例如DATABASE_OBJECT_ACCESS_GROUP,ACKUP_RESTORE_GROUP如删除了所有审计项，则审计自动关闭(Optional) */
func (r *ModifyAuditRequest) SetDrop(drop string) {
    r.Drop = &drop
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyAuditRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyAuditResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyAuditResult `json:"result"`
}

type ModifyAuditResult struct {
}