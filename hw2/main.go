package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Counter struct {
	counter map[string]int
	total   int
}

func New(text string, keyWords []string) *Counter {

	counter := make(map[string]int)

	total := 0

	for _, word := range keyWords {
		counter[word] = strings.Count(text, word)
		total += counter[word]
	}

	return &Counter{
		counter: counter,
		total:   total,
	}
}

func (c *Counter) Find(word string) int {
	if count, ok := c.counter[word]; ok {
		return count
	}

	return 0
}

func main() {
	file, err := os.Open("hw2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var textLines []string

	for scanner.Scan() {
		line := scanner.Text()
		textLines = append(textLines, line)
	}

	text := strings.Join(textLines, "\n")

	text = strings.ToLower(text)

	keyWords := []string{"успех", "цел", "часть", "сложно", "будет"}

	counter := New(text, keyWords)

	for word, count := range counter.counter {
		fmt.Printf("%s: %d\n", word, count)
	}

	fmt.Printf("всего: %d", counter.total)
}
