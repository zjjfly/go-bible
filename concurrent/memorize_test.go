package concurrent

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

var Urls = [...]string{
	"http://www.a9vg.com",
	"https://godoc.org",
	"http://www.a9vg.com",
	"http://gopl.io",
	"http://www.a9vg.com",
	"https://godoc.org",
	"http://www.a9vg.com",
	"http://gopl.io",
}

func TestMemorize(t *testing.T) {
	m := New(HttpGetBody)
	var n sync.WaitGroup
	for _, url := range Urls {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
				n.Done()
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
