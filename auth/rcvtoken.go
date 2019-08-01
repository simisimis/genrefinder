// Package auth handles spotify token retrieval
package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

// GetToken posts to spotify to retrieve token
func GetToken() (string, error) {
	ClientID := os.Getenv("SPOTIFY_CLIENT")
	ClientSecret := os.Getenv("SPOTIFY_SECRET")

	// retrieve base64 encoded app secrets
	secret64 := fmt.Sprintf("Basic %s", getEncodedKeys(ClientID, ClientSecret))
	// prepare to POST
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	data := url.Values{}
	data.Add("grant_type", "client_credentials")

	// Create a POST request to retrieve token
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", secret64)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	rcvToken, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()

	responseData := struct {
		Token string `json:"access_token"`
	}{}
	if err := json.Unmarshal(rcvToken, &responseData); err != nil {
		return "", err
	}
	return responseData.Token, nil
}

func getEncodedKeys(id, secret string) string {
	data := fmt.Sprintf("%v:%v", id, secret)
	return base64.StdEncoding.EncodeToString([]byte(data))
}
