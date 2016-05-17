package compress

import (
	"bytes"
)

const (
	FIRST_QUARTER = 0X200000
	THIRD_QUARTER = 0X600000
	HALF          = 0X400000
	HIGH          = 0X7FFFFF
	INITIAL_READ  = 23
)

type bitOutputBuffer struct {
	ByteBuffer   bytes.Buffer
	CurrentByte  byte
	CurrentBit   byte
}

func (buf *bitOutputBuffer) writeBit(bit byte) {
	buf.CurrentByte = buf.CurrentByte << 1
	buf.CurrentByte += bit 
	buf.CurrentBit += 1

	if buf.CurrentBit == 8 {
		buf.ByteBuffer.WriteByte(buf.CurrentByte)
		buf.CurrentByte = 0
		buf.CurrentBit = 0
	}
}

func (buf *bitOutputBuffer) flush() {
	for {
		if buf.CurrentBit == 0 {
			break
		}

		buf.writeBit(0)
	}
}

func (buf *bitOutputBuffer) bytes() []byte {
	buf.flush()

	return buf.ByteBuffer.Bytes()
}

func Compress(str string) []byte {
	chars := []byte(str)

	bitBuf := bitOutputBuffer{}
	
	var low, mLow, mStep, mScale int
	var freq [257]int
	var current byte

	total := 257
	high := HIGH
	mHigh := HIGH

	for i := 0; i < 257; i++ {
		freq[i] = 1
	}

	for i := 0; i < len(chars) + 1; i++ {
		if i == len(chars) {
			low = total - 1
			high = total
		} else {
			current = chars[i] & 0XFF
			low = 0

			for j := byte(0); j < current; j++ {
				low += freq[j]				
			}

			high = low + freq[current]
		}

		mStep = (mHigh - mLow + 1) / total
		mHigh = (mLow + mStep * high) - 1
		mLow = mLow + mStep*low

		for {
			if mHigh < HALF {
				bitBuf.writeBit(0)
				mLow = mLow * 2
				mHigh = mHigh * 2 + 1

				for ; mScale > 0; mScale-- {
					bitBuf.writeBit(1)
				}
			} else if mLow >= HALF {
				bitBuf.writeBit(1)
				mLow = (mLow - HALF) * 2
				mHigh = (mHigh - HALF) * 2 + 1

				for ; mScale > 0; mScale-- {
					bitBuf.writeBit(0)
				}
			} else {
				break
			}
		}

		for {
			if !((FIRST_QUARTER <= mLow) && (mHigh < THIRD_QUARTER)) {
				break
			}

			mScale += 1
			mLow = (mLow - FIRST_QUARTER) * 2
			mHigh = (mHigh - FIRST_QUARTER) * 2 + 1
		}

		freq[current] += 1
		total += 1
	}

	if(mLow < FIRST_QUARTER) {
		bitBuf.writeBit(0)

		for i := 0; i < mScale+1; i++ {
			bitBuf.writeBit(1)
		}
	} else {
		bitBuf.writeBit(1)
	}

	bitBuf.flush()

	return bitBuf.bytes()
}

type bitInputBuffer struct {
	source      []byte
	bytep       int 
	bitp        uint 
	currentByte byte 
}

func newBitInputBuffer(source []byte) *bitInputBuffer {
	return &bitInputBuffer{source: source, currentByte: source[0]}
}

func (buf *bitInputBuffer) readBit() int {
	result := int((buf.currentByte >> 7) & 1)
	buf.currentByte = buf.currentByte << 1
	buf.bitp++
	
	if buf.bitp == 8 {
		buf.bytep++
		if buf.bytep > len(buf.source) -1 {
			buf.currentByte = 0
		} else {
			buf.currentByte = buf.source[buf.bytep]
			buf.bitp = 0
		}
	}

	return result
}

func Decompress(in []byte) []byte {
	var freq [257]int
	for i:=0; i<257; i++ {
		freq[i] = 1
	}

	var low, current, mStep, mScale, mBuffer int
	high  := HIGH
	mHigh := high 
	mLow  := low
	total := 257
	
	buf   := bytes.Buffer{}
	inbuf := newBitInputBuffer(in)

	for i := 0; i<INITIAL_READ; i++ {
		mBuffer = 2*mBuffer
		mBuffer += inbuf.readBit()
	}

	for {
		/* 1. Retrieve current byte */
		mStep = (mHigh - mLow + 1) / total
		value := (mBuffer - mLow) / mStep
		low = 0
		for current=0; current<256 && low+freq[current] <= value; current++ {
			low += freq[current]
		}

		if current==256 {
			break
		}

		buf.WriteByte(byte(current))
		high = low + freq[current]

		/* 2. Update the decoder */
		mHigh = mLow + mStep * high - 1; // interval open at the top => -1

		/* Update lower bound */
		mLow = mLow + mStep * low;		

		for {
			if mHigh < HALF  {
				mLow = mLow * 2
				mHigh = (mHigh * 2) + 1
				mBuffer = (2 * mBuffer)
			} else if mLow >= HALF {
				mLow = 2 * ( mLow - HALF )
				mHigh = 2 * ( mHigh - HALF ) + 1
				mBuffer = 2 * ( mBuffer - HALF )
			} else {
				break
			}

			mBuffer += inbuf.readBit()
			mScale = 0
		}

		/* e3 mapping */
		for {
			if !( ( FIRST_QUARTER <= mLow ) && ( mHigh < THIRD_QUARTER ) ) {
				break
			}

			mScale++
			mLow = 2 * ( mLow - FIRST_QUARTER )
			mHigh = 2 * ( mHigh - FIRST_QUARTER ) + 1
			mBuffer = 2 * ( mBuffer - FIRST_QUARTER ) 
			mBuffer += inbuf.readBit()
		}
		
		/* 3. Update frequency table */
		freq[current]+=1
		total+=1
	}

	return buf.Bytes()
}