package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var text []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text, scanner.Err()
}

func searchInText(text []string, query string) []string {
	var result []string
	for _, line := range text {
		sentences := strings.Split(line, ". ")
		for _, sentence := range sentences {
			if strings.Contains(sentence, query) {
				result = append(result, sentence)
			}
		}
	}
	return result
}

func main() {
	text, err := readFile("textfile.txt")
	if err != nil {
		fmt.Println("Помилка відкриття файлу:", err)
		return
	}

	var query string
	fmt.Println("Введіть рядок для пошуку:")
	fmt.Scanln(&query)

	results := searchInText(text, query)

	fmt.Println("Результати пошуку:")
	for _, sentence := range results {
		fmt.Println(sentence)
	}
}