package util

import (
	"math/rand"
	"time"
)

// 返回n个随机字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")
	//给result分配一个长度为10的byte数组
	result := make([]byte, n)
	// rand.Seed(time.Now().Unix())
	rand.NewSource(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
