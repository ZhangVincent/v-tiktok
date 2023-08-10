// @description 加密
// @author zkp15
// @date 2023/8/10 9:04
// @version 1.0

package bcrypt

import "golang.org/x/crypto/bcrypt"

func EncodePassword(rawPassword string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	return string(hash)
}

func ValidatePassword(encodePassword, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(inputPassword))
}
