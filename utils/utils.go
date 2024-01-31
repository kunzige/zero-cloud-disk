package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

// 获取UUID

func GetUuid() string {
	// 生成一个新的 UUID
	newUUID := uuid.New()

	// 将 UUID 转换为字符串表示
	uuidStr := newUUID.String()
	return uuidStr
}

// 生成token

func NewJwtToken(userEmail string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["user_email"] = userEmail
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("2418eb42-6991-4039-91c3-7f25873c5ad1"))
}
