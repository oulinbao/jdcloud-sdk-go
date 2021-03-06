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
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeProbeTasksRequest struct {

    core.JDCloudRequest

    /* 探测任务的task_id  */
    ProbeTaskID string `json:"probeTaskID"`

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 根据探测任务的名称查询，支持模糊查询 (Optional) */
    Name *string `json:"name"`

    /* 根据探测任务的类型查询，1、http 2、telnet (Optional) */
    Type *int `json:"type"`
}

/*
 * param probeTaskID: 探测任务的task_id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeProbeTasksRequest(
    probeTaskID string,
) *DescribeProbeTasksRequest {

	return &DescribeProbeTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/am/probeTask",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        ProbeTaskID: probeTaskID,
	}
}

/*
 * param probeTaskID: 探测任务的task_id (Required)
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param name: 根据探测任务的名称查询，支持模糊查询 (Optional)
 * param type_: 根据探测任务的类型查询，1、http 2、telnet (Optional)
 */
func NewDescribeProbeTasksRequestWithAllParams(
    probeTaskID string,
    pageNumber *int,
    pageSize *int,
    name *string,
    type_ *int,
) *DescribeProbeTasksRequest {

    return &DescribeProbeTasksRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/am/probeTask",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        ProbeTaskID: probeTaskID,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Name: name,
        Type: type_,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeProbeTasksRequestWithoutParam() *DescribeProbeTasksRequest {

    return &DescribeProbeTasksRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/am/probeTask",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param probeTaskID: 探测任务的task_id(Required) */
func (r *DescribeProbeTasksRequest) SetProbeTaskID(probeTaskID string) {
    r.ProbeTaskID = probeTaskID
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeProbeTasksRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeProbeTasksRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param name: 根据探测任务的名称查询，支持模糊查询(Optional) */
func (r *DescribeProbeTasksRequest) SetName(name string) {
    r.Name = &name
}

/* param type_: 根据探测任务的类型查询，1、http 2、telnet(Optional) */
func (r *DescribeProbeTasksRequest) SetType(type_ int) {
    r.Type = &type_
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeProbeTasksRequest) GetRegionId() string {
    return ""
}

type DescribeProbeTasksResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeProbeTasksResult `json:"result"`
}

type DescribeProbeTasksResult struct {
    NumberPages int64 `json:"numberPages"`
    NumberRecords int64 `json:"numberRecords"`
    PageNumber int64 `json:"pageNumber"`
    PageSize int64 `json:"pageSize"`
    TaskInfo []monitor.TaskInfo `json:"taskInfo"`
    TotalCount int64 `json:"totalCount"`
}