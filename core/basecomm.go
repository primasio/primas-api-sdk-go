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

package core

const (
	CONST_Version_Error   = "error"
	CONST_Version_21017   = "v3.00.01"
	CONST_Version_Current = CONST_Version_21017

	CONST_Pagesize_default = 20

	CONST_Language_none = "none"
	CONST_Language_zh   = "zh"
	CONST_Language_en   = "en"
	CONST_Language_ja   = "ja"

	CONST_ResultCode_Success               = 0
	CONST_ResultCode_Unknown_Error         = 499
	CONST_ResultCode_ClientError           = 400
	CONST_ResultCode_InvalidData           = 401
	CONST_ResultCode_ParseJSON_Error       = 402
	CONST_ResultCode_ClientSignature_Error = 403
	CONST_ResultCode_InputParamter_Error   = 404
	CONST_ResultCode_InputParamter_Empty   = 405
	CONST_ResultCode_NotFindRecord         = 406
	CONST_ResultCode_Server_Error          = 500
)

type Response struct {
	ResultCode int    `json:"result_code"`
	ResultMsg  string `json:"result_msg"`
}
