package dtcpv1

type ContentReportPost struct {
	Version   string                  `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string                  `json:"type"`                // Fixed to "relation".
	Tag       string                  `json:"tag"`                 // Fixed to "share_report".
	SrcId     string                  `json:"src_id"`              // Account id.
	DestId    string                  `json:"dest_id"`             // Share id.
	Creator   *ContentReportCreator   `json:"creator"`             // Creator.
	Created   int                     `json:"created"`             // Report created time. Unix timestamp.
	Status    string                  `json:"status"`              // Fixed to "created".
	Extra     *ContentReportPostExtra `json:"extra"`               // Extra metadata.
	Signature string                  `json:"signature,omitempty"` // Metadata signature.
}

type ContentReportCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name. For fast creation of new sub accounts.
}

type ContentReportPostExtra struct {
	Content      string `json:"content"`       // base64 encoded report content.
	ContentHash  string `json:"content_hash"`  // Lowercase hex string of the SHA256 hash of the raw content.
	ReportType   string `json:"report_type"`   // Report type.
	ReportStatus string `json:"report_status"` // Fixed to "pending".
}

func NewContentReportPost() *ContentReportPost {
	return &ContentReportPost{
		Version: CONST_DTCP_Version_v1,
		Atype:   CONST_DTCP_Type_Relation,
		Tag:     CONST_DTCP_Tag_Share_report,
		Status:  CONST_DTCP_Status_Created,
	}
}