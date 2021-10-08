package concurrent

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func HttpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		log.Println("cache hit")
		memo.mu.Unlock()
		//key已经在缓存中，但可能实际的值还没有放入，所以要等待ready这个channel关闭
		<-e.ready
	}
	return e.res.value, e.res.err
}
