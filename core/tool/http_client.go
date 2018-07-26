package tool

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
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
