package lsh256

import "hash"

func newContext(size int) hash.Hash {
	return newContextGo(size)
}

func sum(size int, data []byte) [Size]byte {
	return sumGo(size, data)
}
