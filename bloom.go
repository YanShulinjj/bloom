/* ----------------------------------
*  @author suyame 2022-10-10 17:19:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package bloom

import (
	"hash/fnv"
)

var ExampleFuns = []func(string) int{
	func(s string) int {
		return Hash(s) + 100
	},
	func(s string) int {
		return Hash(s) + 200
	},
	func(s string) int {
		return Hash(s) + 300
	},
	func(s string) int {
		return Hash(s) + 400
	},
	func(s string) int {
		return Hash(s) + 500
	},
	func(s string) int {
		return Hash(s) + 600
	},
	func(s string) int {
		return Hash(s) + 700
	},
	func(s string) int {
		return Hash(s) + 800
	},
}

type Bloom struct {
	bitmap    []bool // 位图
	hashFuncs []func(string) int
}

func Hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

func NewBloom(bitlen int, fs []func(string) int) *Bloom {
	return &Bloom{
		bitmap:    make([]bool, bitlen),
		hashFuncs: fs,
	}
}

func (b *Bloom) Add(v string) {
	for _, f := range b.hashFuncs {
		idx := f(v) % len(b.bitmap)
		b.bitmap[idx] = true
	}
}

// Exists 判断集合中是否存在v
// 采用布隆过滤器
// 参考：https://developer.aliyun.com/article/773205
func (b *Bloom) Exists(v string) bool {
	for _, f := range b.hashFuncs {
		idx := f(v) % len(b.bitmap)
		if !b.bitmap[idx] {
			return false
		}
	}
	return true
}
