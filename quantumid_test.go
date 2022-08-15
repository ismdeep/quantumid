package quantumid

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	got := Generate()
	t.Logf("got = %v", got)
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Generate()
	}
}

func Test_bytesToString(t *testing.T) {
	type args struct {
		raw []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				//       15       14       13       12       11       10        9        8        7        6        5        4        3        2        1        0
				// 11111010 00010100 11101000 11100110 11101110 10011100 11111101 00111010 00101010 00101100 11111111 10000101 01101111 10010001 10011110 00110010
				//     0xfa     0x14     0xe8     0xe6     0xee     0x9c     0xfd     0x3a     0x2a     0x2c     0xff     0x85     0x6f     0x91     0x9e     0x32
				//
				//     21     20     19     18     17     16     15     14     13     12     11     10      9      8      7      6      5      4      3      2      1      0
				// 000011 111010 000101 001110 100011 100110 111011 101001 110011 111101 001110 100010 101000 101100 111111 111000 010101 101111 100100 011001 111000 110010
				//      3     58      5     14     35     38     59     41     51     61     14     34     40     44     63     56     21     47     36     25     56     50
				//
				// a b c d e f g h i j  k  l  m  n  o  p  q  r  s  t  u  v  w  x  y  z  A  B  C  D  E  F  G  H  I  J  K  L  M  N  O  P  Q  R  S  T  U  V  W  X  Y  Z  0  1  2  3  4  5  6  7  8  9  _  -
				// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63
				//
				// 50 56 25 36 47 21 56 63 44 40 34 14 61 51 41 59 38 35 14 5 58 3
				//  Y  4  z  K  V  v  4  -  S  O  I  o  9  Z  P  7  M  J  o f  6 d
				//  Y4zKVv4-SOIo9ZP7MJof6d
				raw: []byte{0x32, 0x9e, 0x91, 0x6f, 0x85, 0xff, 0x2c, 0x2a, 0x3a, 0xfd, 0x9c, 0xee, 0xe6, 0xe8, 0x14, 0xfa},
			},
			want: "Y4zKVv4-SOIo9ZP7MJof6d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bytesToString(tt.args.raw); got != tt.want {
				t.Errorf("bytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
