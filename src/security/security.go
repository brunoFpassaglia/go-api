package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPasswd(passwd, hashedPasswd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasswd), []byte(passwd))
}
