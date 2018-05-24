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
    rds "github.com/jdcloud-api/jdcloud-sdk-go/services/rds/models"
)

type DescribeBackupsRequest struct {

    core.JDCloudRequest

    /* 地域代码  */
    RegionId string `json:"regionId"`

    /* 实例ID  */
    InstanceId string `json:"instanceId"`

    /* 查询备份类型，0为手动备份，1为自动备份，不传表示全部. - 测试参数，后续可能被其他参数取代 (Optional) */
    Auto *int `json:"auto"`

    /* 返回backupType等于指定值的备份列表。full为全量备份，diff为增量备份- 测试参数，后续可能被其他参数取代 (Optional) */
    BackupTypeFilter *string `json:"backupTypeFilter"`

    /* 返回dbName等于指定值的备份列表，不传或为空返回全部- 测试参数，后续可能被其他参数取代 (Optional) */
    DbNameFilter *string `json:"dbNameFilter"`

    /* 返回备份开始时间大于该时间的备份列表- 测试参数，后续可能被其他参数取代 (Optional) */
    BackupTimeRangeStartFilter *string `json:"backupTimeRangeStartFilter"`

    /* 返回备份开始时间小于等于该时间的备份列表- 测试参数，后续可能被其他参数取代 (Optional) */
    BackupTimeRangeEndFilter *string `json:"backupTimeRangeEndFilter"`

    /* 显示数据的页码，取值范围：[1,1000)，页码超过总页数时，显示最后一页，用于查询列表的接口  */
    PageNumber int `json:"pageNumber"`

    /* 每页显示的数据条数，取值范围：10/20/30/50/100  */
    PageSize int `json:"pageSize"`
}

/*
 * param regionId: 地域代码 
 * param instanceId: 实例ID 
 * param auto: 查询备份类型，0为手动备份，1为自动备份，不传表示全部. - 测试参数，后续可能被其他参数取代 (Optional)
 * param backupTypeFilter: 返回backupType等于指定值的备份列表。full为全量备份，diff为增量备份- 测试参数，后续可能被其他参数取代 (Optional)
 * param dbNameFilter: 返回dbName等于指定值的备份列表，不传或为空返回全部- 测试参数，后续可能被其他参数取代 (Optional)
 * param backupTimeRangeStartFilter: 返回备份开始时间大于该时间的备份列表- 测试参数，后续可能被其他参数取代 (Optional)
 * param backupTimeRangeEndFilter: 返回备份开始时间小于等于该时间的备份列表- 测试参数，后续可能被其他参数取代 (Optional)
 * param pageNumber: 显示数据的页码，取值范围：[1,1000)，页码超过总页数时，显示最后一页，用于查询列表的接口 
 * param pageSize: 每页显示的数据条数，取值范围：10/20/30/50/100 
 */
func NewDescribeBackupsRequest(
    regionId string,
    instanceId string,
    pageNumber int,
    pageSize int,
) *DescribeBackupsRequest {

	return &DescribeBackupsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/backups",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        PageNumber: pageNumber,
        PageSize: pageSize,
	}
}

func (r *DescribeBackupsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DescribeBackupsRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

func (r *DescribeBackupsRequest) SetAuto(auto int) {
    r.Auto = &auto
}

func (r *DescribeBackupsRequest) SetBackupTypeFilter(backupTypeFilter string) {
    r.BackupTypeFilter = &backupTypeFilter
}

func (r *DescribeBackupsRequest) SetDbNameFilter(dbNameFilter string) {
    r.DbNameFilter = &dbNameFilter
}

func (r *DescribeBackupsRequest) SetBackupTimeRangeStartFilter(backupTimeRangeStartFilter string) {
    r.BackupTimeRangeStartFilter = &backupTimeRangeStartFilter
}

func (r *DescribeBackupsRequest) SetBackupTimeRangeEndFilter(backupTimeRangeEndFilter string) {
    r.BackupTimeRangeEndFilter = &backupTimeRangeEndFilter
}

func (r *DescribeBackupsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = pageNumber
}

func (r *DescribeBackupsRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBackupsRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DescribeBackupsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeBackupsResult `json:"result"`
}

type DescribeBackupsResult struct {
    Backup []rds.Backup `json:"backup"`
    TotalCount int `json:"totalCount"`
}