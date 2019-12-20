//Package elastic is for performing elasticsearch related operations
package elastic

import (
	"github.com/simisimis/genrefinder/spotify"
	"github.com/sirupsen/logrus"

	// ES related
	"context" // get client attributes
	"time"    // set timeout for connection

	// Import the Olivere Golang driver for Elasticsearch 7
	"github.com/olivere/elastic/v7"
)

var log = logrus.WithField("pkg", "elastic")

type artistPost struct {
	ID        string   `json:"id"`
	Genres    []string `json:"genres"`
	Artist    string   `json:"artist"`
	Playlists []string `json:"playlists"`
}

// PostBulkData sends provided data struct to ES
func PostBulkData(genreArtistData map[string]spotify.Artist) {

	// Print ES version
	log.Info("ES client version:", elastic.Version)

	// Create context for API calls
	ctx := context.Background()

	// Initiate ES client instance
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(5*time.Second),
	)
	if err != nil {
		log.Fatal("Set ES client failed, quitting. \n", err)
	} else {
		// Print client information
		log.Println("client:", client)
	}

	//hours, minutes, seconds := time.Now().Clock()
	//indexName := fmt.Sprintf("artists-%02d%02d%02d", hours, minutes, seconds)

	// Set index to post to name
	indexName := "artists-18122019"
	// Declare a string slice with the index name in it
	indices := []string{indexName}

	// Check index
	// Create an instance
	indexExists := elastic.NewIndicesExistsService(client)
	// Set index name to check
	indexExists.Index(indices)
	// Check if index exists
	exist, err := indexExists.Do(ctx)
	if err != nil {
		log.Fatal("Problems querying index:", err)

	} else if exist == false {
		log.Errorf("Index %s does not exist on ES", indexName)
	} else if exist == true {
		log.Printf("Index: '%s' exists.", indexName)
		// Create a new Bulk() object
		bulk := client.Bulk()

		// Put genre artist data into bulk instance
		for id, doc := range genreArtistData {
			playlists := []string{}
			for plist := range doc.Playlist {
				playlists = append(playlists, plist)
			}
			var artistDoc artistPost
			artistDoc.ID = doc.ID
			artistDoc.Genres = doc.Genres
			artistDoc.Artist = doc.Name
			artistDoc.Playlists = playlists
			// Declare a new NewBulkIndexRequest() instance
			req := elastic.NewBulkIndexRequest()

			// Set index name for bulk instance
			req.OpType("index")
			req.Index(indexName)
			req.Id(id)
			req.Doc(artistDoc)

			// Add req entry to bulk
			bulk = bulk.Add(req)
		}
		log.Printf("%d entries are going to be bulk posted", bulk.NumberOfActions())

		// Post bulk to Elasticsearch
		bulkResponse, err := bulk.Do(ctx)

		if err != nil {
			log.Fatal("Failed posting to ES:", err)
		} else {
			// Print response
			response := bulkResponse.Indexed()
			for num, document := range response {
				log.Printf("%03d item: %v", num, document)
			}
		}
	}
}
