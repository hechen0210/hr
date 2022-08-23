package util

import "golang.org/x/crypto/bcrypt"

// 密码加密
func BcryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

// 密码验证
func VerifyPwd(password, bcryptPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(bcryptPwd), []byte(password))
	return err == nil
}
