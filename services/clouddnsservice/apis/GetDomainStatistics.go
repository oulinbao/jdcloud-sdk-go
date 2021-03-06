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

type GetDomainStatisticsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 域名ID  */
    DomainId string `json:"domainId"`

    /* 查询动作，"query"查询流量，"resolve"解析流量  */
    Action string `json:"action"`

    /* 域名  */
    DomainName string `json:"domainName"`

    /* 起始时间, UTC时间例如2017-11-10T23:00:00Z  */
    Start string `json:"start"`

    /* 终止时间, UTC时间例如2017-11-10T23:00:00Z  */
    End string `json:"end"`
}

/*
 * param regionId: Region ID (Required)
 * param domainId: 域名ID (Required)
 * param action: 查询动作，"query"查询流量，"resolve"解析流量 (Required)
 * param domainName: 域名 (Required)
 * param start: 起始时间, UTC时间例如2017-11-10T23:00:00Z (Required)
 * param end: 终止时间, UTC时间例如2017-11-10T23:00:00Z (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetDomainStatisticsRequest(
    regionId string,
    domainId string,
    action string,
    domainName string,
    start string,
    end string,
) *GetDomainStatisticsRequest {

	return &GetDomainStatisticsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain/{domainId}/stat",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        DomainId: domainId,
        Action: action,
        DomainName: domainName,
        Start: start,
        End: end,
	}
}

/*
 * param regionId: Region ID (Required)
 * param domainId: 域名ID (Required)
 * param action: 查询动作，"query"查询流量，"resolve"解析流量 (Required)
 * param domainName: 域名 (Required)
 * param start: 起始时间, UTC时间例如2017-11-10T23:00:00Z (Required)
 * param end: 终止时间, UTC时间例如2017-11-10T23:00:00Z (Required)
 */
func NewGetDomainStatisticsRequestWithAllParams(
    regionId string,
    domainId string,
    action string,
    domainName string,
    start string,
    end string,
) *GetDomainStatisticsRequest {

    return &GetDomainStatisticsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/stat",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        DomainId: domainId,
        Action: action,
        DomainName: domainName,
        Start: start,
        End: end,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetDomainStatisticsRequestWithoutParam() *GetDomainStatisticsRequest {

    return &GetDomainStatisticsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain/{domainId}/stat",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetDomainStatisticsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param domainId: 域名ID(Required) */
func (r *GetDomainStatisticsRequest) SetDomainId(domainId string) {
    r.DomainId = domainId
}

/* param action: 查询动作，"query"查询流量，"resolve"解析流量(Required) */
func (r *GetDomainStatisticsRequest) SetAction(action string) {
    r.Action = action
}

/* param domainName: 域名(Required) */
func (r *GetDomainStatisticsRequest) SetDomainName(domainName string) {
    r.DomainName = domainName
}

/* param start: 起始时间, UTC时间例如2017-11-10T23:00:00Z(Required) */
func (r *GetDomainStatisticsRequest) SetStart(start string) {
    r.Start = start
}

/* param end: 终止时间, UTC时间例如2017-11-10T23:00:00Z(Required) */
func (r *GetDomainStatisticsRequest) SetEnd(end string) {
    r.End = end
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetDomainStatisticsRequest) GetRegionId() string {
    return r.RegionId
}

type GetDomainStatisticsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetDomainStatisticsResult `json:"result"`
}

type GetDomainStatisticsResult struct {
    Time []int `json:"time"`
    Traffic []int `json:"traffic"`
}