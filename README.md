## Spotify playlist genre finder
A placeholder repo for script like app that goes through selected playlist and extracts genres artists are playing in.

This program is looking up SPOTIFY_CLIENT(Client ID) and SPOTIFY_SECRET(Client Secret) environment variables to get a token.
You can create an app in spotify developers dashboard.

Currently genrefinder:
* retrieves token
* makes a call to a hardcoded playlist to get a list of artists.
* retrieves an array of genres per artist.

TODO:
* input/select playlist
* count them
* export them
