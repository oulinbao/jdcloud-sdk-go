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
    vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
)

type DescribeElasticIpRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* ElasticIp ID  */
    ElasticIpId string `json:"elasticIpId"`
}

/*
 * param regionId: Region ID 
 * param elasticIpId: ElasticIp ID 
 */
func NewDescribeElasticIpRequest(
    regionId string,
    elasticIpId string,
) *DescribeElasticIpRequest {

	return &DescribeElasticIpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/elasticIps/{elasticIpId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ElasticIpId: elasticIpId,
	}
}

func (r *DescribeElasticIpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

func (r *DescribeElasticIpRequest) SetElasticIpId(elasticIpId string) {
    r.ElasticIpId = elasticIpId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeElasticIpRequest) GetRegionId() string {
    fieldName := "RegionId"
    reqType := reflect.TypeOf(r)
    value := reflect.ValueOf(r)
    _, ok := reqType.FieldByName(fieldName)
    if ok {
        return value.FieldByName(fieldName).String()
    }

    return ""
}

type DescribeElasticIpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeElasticIpResult `json:"result"`
}

type DescribeElasticIpResult struct {
    ElasticIp vpc.ElasticIp `json:"elasticIp"`
}