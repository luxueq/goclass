package week9

import (
	"bufio"
	"encoding/binary"
	"io"
)


type Conn struct {
	rwc     io.ReadWriteCloser
	r       *bufio.Reader
	w       *bufio.Writer
}

func decode()  {
	var (
		rwc io.ReadWriteCloser
		r   *bufio.Reader
		w   *bufio.Writer
		pkgLengthBytes []byte
	)
	con := newConn(rwc, r, w)

	con.r.Read(pkgLengthBytes)

	pkgLength := binary.BigEndian.Uint32(pkgLengthBytes)

	msg := make([]byte, 0, pkgLength - 4)
	con.r.Read(msg)

	headerLengthBytes := msg[0:2]

	headerLength := binary.BigEndian.Uint32(headerLengthBytes)

	header := msg[0:headerLength - 4]

	version := msg[2:4]

	op := msg[4:8]

	sid := msg[8:12]
	//头部剩余的数据
	extraHeader := msg[12:headerLength - 4]

	body := msg[headerLength - 4:]

	handle(version, op, sid, header, extraHeader, body)
}

func newConn(rwc io.ReadWriteCloser, r *bufio.Reader, w *bufio.Writer) *Conn {
	return &Conn{rwc: rwc, r: r, w: w}
}

func handle(...[]byte)  {

}