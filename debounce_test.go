package debounce

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	count := 0

	fn := func() {
		count++
	}
	debouncedFn := New(fn, 10*time.Millisecond)

	for i := 0; i < 100; i++ {
		debouncedFn()
	}

	time.Sleep(20 * time.Millisecond)
	if count != 1 {
		t.Errorf("Expected count to be 1, got %d", count)
	}
}
