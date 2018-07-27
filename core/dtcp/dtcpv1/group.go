package dtcpv1

type GroupPost struct {
	Version   string            `json:"version"`             // DTCP version. Fixed to "1.0"
	Atype     string            `json:"type"`                // Fixed to "object".
	Tag       string            `json:"tag"`                 // Fixed to "group".
	Title     string            `json:"title"`               // Group title.
	Creator   *GroupPostCreator `json:"creator"`             // Creator.
	Avatar    string            `json:"avatar"`              // An image id used for avatar.
	Abstract  string            `json:"abstract"`            // Group introduction.
	Language  string            `json:"language"`            // Group language. RFC4646 defined locales such as "en-US"
	Category  string            `json:"category"`            // Group categories. Comma separated words list.
	Created   int               `json:"created"`             // Group creation time. Unix timestamp.
	Extra     *GroupPostExtra   `json:"extra"`               // Extra metadata.
	Status    string            `json:"status"`              // Fixed to "created".
	Signature string            `json:"signature,omitempty"` // Metadata signature.
}

type GroupPostCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name. For fast creation of new sub accounts.
}

type GroupPostExtra struct {
	AllowJoin          string   `json:"allow_join"`                     // Joining group control. "all" or "application".
	AllowPost          string   `json:"allow_post"`                     // Posting control. "all", "none", "application".
	AllowPostWhitelist []string `json:"allow_post_whitelist,omitempty"` //An array containing account_ids that can always post in the group.
}

func NewGroupPost() *GroupPost {
	return &GroupPost{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Object,
		Tag:     CONST_DTCP_Tag_Group,
		Status:  CONST_DTCP_Status_Created,
	}
}

type GroupPut struct {
	Version   string         `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string         `json:"type"`                // Fixed to "object".
	Tag       string         `json:"tag"`                 // Fixed to "group".
	ParentDna string         `json:"parent_dna"`          // The latest DNA of the group.
	Status    string         `json:"status"`              // Fixed to "updated".
	Updated   int            `json:"updated"`             // Group update time. Unix timestamp.
	Title     string         `json:"title,omitempty"`     // Group title.
	Avatar    string         `json:"avatar,omitempty"`    // An image id used for avatar.
	Abstract  string         `json:"abstract,omitempty"`  // Group introduction.
	Language  string         `json:"language,omitempty"`  // Group language. RFC4646 defined locales such as "en-US"
	Category  string         `json:"category,omitempty"`  // Group categories. Comma separated words list.
	Extra     *GroupPutExtra `json:"extra,omitempty"`     // Extra metadata.
	Signature string         `json:"signature,omitempty"` // Metadata signature.
}

type GroupPutExtra struct {
	AllowJoin string `json:"allow_join"` // Joining group control. "all" or "application".
	AllowPost string `json:"allow_post"` // Posting control. "all", "none", "application".
}

func NewGroupPut() *GroupPut {
	return &GroupPut{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Object,
		Tag:     CONST_DTCP_Tag_Group,
		Status:  CONST_DTCP_Status_Updated,
	}
}
