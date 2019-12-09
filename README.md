## Spotify playlist genre finder
A placeholder repo for script like app that goes through selected playlist and extracts genres artists are playing in.

This program is looking up SPOTIFY_CLIENT(Client ID) and SPOTIFY_SECRET(Client Secret) environment variables to get a token.
You can create an app in spotify developers dashboard.

Currently genrefinder:
* retrieves token
* uses a playlist that user chooses through select menu
* retrieves a map of artists.
* enriches artists dictionary with a name of playlist artist was found in.
* adds genres artist is playing in to artists dictionary

TODO:
* loop over all playlists of provided username
* implement bulk upload to elasticsearch
