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

func Test_base58(t *testing.T) {
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
				raw: []byte{0x32, 0x9e, 0x91, 0x6f, 0x85, 0xff, 0x2c, 0x2a, 0x3a, 0xfd, 0x9c, 0xee, 0xe6, 0xe8, 0x14, 0xfa},
			},
			want: "7FYMc7GxdpecwijKduEstV",
		},
		{
			name: "",
			args: args{
				raw: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			},
			want: "1111111111111111111111",
		},
		{
			name: "",
			args: args{
				raw: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			},
			want: "YcVfxkQb6JRzqk5kF2tNLv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base58(tt.args.raw); got != tt.want {
				t.Errorf("base58() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base64(t *testing.T) {
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
				raw: []byte{0x32, 0x9e, 0x91, 0x6f, 0x85, 0xff, 0x2c, 0x2a, 0x3a, 0xfd, 0x9c, 0xee, 0xe6, 0xe8, 0x14, 0xfa},
			},
			want: "msOZjKszgcXDxndvaYD4u2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base64(tt.args.raw); got != tt.want {
				t.Errorf("base64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase58(t *testing.T) {
	t.Logf("Base58() got = %v", Base58())
}

func BenchmarkBase58(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base58()
	}
}

func BenchmarkBase64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base64()
	}
}

func TestBase58Order(t *testing.T) {
	// 在16字节完整范围内选择10个不同区间，每个区间100万个数据进行顺序测试
	const testCount = 1000000
	const intervalCount = 10

	// 16字节数据的最大值是 2^128 - 1，我们将其分成10个区间
	// 使用高位8字节来划分区间
	for interval := 0; interval < intervalCount; interval++ {
		// 计算起始位置：将高位8字节设置为 interval * (2^64 / 10)
		startHigh := uint64(interval) * (1 << 61) // 2^64 / 10 约等于 2^61 * 1.6，这里简化为 2^61

		// 构造起始16字节数据
		startBytes := make([]byte, 16)
		for i := 0; i < 8; i++ {
			startBytes[i] = byte(startHigh >> (8 * (7 - i)))
		}

		t.Logf("Testing interval %d/%d, starting from 0x%x", interval+1, intervalCount, startBytes)

		var results []string

		for i := 0; i < testCount; i++ {
			// 生成当前数据：在起始位置的基础上加上i
			currentBytes := make([]byte, 16)
			copy(currentBytes, startBytes)

			// 在低位加上i
			carry := uint64(i)
			for j := 15; j >= 0 && carry > 0; j-- {
				sum := uint64(currentBytes[j]) + carry
				currentBytes[j] = byte(sum & 0xFF)
				carry = sum >> 8
			}

			result := base58(currentBytes)
			results = append(results, result)

			// 每10万条输出一次进度
			if i%100000 == 0 {
				t.Logf("  Progress: %d/%d", i, testCount)
			}
		}

		// 验证当前区间内base58编码结果的字典序与字节序一致
		for i := 1; i < len(results); i++ {
			if results[i-1] >= results[i] {
				t.Errorf("Order mismatch in interval %d at index %d: %s should be < %s", interval+1, i, results[i-1], results[i])
				return
			}
		}

		t.Logf("Interval %d passed: %d cases in correct order", interval+1, testCount)
	}

	t.Logf("Successfully tested %d intervals with %d cases each, all in correct order", intervalCount, testCount)
}
