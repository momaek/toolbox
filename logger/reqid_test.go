package logger

import (
	"testing"
)

func TestGenReqID(t *testing.T) {
	reqID1 := GenReqID()
	reqID2 := GenReqID()
	if reqID1 == reqID2 {
		t.Fatal("Gen same reqid")
	}
}
