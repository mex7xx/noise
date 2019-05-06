package noise

import (
	"github.com/pkg/errors"
)

var ErrDisconnect = errors.New("disconnect requested")

func continuously(fn func(stop <-chan struct{}) error) func(stop <-chan struct{}) error {
	return func(stop <-chan struct{}) error {
		for {
			if err := fn(stop); err != nil {
				return err
			}
		}
	}
}
