package id

import (
	"testing"
)

func TestParse(t *testing.T) {
	realArray := []uint64{1, 2}
	for i, v := range []string{"f1", "d2"} {
		opId, err := ParseOpId(v)
		if err != nil {
			t.Errorf(err.Error())
		}

		if opId.Real != realArray[i] {
			t.Errorf("解析出现错误: %q", opId.Id)
		}
	}
}

func TestParseWithArray(t *testing.T) {
	opIds, err := ParseOpIds([]string{"f1", "d2"})
	if err != nil {
		t.Errorf(err.Error())
	}

	realArray := []uint64{1, 2}
	for i, v := range opIds {
		if v.Real != realArray[i] {
			t.Errorf("解析出现错误: %q", v.Id)
		}
	}
}

func TestParseWithNotSupportedExpr(t *testing.T) {
	var (
		str    = "hhh"
		errMsg = "不支持的表达式"
	)

	_, err := ParseOpId(str)
	if err == nil || err.Error() != errMsg {
		t.Errorf("解析出现错误: %q", str)
	}
}

func TestParseWithRootDir(t *testing.T) {
	var str = Personal

	opId, err := ParseOpId(str)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !opId.IsRoot || !opId.IsDir {
		t.Errorf("解析出现错误: %q", str)
	}
}
