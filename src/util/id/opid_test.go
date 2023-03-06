package id

import (
	"testing"
)

func TestParseOpId(t *testing.T) {
	realArray := []uint64{1, 2}
	for i, v := range []string{"f1", "d2"} {
		opId, err := ParseOpId(v)
		if err != nil {
			t.Errorf(err.Error())
		}

		if opId.Real != realArray[i] {
			t.Errorf("解析出现错误: %q, ", opId.Id)
		}
	}
}

func TestParseOpIds(t *testing.T) {
	opIds, err := ParseOpIds("f1,d2")
	if err != nil {
		t.Errorf(err.Error())
	}

	realArray := []uint64{1, 2}
	for i, v := range opIds {
		if v.Real != realArray[i] {
			t.Errorf("解析出现错误: %q, ", v.Id)
		}
	}
}
