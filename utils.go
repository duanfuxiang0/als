package als

import "math/rand"

func RandByte(r *rand.Rand) byte {
	v := r.Intn(256)
	for v <= 0 {
		v = r.Intn(256)
	}
	return byte(v)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")


func RandStringRunes(r *rand.Rand) string {
	n := r.Intn(256)
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
