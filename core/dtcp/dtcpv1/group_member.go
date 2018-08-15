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
		Atype:   CONST_DTCP_Type_Relation,
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

// Get
type GroupMemberGet struct {
	Id            string                 `json:"id"`              // Group member id.
	SrcId         string                 `json:"src_id"`          // Account id. Root account id in the case of Sub account.
	DestId        string                 `json:"dest_id"`         // Group id.
	Creator       *GroupMemberGetCreator `json:"creator"`         // Creator.
	Created       int                    `json:"created"`         // Member joining time. Unix timestamp.
	Updated       int                    `json:"updated"`         // Member updating time. Unix timestamp.
	Extra         *GroupMemberGetExtra   `json:"extra,omitempty"` // Extra metadata.
	Signature     string                 `json:"signature"`       // Metadata signature.
	Dna           string                 `json:"dna"`             // Group member DNA.
	TransactionId string                 `json:"transaction_id"`  // Latest transaction id.
	Account       *AccountGet            `json:"account"`         // Related member account.
}

type GroupMemberGetCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	AccountName    string `json:"account_name"`               // Account name.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name. For fast creation of new sub accounts.
}

type GroupMemberGetExtra struct {
	ApplicationStatus string `json:"application_status"` // "pending", "approved" or "declined".
}

// Delete
type GroupMemberDelete struct {
	Version   string                    `json:"version"`    // DTCP version. Fixed to "1.0".
	Atype     string                    `json:"type"`       // Fixed to "relation".
	Tag       string                    `json:"tag"`        // Fixed to "group_member".
	ParentDna string                    `json:"parent_dna"` // Latest group member DNA.
	Status    string                    `json:"status"`     // "deleted".
	Updated   int                       `json:"updated"`    // Member quiting time. Unix timestamp.
	Creator   *GroupMemberDeleteCreator `json:"creator"`    // Creator.
	Signature string                    `json:"signature"`  // Metadata signature.
}

type GroupMemberDeleteCreator struct {
	AccountId    string `json:"account_id"`               // Account id. Root account id in the case of Sub account posting.
	SubAccountId string `json:"sub_account_id,omitempty"` // Sub account id. Refer to Sub account for details.
}

func NewGroupMemberDelete() *GroupMemberDelete {
	return &GroupMemberDelete{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Group_member,
		Status:  CONST_DTCP_Status_Deleted,
	}
}
