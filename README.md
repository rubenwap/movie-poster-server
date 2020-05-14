# Movie Posters

Simple server powered by the OMDB API. It returns the poster for any movie that you pass as query parameter. It also holds the rest of movie information, but for the moment it's not displayed (easy enough to add it)

Usage: 

Get it started compiling with `go build` and running the resulting executable. Then simply go to:

    localhost:8080/movie/{title of the movie}

and you will get the title, year of release and poster. 

## Try it out in Heroku! 

https://movieposters.herokuapp.com/movie/Akira