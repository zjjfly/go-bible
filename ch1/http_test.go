package ch1

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestHttpGet(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	written, err := io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
	fmt.Printf("http status code:%d\n", resp.StatusCode)
	if err != nil {
		fmt.Printf("read %s:%v\n", "http://www.baidu.com", err)
		os.Exit(1)
	}
	fmt.Printf("response length:%d\n", written)
}

func TestConcurrentHttpGet(t *testing.T) {
	start := time.Now()
	ch := make(chan string)
	urls := [5]string{
		"http://www.baidu.com",
		"http://www.a9vg.com",
		"http://www.sina.com",
		"http://www.zhihu.com",
		"http://www.douban.com",
	}
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	written, err := io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, written, url)
}

func TestHttpServer(t *testing.T) {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	lissajous(w)
}
