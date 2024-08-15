package mande

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AkifhanIlgaz/credible-mandela-api/utils/constants"
)

const (
	TrustDropsUrl string = "https://app.mande.network/subgraphs/name/TrustDrops"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() Client {
	return Client{
		httpClient: &http.Client{
			Timeout: constants.Timeout,
		},
	}
}

func (mc Client) GetCredScoreOfUser(address string) (float64, error) {
	credQueryRequest := generateCredQueryRequest(address)
	reqBody, err := json.Marshal(&credQueryRequest)
	if err != nil {
		return 0, fmt.Errorf("get cred score of user: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, TrustDropsUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return 0, fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := mc.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var credQueryResponse credQueryResponse
	if err := json.NewDecoder(resp.Body).Decode(&credQueryResponse); err != nil {
		return 0, fmt.Errorf("could not decode response body: %w", err)
	}

	credString := credQueryResponse.Data.User.CredScoreAccrued
	if len(credString) == 0 {
		return 0, fmt.Errorf("user with address %s not found", address)
	}

	credScore, err := strconv.ParseFloat(credString, 32)
	if err != nil {
		return 0, fmt.Errorf("failed to convert CRED score of %v", address)
	}

	return credScore / 100.0, nil
}
