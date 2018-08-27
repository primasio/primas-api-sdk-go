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

package content

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type PostContentResult struct {
	Id  string `json:"id"`  // The id of the content.
	Dna string `json:"dna"` // The DNA of the content.
}

type PostContentResponse struct {
	core.Response
	Data *PostContentResult `json:"data"`
}

func PostContent_SignatureStr(tag, title, account_id, sub_account_id, sub_account_name, abstract, language, category string,
	created int, content string, license *dtcpv1.License) (string, *dtcpv1.ContentPost, error) {
	var newCreator dtcpv1.ContentPostCreator
	newCreator.AccountId = account_id
	newCreator.SubAccountId = sub_account_id
	newCreator.SubAccountName = sub_account_name

	var newPostContent *dtcpv1.ContentPost
	if tag == dtcpv1.CONST_DTCP_Tag_Article {
		newPostContent = dtcpv1.NewContentPost_Aritcle()
	}
	if tag == dtcpv1.CONST_DTCP_Tag_Image {
		newPostContent = dtcpv1.NewContentPost_Image()
	}
	newPostContent.Title = title
	newPostContent.Creator = &newCreator
	newPostContent.Abstract = abstract
	newPostContent.Language = language
	newPostContent.Category = category
	newPostContent.Created = created
	newPostContent.Content = content
	newPostContent.ContentHash = dtcpv1.HashValue([]byte(content))
	newPostContent.License = license

	err := postContent_check(newPostContent)
	if err != nil {
		return "", nil, err
	}

	sigSoure, err := tool.StructToSignatureByExclude(newPostContent, "content")
	if err != nil {
		return "", nil, err
	}

	return sigSoure, newPostContent, nil
}

func PostContent_Aritcle(signature string, preObj *dtcpv1.ContentPost) (*PostContentResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	err := postContent_check(preObj)
	if err != nil {
		return nil, err
	}

	// use json
	preObj.Content = base64.StdEncoding.EncodeToString([]byte(preObj.Content))
	preObj.Signature = signature

	url := config.Gogal_Server + `/content`

	requestBody, err := json.Marshal(preObj)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Post(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	var responseObj PostContentResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func PostContent_ImageUrlencoded(signature string, preObj *dtcpv1.ContentPost) (*PostContentResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	preObj.Signature = signature

	err := postContent_check(preObj)
	if err != nil {
		return nil, err
	}

	url := config.Gogal_Server + `/content`

	datas := make(map[string]string, 0)
	datas["version"] = preObj.Version
	datas["type"] = preObj.Atype
	datas["tag"] = preObj.Tag
	datas["title"] = preObj.Title
	creatorArr, err := json.Marshal(preObj.Creator)
	if err != nil {
		return nil, err
	}
	datas["creator"] = string(creatorArr)
	datas["abstract"] = preObj.Abstract
	datas["language"] = preObj.Language
	datas["category"] = preObj.Category
	datas["created"] = strconv.Itoa(preObj.Created)
	datas["content"] = base64.StdEncoding.EncodeToString([]byte(preObj.Content))
	datas["content_hash"] = preObj.ContentHash
	if preObj.License != nil {
		licenseArr, err := json.Marshal(preObj.License)
		if err != nil {
			return nil, err
		}

		datas["license"] = string(licenseArr)
	}
	datas["status"] = preObj.Status
	datas["signature"] = signature

	response, err := tool.Http_PostFormUrlencoded(url, datas)
	if err != nil {
		return nil, err
	}

	var responseObj PostContentResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func PostContent_ImageMultipartForm(signature string, preObj *dtcpv1.ContentPost, imgPath string) (*PostContentResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	preObj.Signature = signature

	err := postContent_check(preObj)
	if err != nil {
		return nil, err
	}

	url := config.Gogal_Server + `/content`

	datas := make(map[string]string, 0)
	datas["version"] = preObj.Version
	datas["type"] = preObj.Atype
	datas["tag"] = preObj.Tag
	datas["title"] = preObj.Title
	creatorArr, err := json.Marshal(preObj.Creator)
	if err != nil {
		return nil, err
	}
	datas["creator"] = string(creatorArr)
	datas["abstract"] = preObj.Abstract
	datas["language"] = preObj.Language
	datas["category"] = preObj.Category
	datas["created"] = strconv.Itoa(preObj.Created)
	datas["content"] = base64.URLEncoding.EncodeToString([]byte(preObj.Content))
	datas["content_hash"] = preObj.ContentHash
	if preObj.License != nil {
		licenseArr, err := json.Marshal(preObj.License)
		if err != nil {
			return nil, err
		}

		datas["license"] = string(licenseArr)
	}
	datas["status"] = preObj.Status
	datas["signature"] = signature

	response, err := tool.Http_PostFormMultipartForm(url, datas, imgPath)
	if err != nil {
		return nil, err
	}

	var responseObj PostContentResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func postContent_check(preObj *dtcpv1.ContentPost) error {
	if preObj.Version != dtcpv1.CONST_DTCP_Version_v1 {
		return errors.New("parameter version error")
	}

	if preObj.Atype != dtcpv1.CONST_DTCP_Type_Object {
		return errors.New("parameter type error")
	}

	if preObj.Tag != dtcpv1.CONST_DTCP_Tag_Article && preObj.Tag != dtcpv1.CONST_DTCP_Tag_Image {
		return errors.New("parameter tag error")
	}

	if preObj.Title == "" {
		return errors.New("parameter title error")
	}

	if preObj.Creator == nil {
		return errors.New("parameter creator error")
	}

	if preObj.Language == "" {
		return errors.New("parameter language error")
	}

	if preObj.Category == "" {
		return errors.New("parameter category error")
	}

	if preObj.Created <= 0 {
		return errors.New("parameter created error")
	}

	if preObj.Content == "" {
		return errors.New("parameter content error")
	}

	if preObj.Status != dtcpv1.CONST_DTCP_Status_Created {
		return errors.New("parameter status error")
	}

	return nil
}
