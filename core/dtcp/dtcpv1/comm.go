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

import (
	"encoding/hex"
	"errors"

	"github.com/primasio/go-ethereum/crypto"
)

const (
	CONST_DTCP_Version_v1 = "1.0"

	CONST_DTCP_Type_Object   = "object"
	CONST_DTCP_Type_Relation = "relation"

	CONST_DTCP_Status_Created = "created"
	CONST_DTCP_Status_Updated = "updated"
	CONST_DTCP_Status_Deleted = "deleted"

	CONST_DTCP_Tag_Account                = "account"
	CONST_DTCP_Tag_Group                  = "group"
	CONST_DTCP_Tag_Group_member           = "group_member"
	CONST_DTCP_Tag_Group_member_whitelist = "group_member_whitelist"
	CONST_DTCP_Tag_Group_share            = "group_share"
	CONST_DTCP_Tag_Article                = "article"
	CONST_DTCP_Tag_Image                  = "image"
	CONST_DTCP_Tag_Share_report           = "share_report"
	CONST_DTCP_Tag_Share_like             = "share_like"
	CONST_DTCP_Tag_Share_comment          = "share_comment"

	CONST_DTCP_ApplicationStatus_Pending  = "pending"
	CONST_DTCP_ApplicationStatus_Approved = "approved"
	CONST_DTCP_ApplicationStatus_Declined = "declined"

	CONST_DTCP_Group_AllowJoin_Type_all         = "all"
	CONST_DTCP_Group_AllowJoin_Type_application = "application"

	CONST_DTCP_Group_AllowPost_Type_all         = "all"
	CONST_DTCP_Group_AllowPost_Type_none        = "none"
	CONST_DTCP_Group_AllowPost_Type_application = "application"

	CONST_DTCP_Group_AccountRole_Type_Owner     = "owner"
	CONST_DTCP_Group_AccountRole_Type_Member    = "member"
	CONST_DTCP_Group_AccountRole_Type_Applicant = "applicant"
	CONST_DTCP_Group_AccountRole_Type_None      = "none"

	CONST_DTCP_Image_Type_Image = "image"
	CONST_DTCP_Image_Type_Audio = "audio"
	CONST_DTCP_Image_Type_Video = "video"
)

func NewDna(signature string) (string, error) {
	if signature == "" {
		return "", errors.New("parameter is empty")
	}

	digest := crypto.Keccak256([]byte(signature))

	return hex.EncodeToString(digest), nil
}

func NewId(Dna string) (string, error) {
	if Dna == "" {
		return "", errors.New("parameter is empty")
	}

	digest := crypto.Keccak256([]byte(Dna))

	return hex.EncodeToString(digest), nil
}

func HashValue(value []byte) string {
	digest := crypto.Keccak256(value)

	return hex.EncodeToString(digest)
}
