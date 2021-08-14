# covid19-india
Api Server to get Covid cases in a user's State & in India on the basis of their geo-location.

## Running the Server
- Install the Go dependencies
- Expose environment variables

Variable | Description
---| --- |
PORT | Port of server
MONGODB_URI | MongoDb connection URI
HERE_MAPS_API_KEY | HERE Map's reverse geo coding API Key ([here](https://developer.here.com/documentation/geocoder/dev_guide/topics/resource-reverse-geocode.html))
- Run the file [main.go](https://github.com/Arpit007/covid19-india/blob/master/cmd/covid19-india/main.go)
