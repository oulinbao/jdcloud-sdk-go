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
    billing "github.com/jdcloud-api/jdcloud-sdk-go/services/billing/models"
)

type QueryConsumptionOverViewRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /*   */
    TimeType int `json:"timeType"`

    /*   */
    StartTime string `json:"startTime"`

    /*   */
    EndTime string `json:"endTime"`
}

/*
 * param regionId:  (Required)
 * param timeType:  (Required)
 * param startTime:  (Required)
 * param endTime:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryConsumptionOverViewRequest(
    regionId string,
    timeType int,
    startTime string,
    endTime string,
) *QueryConsumptionOverViewRequest {

	return &QueryConsumptionOverViewRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/consumeOverView",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TimeType: timeType,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId:  (Required)
 * param timeType:  (Required)
 * param startTime:  (Required)
 * param endTime:  (Required)
 */
func NewQueryConsumptionOverViewRequestWithAllParams(
    regionId string,
    timeType int,
    startTime string,
    endTime string,
) *QueryConsumptionOverViewRequest {

    return &QueryConsumptionOverViewRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/consumeOverView",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TimeType: timeType,
        StartTime: startTime,
        EndTime: endTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryConsumptionOverViewRequestWithoutParam() *QueryConsumptionOverViewRequest {

    return &QueryConsumptionOverViewRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/consumeOverView",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *QueryConsumptionOverViewRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param timeType: (Required) */
func (r *QueryConsumptionOverViewRequest) SetTimeType(timeType int) {
    r.TimeType = timeType
}

/* param startTime: (Required) */
func (r *QueryConsumptionOverViewRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: (Required) */
func (r *QueryConsumptionOverViewRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryConsumptionOverViewRequest) GetRegionId() string {
    return r.RegionId
}

type QueryConsumptionOverViewResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryConsumptionOverViewResult `json:"result"`
}

type QueryConsumptionOverViewResult struct {
    Pin string `json:"pin"`
    StartTime string `json:"startTime"`
    EndTime string `json:"endTime"`
    ActualFee int `json:"actualFee"`
    CashPayFee int `json:"cashPayFee"`
    BalancePayFee int `json:"balancePayFee"`
    CashCouponPayFee int `json:"cashCouponPayFee"`
    ArrearFee int `json:"arrearFee"`
    ConsumptionProductVoList []billing.ConsumptionProduct `json:"consumptionProductVoList"`
    ConsumptionList interface{} `json:"consumptionList"`
}