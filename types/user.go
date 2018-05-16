/*

 Copyright 2017 Loopring Project Ltd (Loopring Foundation).

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.

*/

package types

import "github.com/Loopring/go-ethereum/common"

// 白名单用户允许广播，但是订单不允许提供给miner

//go:generate gencodec -type WhiteListUser -out gen_white_list_user_json.go
type WhiteListUser struct {
	Owner      common.Address `json:"owner"`
	CreateTime int64          `json:"create_time"`
}
