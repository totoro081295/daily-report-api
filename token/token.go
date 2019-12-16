package token

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type tokenHandler struct{}

// NewTokenHandler mount token handler
func NewTokenHandler() Handler {
	return &tokenHandler{}
}

// Handler token handler interface
type Handler interface {
	RandToken() string
	GenerateJWT(id uuid.UUID, valid bool) (string, error)
	LoadJWTPublicKeys() (*rsa.PublicKey, error)
}

// GenerateJWT トークン生成
func (h *tokenHandler) GenerateJWT(id uuid.UUID, valid bool) (string, error) {
	t := jwt.New(jwt.SigningMethodRS256)

	// Set claims
	claims := t.Claims.(jwt.MapClaims)
	claims["iss"] = "daily-report"
	claims["scopes"] = "api:access"
	claims["sub"] = id
	atoi, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRES_MIN"))
	claims["exp"] = time.Now().Add(time.Second * time.Duration(atoi)).Unix() // 30分
	claims["iat"] = time.Now().Unix()
	if valid {
		claims["aud"] = valid
	}

	// カレントディレクトり取得
	p, _ := os.Getwd()
	fileKey, err := ioutil.ReadFile(p + "/assets/jwt-key.rsa")
	if err != nil {
		return "", err
	}
	// 秘密鍵をパース
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(fileKey)
	if err != nil {
		return "", err
	}
	// Generate encoded token and send it as response.
	tokenString, err := t.SignedString(privKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
func (h *tokenHandler) LoadJWTPublicKeys() (*rsa.PublicKey, error) {
	keyFile, err := ioutil.ReadFile("./assets/jwt-key.rsa.pub")
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to load key %s: %s", keyFile, err)
	}

	return key, nil
}

var randSrc = rand.NewSource(time.Now().UnixNano())

const (
	rsLetters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rsLetterIdxBits = 6
	rsLetterIdxMask = 1<<rsLetterIdxBits - 1
	rsLetterIdxMax  = 63 / rsLetterIdxBits
)

// RandToken 64文字のトークン生成
func (h *tokenHandler) RandToken() string {
	rand := make([]byte, 64)
	cache, remain := randSrc.Int63(), rsLetterIdxMax
	for i := 64 - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rsLetterIdxMax
		}
		idx := int(cache & rsLetterIdxMask)
		if idx < len(rsLetters) {
			rand[i] = rsLetters[idx]
			i--
		}
		cache >>= rsLetterIdxBits
		remain--
	}
	return string(rand)
}
