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


type ClusterDetailInfo struct {

    /* Master节点数量 (Optional) */
    MaterNum int `json:"materNum"`

    /* Master节点CPU (Optional) */
    MasterCore int `json:"masterCore"`

    /* Master节点内存（推荐至少8G内存，否则服务可能会部署失败） (Optional) */
    MasterMemory int `json:"masterMemory"`

    /* "Master节点云盘类型，可传类型为（以下以“/”分割各类型）"
"NBD/NBD_SATA"
"分别代表：性能型/容量型"
 (Optional) */
    MasterDiskType string `json:"masterDiskType"`

    /* Master节点云盘容量，必须是10的整数倍，且大于20小于3000 (Optional) */
    MasterDiskVolume int `json:"masterDiskVolume"`

    /* Master节点规格，比如：g.n1.xlarge，更多规格请参考[文档](https://www.jdcloud.com/help/detail/296/isCatalog/1) (Optional) */
    MasterInstanceType string `json:"masterInstanceType"`

    /* master节点实例信息 (Optional) */
    MasterInstanceInfo string `json:"masterInstanceInfo"`

    /* Slave节点数量 (Optional) */
    SlaveNum int `json:"slaveNum"`

    /* Slave节点CPU (Optional) */
    SlaveCore int `json:"slaveCore"`

    /* Slave节点内存(推荐至少4G内存，否则服务可能会部署失败) (Optional) */
    SlaveMemory int `json:"slaveMemory"`

    /* "Slave节点云盘类型，可传类型为（以下以“/”分割各类型）"
"NBD/NBD_SATA"
"分别代表：性能型/容量型"
 (Optional) */
    SlaveDiskType string `json:"slaveDiskType"`

    /* Slave节点云盘容量，必须是10的整数倍，且大于20小于3000 (Optional) */
    SlaveDiskVolume int `json:"slaveDiskVolume"`

    /* Slave节点规格，比如：g.n1.xlarge，更多规格请参考[文档](https://www.jdcloud.com/help/detail/296/isCatalog/1) (Optional) */
    SlaveInstanceType string `json:"slaveInstanceType"`

    /* Slave节点实例信息 (Optional) */
    SlaveInstanceInfo string `json:"slaveInstanceInfo"`
}
