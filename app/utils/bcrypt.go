package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"eyesStars/app/common"
	"eyesStars/global"
	"golang.org/x/crypto/bcrypt"
	"log"
)

/**
 * @author eyesYeager
 * @date 2023/1/29 10:35
 */

/*
 ***************************************************************************
                                  单向散列
 ***************************************************************************
*/

// BCryptPsdMake 密码加密
func BCryptPsdMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// BCryptPsdCheck 密码校验
func BCryptPsdCheck(pwd []byte, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	return err == nil
}

/*
 ***************************************************************************
                                  双向加密
 ***************************************************************************
*/

// AesEncrypt AES(CBC)加密
func AesEncrypt(str string) string {
	origData := []byte(str)
	aesKey := global.Config.Program.AesKey
	k := []byte(aesKey)
	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()
	origData = sPKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	crypt := make([]byte, len(origData))
	blockMode.CryptBlocks(crypt, origData)
	return base64.StdEncoding.EncodeToString(crypt)
}

// AesDecrypt AES(CBC)解密
func AesDecrypt(str string) (err error, decrypt string) {
	cratedByte, _ := base64.StdEncoding.DecodeString(str)
	aesKey := global.Config.Program.AesKey
	k := []byte(aesKey)

	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	orig := make([]byte, len(cratedByte))
	blockMode.CryptBlocks(orig, cratedByte)
	orig = sPKCS7UnPadding(orig)

	defer func() {
		if aesErr := recover(); aesErr != nil {
			err = common.CustomError{}.SetErrorMsg("AES decryption error")
		}
	}()
	return err, string(orig)
}

// 补码
func sPKCS7Padding(ciphered []byte, blockSize int) []byte {
	padding := blockSize - len(ciphered)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphered, padText...)
}

// 去码
func sPKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
