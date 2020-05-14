#!/bin/bash
docker run -it -p 8080:8080 -e OMDB_URL=http://www.omdbapi.com -e OMDB_KEY=334b48a9 ruben/movieposter:0.1