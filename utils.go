package als

import "math/rand"


var letterBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
const DefaultStrLen = 50
const LettersLen = 52

func RandByte(r *rand.Rand) byte {
	v := r.Intn(DefaultStrLen)
	for v == 0 {
		v = r.Intn(DefaultStrLen)
	}
	return byte(v)
}

func RandBytes(r *rand.Rand, max int) (int, []byte) {
	if max > 256 || max <= 0 {
		max = DefaultStrLen
	}
	n := r.Intn(max)
	for n == 0 {
		n = r.Intn(max)
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(LettersLen)]
	}
	return n, b
}
