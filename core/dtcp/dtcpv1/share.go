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

package dtcpv1

import "github.com/shopspring/decimal"

type Share struct {
	Id            string        `json:"id"`              // Share id.
	SrcId         string        `json:"src_id"`          // Content id.
	DestId        string        `json:"dest_id"`         // Group id.
	Creator       *ShareCreator `json:"creator"`         // Creator.
	Created       int           `json:"created"`         // Share created time. Unix timestamp.
	Updated       int           `json:"updated"`         // Share updated time. Unix timestamp.
	Extra         *ShareExtra   `json:"extra,omitempty"` // Extra metadata.
	Signature     string        `json:"signature"`       // Metadata signature.
	Dna           string        `json:"dna"`             // Latest share DNA.
	TransactionId string        `json:"transaction_id"`  // Latest transaction id.
}

type ShareCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	AccountName    string `json:"account_name"`               // Account name.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name.
}

type ShareExtra struct {
	ShareId       string            `json:"share_id"`           // Parent share id.
	LikesTotal    int               `json:"likes_total"`        // Total likes number.
	CommentsTotal int               `json:"comments_total"`     // Total comments number.
	SharesTotal   int               `json:"shares_total"`       // Total shares number.
	PstTotal      decimal.Decimal   `json:"pst_total"`          // Total PST earned.
	PstUpdated    int               `json:"pst_updated"`        // Last PST updated time. Unix timestamp.
	IsLiked       bool              `json:"is_liked,omitempty"` // Whether current account liked this share.
	Content       *ContentGet       `json:"content"`            // Share related content.
	Report        *ContentReportGet `json:"report,omitempty"`   // Report metadata.
	Hp            int               `json:"hp"`                 // hp value.
}
