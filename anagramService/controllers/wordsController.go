package controllers

import (
	"../models"
	u "../utils"
	"encoding/json"
	"log"
	"net/http"
)

// Handler for loading words to db
var LoadWords = func(w http.ResponseWriter, r *http.Request) {
	var newWords []string
	err := json.NewDecoder(r.Body).Decode(&newWords)
	if err != nil {
		u.RespondError(w, u.Message(false, "Error while decoding request body"))
		return
	}

	existWordsCounter := 0

	for _, v := range newWords {
		word := &models.Words{Word: v}
		if isWordExist := word.Create(); isWordExist {
			existWordsCounter++
		}
	}
	log.Printf(
		"%d word(s) were loaded.\n%d of them were new for our base and were added to it.\n%d word(s) were skipped as duplicates\n",
		len(newWords),
		existWordsCounter,
		len(newWords) - existWordsCounter)
}

// Handler for get anagrams request
var GetAnagramsFor = func(w http.ResponseWriter, r *http.Request) {
	word, ok := r.URL.Query()["word"]

	if !ok || len(word) < 1 {
		log.Println("Url Param 'word' is missing")
		return
	}

	anagrams := models.GetAnagrams(word[0])
	u.Respond(w, anagrams)
}
