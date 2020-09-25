package logger

import (
	"encoding/base64"
	"encoding/binary"
	"time"
)

var (
	pid = uint32(time.Now().UnixNano() % 4294967291)
)

// GenReqID generate request id
func GenReqID() string {
	var b [12]byte
	binary.LittleEndian.PutUint32(b[:], pid)
	binary.LittleEndian.PutUint64(b[4:], uint64(time.Now().UnixNano()))
	return base64.URLEncoding.EncodeToString(b[:])
}
