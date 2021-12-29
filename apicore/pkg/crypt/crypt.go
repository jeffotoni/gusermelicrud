// Back-End in Go server
// @jeffotoni

package crypt

import (
	// #nosec
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/jeffotoni/gusermeli/apicore/pkg/fmts"

	"golang.org/x/crypto/bcrypt"
)

var (
	SHA1_SALT = "$10.ytcjjju#@90001.x9wo#@356xw2."
)

//Token..
func Token(key string) (string, error) {
	if len(key) == 0 {
		return "", errors.New("key cannot be empty")
	}
	hash := strconv.Itoa(Random(1, 1000))
	hash_key := fmts.ConcatStr(key, "$#%_", hash)
	return GSha1(hash_key), nil
}

//Random ..
func Random(min, max int) int {
	if max < min {
		return 0
	}

	rand.Seed(time.Now().UnixNano())
	// #nosec
	rdn := rand.Intn(max-min) + min
	return rdn
}

//Blowfish gera o hash da senha
func Blowfish(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password cannot be empty")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckBlowfish checa o hash com a senha
func CheckBlowfish(password, hash string) error {
	if len(password) == 0 {
		return errors.New("password cannot be empty")
	}
	if len(hash) == 0 {
		return errors.New("hash cannot be empty")
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GSha1(key string) string {
	if len(key) == 0 {
		return ""
	}
	data := []byte(fmts.ConcatStr(key, SHA1_SALT))
	// #nosec
	sh1 := sha1.Sum(data)
	return fmt.Sprintf("%x", sh1)
}
