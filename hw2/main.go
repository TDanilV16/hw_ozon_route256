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
func merge(inPair []chan Pair) chan Pair {
	var wg sync.WaitGroup

	resultCh := make(chan Pair)

	wg.Add(len(inPair))

	for _, ch := range inPair {
		go func(ch chan Pair) {
			defer wg.Done()
			for pair := range ch {
				resultCh <- pair
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	return resultCh
}

func write(resultCh <-chan Pair) {
	var total int

	cnt := make(map[string]int)

	for pair := range resultCh {
		cnt[pair.word] += pair.count
	}

	for word, count := range cnt {
		total += count
		fmt.Printf("%s: %d\n", word, count)
	}

	fmt.Printf("всего: %d", total)
}
