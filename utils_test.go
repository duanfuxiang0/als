package als

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandBytes(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len, st := RandBytes(r)
	t.Logf("%d, %s", len, string(st))
}
