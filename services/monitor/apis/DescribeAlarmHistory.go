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

type DescribeAlarmHistoryRequest struct {

    core.JDCloudRequest

    /* 地域 Id  */
    RegionId string `json:"regionId"`

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 产品线 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 资源Id (Optional) */
    ResourceId *string `json:"resourceId"`

    /* resourceId列表 (Optional) */
    ResourceIdList []string `json:"resourceIdList"`

    /* 规则Id (Optional) */
    AlarmId *string `json:"alarmId"`

    /* 正在报警, 取值为1 (Optional) */
    Alarming *int `json:"alarming"`

    /* 产品线列表 (Optional) */
    ServiceCodeList []string `json:"serviceCodeList"`

    /* 开始时间 (Optional) */
    StartTime *string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime *string `json:"endTime"`

    /* 规则类型,默认查询1， 1表示资源监控，6表示站点监控,7表示可用性监控 (Optional) */
    RuleType *int `json:"ruleType"`

    /* 服务码或资源Id列表
filter name 为serviceCodes表示查询多个产品线的规则
filter name 为resourceIds表示查询多个资源的规则 (Optional) */
    Filters []monitor.Filter `json:"filters"`
}

/*
 * param regionId: 地域 Id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeAlarmHistoryRequest(
    regionId string,
) *DescribeAlarmHistoryRequest {

	return &DescribeAlarmHistoryRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/alarmHistory",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域 Id (Required)
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param serviceCode: 产品线 (Optional)
 * param resourceId: 资源Id (Optional)
 * param resourceIdList: resourceId列表 (Optional)
 * param alarmId: 规则Id (Optional)
 * param alarming: 正在报警, 取值为1 (Optional)
 * param serviceCodeList: 产品线列表 (Optional)
 * param startTime: 开始时间 (Optional)
 * param endTime: 结束时间 (Optional)
 * param ruleType: 规则类型,默认查询1， 1表示资源监控，6表示站点监控,7表示可用性监控 (Optional)
 * param filters: 服务码或资源Id列表
filter name 为serviceCodes表示查询多个产品线的规则
filter name 为resourceIds表示查询多个资源的规则 (Optional)
 */
func NewDescribeAlarmHistoryRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    serviceCode *string,
    resourceId *string,
    resourceIdList []string,
    alarmId *string,
    alarming *int,
    serviceCodeList []string,
    startTime *string,
    endTime *string,
    ruleType *int,
    filters []monitor.Filter,
) *DescribeAlarmHistoryRequest {

    return &DescribeAlarmHistoryRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarmHistory",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        ServiceCode: serviceCode,
        ResourceId: resourceId,
        ResourceIdList: resourceIdList,
        AlarmId: alarmId,
        Alarming: alarming,
        ServiceCodeList: serviceCodeList,
        StartTime: startTime,
        EndTime: endTime,
        RuleType: ruleType,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeAlarmHistoryRequestWithoutParam() *DescribeAlarmHistoryRequest {

    return &DescribeAlarmHistoryRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/alarmHistory",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 Id(Required) */
func (r *DescribeAlarmHistoryRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeAlarmHistoryRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeAlarmHistoryRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param serviceCode: 产品线(Optional) */
func (r *DescribeAlarmHistoryRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param resourceId: 资源Id(Optional) */
func (r *DescribeAlarmHistoryRequest) SetResourceId(resourceId string) {
    r.ResourceId = &resourceId
}

/* param resourceIdList: resourceId列表(Optional) */
func (r *DescribeAlarmHistoryRequest) SetResourceIdList(resourceIdList []string) {
    r.ResourceIdList = resourceIdList
}

/* param alarmId: 规则Id(Optional) */
func (r *DescribeAlarmHistoryRequest) SetAlarmId(alarmId string) {
    r.AlarmId = &alarmId
}

/* param alarming: 正在报警, 取值为1(Optional) */
func (r *DescribeAlarmHistoryRequest) SetAlarming(alarming int) {
    r.Alarming = &alarming
}

/* param serviceCodeList: 产品线列表(Optional) */
func (r *DescribeAlarmHistoryRequest) SetServiceCodeList(serviceCodeList []string) {
    r.ServiceCodeList = serviceCodeList
}

/* param startTime: 开始时间(Optional) */
func (r *DescribeAlarmHistoryRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 结束时间(Optional) */
func (r *DescribeAlarmHistoryRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param ruleType: 规则类型,默认查询1， 1表示资源监控，6表示站点监控,7表示可用性监控(Optional) */
func (r *DescribeAlarmHistoryRequest) SetRuleType(ruleType int) {
    r.RuleType = &ruleType
}

/* param filters: 服务码或资源Id列表
filter name 为serviceCodes表示查询多个产品线的规则
filter name 为resourceIds表示查询多个资源的规则(Optional) */
func (r *DescribeAlarmHistoryRequest) SetFilters(filters []monitor.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeAlarmHistoryRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeAlarmHistoryResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeAlarmHistoryResult `json:"result"`
}

type DescribeAlarmHistoryResult struct {
    AlarmHistoryList []monitor.DescribedAlarmHistory `json:"alarmHistoryList"`
    Total int64 `json:"total"`
}