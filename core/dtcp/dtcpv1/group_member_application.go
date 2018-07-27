package dtcpv1

// created
type GroupMemberApplicationPost struct {
	Version   string                     `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string                     `json:"type"`                // Fixed to "relation".
	Tag       string                     `json:"tag"`                 // Fixed to "group_member".
	SrcId     string                     `json:"src_id"`              // Account id. Root account id in the case of Sub account.
	DestId    string                     `json:"dest_id"`             // Group id.
	Created   int                        `json:"created"`             // Member joining time. Unix timestamp.
	Creator   *GroupMemberAppPostCreator `json:"creator"`             // Creator. Group owner.
	Status    string                     `json:"status"`              // Fixed to "created".
	Extra     *GroupMemberAppPostExtra   `json:"extra,omitempty"`     // Extra metadata.
	Signature string                     `json:"signature,omitempty"` // Metadata signature.
}

type GroupMemberAppPostCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name. For fast creation of new sub accounts.
}

type GroupMemberAppPostExtra struct {
	ApplicationStatus string `json:"application_status"` // For group requiring application. Fill "pending".
	ApplicationExpire int    `json:"application_expire"` // Application expiration time.
}

func NewGroupMemberApplicationPost() *GroupMemberApplicationPost {
	return &GroupMemberApplicationPost{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Group_member,
		Status:  CONST_DTCP_Status_Created,
	}
}

// updated
type GroupMemberApplicationPut struct {
	Version   string                    `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string                    `json:"type"`                // Fixed to "relation".
	Tag       string                    `json:"tag"`                 // Fixed to "group_member".
	ParentDna string                    `json:"parent_dna"`          // Latest group member DNA.
	Status    string                    `json:"status"`              // Fixed to "updated".
	Updated   int                       `json:"updated"`             // Member updating time. Unix timestamp.
	Creator   *GroupMemberAppPutCreator `json:"creator"`             // Creator. Group owner.
	Extra     *GroupMemberAppPutExtra   `json:"extra"`               // Extra metadata.
	Signature string                    `json:"signature,omitempty"` // Metadata signature.
}

type GroupMemberAppPutCreator struct {
	AccountId    string `json:"account_id"`               // Account id. Root account id in the case of Sub account posting.
	SubAccountId string `json:"sub_account_id,omitempty"` // Sub account id. Refer to Sub account for details.
}

type GroupMemberAppPutExtra struct {
	ApplicationStatus string `json:"application_status"` // "approved" or "declined".
}

func NewGroupMemberApplicationPut() *GroupMemberApplicationPut {
	return &GroupMemberApplicationPut{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Group_member,
		Status:  CONST_DTCP_Status_Updated,
	}
}
