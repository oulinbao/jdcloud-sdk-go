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


type EmrWorkflow struct {

    /*  (Optional) */
    Id *int `json:"id"`

    /* 工作流ID (Optional) */
    WorkflowId *string `json:"workflowId"`

    /* 工作流名称 (Optional) */
    WorkflowName *string `json:"workflowName"`

    /* 工作流状态 (Optional) */
    Status *string `json:"status"`

    /* 用户名 (Optional) */
    Userpin *string `json:"userpin"`

    /* 工作流创建时间 (Optional) */
    CreateTime *string `json:"createTime"`

    /* 数据中心，即regionId (Optional) */
    DataCenter *string `json:"dataCenter"`

    /* 上一次修改时间 (Optional) */
    ModifyTime *string `json:"modifyTime"`

    /* "是否自身依赖"
"1：自身依赖(默认)，0：非依赖"
 (Optional) */
    IsSelfDependence *bool `json:"isSelfDependence"`

    /* "0:即时任务(默认)"
"1：周期性任务"
"2：定时任务"
 (Optional) */
    TaskScheduleType *int `json:"taskScheduleType"`

    /* 失败后是否发送通知 (Optional) */
    IsSendMsg *bool `json:"isSendMsg"`
}
