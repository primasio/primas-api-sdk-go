package dtcpv1

type ContentPost struct {
	Version     string              `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype       string              `json:"type"`                // Fixed to "object".
	Tag         string              `json:"tag"`                 // Content type. Currently "article", "image" are supported.
	Title       string              `json:"title"`               // Content title.
	Creator     *ContentPostCreator `json:"creator"`             // Creator.
	Abstract    string              `json:"abstract"`            // Content abstract.
	Language    string              `json:"language"`            // Content language. RFC4646 defined locales such as "en-US"
	Category    string              `json:"category"`            // Content categories. Comma separated words list.
	Created     int                 `json:"created"`             // Content creation time. Unix timestamp.
	Content     string              `json:"content"`             // Raw content in base64 encoded format.
	ContentHash string              `json:"content_hash"`        // Lowercase hex string of the SHA256 hash of the raw content.
	License     *License            `json:"license,omitempty"`   // Content authorization license. "none" is used if empty.
	Status      string              `json:"status"`              // Fixed to "created".
	Signature   string              `json:"signature,omitempty"` // Metadata signature.
}

type ContentPostCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name. For fast creation of new sub accounts.
}

func NewContentPost_Aritcle() *ContentPost {
	return &ContentPost{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Object,
		Tag:     CONST_DTCP_Tag_Article,
		Status:  CONST_DTCP_Status_Created,
	}
}

func NewContentPost_Iamge() *ContentPost {
	return &ContentPost{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Object,
		Tag:     CONST_DTCP_Tag_Image,
		Status:  CONST_DTCP_Status_Created,
	}
}

type ContentPut struct {
	Version     string   `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype       string   `json:"type"`                // Fixed to "object".
	Tag         string   `json:"tag"`                 // Content type. Currently "article", "image" are supported.
	ParentDna   string   `json:"parent_dna"`          // Latest DNA of the content.
	Status      string   `json:"status"`              // Fixed to "updated".
	Updated     int      `json:"updated"`             // Content updating time. Unix timestamp.
	Title       string   `json:"title"`               // Content title.
	Abstract    string   `json:"abstract"`            // Content abstract.
	Category    string   `json:"category"`            // Content categories. Comma separated words list.
	Content     string   `json:"content"`             // Raw content in base64 encoded format.
	ContentHash string   `json:"content_hash"`        // Lowercase hex string of the SHA256 hash of the raw content.
	License     *License `json:"license"`             // Content authorization license. "none" is used if empty.
	Signature   string   `json:"signature,omitempty"` // Metadata signature.
}

func NewContentPut_Aritcle() *ContentPut {
	return &ContentPut{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Object,
		Tag:     CONST_DTCP_Tag_Article,
		Status:  CONST_DTCP_Status_Updated,
	}
}

func NewContentPut_Iamge() *ContentPut {
	return &ContentPut{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Object,
		Tag:     CONST_DTCP_Tag_Image,
		Status:  CONST_DTCP_Status_Updated,
	}
}
