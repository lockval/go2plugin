package go2plugin

import (
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

var (
	src = rand.NewSource(time.Now().UnixNano())
)

const (
	// // base92
	// letterIdxBits = 7 // 7 bits to represent a letter index

	// base62
	letterIdxBits = 6 // 7 bits to represent a letter index

	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	// base62
	letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	// // base80
	// letterBytes = "%*+,-./0123456789:=?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_abcdefghijklmnopqrstuvwxyz{}~"
)

// Base62UUIDLen Base62UUIDLen
const Base62UUIDLen = 22

// RandStringBytesMaskImprSrcSB RandStringBytesMaskImprSrcSB
func RandStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

// RandStringBytesMaskImprSrcUnsafe RandStringBytesMaskImprSrcUnsafe
func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// NewBase62UUID NewBase62UUID
func NewBase62UUID() string {
	return RandStringBytesMaskImprSrcUnsafe(Base62UUIDLen)
}

// // Base80UUIDLen Base80UUIDLen
// const Base80UUIDLen = 21

// // NewBase80UUID NewBase80UUID
// func NewBase80UUID() string {
// 	return RandStringBytesMaskImprSrcUnsafe(Base80UUIDLen)
// }
