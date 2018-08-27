/*
 * Copyright 2018 Primas Lab Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dtcpv1

type ContentReportPost struct {
	Version   string                    `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype     string                    `json:"type"`                // Fixed to "relation".
	Tag       string                    `json:"tag"`                 // Fixed to "share_report".
	SrcId     string                    `json:"src_id"`              // Account id.
	DestId    string                    `json:"dest_id"`             // Share id.
	Creator   *ContentReportPostCreator `json:"creator"`             // Creator.
	Created   int                       `json:"created"`             // Report created time. Unix timestamp.
	Status    string                    `json:"status"`              // Fixed to "created".
	Extra     *ContentReportPostExtra   `json:"extra"`               // Extra metadata.
	Signature string                    `json:"signature,omitempty"` // Metadata signature.
}

type ContentReportPostCreator struct {
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

// Get
type ContentReportGet struct {
	Version       string                   `json:"version"`             // DTCP version. Fixed to "1.0".
	Atype         string                   `json:"type"`                // Fixed to "relation".
	Tag           string                   `json:"tag"`                 // Fixed to "share_report".
	SrcId         string                   `json:"src_id"`              // Account id.
	DestId        string                   `json:"dest_id"`             // Share id.
	Creator       *ContentReportGetCreator `json:"creator"`             // Creator.
	Created       int                      `json:"created"`             // Report created time. Unix timestamp.
	Updated       int                      `json:"updated"`             //Report updated time. Unix timestamp.
	Extra         *ContentReportGetExtra   `json:"extra"`               // Extra metadata.
	Signature     string                   `json:"signature,omitempty"` // Metadata signature.
	TransactionId string                   `json:"transaction_id"`      // Latest transaction id.
}

type ContentReportGetCreator struct {
	AccountId      string `json:"account_id"`                 // Account id. Root account id in the case of Sub account posting.
	AccountName    string `json:"account_name"`               // Account name.
	SubAccountId   string `json:"sub_account_id,omitempty"`   // Sub account id. Refer to Sub account for details.
	SubAccountName string `json:"sub_account_name,omitempty"` // Sub account name. For fast creation of new sub accounts.
}

type ContentReportGetExtra struct {
	Content      string `json:"content"`       // base64 encoded report content.
	ContentHash  string `json:"content_hash"`  // Lowercase hex string of the SHA256 hash of the raw content.
	ReportType   string `json:"report_type"`   // Report type.
	ReportStatus string `json:"report_status"` // Fixed to "pending".
}
