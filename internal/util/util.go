package util

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	math_rand "math/rand"
)

func P[T any](v T) *T {
	return &v
}

var (
	rnd *rand.Rand
)

func init() {
	b := make([]byte, 8)
	crypto_rand.Read(b)
	rnd = math_rand.New(math_rand.NewSource(int64(binary.LittleEndian.Uint64(b))))
}

func Rnd6() string {
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}
