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


type RoutePropagation struct {

    /* 路由传播Id (Optional) */
    PropagationId string `json:"propagationId"`

    /* 边界网关Id (Optional) */
    BgwId string `json:"bgwId"`

    /* 路由传播范围，指定路由传播网段，CIDR格式，多个CIDR之间以英文逗号“,”分隔，0.0.0.0/0表示接受所有传播路由，设置特定网段就只能接收该网段范围内或子网段的路由传播 (Optional) */
    PropagationCidrs string `json:"propagationCidrs"`
}
