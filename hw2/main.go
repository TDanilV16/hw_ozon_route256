package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

type Pair struct {
	word  string
	count int
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	keyWords := []string{"успех", "цел", "часть", "сложно", "будет"}

	in := make(chan string)
	out := make([]chan Pair, 2)

	go read(file, in)

	for i := 0; i < len(out); i++ {
		out[i] = make(chan Pair)
		go countWords(keyWords, in, out[i])
	}

	counter := merge(out)

	write(counter)
}

func read(r io.Reader, in chan<- string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		in <- strings.ToLower(scanner.Text())
	}
	close(in)
}

func countWords(keyWords []string, in <-chan string, outPairs chan<- Pair) {

	for line := range in {
		for _, word := range keyWords {
			wordCountInLine := strings.Count(line, word)
			p := Pair{word, wordCountInLine}
			outPairs <- p
		}
	}

	close(outPairs)
}
func merge(inPair []chan Pair) map[string]int {
	var wg sync.WaitGroup
	var mu sync.Mutex

	counter := make(map[string]int)

	wg.Add(len(inPair))

	for _, ch := range inPair {
		go func(ch chan Pair) {
			defer wg.Done()
			for pair := range ch {
				mu.Lock()
				counter[pair.word] += pair.count
				mu.Unlock()
			}
		}(ch)
	}

	wg.Wait()

	return counter
}

func write(counter map[string]int) {
	var total int

	for word, count := range counter {
		total += count
		fmt.Printf("%s: %d\n", word, count)
	}

	fmt.Printf("всего: %d", total)
}
