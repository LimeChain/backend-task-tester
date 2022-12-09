package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Transaction struct {
	TransactionHash   string `json:"transactionHash"`
	TransactionStatus int    `json:"transactionStatus"`
	BlockHash         string `json:"blockHash"`
	BlockNumber       int    `json:"blockNumber"`
	From              string `json:"from"`
	To                string `json:"to"`
	ContractAddress   string `json:"contractAddress"`
	LogsCount         int    `json:"logsCount"`
	Input             string `json:"input"`
	Value             string `json:"value"`
}

type TransactionResponse struct {
	Transactions []Transaction
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}

type LimeClient struct {
	endpointURL string
}

func (c *LimeClient) GetAll() (*TransactionResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/lime/all", c.endpointURL))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var txResponse TransactionResponse

	if err := json.NewDecoder(resp.Body).Decode(&txResponse); err != nil {
		return nil, err
	}

	return &txResponse, nil
}

func (c *LimeClient) GetEth(rlpString, authToken string) (*TransactionResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/lime/eth/%s", c.endpointURL, rlpString), nil)
	if err != nil {
		return nil, err
	}

	if authToken != "" {
		req.Header.Set("AUTH_TOKEN", authToken)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var txResponse TransactionResponse

	if err := json.NewDecoder(resp.Body).Decode(&txResponse); err != nil {
		return nil, err
	}

	return &txResponse, nil
}

func (c *LimeClient) GetMy(authToken string) (*TransactionResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/lime/my", c.endpointURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("AUTH_TOKEN", authToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var txResponse TransactionResponse

	if err := json.NewDecoder(resp.Body).Decode(&txResponse); err != nil {
		return nil, err
	}

	return &txResponse, nil
}

func (c *LimeClient) PostAuthenticate(username, password string) (*AuthenticateResponse, error) {
	postBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/lime/authenticate", c.endpointURL), "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var authResp AuthenticateResponse

	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, err
	}

	return &authResp, nil
}

func NewLimeAPIClient(endpoint string) *LimeClient {
	return &LimeClient{
		endpointURL: endpoint,
	}
}
