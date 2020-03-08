package als

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandStringRunes(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	st := RandStringRunes(r)
	fmt.Println(st)
}
