package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Row struct {
	Id int
}

func main() {
	text := getAllText()
	word := "wisdom"
	words := textToWords(text)
	frequency := wordFrequency(words)
	fmt.Println(frequency)
	mapWords := indexText(words)
	fmt.Println(mapWords)
	wordRows := mapWords[word]
	fmt.Println(wordRows)
}

func getAllText() []string {
	myfile, err := os.Open("text.txt")
	text := []string{}
	if err != nil {
		fmt.Println("Error opening file:", err)
		return text
	}
	defer myfile.Close()

	scanner := bufio.NewScanner(myfile)

	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
	return text
}

func textToWords(text []string) []string {
	wordsSlice := []string{}
	for _, v := range text {
		w := strings.Split(v, " ")
		wordsSlice = append(wordsSlice, w...)

	}
	return wordsSlice
}

func wordFrequency(words []string) map[string]int {
	frequency := make(map[string]int)
	for _, v := range words {
		frequency[v]++
	}
	return frequency
}

func indexText(words []string) map[string][]Row {
	result := make(map[string][]Row)

	for i, word := range words {
		word = strings.ToLower(word)
		result[word] = append(result[word], Row{Id: i})
	}

	return result
}
