package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	readNumbers()
	readText()
}

func readNumbers() {
	path := "files/1689007675141_numbers.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Помилка відкриття файлу:", err)
		return
	}
	defer file.Close()
	re := regexp.MustCompile(`\(?\d{3}\)?[-. ]?\d{3}[-. ]?\d{4}`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			fmt.Println("Знайдено телефонний номер:", match)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Помилка читання файлу:", err)
	}
}

func readText() {
	path := "1689007676028_text.txt"
}
