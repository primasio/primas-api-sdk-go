package dtcpv1

import "github.com/shopspring/decimal"

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

func NewContentPost_Image() *ContentPost {
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

// Get
type ContentGet struct {
	Id            string             `json:"id"`                // Content id.
	Tag           string             `json:"tag"`               // Content tag. Currently "article", "image" are supported.
	Title         string             `json:"title"`             // Content title.
	Creator       *ContentGetCreator `json:"creator"`           // Creator of the content.
	Abstract      string             `json:"abstract"`          // Content abstract.
	Language      string             `json:"language"`          // Content language. RFC4646 defined locales such as "en-US"
	Category      string             `json:"category"`          // Content categories. Comma separated words list.
	Created       int                `json:"created"`           // Content creation time. Unix timestamp.
	Updated       int                `json:"updated"`           // Content last updating time. Unix timestamp.
	Content       string             `json:"content"`           // Content URI. In the case of IPFS, a link starts with "ipfs://"
	ContentHash   string             `json:"content_hash"`      // Lowercase hex string of the SHA256 hash of the raw content.
	License       *License           `json:"license,omitempty"` // Content authorization license.
	Signature     string             `json:"signature"`         // Metadata signature.
	Dna           string             `json:"dna"`               // Content DNA.
	Extra         *ContentGetExtra   `json:"extra"`             // Extra content metadata.
	TransactionId string             `json:"transaction_id"`    // Transaction id.
}

type ContentGetCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	AccountName    string `json:"account_name"`               // Account name. Root account name in the case of Sub account posting.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name.
}

type ContentGetExtra struct {
	Ext        string                   `json:"ext,omitempty"`     // Image format, such as 'png', 'jpg'. Image only.
	Width      int                      `json:"width,omitempty"`   // Image width in pixels. Image only.
	Height     int                      `json:"height,omitempty"`  // Image height in pixels. Image only.
	Size       int                      `json:"size,omitempty"`    // Image size in bytes. Image only.
	PstTotal   decimal.Decimal          `json:"pst_total"`         // Total PST earned.
	PstUpdated int                      `json:"pst_updated"`       // Last PST updated time. Unix timestamp.
	Objects    []ContentGetExtraObjects `json:"objects,omitempty"` // A list of images, videos, audios contained in the content. Article only.
}

type ContentGetExtraObjects struct {
	Id    string                 `json:"id"`    // Object id.
	Atype string                 `json:"type"`  // "image", "audio" or "video".
	Extra *ContentGetExtraObject `json:"extra"` // Extra metadata.
}

type ContentGetExtraObject struct {
	Ext    string `json:"ext,omitempty"`    // Image format, such as 'png', 'jpg'. Image only.
	Width  int    `json:"width,omitempty"`  // Image width in pixels. Image only.
	Height int    `json:"height,omitempty"` // Image height in pixels. Image only.
	Size   int    `json:"size,omitempty"`   // Image size in bytes. Image only.
}
