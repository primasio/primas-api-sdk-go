/*
 * Copyright 2018 Primas Lab Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tool

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/shopspring/decimal"
)

type TestTokenPreLock struct {
	ID          string           `json:"id"`
	CreatedAt   int              `json:"created_at,omitempty"`
	UserAddress string           `json:"user_address"`
	NodeAddress *string          `json:"node_address,omitempty"`
	Amount      *decimal.Decimal `json:"amount,omitempty"`
	NodeFee     decimal.Decimal  `json:"node_fee,omitempty"`
	TxHash      string           `json:"tx_hash"`
	AccountId   string           `json:"account_id,omitempty"`
	Creator     *CreatorObj      `json:"creator,omitempty"`
	Tags        []*string        `json:"tags"`
	SubsPtr     []*SubAccountObj `json:"subs_ptr,omitempty"`
	Subs        []SubAccountObj  `json:"subs,omitempty"`
	FullName    string
}

type CreatorObj struct {
	SubAccount  *SubAccountObj `json:"sub_account,omitempty"`
	AccountName string         `json:"account_name"`
	AccountId   string         `json:"account_id"`
}

type SubAccountObj struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func TestTokenPreLock_JsonSerialize(t *testing.T) {
	curSubAccountObj := SubAccountObj{
		ID:   "Sub01",
		Name: "movi",
	}
	curCreator := CreatorObj{
		AccountName: "kevin",
		AccountId:   "NO134134",
		SubAccount:  &curSubAccountObj,
	}
	curSubAccountObj_1 := SubAccountObj{
		ID:   "Sub01_1",
		Name: "movi_1",
	}
	curSubAccountObj_2 := SubAccountObj{
		ID:   "Sub01_2",
		Name: "movi_2",
	}
	subsPtr := make([]*SubAccountObj, 0)
	subsPtr = append(subsPtr, &curSubAccountObj_1)
	subsPtr = append(subsPtr, &curSubAccountObj_2)

	subs := make([]SubAccountObj, 0)
	subs = append(subs, curSubAccountObj_1)
	subs = append(subs, curSubAccountObj_2)

	curTags := make([]*string, 0)
	as_1 := "wa"
	as_2 := "ab"
	curTags = append(curTags, &as_1)
	curTags = append(curTags, &as_2)
	nodeAddress := "0xd75407ad8cabeeebfed78c4f3794208b3339fbf4"
	amount := decimal.NewFromFloat(10000000000000000000000000000000000000000000)
	tokenPreLock := TestTokenPreLock{
		ID:          "bb48c151f509a1064e574dceb95e3b39eb6a1dbd9dfefa75386dc2250fdeac80",
		CreatedAt:   1527758683,
		UserAddress: "0x251af16381A9908709DdaaA9185FB21f2811C322",
		NodeAddress: &nodeAddress,
		Amount:      &amount,
		NodeFee:     decimal.NewFromFloat(4),
		TxHash:      "aa48c151f509a1064e574dceb95e3b39eb6a1dbd9dfefa75386dc2250fdeac80",
		AccountId:   "0x111af16381A9908709DdaaA9185FB21f2811C3dd",
		Creator:     &curCreator,
		SubsPtr:     subsPtr,
		Subs:        subs,
		Tags:        curTags,
		FullName:    "Full you<>&",
	}

	beforeValue, err := json.Marshal(tokenPreLock)
	if err != nil {
		t.Errorf("json.Marshal error:%v", err.Error())
		return
	}
	t.Logf("TokenPreLock_JsonSerialize before:%v", string(beforeValue))

	resultValue, err := StructToSignature(tokenPreLock)
	if err != nil {
		t.Errorf("TokenPreLock_JsonSerialize error:%v", err.Error())
		return
	}

	t.Logf("TokenPreLock_JsonSerialize after:%v", resultValue)
}

func TestRecoverPublickey(t *testing.T) {
	msg := `{"amount":123,"created":1532527550,"node_fee":123,"node_id":"58f47077984e5daa4d2ea46f2e689177a1655c1321544e69f851530a789e9fd7"}`
	sig := `644577bf6d045815e433a5fcc0b8be4650814e0da7f377a5f252b5d6c3fc615b11cd6d2fd0afe5bfdc8c1e55de6469e8ae9fe1cdbd053be54e6db6c2e4efe45b01`

	getAddress, err := RecoverPublickey(msg, sig)
	if err != nil {
		t.Errorf("RecoverPublickey error:%v", err.Error())
		return
	}

	log.Println(getAddress)
	if getAddress != "0xD75407aD8caBEeeBFeD78C4F3794208b3339fbF4" {
		t.Errorf("RecoverPublickey failed ")
		return
	}
}
