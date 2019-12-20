// Package spotify is for storing all the different calls to spotify apis
package spotify

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var log = logrus.WithField("pkg", "genres")

// genresData is a struct to receive json data about artist genres
type genresData struct {
	Name   string   `json:"name"`
	Genres []string `json:"genres"`
}

// GetGenreMap retrieves an array of genres per artist
func GetGenreMap(token string, artists map[string]Artist) (map[string]Artist, error) {
	tokenHdr := "Bearer " + token

	for id, artist := range artists {
		genreList, err := getGenres(tokenHdr, artist.Href)
		//		time.Sleep(5 * time.Millisecond)
		if err != nil {
			return artists, err
		}
		artist.Genres = genreList
		artists[id] = artist
	}
	return artists, nil
}

// getGenres retrieves a list of genres for given artist
func getGenres(tokenHdr, url string) ([]string, error) {

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	contentTypeHdr := fmt.Sprint("application/json")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []string{}, err
	}
	req.Header.Add("Accept", contentTypeHdr)
	req.Header.Add("Content-Type", contentTypeHdr)
	req.Header.Add("Authorization", tokenHdr)

	resp, err := client.Do(req)
	// Check is rate-limit has been triggered and sleep for time required
	if resp.StatusCode == 429 {
		secondsWait, _ := strconv.Atoi(resp.Header["Retry-After"][0])
		log.Printf("Sleeping %d seconds because of rate limit", secondsWait)
		time.Sleep(time.Duration(secondsWait) * time.Second)
	}
	// Add little pause between calls
	time.Sleep(5 * time.Millisecond)

	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []string{}, err
	}
	genres := genresData{}

	if err := json.Unmarshal(body, &genres); err != nil {
		return []string{}, err
	}
	return genres.Genres, nil

}
