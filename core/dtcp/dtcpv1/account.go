package dtcpv1

type AccountPost struct {
	Version   string              `json:"version"` // DTCP version. Fixed to "1.0".
	Atype     string              `json:"type"`    // Fixed to "object".
	Tag       string              `json:"tag"`     // Fixed to "account".
	Name      string              `json:"name"`    // Name.
	Address   string              `json:"address"`
	Abstract  string              `json:"abstract,omitempty"`  // Description.
	Avatar    string              `json:"avatar,omitempty"`    // An image id used for avatar.
	Creator   *AccountPostCreator `json:"creator,omitempty"`   // Creator. Required when creating sub account.
	Created   int                 `json:"created"`             // Account creation time. Unix timestamp.
	Extra     *AccountPostExtra   `json:"extra,omitempty"`     // Extra metadata.
	Status    string              `json:"status"`              // Fixed to "created".
	Signature string              `json:"signature,omitempty"` // Metadata signature.
}

type AccountPostCreator struct {
	AccountId    string `json:"account_id"`     // Root account id.
	SubAccountId string `json:"sub_account_id"` // Sub account id. This id is provided by the third-party application. Usually the id in the application system is used directly.
}

type AccountPostExtra struct {
	Hash string `json:"hash"`
}

func NewAccountPost() *AccountPost {
	return &AccountPost{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Object,
		Tag:     CONST_DTCP_Tag_Account,
		Status:  CONST_DTCP_Status_Created,
	}
}

type AccountPut struct {
	Version   string             `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string             `json:"type"`                // Fixed to "object".
	Tag       string             `json:"tag"`                 // Fixed to "account".
	ParentDna string             `json:"parent_dna"`          // The latest DNA of the account
	Updated   int                `json:"creaupdatedted"`      // Updating time. Unix timestamp.
	Name      string             `json:"name"`                // Name.
	Abstract  string             `json:"abstract,omitempty"`  // Description.
	Avatar    string             `json:"avatar,omitempty"`    // An image id used for avatar.
	Creator   *AccountPutCreator `json:"creator,omitempty"`   // Creator. Required when creating sub account.
	Extra     *AccountPutExtra   `json:"extra,omitempty"`     // Extra metadata.
	Status    string             `json:"status"`              // Fixed to "updated".
	Signature string             `json:"signature,omitempty"` // Metadata signature.
}

type AccountPutCreator struct {
	AccountId    string `json:"account_id"`     // Root account id.
	SubAccountId string `json:"sub_account_id"` // Sub account id. This id is provided by the third-party application. Usually the id in the application system is used directly.
}

type AccountPutExtra struct {
	Hash string `json:"hash"`
}

func NewAccountPut() *AccountPut {
	return &AccountPut{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Object,
		Tag:     CONST_DTCP_Tag_Account,
		Status:  CONST_DTCP_Status_Updated,
	}
}

// get
type AccountGet struct {
	Id            string             `json:"id"`                 // Account id.
	Address       string             `json:"address"`            // Account address.
	Title         string             `json:"title"`              // Account name.
	Abstract      string             `json:"abstract,omitempty"` // Description.
	Avatar        string             `json:"avatar,omitempty"`   // An image id used for avatar.
	Creator       *AccountGetCreator `json:"creator"`            // Creator of the sub account.
	Created       int                `json:"created"`            // Account creation time. Unix timestamp.
	Updated       int                `json:"updated"`            // Account last updating time. Unix timestamp.
	Extra         *AccountGetExtra   `json:"extra,omitempty"`    // Extra metadata.
	Signature     string             `json:"signature"`          // Metadata signature.
	Dna           string             `json:"dna"`                // DNA of the account.
	Credits       int                `json:"credits"`            // Current credits.
	TransactionId string             `json:"transaction_id"`     // Latest transaction id.
}

type AccountGetCreator struct {
	AccountId   string `json:"account_id"`   // Root account id.
	AccountName string `json:"account_name"` // Root account name.
}

type AccountGetExtra struct {
	Hash string `json:"hash"` // In the case of proof of existence of secret data. The hash can be filled in this field.
}
