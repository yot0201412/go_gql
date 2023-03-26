package graph

import "strconv"

// 文字列をint32に変換する
func ToInt32(s string) int32 {
	id, _ := strconv.Atoi(s)
	return int32(id)
}

// Int32を文字列に変換する
func ToString(i int32) string {
	return strconv.Itoa(int(i))
}
