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
    "reflect"
    mps "github.com/jdcloud-api/jdcloud-sdk-go/services/mps/models"
)

type CreateThumbnailTaskRequest struct {

    core.JDCloudRequest

    /* region id  */
    RegionId string `json:"regionId"`

    /* 任务ID (readonly) (Optional) */
    TaskID *string `json:"taskID"`

    /* 状态 (SUCCESS, ERROR, PENDDING, RUNNING) (readonly) (Optional) */
    Status *string `json:"status"`

    /* 错误码 (readonly) (Optional) */
    ErrorCode *int `json:"errorCode"`

    /* 任务创建时间 时间格式(GMT): yyyy-MM-dd’T’HH:mm:ss.SSS’Z’  (readonly) (Optional) */
    CreatedTime *string `json:"createdTime"`

    /* 任务创建时间 时间格式(GMT): yyyy-MM-dd’T’HH:mm:ss.SSS’Z’  (readonly) (Optional) */
    LastUpdatedTime *string `json:"lastUpdatedTime"`

    /*   */
    Source *mps.ThumbnailTaskSource `json:"source"`

    /*   */
    Target *mps.ThumbnailTaskTarget `json:"target"`

    /*  (Optional) */
    Rule *mps.ThumbnailTaskRule `json:"rule"`
}

/*
 * param regionId: region id 
 * param taskID: 任务ID (readonly) (Optional)
 * param status: 状态 (SUCCESS, ERROR, PENDDING, RUNNING) (readonly) (Optional)
 * param errorCode: 错误码 (readonly) (Optional)
 * param createdTime: 任务创建时间 时间格式(GMT): yyyy-MM-dd’T’HH:mm:ss.SSS’Z’  (readonly) (Optional)
 * param lastUpdatedTime: 任务创建时间 时间格式(GMT): yyyy-MM-dd’T’HH:mm:ss.SSS’Z’  (readonly) (Optional)
 * param source:  
 * param target:  
 * param rule:  (Optional)
 */
func NewCreateThumbnailTaskRequest(
    regionId string,
    source *mps.ThumbnailTaskSource,
    target *mps.ThumbnailTaskTarget,
) *CreateThumbnailTaskRequest {

	return &CreateThumbnailTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/thumbnail",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Source: source,
        Target: target,
	}
}

func (r *CreateThumbnailTaskRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *CreateThumbnailTaskRequest) SetTaskID(taskID string) {
    r.TaskID = &taskID
}

func (r *CreateThumbnailTaskRequest) SetStatus(status string) {
    r.Status = &status
}

func (r *CreateThumbnailTaskRequest) SetErrorCode(errorCode int) {
    r.ErrorCode = &errorCode
}

func (r *CreateThumbnailTaskRequest) SetCreatedTime(createdTime string) {
    r.CreatedTime = &createdTime
}

func (r *CreateThumbnailTaskRequest) SetLastUpdatedTime(lastUpdatedTime string) {
    r.LastUpdatedTime = &lastUpdatedTime
}

func (r *CreateThumbnailTaskRequest) SetSource(source *mps.ThumbnailTaskSource) {
    r.Source = source
}

func (r *CreateThumbnailTaskRequest) SetTarget(target *mps.ThumbnailTaskTarget) {
    r.Target = target
}

func (r *CreateThumbnailTaskRequest) SetRule(rule *mps.ThumbnailTaskRule) {
    r.Rule = rule
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateThumbnailTaskRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type CreateThumbnailTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateThumbnailTaskResult `json:"result"`
}

type CreateThumbnailTaskResult struct {
    TaskID string `json:"taskID"`
}