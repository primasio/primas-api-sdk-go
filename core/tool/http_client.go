package tool

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Http_Get(url string, queryParams map[string]interface{}) ([]byte, error) {
	urlQuery := ``
	for itemName, itemValue := range queryParams {
		urlQuery = urlQuery + `&` + itemName + `=` + ConvertToString(itemValue)
	}
	if urlQuery != `` {
		urlQuery = strings.Trim(urlQuery, "&")
		url = url + `?` + urlQuery
	}

	response, err := http.Get(url)
	if err != nil {
		locerr := errors.New("Maybe server not find, error:" + err.Error())
		log.Printf(locerr.Error())

		return nil, locerr
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("%s", err)
			return nil, err
		}

		if response.StatusCode != 200 {
			return nil, errors.New("response StatusCode error:%v" + strconv.Itoa(response.StatusCode))
		}

		if len(contents) == 0 {
			return nil, errors.New("response body is empty")
		}

		return contents, nil
	}

	return nil, nil
}

func Http_Get_Direct(url string, queryParams map[string]interface{}) (*http.Response, error) {
	urlQuery := ``
	for itemName, itemValue := range queryParams {
		urlQuery = urlQuery + `&` + itemName + `=` + ConvertToString(itemValue)
	}
	if urlQuery != `` {
		urlQuery = strings.Trim(urlQuery, "&")
		url = url + `?` + urlQuery
	}

	response, err := http.Get(url)
	if err != nil {
		locerr := errors.New("Maybe server not find, error:" + err.Error())
		log.Printf(locerr.Error())

		return nil, locerr
	}

	return response, nil
}

func Http_Post(url string, jsonData string) ([]byte, error) {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}

func Http_PostFormUrlencoded(urlAddr string, params map[string]string) ([]byte, error) {
	formDatas := url.Values{}
	for key, val := range params {
		tempArr := make([]string, 0)
		tempArr = append(tempArr, val)
		formDatas[key] = tempArr
	}

	log.Println("formDatas:", formDatas)

	// application/x-www-form-urlencoded
	resp, err := http.PostForm(urlAddr, formDatas)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	defer resp.Body.Close()

	resultBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return resultBody, nil
}

func Http_PostFormMultipartForm(urlAddr string, params map[string]string, imgPath string) ([]byte, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("content", filepath.Base(imgPath))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	//log.Println("body:", body)

	req, err := http.NewRequest("POST", urlAddr, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	defer resp.Body.Close()

	resultBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return resultBody, nil
}

func Http_Put(url string, jsonData string) ([]byte, error) {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}

func Http_Delete(url string, jsonData string) ([]byte, error) {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}
