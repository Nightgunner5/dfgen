package dfgen

import (
	"bufio"
	"io"
)

/*

Integers are encoded in little-endian using 1, 2, 4, or 8 bytes each.

Lists contain a 4 byte integer denoting length, followed by the contents of the list. (Lists are not included in this library because they can be recursive lists of lists of lists)

Strings contain a 2 byte integer denoting length, followed by the bytes forming the string. Strings are CP437 with no null terminator.

*/
type Reader struct {
	R *bufio.Reader
}

func (r Reader) ReadByte() (ret uint8, err error) {
	var buf [1]byte
	_, err = io.ReadFull(r.R, buf[:])

	ret = uint8(buf[0])

	return
}

func (r Reader) ReadShort() (ret uint16, err error) {
	var buf [2]byte
	_, err = io.ReadFull(r.R, buf[:])

	ret = uint16(buf[0]) | uint16(buf[1])<<8

	return
}

func (r Reader) ReadLong() (ret uint32, err error) {
	var buf [4]byte
	_, err = io.ReadFull(r.R, buf[:])

	ret = uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24

	return
}

func (r Reader) ReadLongLong() (ret uint64, err error) {
	var buf [8]byte
	_, err = io.ReadFull(r.R, buf[:])

	ret = uint64(buf[0]) | uint64(buf[1])<<8 | uint64(buf[2])<<16 | uint64(buf[3])<<24 | uint64(buf[4])<<32 | uint64(buf[5])<<40 | uint64(buf[6])<<48 | uint64(buf[7])<<56

	return
}

var cp437 = []rune("\000☺☻♥♦♣♠•◘○◙♂♀♪♫☼►◄↕‼¶§▬↨↑↓→←∟↔▲▼ !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~⌂ÇüéâäàåçêëèïîìÄÅÉæÆôöòûùÿÖÜ¢£¥₧ƒáíóúñÑªº¿⌐¬½¼¡«»░▒▓│┤╡╢╖╕╣║╗╝╜╛┐└┴┬├─┼╞╟╚╔╩╦╠═╬╧╨╤╥╙╘╒╓╫╪┘┌█▄▌▐▀αßΓπΣσµτΦΘΩδ∞φε∩≡±≥≤⌠⌡÷≈°∙·√ⁿ²■\xA0")

func (r Reader) ReadString() (ret string, err error) {
	length, err := r.ReadShort()

	if err == nil {
		buf := make([]byte, length)
		rbuf := make([]rune, length)
		_, err = io.ReadFull(r.R, buf)

		for i := range buf {
			rbuf[i] = cp437[buf[i]]
		}

		ret = string(rbuf)
	}

	return
}
