package enum

//KeyMap 输出模型
type KeyMap struct {
	Key string `json:"k"`
	Val int    `json:"v"`
}

//KeyMapSlice keymap sort
type KeyMapSlice []KeyMap

func (s KeyMapSlice) Len() int {
	return len(s)
}

func (s KeyMapSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s KeyMapSlice) Less(i, j int) bool {
	return s[i].Val < s[j].Val
}
