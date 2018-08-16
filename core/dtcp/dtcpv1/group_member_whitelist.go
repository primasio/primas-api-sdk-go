package dtcpv1

// created
type GroupMemberWhitelistPost struct {
	Version   string                       `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string                       `json:"type"`                // Fixed to "relation".
	Tag       string                       `json:"tag"`                 // Fixed to "group_member_whitelist".
	SrcId     string                       `json:"src_id"`              // Account id.
	DestId    string                       `json:"dest_id"`             // Group id.
	Creator   *GroupMemberWhitePostCreator `json:"creator"`             // Creator.
	Created   int                          `json:"created"`             // Whitelist creating time. Unix timestamp.
	Status    string                       `json:"status"`              // Fixed to "created".
	Extra     *GroupMemberWhitePostExtra   `json:"extra"`               // Extra metadata.
	Signature string                       `json:"signature,omitempty"` // Metadata signature.
}

type GroupMemberWhitePostCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name. For fast creation of new sub accounts.
}

type GroupMemberWhitePostExtra struct {
	ApplicationStatus string `json:"application_status"` // Fixed to "pending".
}

func NewGroupMemberWhitelistPost() *GroupMemberWhitelistPost {
	return &GroupMemberWhitelistPost{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Group_member_whitelist,
		Status:  CONST_DTCP_Status_Created,
	}
}

// updated
type GroupMemberWhitelistPut struct {
	Version   string                      `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string                      `json:"type"`                // Fixed to "relation".
	Tag       string                      `json:"tag"`                 // Fixed to "group_member_whitelist".
	ParentDna string                      `json:"parent_dna"`          // Latest whitelist DNA.
	Status    string                      `json:"status"`              // Fixed to "updated".
	Updated   int                         `json:"updated"`             // Whitelist updating time. Unix timestamp.
	Creator   *GroupMemberWhitePutCreator `json:"creator"`             // Creator.
	Extra     *GroupMemberWhitePutExtra   `json:"extra"`               // Extra metadata.
	Signature string                      `json:"signature,omitempty"` // Metadata signature.
}

type GroupMemberWhitePutCreator struct {
	AccountId    string `json:"account_id"`               // Account id. Root account id in the case of Sub account posting.
	SubAccountId string `json:"sub_account_id,omitempty"` // Sub account id. Refer to Sub account for details.
}

type GroupMemberWhitePutExtra struct {
	ApplicationStatus string `json:"application_status"` // "approved" or "declined".
}

func NewGroupMemberWhitelistPut() *GroupMemberWhitelistPut {
	return &GroupMemberWhitelistPut{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Group_member_whitelist,
		Status:  CONST_DTCP_Status_Updated,
	}
}

// get
type GroupMemberWhitelistGet struct {
	Id        string                      `json:"id"`                  // Group member id.
	SrcId     string                      `json:"src_id"`              // Account id. Root account id in the case of Sub account.
	DestId    string                      `json:"dest_id"`             // Group id.
	Creator   *GroupMemberWhiteGetCreator `json:"creator"`             // Creator.
	Created   int                         `json:"created"`             // Member joining time. Unix timestamp.
	Status    string                      `json:"status"`              // Fixed to "created".
	Extra     *GroupMemberWhiteGetExtra   `json:"extra,omitempty"`     // Extra metadata.
	Signature string                      `json:"signature,omitempty"` // Metadata signature.
	DNA       string                      `json:"dna"`                 // Group member DNA.
	Account   *AccountGet                 `json:"account"`             // Related member account.
}

type GroupMemberWhiteGetCreator struct {
	AccountId      string `json:"account_id"`
	AccountName    string `json:"account_name"`
	SubAccountId   string `json:"sub_account_id,omitempty"`
	SubAccountName string `json:"sub_account_name,omitempty"`
}

type GroupMemberWhiteGetExtra struct {
	ApplicationStatus string `json:"application_status"`
}

// detele
type GroupMemberWhitelistDelete struct {
	Version   string                             `json:"version"`    // DTCP version. Fixed to "1.0".
	Atype     string                             `json:"atype"`      // Fixed to "relation".
	Tag       string                             `json:"tag"`        // Fixed to "group_member_whitelist".
	ParentDna string                             `json:"parent_dna"` // Latest whitelist DNA.
	Status    string                             `json:"status"`     // Fixed to "deleted".
	Updated   int                                `json:"updated"`    // Whitelist updating time. Unix timestamp.
	Creator   *GroupMemberWhitelistDeleteCreator `json:"creator"`    // Creator.
	Signature string                             `json:"signature"`  // Metadata signature.
}

type GroupMemberWhitelistDeleteCreator struct {
	AccountId    string `json:"account_id"`               // Account id. Root account id in the case of Sub account posting.
	SubAccountId string `json:"sub_account_id,omitempty"` // Sub account id. Refer to Sub account for details.
}

func NewGroupMemberWhitelistDelete() *GroupMemberWhitelistDelete {
	return &GroupMemberWhitelistDelete{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Group_member_whitelist,
		Status:  CONST_DTCP_Status_Deleted,
	}
}
