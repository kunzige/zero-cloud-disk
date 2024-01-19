package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// 加密

func Encrypt(message string) string {

	h := sha1.New()
	h.Write([]byte(message))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}

// 生成随机验证码

func RandomCode() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	random_code := ""
	rand_int := 0
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		rand_int = rand.Intn(62)
		random_code += str[rand_int : rand_int+1]
	}
	return random_code
}
