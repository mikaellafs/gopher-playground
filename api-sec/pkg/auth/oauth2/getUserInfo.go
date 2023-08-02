package oauth2

import (
	"encoding/json"
	"fmt"
	"gopher-playground/api-sec/pkg/env"
	"net/http"
	"os"
	"time"
)

func GetUserInfo(accessToken string) (*UserInfo, error) {
	// Create an HTTP client with a timeout
	client := &http.Client{Timeout: 10 * time.Second}

	// Make a request to the Google API to get user info
	url := os.Getenv(env.OAUTH2_USER_INFO_ENDPOINT)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info, status code: %d", resp.StatusCode)
	}

	var userInfo UserInfo
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}
