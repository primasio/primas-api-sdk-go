package dtcpv1

type GroupMemberPost struct {
	Version   string                  `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string                  `json:"type"`                // Fixed to "relation".
	Tag       string                  `json:"tag"`                 // Fixed to "group_member".
	SrcId     string                  `json:"src_id"`              // Account id. Root account id in the case of Sub account.
	DestId    string                  `json:"dest_id"`             // Group id.
	Creator   *GroupMemberPostCreator `json:"creator"`             // Creator.
	Created   int                     `json:"created"`             // Member joining time. Unix timestamp.
	Status    string                  `json:"status"`              // Fixed to "created".
	Extra     *GroupMemberPostExtra   `json:"extra,omitempty"`     // Extra metadata.
	Signature string                  `json:"signature,omitempty"` // Metadata signature.
}

type GroupMemberPostCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name. For fast creation of new sub accounts.
}

type GroupMemberPostExtra struct {
	ApplicationStatus string `json:"application_status"` // For group requiring application. Fill "pending".
	ApplicationExpire int    `json:"application_expire"` // Application expiration time.
}

func NewGroupMemberPost() *GroupMemberPost {
	return &GroupMemberPost{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Object,
		Tag:     CONST_DTCP_Tag_Group_member,
		Status:  CONST_DTCP_Status_Created,
	}
}

type GroupMemberPut struct {
	Version   string                 `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string                 `json:"type"`                // Fixed to "relation".
	Tag       string                 `json:"tag"`                 // Fixed to "group_member".
	ParentDna string                 `json:"parent_dna"`          // Latest group member DNA.
	Status    string                 `json:"status"`              // Fixed to "updated".
	Updated   int                    `json:"updated"`             // Member updating time. Unix timestamp.
	Creator   *GroupMemberPutCreator `json:"creator"`             // Creator. Group owner.
	Extra     *GroupMemberPutExtra   `json:"extra"`               // Extra metadata.
	Signature string                 `json:"signature,omitempty"` // Metadata signature.
}

type GroupMemberPutCreator struct {
	AccountId    string `json:"account_id"`               // 	Account id. Root account id in the case of Sub account posting.
	SubAccountId string `json:"sub_account_id,omitempty"` // Sub account id. Refer to Sub account for details.
}

type GroupMemberPutExtra struct {
	ApplicationStatus string `json:"application_status"` // "approved" or "declined".
}

func NewGroupMemberPut() *GroupMemberPut {
	return &GroupMemberPut{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Object,
		Tag:     CONST_DTCP_Tag_Group_member,
		Status:  CONST_DTCP_Status_Updated,
	}
}
