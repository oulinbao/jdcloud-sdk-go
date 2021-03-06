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


type RRInfo struct {

    /* 创建者 (Optional) */
    Creator string `json:"creator"`

    /* 线路名称 (Optional) */
    ViewName string `json:"viewName"`

    /* 域名解析的唯一ID (Optional) */
    Id int `json:"id"`

    /* 主机记录 (Optional) */
    HostRecord string `json:"hostRecord"`

    /* 解析记录的值 (Optional) */
    HostValue string `json:"hostValue"`

    /* 是否是京东云资源 (Optional) */
    JcloudRes bool `json:"jcloudRes"`

    /* 优先级，只存在于某些解析记录类型 (Optional) */
    MxPriority int `json:"mxPriority"`

    /* 端口，只存在于某些解析记录类型 (Optional) */
    Port int `json:"port"`

    /* 解析记录的生存时间 (Optional) */
    Ttl int `json:"ttl"`

    /* 解析记录的类型 (Optional) */
    Type string `json:"type"`

    /* 解析记录的权重 (Optional) */
    Weight int `json:"weight"`

    /* 解析线路的ID (Optional) */
    ViewValue []int `json:"viewValue"`

    /* 解析记录的状态 (Optional) */
    ResolvingStatus string `json:"resolvingStatus"`

    /* 解析记录更新的时间 (Optional) */
    UpdateTime int64 `json:"updateTime"`
}
