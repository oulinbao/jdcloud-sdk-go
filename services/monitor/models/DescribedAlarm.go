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

package models


type DescribedAlarm struct {

    /* 计算单位 (Optional) */
    CalculateUnit string `json:"calculateUnit"`

    /* 统计方法：平均值=avg、最大值=max、最小值=min (Optional) */
    Calculation string `json:"calculation"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 降采样方法 (Optional) */
    DownSample string `json:"downSample"`

    /* 是否启用 (Optional) */
    Enabled int64 `json:"enabled"`

    /* 报警规则ID (Optional) */
    Id string `json:"id"`

    /* 监控项 (Optional) */
    Metric string `json:"metric"`

    /* 监控项名称 (Optional) */
    MetricName string `json:"metricName"`

    /*  (Optional) */
    NoticeLevel NoticeLevel `json:"noticeLevel"`

    /* 告警周期 (Optional) */
    NoticePeriod int64 `json:"noticePeriod"`

    /* gt, gte, lt, lte, eq, ne (Optional) */
    Operation string `json:"operation"`

    /* 统计周期（单位：分钟） (Optional) */
    Period int64 `json:"period"`

    /* 地域信息 (Optional) */
    Region string `json:"region"`

    /* 资源id (Optional) */
    ResourceId string `json:"resourceId"`

    /* 产品线编码 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 监控项状态：1正常，2告警，4数据不足 (Optional) */
    Status int64 `json:"status"`

    /* 标签 (Optional) */
    Tags interface{} `json:"tags"`

    /* 告警阈值 (Optional) */
    Threshold float64 `json:"threshold"`

    /* 告警次数 (Optional) */
    Times int64 `json:"times"`
}