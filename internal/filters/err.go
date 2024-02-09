package filters

import (
	"github.com/ibiscum/Effective-Concurrency-in-Go/internal/store"
)

func ErrFilter(in <-chan store.Entry) (<-chan store.Entry, <-chan error) {
	outCh := make(chan store.Entry)
	errCh := make(chan error)
	go func() {
		defer close(outCh)
		defer close(errCh)
		for entry := range in {
			if entry.Error != nil {
				errCh <- entry.Error
			} else {
				outCh <- entry
			}
		}
	}()
	return outCh, errCh
}
