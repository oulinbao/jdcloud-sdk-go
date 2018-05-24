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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/apis"
    "encoding/json"
    "errors"
)

type VpcClient struct {
    core.JDCloudClient
}

func NewVpcClient(credential *core.Credential) *VpcClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("vpc.jdcloud-api.com")

    return &VpcClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "vpc",
            Revision:    "0.2.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *VpcClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *VpcClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 查询弹性ip列表 */
func (c *VpcClient) DescribeElasticIps(request *vpc.DescribeElasticIpsRequest) (*vpc.DescribeElasticIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DescribeElasticIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 给网卡分配secondaryIp接口 */
func (c *VpcClient) AssignSecondaryIps(request *vpc.AssignSecondaryIpsRequest) (*vpc.AssignSecondaryIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.AssignSecondaryIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询VpcPeering资源详情 */
func (c *VpcClient) DescribeVpcPeering(request *vpc.DescribeVpcPeeringRequest) (*vpc.DescribeVpcPeeringResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DescribeVpcPeeringResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 给网卡解绑弹性Ip接口 */
func (c *VpcClient) DisassociateElasticIp(request *vpc.DisassociateElasticIpRequest) (*vpc.DisassociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DisassociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除弹性Ip */
func (c *VpcClient) DeleteElasticIp(request *vpc.DeleteElasticIpRequest) (*vpc.DeleteElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DeleteElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建一个或者多个弹性Ip */
func (c *VpcClient) CreateElasticIps(request *vpc.CreateElasticIpsRequest) (*vpc.CreateElasticIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.CreateElasticIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 给网卡删除secondaryIp接口 */
func (c *VpcClient) UnassignSecondaryIps(request *vpc.UnassignSecondaryIpsRequest) (*vpc.UnassignSecondaryIpsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.UnassignSecondaryIpsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询Vpc信息详情 */
func (c *VpcClient) DescribeVpc(request *vpc.DescribeVpcRequest) (*vpc.DescribeVpcResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DescribeVpcResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询VpcPeering资源列表 */
func (c *VpcClient) DescribeVpcPeerings(request *vpc.DescribeVpcPeeringsRequest) (*vpc.DescribeVpcPeeringsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DescribeVpcPeeringsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询安全组列表 */
func (c *VpcClient) DescribeNetworkSecurityGroups(request *vpc.DescribeNetworkSecurityGroupsRequest) (*vpc.DescribeNetworkSecurityGroupsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DescribeNetworkSecurityGroupsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询子网列表 */
func (c *VpcClient) DescribeSubnets(request *vpc.DescribeSubnetsRequest) (*vpc.DescribeSubnetsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DescribeSubnetsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询子网信息详情 */
func (c *VpcClient) DescribeSubnet(request *vpc.DescribeSubnetRequest) (*vpc.DescribeSubnetResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DescribeSubnetResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 创建VpcPeering接口 */
func (c *VpcClient) CreateVpcPeering(request *vpc.CreateVpcPeeringRequest) (*vpc.CreateVpcPeeringResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.CreateVpcPeeringResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询安全组信息详情 */
func (c *VpcClient) DescribeNetworkSecurityGroup(request *vpc.DescribeNetworkSecurityGroupRequest) (*vpc.DescribeNetworkSecurityGroupResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DescribeNetworkSecurityGroupResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 修改VpcPeering接口 */
func (c *VpcClient) ModifyVpcPeering(request *vpc.ModifyVpcPeeringRequest) (*vpc.ModifyVpcPeeringResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.ModifyVpcPeeringResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 给网卡绑定弹性Ip接口 */
func (c *VpcClient) AssociateElasticIp(request *vpc.AssociateElasticIpRequest) (*vpc.AssociateElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.AssociateElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 删除VpcPeering接口 */
func (c *VpcClient) DeleteVpcPeering(request *vpc.DeleteVpcPeeringRequest) (*vpc.DeleteVpcPeeringResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DeleteVpcPeeringResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* ElasticIp资源信息详情 */
func (c *VpcClient) DescribeElasticIp(request *vpc.DescribeElasticIpRequest) (*vpc.DescribeElasticIpResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DescribeElasticIpResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}

/* 查询私有网络列表 */
func (c *VpcClient) DescribeVpcs(request *vpc.DescribeVpcsRequest) (*vpc.DescribeVpcsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &vpc.DescribeVpcsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        return nil, err
    }

    return jdResp, err
}
