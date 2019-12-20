## Spotify playlist genre finder
A placeholder repo for script like app that goes through selected playlist and extracts genres artists are playing in.

This program is looking up SPOTIFY_CLIENT(Client ID) and SPOTIFY_SECRET(Client Secret) environment variables to get a token.
You can create an app in spotify developers dashboard.

Currently genrefinder:
* retrieves token
* Scenario1(default):
  * Goes through all provided user public playlists
* Scenario2:
  * Uses a playlist that user chooses through select menu
* generates a map of artists.
* enriches artists map with names of playlists artist was found in.
* adds genres artist is playing in to artists map
* does a bulk upload to elasticsearch server running on http://localhost:9200

NOTE: elasticsearch index name and host details are hardcoded in the code

### Screenshots
Kibana
![Alt text](/screenshots/20-12-19_09_05_scrot.png?raw=true "Kibana screenshot")
term
![Alt text](/screenshots/20-12-19_16_00_scrot.png?raw=true "genrefinder in action")

```
DELETE artists-18122019

PUT artists-18122019
{
  "settings": {
    "number_of_replicas": 0
  }
}

POST _aliases
{
  "actions": [
    {
      "add": {
        "index": "artists-18122019",
        "alias": "artists"
      }
    }
  ]
}

PUT artists/_mapping
{
  "properties": {
    "artist": {
      "type": "text",
      "analyzer":"standard",
          "fielddata":true,
      "fields": {
        "raw": {
          "type": "keyword"
        }
      }
    },
    "genres": {
      "type": "text",
      "analyzer":"standard",
      "fielddata":true,
      "fields": {
        "raw": {
          "type": "keyword"
        }
      }
    },
    "playlists": {
      "type": "text",
      "analyzer":"standard",
      "fielddata":true,
      "fields": {
        "raw": {
          "type": "keyword"
        }
      }
    }
  }
}
```
