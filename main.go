package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/simisimis/genrefinder/apicalls"
	"github.com/simisimis/genrefinder/auth"
)

type genreDesc struct {
	repeats int
	artists []string
}

func main() {
	// retrieve token
	token, err := auth.GetToken()
	var artistList []string
	// retrieve artists from playlist songs
	artistList, err = apicalls.GetArtists(token)

	if err != nil {
		panic(err)
	}
	// retrieve genres per artist
	resultGenreData, err := apicalls.GetGenreMap(token, artistList)
	if err != nil {
		panic(err)
	}

	genresMap := make(map[string]map[string]interface{})

	for name, genres := range resultGenreData {
		for _, genre := range genres {
			if _, ok := genresMap[genre]; ok {
				genresMap[genre]["repeats"] = genresMap[genre]["repeats"].(int) + 1
				genresMap[genre]["artists"] = append(genresMap[genre]["artists"].([]string), name)
			} else {
				genresMap[genre] = map[string]interface{}{"repeats": 1, "artists": []string{name}}
			}

		}
	}

	genresJSON, err := json.Marshal(genresMap)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("genres_out.json", genresJSON, 0644)
	if err != nil {
		panic(err)
	}
}
