package debounce

// references:
//	https://studygolang.com/articles/34169
//  https://github.com/bep/debounce/blob/master/debounce.go
//  https://nathanleclaire.com/blog/2014/08/03/write-a-function-similar-to-underscore-dot-jss-debounce-in-golang/
//	https://drailing.net/2018/01/debounce-function-for-golang/

import (
	"sync"
	"time"
)

// New returns a new debounced function which will postpone its execution
func New(fn func(), interval time.Duration) func() error {
	var mu sync.Mutex
	var timer *time.Timer

	return func() error {
		mu.Lock()
		defer mu.Unlock()

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(interval, fn)

		return nil
	}
}
