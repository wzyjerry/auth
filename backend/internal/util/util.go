package util

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	math_rand "math/rand"

	"github.com/google/uuid"
	"github.com/google/wire"
)

func P[T any](v T) *T {
	return &v
}

var (
	rnd         *rand.Rand
	ProviderSet = wire.NewSet(NewTokenHelper, NewAliyunHelper)
)

func init() {
	b := make([]byte, 8)
	crypto_rand.Read(b)
	rnd = math_rand.New(math_rand.NewSource(int64(binary.LittleEndian.Uint64(b))))
}

func Rnd6() string {
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

func RndFloat64() float64 {
	return rnd.Float64()
}

func NewUUID() string {
	return uuid.NewString()
}
