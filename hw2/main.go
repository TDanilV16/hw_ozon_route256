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

func New(textLines []string, keyWords []string) *Counter {

	c := make(chan map[string]int)

	go countWordsChan(textLines[:len(textLines)/2], keyWords, c)
	go countWordsChan(textLines[len(textLines)/2:], keyWords, c)

	first, second := <-c, <-c

	for word, count := range first {
		if _, ok := second[word]; !ok {
			second[word] = count
		} else {
			second[word] += count
		}
	}

	counter := second

	total := 0

	for _, count := range counter {
		total += count
	}
	return &Counter{
		counter: counter,
		total:   total,
	}
}

func countWords(lines []string, keyWords []string) map[string]int {
	text := strings.Join(lines, "\n")
	text = strings.ToLower(text)

	counter := make(map[string]int)

	for _, word := range keyWords {
		counter[word] = strings.Count(text, word)
	}

	return counter
}

func countWordsChan(lines []string, keyWords []string, c chan map[string]int) {
	c <- countWords(lines, keyWords)
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

	keyWords := []string{"успех", "цел", "часть", "сложно", "будет"}

	counter := New(textLines, keyWords)

	for word, count := range counter.counter {
		fmt.Printf("%s: %d\n", word, count)
	}

	fmt.Printf("всего: %d", counter.total)
}
