package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/yokyann/PC3R_tetrio/server/models"
)

var uri = "https://ch.tetr.io/api/users"

func FetchPlayerUsername(username string) (string, error) {
	apiUrl := uri + "/" + username

	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var playerInfoResp PlayerInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&playerInfoResp); err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %w", err)
	}

	if !playerInfoResp.Success {
		return "", fmt.Errorf("API request unsuccessful: %s", playerInfoResp.Error)
	}

	return playerInfoResp.Data.Player.Username, nil
}
