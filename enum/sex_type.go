package enum

import (
	"fmt"
	"sort"
)

// SexType 性别
type SexType int

const (
	SexTypeMale SexType = iota + 1
	SexTypeFemale
	SexTypeNone
)

var sexTypeText = map[SexType]string{
	SexTypeMale:   "男",
	SexTypeFemale: "女",
	SexTypeNone:   "未知",
}

//Chinese 状态中文
func (m SexType) Chinese() string {
	return sexTypeText[m]
}

//List 列表输出
func (m SexType) List() []KeyMap {
	km := make([]KeyMap, 0)
	for k, v := range sexTypeText {
		km = append(km, KeyMap{Key: fmt.Sprintf("%v", v), Val: int(k)})
	}
	sort.Sort(KeyMapSlice(km))
	return km
}
