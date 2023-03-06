package id

import (
	"regexp"
	"strconv"
	"strings"
)

// 目录或文件的操作标识
type OpId struct {
	Id     string // 原始的标识字符串
	Real   uint64 // 真实的数据库标识
	IsDir  bool   // 是否目录
	IsRoot bool   // 是否根目录
}

const (
	regex = "^([f|d])(\\d+)$"
)

// 解析操作标识
func ParseOpId(str string) (OpId, error) {
	pattern := regexp.MustCompile(regex)

	g := pattern.FindStringSubmatch(str)
	t := g[1]
	id := g[2]

	isDir := t == "d"
	real, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		return OpId{}, err
	}

	return OpId{Id: str, Real: real, IsDir: isDir, IsRoot: false}, nil
}

// 解析操作标识
func ParseOpIds(str string) ([]OpId, error) {
	if !strings.Contains(",", str) {
		return []OpId{}, nil
	}

	var opIds []OpId

	ss := strings.Split(str, ",")
	for _, v := range ss {
		opId, err := ParseOpId(v)
		if err != nil {
			return []OpId{}, nil
		}

		opIds = append(opIds, opId)
	}

	return opIds, nil
}
