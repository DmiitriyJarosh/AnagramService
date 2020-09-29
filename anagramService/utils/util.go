package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Message to show parsing error on load
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{} {"status" : status, "message" : message}
}

// Method for adding our data to respond
func Respond(w http.ResponseWriter, data []string)  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Method for adding error message to respond
func RespondError(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Method for counting letters in word and saving it in string
func CountLetters(word string) string {
	word = strings.ToLower(word)
	var counterArray [26]int
	for _, v := range word {
		counterArray[v - 'a']++
	}
	s := ""

	for i, v := range counterArray {
		s += fmt.Sprintf("%s%d", string(i + 'a'), v)
	}
	return s
}