package myid

import (
	"math/rand"
	"time"
)

func init() { rand.Seed(time.Now().UnixNano()) }

// Generate generates ID.
//
// ID Spec:
//   length=8, [0-9,a-z,A-Z]
//
// TODO: Better method?
//
func Generate() string {
	const str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// random returns random int in [0,n).
	random := func(n int) int { return rand.Intn(n) }

	return string([]byte{
		str[random(len(str))], str[random(len(str))],
		str[random(len(str))], str[random(len(str))],
		str[random(len(str))], str[random(len(str))],
		str[random(len(str))], str[random(len(str))],
	})
}
