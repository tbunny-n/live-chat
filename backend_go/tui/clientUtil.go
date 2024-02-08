package tui

import (
	"math/rand"
)

var wordList = []string{
	"moonlit",
	"sunlit",
	"starlit",
	"cloudy",
	"rainy",
	"sunny",
	"stormy",
	"windy",
	"breezy",
	"foggy",
	"misty",
	"icy",
	"snowy",
	"gloomy",

	"forest",
	"jungle",
	"desert",
	"tundra",
	"savannah",
	"grassland",
}

func generateUsername() string {
	// Choose two random words from the word list
	wordOne := wordList[rand.Intn(len(wordList))]
	wordTwo := wordList[rand.Intn(len(wordList))]

	return wordOne + wordTwo
}
