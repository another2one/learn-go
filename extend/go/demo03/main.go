// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	cache map[string]result
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

var mu sync.Mutex

// NOTE: not concurrency-safe!
func (memo *Memo) Get(key string) (interface{}, error) {
	mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	mu.Unlock()
	return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
	fmt.Println("get url .......")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() []string {
	return []string{
		"https://www.baidu.com",
		"https://www.lewaimai.com",
	}
}

var wg sync.WaitGroup

func main() {
	m := New(httpGetBody)
	totalStart := time.Now()
	for i := 0; i < 3; i++ {
		for _, url := range incomingURLs() {
			wg.Add(1)
			go func(url string) {
				start := time.Now()
				value, err := m.Get(url)
				if err != nil {
					log.Print(err)
				}
				fmt.Printf("%s, %s, %d bytes\n",
					url, time.Since(start), len(value.([]byte)))
				wg.Done()
			}(url)
		}
	}
	wg.Wait()
	fmt.Printf("---total use: %s bytes\n", time.Since(totalStart))
}
