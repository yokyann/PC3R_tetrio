package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserInfoResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    struct {
		User struct {
			Username string `json:"username"`
		} `json:"user"`
	} `json:"data,omitempty"`
}

func FetchUsername(username string) (string, error) {
	apiUrl := fmt.Sprintf("https://ch.tetr.io/api/users/%s", username)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var userInfoResp UserInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&userInfoResp); err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %w", err)
	}

	if !userInfoResp.Success {
		return "", fmt.Errorf("API request unsuccessful: %s", userInfoResp.Error)
	}

	return userInfoResp.Data.User.Username, nil
}
	