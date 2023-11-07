package util

import "golang.org/x/crypto/bcrypt"

/**
*@Description:
*@Author: BZ
*@date: 2023/11/4 16:10
*@Version: V1.0
 */

func GetHashPassword(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(fromPassword), err

}

func ComparePassword(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
