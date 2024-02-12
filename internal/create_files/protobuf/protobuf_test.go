package protobuf

import "testing"

func TestFindLastGoodPos(t *testing.T) {
	s := "123456\n789"
	pos1 := FindLastGoodPos(s, "5")
	if pos1 < 0 {
		t.Error("FindLastGoodPos() error: ", pos1)
	}
	s2 := s[pos1:]
	if s2 != "789" {
		t.Error("FindLastGoodPos() error: ", s2)
	}
}
