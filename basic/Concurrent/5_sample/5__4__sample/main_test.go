package ___4__sample

import (
	"bytes"
	"log"
	"testing"
)

func Test_Main2(t *testing.T) {
	var s bytes.Buffer
	log.SetOutput(&s)
	log.SetFlags(0)
	main()

	if s.String() != "3\n" {
		t.Error(s.String())
	}
}
