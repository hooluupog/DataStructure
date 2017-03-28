package bitmap

type BitMap struct {
	//                  3   2   1   0
	block   []byte // |0|1|0|1|1|0|0|0|      e.g (radix == 2)
	r       uint64 //radix, 1 digit == r bits.
	lowest  uint64 // first flag bit
	highest uint64 // last flag bit
}

func (bm *BitMap) NewBitMap(size, r uint64) *BitMap {
	if r > 4 {
		panic("Radix should not bigger than 4.")
	}
	radix := r
	if r == 3 { // Do not use 3 bits to avoid unaligned memory access.
		radix = 4
	}
	length := size * radix / 8
	if size*radix%8 != 0 {
		length = length + 1
	}
	bm = &BitMap{
		block: make([]byte, length), r: radix,
		lowest: length - 1, highest: 0,
	}
	return bm
}

func (bm *BitMap) Get(idx uint64) string {
	buf := make([]byte, bm.r)
	bIdx, offset := bm.getAddress(idx)
	shift := bm.block[bIdx] >> (offset * bm.r)
	for i := uint64(0); i < bm.r; i++ {
		buf[i] = (shift & 0x1) + 48
		shift = shift >> 1
	}
	return string(buf)
}

func (bm *BitMap) Set(idx uint64, s string) {
	if uint64(len(s)) != bm.r {
		panic("Illegal string length.")
	}
	bIdx, offset := bm.getAddress(idx)
	if bIdx >= uint64(len(bm.block)) {
		panic("index out of range.")
	}
	if bIdx < bm.lowest {
		bm.lowest = bIdx
	}
	if bIdx > bm.highest {
		bm.highest = bIdx
	}
	shift := byte(0x1 << (offset * bm.r))
	for _, v := range s {
		if v == '0' { // clear bit.
			bm.block[bIdx] &^= shift
		} else if v == '1' {
			bm.block[bIdx] |= shift
		} else {
			panic("Illegal string.")
		}
		shift = shift << 1
	}
}

func (bm BitMap) Block() []byte {
	if bm.lowest > bm.highest {
		return bm.block
	}
	return bm.block[bm.lowest : bm.highest+1]
}

func (bm *BitMap) getAddress(idx uint64) (uint64, uint64) {
	return idx * bm.r / 8, idx % (8 / bm.r)
}

func ReverseString(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	return string(b)
}
