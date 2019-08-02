package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/simisimis/genrefinder/apicalls"
	"github.com/simisimis/genrefinder/auth"
)

type genreDesc struct {
	repeats int
	artists map[string]bool
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
	resultGenreMap, err := apicalls.GetGenreMap(token, artistList)
	if err != nil {
		panic(err)
	}

	genresMap := make(map[string]*genreDesc)

	for name, genres := range resultGenreMap {
		for _, genre := range genres {
			if _, ok := genresMap[genre]; ok {
				genresMap[genre].repeats++
				genresMap[genre].artists[name] = true
			} else {
				genresMap[genre] = &genreDesc{1, map[string]bool{name: true}}
			}

		}
	}
	// for genre, desc := range genresMap {
	// 	var artistList []string
	// 	for artist := range desc.artists {
	// 		artistList = append(artistList, artist)
	// 	}
	// 	fmt.Printf("genre: %s, times repeats:%v, artists playing: %+q \n", genre, desc.repeats, artistList)
	// }
	printRes := make(map[string]string)
	for genre, desc := range genresMap {
		var artistList []string
		for artist := range desc.artists {
			artistList = append(artistList, artist)
		}
		printRes[genre] = fmt.Sprintf("{ repeats:%d, artists:[%s]", desc.repeats, strings.Join(artistList, ", "))
	}
	genresJSON, err := json.Marshal(printRes)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("genres_out.json", genresJSON, 0644)
	if err != nil {
		panic(err)
	}
}
