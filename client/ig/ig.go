package ig

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type IGClient struct {
	URL            string
	APIKey         string
	Session        *Session
	CST            string
	XSecurityToken string
}

func (ig *IGClient) newHeader(headerMap map[string]string) http.Header {
	header := http.Header{}
	header.Add("Content-Type", "application/json; charset=UTF-8")
	header.Add("Accept", "application/json; charset=UTF-8")
	header.Add("X-IG-API-KEY", ig.APIKey)
	if ig.Session != nil {
		header.Add("Authorization", fmt.Sprintf("%s %s", ig.Session.OauthToken.TokenType, ig.Session.OauthToken.AccessToken))
		header.Add("IG-ACCOUNT-ID", ig.Session.AccountId)
	}
	for k, v := range headerMap {
		header.Add(k, v)
	}
	return header
}

func request(method, url string, header http.Header, payload []byte) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ig *IGClient) delete(url string, headerMap map[string]string) error {
	header := ig.newHeader(headerMap)
	_, err := request("GET", url, header, nil)
	return err
}

func (ig *IGClient) get(url string, headerMap map[string]string) ([]byte, error) {
	header := ig.newHeader(headerMap)

	resp, err := request("GET", url, header, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyResp, _ := ioutil.ReadAll(resp.Body)
	return bodyResp, nil
}

func (ig *IGClient) Login(identifier, password string) error {
	url := ig.URL + "/session"
	body := map[string]string{"identifier": identifier, "password": password}
	jsonString, err := json.Marshal(body)
	if err != nil {
		return err
	}

	header := ig.newHeader(map[string]string{"VERSION": "3"})

	resp, err := request("POST", url, header, jsonString)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyResp, _ := ioutil.ReadAll(resp.Body)
	if resp.Status != "200 OK" {
		return errors.New(fmt.Sprintf("ERROR: %s: %s", resp.Status, string(bodyResp)))
	}
	var sessionRep Session
	err = json.Unmarshal(bodyResp, &sessionRep)
	if err != nil {
		return err
	}

	ig.Session = &sessionRep
	return nil
}

func (ig *IGClient) Logout() error {
	url := ig.URL + "/session"
	err := ig.delete(url, nil)
	if err != nil {
		return err
	}
	return nil
}

func (ig *IGClient) GetSession() error {
	params := url.Values{}
	params.Add("fetchSessionTokens", "true")
	query := params.Encode()
	url := ig.URL + "/session?" + query
	header := ig.newHeader(map[string]string{"VERSION": "1"})
	resp, err := request("GET", url, header, nil)
	if err != nil {
		return err
	}
	ig.XSecurityToken = resp.Header["X-Security-Token"][0]
	ig.CST = resp.Header["Cst"][0]
	for k, h := range resp.Header {
		log.Println(k, h)
	}
	return nil
}

func (ig *IGClient) GetAccounts() ([]Account, error) {
	url := ig.URL + "/accounts"
	bodyResp, err := ig.get(url, nil)
	if err != nil {
		return nil, err
	}
	var respStruct struct {
		Accounts []Account `json:"accounts"`
	}
	err = json.Unmarshal(bodyResp, &respStruct)
	if err != nil {
		return nil, err
	}

	return respStruct.Accounts, nil
}

func (ig *IGClient) GetMarketNavigation(nodeId string) ([]MarketData, []MarketNode, error) {
	url := ig.URL + "/marketnavigation/" + nodeId
	bodyResp, err := ig.get(url, nil)
	if err != nil {
		return nil, nil, err
	}
	var respStruct struct {
		Markets []MarketData `json:"markets"`
		Nodes   []MarketNode `json:"nodes"`
	}
	err = json.Unmarshal(bodyResp, &respStruct)
	if err != nil {
		return nil, nil, err
	}

	return respStruct.Markets, respStruct.Nodes, nil
}

func (ig *IGClient) GetMarketDetails(epic string) (*MarketDetail, error) {
	url := ig.URL + "/markets/" + epic
	bodyResp, err := ig.get(url, map[string]string{"VERSION": "3"})
	if err != nil {
		return nil, err
	}
	var respStruct MarketDetail
	err = json.Unmarshal(bodyResp, &respStruct)
	if err != nil {
		return nil, err
	}

	return &respStruct, nil
}

func (ig *IGClient) GetPrices(epic, resolution, from, to, max, pageSize, pageNumber string) (*PriceList, error) {

	params := url.Values{}
	if resolution != "" {
		params.Add("resolution", resolution)
	}
	if from != "" {
		params.Add("from", from)
		if to != "" {
			params.Add("to", to)
		}
	} else if max != "" {
		params.Add("max", max)
	}
	if pageSize != "" {
		params.Add("pageSize", pageSize)
	}
	if pageNumber != "" {
		params.Add("pageNumber", pageNumber)
	}
	query := params.Encode()

	url := ig.URL + "/prices/" + epic
	if query != "" {
		url += "?" + query
	}
	bodyResp, err := ig.get(url, map[string]string{"VERSION": "3"})
	if err != nil {
		return nil, err
	}
	var respStruct PriceList
	err = json.Unmarshal(bodyResp, &respStruct)
	if err != nil {
		return nil, err
	}

	return &respStruct, nil
}
