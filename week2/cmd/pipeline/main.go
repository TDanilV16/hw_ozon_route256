package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
	"week2/logger"
)

func main() {
	ctx := context.Background()

	ctx = logger.ContextWithTimeStamp(ctx, time.Now())

	if err := run(ctx); err != nil {
		logger.Errorf(ctx, "%value", err)
		os.Exit(1)
	}

	logger.Infof(ctx, "exit")
}

type Line struct {
	value string
	idx   int
}

func run(ctx context.Context) error {

	r, err := os.Open("input")
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.OpenFile("output", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer w.Close()

	in := make(chan Line)
	out := make([]chan Line, runtime.NumCPU())

	go read(ctx, r, in)

	for i := 0; i < len(out); i++ {
		out[i] = make(chan Line)
		go process(ctx, in, out[i])

	}

	write(ctx, w, merge(out))

	return nil
}

func merge(in []chan Line) chan Line {
	out := make(chan Line)

	var wg sync.WaitGroup

	wg.Add(len(in))

	for _, ch := range in {
		go func(ch chan Line) {
			defer wg.Done()
			for l := range ch {
				out <- l
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func read(ctx context.Context, r io.Reader, out chan<- Line) {
	s := bufio.NewScanner(r)

	idx := 1
	for s.Scan() {
		line := s.Text()
		out <- Line{
			line,
			idx,
		}

		idx++
	}

	close(out)
}

func process(ctx context.Context, in <-chan Line, out chan<- Line) {
	for l := range in {
		words := strings.Split(l.value, " ")
		for i := range words {
			words[i] = strings.Title(words[i])
		}
		out <- Line{strings.Join(words, " "), l.idx}
	}
	close(out)
}

func write(ctx context.Context, w io.Writer, in <-chan Line) {
	var cache []Line

	for l := range in {
		cache = append(cache, l)
	}

	sort.Slice(cache, func(i, j int) bool {
		return cache[i].idx < cache[j].idx
	})

	for _, l := range cache {
		fmt.Fprintf(w, "%d: %s\n", l.idx, l.value)
	}

}
