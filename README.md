# AnagramService
It is a simple service which can find anagrams for loaded words.

## Requirements
This service use `Postgre` as database, so you have to install it and setup `.env` file in project. Also you will need some packets for Golang:
* `go get github.com/gorilla/mux` - Router for REST API
* `go get github.com/joho/godotenv` - .env file support
* `go get gorm.io/gorm` - ORM for Golang
* `go get gorm.io/driver/postgres` - ORM driver for postgre

## Usage
To run service use command `go run main.go` in directory `anagramService`.

Use `localhost:8080/load` to load words. They have to be packed as array in Json.

Use `localhost:8080/get?word=<you word here>` to find anagrams for the word.

Examples:
* `curl localhost:8080/load -d '["foobar", "aabb", "baba", "boofar", "test"]'`
* `curl localhost:8080/get?word=foobar`

## Main idea
Main idea of service is to use hash index of `Postgre` for fast search via words for possible anagrams. As anagrams are common in amount of the same letters, so for each words we count amount of each letter in it and save this data in string. For example: `"a1b1c2d0e0..."` (all other letters followed by zeroes) for `"abcc"`.

When we want to find anagrams we just search in hash index for such string which was calculated for word in request.
