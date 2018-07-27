package dtcpv1

import (
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
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
