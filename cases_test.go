package main

var testCasesBuildParams = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "Single word title",
		input:       "bambi",
		expected:    "http://www.omdbapi.com?apikey=334b48a9&plot=full&t=bambi",
	},
		{
		description: "Multi word title",
		input:       "the dark knight",
		expected:    "http://www.omdbapi.com?apikey=334b48a9&plot=full&t=the+dark+knight",
	},
}
