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


type CreateSubUserInfo struct {

    /* 子账号用户名，4~20位数字、字母、中文、下划线、中划线  */
    Name string `json:"name"`

    /* 描述，0~256个字符 (Optional) */
    Description *string `json:"description"`

    /* 密码，6~20位，至少包含一个字母，至少包含一个数字或半角符号  */
    Password string `json:"password"`

    /* 手机号码，区号-手机号，目前只支持0086-中国手机号码  */
    Phone string `json:"phone"`

    /* 邮箱  */
    Email string `json:"email"`

    /* 确认密码  */
    PasswordConfirm string `json:"passwordConfirm"`

    /* 是否创建accessKey  */
    CreateAk bool `json:"createAk"`
}
