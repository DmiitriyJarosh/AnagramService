package models

import (
	u "../utils"
	"gorm.io/gorm"
	"log"
)

// Struct for word and its letters
type Words struct {
	gorm.Model
	LetterCounter string `gorm:"index:idx_letters,type:hash"`
	Word string
}

// Check if this word exists in db
func (word *Words) Validate() bool {
	temp := &Words{}
	err := GetDB().Table("words").Where("word = ?", word.Word).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}
	if temp.Word != "" {
		return false
	}

	return true
}

func (word *Words) Create() bool {
	if !word.Validate() {
		return false
	}

	word.LetterCounter = u.CountLetters(word.Word)

	// Add word to our db
	GetDB().Create(word)

	if word.ID <= 0 {
		return false
	}

	return true
}

// Get slice of anagrams for word
func GetAnagrams(word string) []string {
	letters := u.CountLetters(word)
	anagrams := make([]*Words, 0)
	err := GetDB().Table("words").Where("letter_counter = ?", letters).Find(&anagrams).Error
	if err != nil {
		log.Print(err)
		return nil
	}

	return WordsToStrings(anagrams)
}

func WordsToStrings(words []*Words) []string {
	wordStrings := make([]string, len(words))
	for i, v := range words {
		wordStrings[i] = v.Word
	}
	return wordStrings
}
