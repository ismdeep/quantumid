package quantumid

import (
	"crypto/rand"
	"time"
)

const (
	base64Alphabet = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
	base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

func base64(raw []byte) string {
	items := make([]byte, 22)

	items[0] = base64Alphabet[raw[0]<<2>>2]
	items[1] = base64Alphabet[raw[1]<<4>>2+raw[0]>>6]
	items[2] = base64Alphabet[raw[2]<<6>>2+raw[1]>>4]
	items[3] = base64Alphabet[raw[2]>>2]

	items[4] = base64Alphabet[raw[3]<<2>>2]
	items[5] = base64Alphabet[raw[4]<<4>>2+raw[3]>>6]
	items[6] = base64Alphabet[raw[5]<<6>>2+raw[4]>>4]
	items[7] = base64Alphabet[raw[5]>>2]

	items[8] = base64Alphabet[raw[6]<<2>>2]
	items[9] = base64Alphabet[raw[7]<<4>>2+raw[6]>>6]
	items[10] = base64Alphabet[raw[8]<<6>>2+raw[7]>>4]
	items[11] = base64Alphabet[raw[8]>>2]

	items[12] = base64Alphabet[raw[9]<<2>>2]
	items[13] = base64Alphabet[raw[10]<<4>>2+raw[9]>>6]
	items[14] = base64Alphabet[raw[11]<<6>>2+raw[10]>>4]
	items[15] = base64Alphabet[raw[11]>>2]

	items[16] = base64Alphabet[raw[12]<<2>>2]
	items[17] = base64Alphabet[raw[13]<<4>>2+raw[12]>>6]
	items[18] = base64Alphabet[raw[14]<<6>>2+raw[13]>>4]
	items[19] = base64Alphabet[raw[14]>>2]

	items[20] = base64Alphabet[raw[15]<<2>>2]
	items[21] = base64Alphabet[raw[15]>>6]
	return string(items)
}

func base58(raw []byte) string {
	num := make([]byte, len(raw))
	copy(num, raw)

	var result []byte
	for len(num) > 0 {
		rem := 0
		for i := 0; i < len(num); i++ {
			temp := rem*256 + int(num[i])
			num[i] = byte(temp / 58)
			rem = temp % 58
		}
		result = append([]byte{base58Alphabet[rem]}, result...)

		for len(num) > 0 && num[0] == 0 {
			num = num[1:]
		}
	}

	for _, b := range raw {
		if b != 0 {
			break
		}
		result = append([]byte{base58Alphabet[0]}, result...)
	}

	// 确保固定长度22位，前面用base58Alphabet[0]补全
	for len(result) < 22 {
		result = append([]byte{base58Alphabet[0]}, result...)
	}

	return string(result)
}

// Generate function
// Deprecated: Use Base64() instead.
func Generate() string {
	return Base64()
}

func Base64() string {
	raw := make([]byte, 16)
	s := time.Now().UnixNano()
	for i := 0; i < 8; i++ {
		raw[7-i] = byte(s % 256)
		s >>= 8
	}
	_, _ = rand.Read(raw[8:16])

	return base64(raw)
}

func Base58() string {
	raw := make([]byte, 16)
	s := time.Now().UnixNano()
	for i := 0; i < 8; i++ {
		raw[7-i] = byte(s % 256)
		s >>= 8
	}
	_, _ = rand.Read(raw[8:16])

	return base58(raw)
}
