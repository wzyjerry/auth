package util

import (
	crypto_rand "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/rand"
	math_rand "math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"

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

func ImgBase64(img []byte) string {
	mimeType := http.DetectContentType(img)
	var avatar strings.Builder
	avatar.WriteString("data:")
	avatar.WriteString(mimeType)
	avatar.WriteString(";base64,")
	avatar.Write([]byte(base64.StdEncoding.EncodeToString(img)))
	return avatar.String()
}

func IsLoopback(hostname string) bool {
	ip := net.ParseIP(hostname)
	if ip == nil {
		if hostname == "localhost" {
			return true
		}
	}
	return ip.IsLoopback()
}

// 验证actual是否与expected相同或为子域名
// 对于环回地址，允许不安全版本协议和任意端口号
// 否则，只允许https协议，且要求端口匹配
func VerifyRedirectUri(expected string, actual string) bool {
	expectedUrl, err := url.Parse(expected)
	if err != nil {
		return false
	}
	actualUrl, err := url.Parse(actual)
	if err != nil {
		return false
	}
	if IsLoopback(expectedUrl.Hostname()) {
		if IsLoopback(actualUrl.Hostname()) {
			return strings.HasPrefix(actualUrl.Path, expectedUrl.Path)
		}
		return false
	}
	if !strings.EqualFold(expectedUrl.Scheme, "https") && !strings.EqualFold(actualUrl.Scheme, "https") {
		return false
	}
	if expectedUrl.Host != actualUrl.Host {
		return false
	}
	return strings.HasPrefix(actualUrl.Path, expectedUrl.Path)
}
