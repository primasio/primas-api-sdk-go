package dtcpv1

type GroupSharePost struct {
	Version   string                 `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string                 `json:"type"`                // Fixed to "relation".
	Tag       string                 `json:"tag"`                 // Fixed to "group_share".
	SrcId     string                 `json:"src_id"`              // Content id.
	DestId    string                 `json:"dest_id"`             // Group id.
	Creator   *GroupSharePostCreator `json:"creator"`             // Creator.
	Created   int                    `json:"created"`             // Share created time. Unix timestamp.
	Status    string                 `json:"status"`              // Fixed to "created".
	Extra     *GroupSharePostExtra   `json:"extra,omitempty"`     // Extra metadata.
	Signature string                 `json:"signature,omitempty"` // Metadata signature.
}

type GroupSharePostCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name. For fast creation of new sub accounts.
}

type GroupSharePostExtra struct {
	ShareId           string `json:"share_id,omitempty"`           // Parent share id.
	ApplicationStatus string `json:"application_status,omitempty"` // For group requiring application. Fill "pending".
	ApplicationExpire int    `json:"application_expire,omitempty"` // Application expiration time.
}

func NewGroupSharePost() *GroupSharePost {
	return &GroupSharePost{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Group_share,
		Status:  CONST_DTCP_Status_Created,
	}
}

type GroupShareAppPut struct {
	Version   string                   `json:"version"`             // DTCP version. Fixed to "1.0"
	Atype     string                   `json:"type"`                // Fixed to "relation".
	Tag       string                   `json:"tag"`                 // Fixed to "group_share".
	ParentDna string                   `json:"parent_dna"`          // Latest share DNA.
	Status    string                   `json:"status"`              // Fixed to "updated".
	Updated   int                      `json:"updated"`             // Share updated time. Unix timestamp.
	Creator   *GroupShareAppPutCreator `json:"creator"`             // Creator. Group owner
	Extra     *GroupShareAppPutExtra   `json:"extra"`               // Extra metadata.
	Signature string                   `json:"signature,omitempty"` // Metadata signature.
}

type GroupShareAppPutCreator struct {
	AccountId    string `json:"account_id"`               // Account id. Root account id in the case of Sub account posting.
	SubAccountId string `json:"sub_account_id,omitempty"` // Sub account id. Refer to Sub account for details.
}

type GroupShareAppPutExtra struct {
	ApplicationStatus string `json:"application_status"` // "approved" or "declined"
}

func NewGroupShareAppPut() *GroupShareAppPut {
	return &GroupShareAppPut{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Group_share,
		Status:  CONST_DTCP_Status_Updated,
	}
}

// Delete
type GroupShareDelete struct {
	Version   string                   `json:"version"`           // DTCP version. Fixed to "1.0".
	Atype     string                   `json:"type"`              // Fixed to "relation".
	Tag       string                   `json:"tag"`               // Fixed to "group_share".
	ParentDna string                   `json:"parent_dna"`        // Latest share DNA.
	Status    string                   `json:"status"`            // Fixed to "deleted".
	Updated   int                      `json:"updated"`           // Share updated time. Unix timestamp.
	Creator   *GroupShareDeleteCreator `json:"creator,omitempty"` // Creator. Group owner.
	Signature string                   `json:"signature"`         // Metadata signature.
}

type GroupShareDeleteCreator struct {
	AccountId    string `json:"account_id"`               // Account id. Root account id in the case of Sub account posting.
	SubAccountId string `json:"sub_account_id,omitempty"` // Sub account id. Refer to Sub account for details.
}

func NewGroupShareDelete() *GroupShareDelete {
	return &GroupShareDelete{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Group_share,
		Status:  CONST_DTCP_Status_Deleted,
	}
}
