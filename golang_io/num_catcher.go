package golang_io

import (
	"io"
)

type NumCatcher struct {
	source string
	index int
}

func NewNumCatcher(src string) *NumCatcher {
	return &NumCatcher{source: src}
}

func filter(r byte) byte {
	if r >= '0' && r <= '9' {
		return r
	}
	return 0
}

func (nc *NumCatcher) Read(p []byte) (int, error) {
	// n: 待返回读取长度
	// bound: 本次实际读取长度
	// rest: 剩余未读长度
	var (
		n, bound, rest int
	)
	if nc.index >= len(nc.source) {
		return 0, io.EOF
	}

	// 确定界限
	rest = len(nc.source) - nc.index
	if rest >= len(p) {
		bound = len(p)
	} else {
		bound = rest
	}

	// 开始过滤
	buf := make([]byte, bound)
	for n < bound {
		if char := filter(nc.source[nc.index]); char != 0 {
			buf[n] = char
		}
		n++
		nc.index++
	}
	copy(p, buf)
	return n, nil
}