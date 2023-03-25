package graph

import "strconv"

// 文字列をint32に変換する
func ToInt32(s string) int32 {
	id, _ := strconv.Atoi(s)
	return int32(id)
}
