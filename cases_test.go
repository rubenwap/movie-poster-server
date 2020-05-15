package main
import "os"

var testCasesBuildParams = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "It should build correctly URL with unique word titles",
		input:       "bambi",
		expected:    "http://www.omdbapi.com?apikey="+ os.Getenv("OMDB_KEY") +"&plot=full&t=bambi",
	},
	{
		description: "It should build correctly URL with multi word titles",
		input:       "the dark knight",
		expected:    "http://www.omdbapi.com?apikey="+ os.Getenv("OMDB_KEY") +"&plot=full&t=the+dark+knight",
	},
}

var testCasesGetMovieResponse = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "It should return a correct json response from the server",
		input:       "http://www.omdbapi.com?apikey="+ os.Getenv("OMDB_KEY") +"&plot=full&t=bambi",
		expected:    `{"Title":"Bambi","Year":"1942","Rated":"G","Released":"21 Aug 1942","Runtime":"70 min","Genre":"Animation, Drama, Family","Director":"James Algar, Samuel Armstrong, David Hand, Graham Heid, Bill Roberts, Paul Satterfield, Norman Wright, Arthur Davis, Clyde Geronimi","Writer":"Felix Salten (from the story by), Perce Pearce (story direction), Larry Morey (story adaptation), Vernon Stallings (story development), Mel Shaw (story development), Carl Fallberg (story development), Chuck Couch (story development), Ralph Wright (story development)","Actors":"Hardie Albright, Stan Alexander, Bobette Audrey, Peter Behn","Plot":"The animated story of Bambi, a young deer hailed as the 'Prince of the Forest' at his birth. As Bambi grows, he makes friends with the other animals of the forest, learns the skills needed to survive, and even finds love. One day, however, the hunters come, and Bambi must learn to be as brave as his father if he is to lead the other deer to safety.","Language":"English","Country":"USA","Awards":"Nominated for 3 Oscars. Another 6 wins & 3 nominations.","Poster":"https://m.media-amazon.com/images/M/MV5BMTY1NzM4NDg5MV5BMl5BanBnXkFtZTgwMjI1MTkzMjE@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"7.3/10"},{"Source":"Rotten Tomatoes","Value":"90%"},{"Source":"Metacritic","Value":"91/100"}],"Metascore":"91","imdbRating":"7.3","imdbVotes":"125,824","imdbID":"tt0034492","Type":"movie","DVD":"01 Mar 2005","BoxOffice":"N/A","Production":"RKO Radio Pictures","Website":"N/A","Response":"True"}`,
	},
}

var testCasesUnmarshalMovie = []struct {
	description string
	input       string
	expected    Movie
}{
	{
		description: "It should return a correct struct marshaled from the json response",
		input:    `{"Title":"Bambi","Year":"1942","Rated":"G","Released":"21 Aug 1942","Runtime":"70 min","Genre":"Animation, Drama, Family","Director":"James Algar, Samuel Armstrong, David Hand, Graham Heid, Bill Roberts, Paul Satterfield, Norman Wright, Arthur Davis, Clyde Geronimi","Writer":"Felix Salten (from the story by), Perce Pearce (story direction), Larry Morey (story adaptation), Vernon Stallings (story development), Mel Shaw (story development), Carl Fallberg (story development), Chuck Couch (story development), Ralph Wright (story development)","Actors":"Hardie Albright, Stan Alexander, Bobette Audrey, Peter Behn","Plot":"The animated story of Bambi, a young deer hailed as the 'Prince of the Forest' at his birth. As Bambi grows, he makes friends with the other animals of the forest, learns the skills needed to survive, and even finds love. One day, however, the hunters come, and Bambi must learn to be as brave as his father if he is to lead the other deer to safety.","Language":"English","Country":"USA","Awards":"Nominated for 3 Oscars. Another 6 wins & 3 nominations.","Poster":"https://m.media-amazon.com/images/M/MV5BMTY1NzM4NDg5MV5BMl5BanBnXkFtZTgwMjI1MTkzMjE@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"7.3/10"},{"Source":"Rotten Tomatoes","Value":"90%"},{"Source":"Metacritic","Value":"91/100"}],"Metascore":"91","imdbRating":"7.3","imdbVotes":"125,824","imdbID":"tt0034492","Type":"movie","DVD":"01 Mar 2005","BoxOffice":"N/A","Production":"RKO Radio Pictures","Website":"N/A","Response":"True"}`,
		expected: Movie{"Bambi", "1942", "G", "21 Aug 1942", "70 min", "Animation, Drama, Family", "James Algar, Samuel Armstrong, David Hand, Graham Heid, Bill Roberts, Paul Satterfield, Norman Wright, Arthur Davis, Clyde Geronimi","Felix Salten (from the story by), Perce Pearce (story direction), Larry Morey (story adaptation), Vernon Stallings (story development), Mel Shaw (story development), Carl Fallberg (story development), Chuck Couch (story development), Ralph Wright (story development)", "Hardie Albright, Stan Alexander, Bobette Audrey, Peter Behn", "The animated story of Bambi, a young deer hailed as the 'Prince of the Forest' at his birth. As Bambi grows, he makes friends with the other animals of the forest, learns the skills needed to survive, and even finds love. One day, however, the hunters come, and Bambi must learn to be as brave as his father if he is to lead the other deer to safety.", "English", "USA","Nominated for 3 Oscars. Another 6 wins & 3 nominations.", "https://m.media-amazon.com/images/M/MV5BMTY1NzM4NDg5MV5BMl5BanBnXkFtZTgwMjI1MTkzMjE@._V1_SX300.jpg", []Rating{Rating{"Internet Movie Database", "7.3/10"}, Rating{"Rotten Tomatoes", "90%"}, Rating{"Metacritic", "91/100"}}, "91", "7.3", "125,824", "tt0034492", "movie", "01 Mar 2005","N/A","RKO Radio Pictures","N/A","True"},
	},
}