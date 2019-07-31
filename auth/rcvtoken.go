// Package auth handles spotify token retrieval
package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// AppSecrets holds login details
type AppSecrets struct {
	ClientID     string
	ClientSecret string
	AccessToken  string `json:"access_token"`
}

// GetToken posts to spotify to retrieve token
func (spotify *AppSecrets) GetToken() string {
	spotify.ClientID = os.Getenv("SPOTIFY_CLIENT")
	spotify.ClientSecret = os.Getenv("SPOTIFY_SECRET")

	// retrieve base64 encoded app secrets
	secret64 := fmt.Sprintf("Basic %s", spotify.getEncodedKeys())
	fmt.Println(secret64)
	// prepare to POST
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	data := url.Values{}
	data.Add("grant_type", "client_credentials")

	// Create a POST request to retrieve token
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Authorization", secret64)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	rcvToken, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	fmt.Println(string(rcvToken))

	if err == nil {
		json.Unmarshal(rcvToken, &spotify)
	}
	fmt.Println(spotify.AccessToken)
	return spotify.AccessToken
}

func (spotify *AppSecrets) getEncodedKeys() string {
	data := fmt.Sprintf("%v:%v", spotify.ClientID, spotify.ClientSecret)
	return base64.StdEncoding.EncodeToString([]byte(data))
}
