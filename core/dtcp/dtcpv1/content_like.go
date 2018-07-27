package dtcpv1

type ContentLikePost struct {
	Version   string                  `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string                  `json:"type"`                // Fixed to "relation".
	Tag       string                  `json:"tag"`                 // Fixed to "share_like".
	SrcId     string                  `json:"src_id"`              // Account id.
	DestId    string                  `json:"dest_id"`             // Share id.
	Creator   *ContentLikePostCreator `json:"creator"`             // Creator.
	Created   int                     `json:"created"`             // Like created time. Unix timestamp.
	Status    string                  `json:"status"`              // Fixed to "created".
	Signature string                  `json:"signature,omitempty"` // Metadata signature.
}

type ContentLikePostCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name. For fast creation of new sub accounts.
}

func NewContentLikePost() *ContentLikePost {
	return &ContentLikePost{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Share_like,
		Status:  CONST_DTCP_Status_Created,
	}
}

type ContentLikeDelete struct {
	Version   string `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string `json:"type"`                // Fixed to "relation".
	Tag       string `json:"tag"`                 // Fixed to "share_like".
	ParentDna string `json:"parent_dna"`          // Latest DNA of the like.
	Updated   int    `json:"updated"`             // Like updated time. Unix timestamp.
	Status    string `json:"status"`              // Fixed to "deleted".
	Signature string `json:"signature,omitempty"` // Metadata signature.
}

func NewContentLikeDelete() *ContentLikeDelete {
	return &ContentLikeDelete{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Share_like,
		Status:  CONST_DTCP_Status_Deleted,
	}
}
