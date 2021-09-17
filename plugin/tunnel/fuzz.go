// +build gofuzz

package tunnel

import (
	"github.com/coredns/coredns/plugin/pkg/fuzz"
)

// Fuzz fuzzes cache.
func Fuzz(data []byte) int {
	w := Tunnel{}
	return fuzz.Do(w, data)
}
