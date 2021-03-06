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

type DescribeTemplatesRequest struct {

    core.JDCloudRequest

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 产品线标识 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* 模板名称 (Optional) */
    TemplateName *string `json:"templateName"`

    /* 模板类型，区分默认模板和用户自定义模板：1表示默认模板，2表示用户自定义模板  */
    TemplateType int `json:"templateType"`
}

/*
 * param templateType: 模板类型，区分默认模板和用户自定义模板：1表示默认模板，2表示用户自定义模板 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeTemplatesRequest(
    templateType int,
) *DescribeTemplatesRequest {

	return &DescribeTemplatesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/template",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        TemplateType: templateType,
	}
}

/*
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param serviceCode: 产品线标识 (Optional)
 * param templateName: 模板名称 (Optional)
 * param templateType: 模板类型，区分默认模板和用户自定义模板：1表示默认模板，2表示用户自定义模板 (Required)
 */
func NewDescribeTemplatesRequestWithAllParams(
    pageNumber *int,
    pageSize *int,
    serviceCode *string,
    templateName *string,
    templateType int,
) *DescribeTemplatesRequest {

    return &DescribeTemplatesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/template",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PageNumber: pageNumber,
        PageSize: pageSize,
        ServiceCode: serviceCode,
        TemplateName: templateName,
        TemplateType: templateType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeTemplatesRequestWithoutParam() *DescribeTemplatesRequest {

    return &DescribeTemplatesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/template",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeTemplatesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeTemplatesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param serviceCode: 产品线标识(Optional) */
func (r *DescribeTemplatesRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param templateName: 模板名称(Optional) */
func (r *DescribeTemplatesRequest) SetTemplateName(templateName string) {
    r.TemplateName = &templateName
}

/* param templateType: 模板类型，区分默认模板和用户自定义模板：1表示默认模板，2表示用户自定义模板(Required) */
func (r *DescribeTemplatesRequest) SetTemplateType(templateType int) {
    r.TemplateType = templateType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeTemplatesRequest) GetRegionId() string {
    return ""
}

type DescribeTemplatesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeTemplatesResult `json:"result"`
}

type DescribeTemplatesResult struct {
    NumberPages int64 `json:"numberPages"`
    NumberRecords int64 `json:"numberRecords"`
    PageNumber int64 `json:"pageNumber"`
    PageSize int64 `json:"pageSize"`
    TemplateCount int64 `json:"templateCount"`
    TemplateList []monitor.TemplateVo `json:"templateList"`
}