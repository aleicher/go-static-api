# Go Webserver to serve static JSON files

Start with docker-compose

`docker-compose up`

then test the example api:

`curl localhost:3000/todos | jq '.'`
