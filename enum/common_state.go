package enum

import (
	"fmt"
	"sort"
)

// CommonState 通用状态
//	适用于只有 -1 0 两种状态
type CommonState int

const (
	// CommonStateDisable 禁用
	CommonStateDisable CommonState = iota - 1

	// CommonStateEnable 启用
	CommonStateEnable
)

var commonStateText = map[CommonState]string{
	CommonStateDisable: "禁用",
	CommonStateEnable:  "启用",
}

//Chinese 状态中文
func (m CommonState) Chinese() string {
	return commonStateText[m]
}

//List 列表输出
func (m CommonState) List() []KeyMap {
	km := make([]KeyMap, 0)
	for k, v := range commonStateText {
		km = append(km, KeyMap{Key: fmt.Sprintf("%v", v), Val: int(k)})
	}
	sort.Sort(KeyMapSlice(km))
	return km
}
