package id

import (
	"errors"
	"regexp"
	"strconv"
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
	// 如果是根目录
	if str == Personal {
		return OpId{Id: str, Real: 0, IsDir: true, IsRoot: true}, nil
	}

	pattern := regexp.MustCompile(regex)
	if !pattern.MatchString(str) {
		return OpId{}, errors.New("不支持的表达式")
	}

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
func ParseOpIds(ss []string) ([]OpId, error) {
	var opIds []OpId

	for _, v := range ss {
		opId, err := ParseOpId(v)
		if err != nil {
			return nil, err
		}

		opIds = append(opIds, opId)
	}

	return opIds, nil
}
