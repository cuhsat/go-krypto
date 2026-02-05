package internal

func Memclr(b []byte) {
	if len(b) == 0 {
		return
	}
	b[0] = 0
	for bp := 1; bp < len(b); bp *= 2 {
		copy(b[bp:], b[:bp])
	}
}

func MemclrU32(b []uint32) {
	if len(b) == 0 {
		return
	}
	b[0] = 0
	for bp := 1; bp < len(b); bp *= 2 {
		copy(b[bp:], b[:bp])
	}
}

func MemclrU64(b []uint64) {
	if len(b) == 0 {
		return
	}
	b[0] = 0
	for bp := 1; bp < len(b); bp *= 2 {
		copy(b[bp:], b[:bp])
	}
}
