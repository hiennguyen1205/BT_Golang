package readFile

import (
	"log"
	"os"
)

//political_thought_works_corpus.csv
func ReadFile() *os.File {
	file, err := os.Open("testt.csv")
	if err != nil {
		log.Fatalf("failed to open")
	}
	return file
}
