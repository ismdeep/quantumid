package quantumid

import (
	"crypto/rand"
	"time"
)

const base = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"

func bytesToString(raw []byte) string {
	items := make([]byte, 22)

	items[0] = base[raw[0]<<2>>2]
	items[1] = base[raw[1]<<4>>2+raw[0]>>6]
	items[2] = base[raw[2]<<6>>2+raw[1]>>4]
	items[3] = base[raw[2]>>2]

	items[4] = base[raw[3]<<2>>2]
	items[5] = base[raw[4]<<4>>2+raw[3]>>6]
	items[6] = base[raw[5]<<6>>2+raw[4]>>4]
	items[7] = base[raw[5]>>2]

	items[8] = base[raw[6]<<2>>2]
	items[9] = base[raw[7]<<4>>2+raw[6]>>6]
	items[10] = base[raw[8]<<6>>2+raw[7]>>4]
	items[11] = base[raw[8]>>2]

	items[12] = base[raw[9]<<2>>2]
	items[13] = base[raw[10]<<4>>2+raw[9]>>6]
	items[14] = base[raw[11]<<6>>2+raw[10]>>4]
	items[15] = base[raw[11]>>2]

	items[16] = base[raw[12]<<2>>2]
	items[17] = base[raw[13]<<4>>2+raw[12]>>6]
	items[18] = base[raw[14]<<6>>2+raw[13]>>4]
	items[19] = base[raw[14]>>2]

	items[20] = base[raw[15]<<2>>2]
	items[21] = base[raw[15]>>6]
	return string(items)
}

// Generate function
func Generate() string {
	raw := make([]byte, 16)
	s := time.Now().UnixNano()
	for i := 0; i < 8; i++ {
		raw[7-i] = byte(s % 256)
		s >>= 8
	}
	_, _ = rand.Read(raw[8:16])

	return bytesToString(raw)
}
