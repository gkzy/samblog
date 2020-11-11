package enum

import (
	"fmt"
	"sort"
)

// ArticleType 文章开放类型
type ArticleType int

const (
	// CommonStateDisable 禁用
	ArticleTypePrivate ArticleType = iota + 1

	ArticleTypeInside

	ArticleTypePublic
)

var articleTypeText = map[ArticleType]string{
	ArticleTypePrivate: "私人",
	ArticleTypeInside:  "内部",
	ArticleTypePublic:  "公开",
}

//Chinese 状态中文
func (m ArticleType) Chinese() string {
	return articleTypeText[m]
}

//List 列表输出
func (m ArticleType) List() []KeyMap {
	km := make([]KeyMap, 0)
	for k, v := range articleTypeText {
		km = append(km, KeyMap{Key: fmt.Sprintf("%v", v), Val: int(k)})
	}
	sort.Sort(KeyMapSlice(km))
	return km
}
