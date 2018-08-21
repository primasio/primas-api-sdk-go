package dtcpv1

type ContentCommentPost struct {
	Version   string                     `json:"version"`             // DTCP version. Fixed to "1.0"
	Atype     string                     `json:"type"`                // Fixed to "relation".
	Tag       string                     `json:"tag"`                 // Fixed to "share_comment".
	SrcId     string                     `json:"src_id"`              // Account id.
	DestId    string                     `json:"dest_id"`             // Share id.
	Hp        int                        `json:"hp"`                  // hp value
	Creator   *ContentCommentPostCreator `json:"creator"`             // Creator.
	Created   int                        `json:"created"`             // Comment created time. Unix timestamp.
	Status    string                     `json:"status"`              // Fixed to "created".
	Extra     *ContentCommentPostExtra   `json:"extra"`               // Extra metadata.
	Signature string                     `json:"signature,omitempty"` // Metadata signature.
}

type ContentCommentPostCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name. For fast creation of new sub accounts.
}

type ContentCommentPostExtra struct {
	ParentCommentId string `json:"parent_comment_id,omitempty"` // Parent comment id.
	Content         string `json:"content"`                     // Comment content.
	ContentHash     string `json:"content_hash"`                // Lowercase hex string of the SHA256 hash of the raw content.
}

func NewContentCommentPost() *ContentCommentPost {
	return &ContentCommentPost{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Share_comment,
		Status:  CONST_DTCP_Status_Created,
	}
}

type ContentCommentPut struct {
	Version   string                    `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string                    `json:"type"`                // Fixed to "relation".
	Tag       string                    `json:"tag"`                 // Fixed to "share_comment".
	ParentDna string                    `json:"parent_dna"`          // Latest comment DNA.
	Creator   *ContentCommentPutCreator `json:"creator"`             // Creator.
	Updated   int                       `json:"updated"`             // Comment updated time. Unix timestamp.
	Status    string                    `json:"status"`              // Fixed to "updated".
	Extra     *ContentCommentPutExtra   `json:"extra"`               // Extra metadata.
	Signature string                    `json:"signature,omitempty"` // Metadata signature.
}

type ContentCommentPutCreator struct {
	AccountId    string `json:"account_id"`               // Account id. Root account id in the case of Sub account posting.
	SubAccountId string `json:"sub_account_id,omitempty"` // Sub account id. Refer to Sub account for details.
}

type ContentCommentPutExtra struct {
	Content     string `json:"content"`      //Comment content.
	ContentHash string `json:"content_hash"` // Lowercase hex string of the SHA256 hash of the raw content.
}

func NewContentCommentPut() *ContentCommentPut {
	return &ContentCommentPut{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Share_comment,
		Status:  CONST_DTCP_Status_Updated,
	}
}

// Get
type ContentCommentGet struct {
	Id            string                    `json:"id"`             // Comment id.
	SrcId         string                    `json:"src_id"`         // Account id.
	DestId        string                    `json:"dest_id"`        // Share id.
	Hp            int                       `json:"hp"`             // hp value
	Creator       *ContentCommentGetCreator `json:"creator"`        // Creator.
	Created       int                       `json:"created"`        // Comment created time. Unix timestamp.
	Updated       int                       `json:"updated"`        // Comment created time. Unix timestamp.
	Extra         *ContentCommentGetExtra   `json:"extra"`          // Extra metadata.
	Signature     string                    `json:"signature"`      // Metadata signature.
	TransactionId string                    `json:"transaction_id"` // Latest transaction id.
}

type ContentCommentGetCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	AccountName    string `json:"account_name"`               // Account name.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name.
}

type ContentCommentGetExtra struct {
	Content  string              `json:"content"`  // Comment content.
	Comments []ContentCommentGet `json:"comments"` // Replying comments overview.
}
