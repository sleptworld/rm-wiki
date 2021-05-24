package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"fmt"
	"github.com/goinggo/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
)

func mapToStruct(m map[string]interface{},s interface{}) (error){
	if err := mapstructure.Decode(m,s);err != nil{
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}

func IsContain(items []string , item string) bool{
	for _,eachitem := range items{
		if eachitem == item{
			return true
		}
	}
	return false
}

func FileIsExist(path string) bool{
	_,err := os.Stat(path)
	if err != nil{
		if os.IsExist(err){
			return true
		}
		return false
	}
	return true
}

func PwdEncrypt(p string,aesKey []byte) ([]byte,error){
	s512 := sha512.Sum512([]byte(p))
	b,err := bcrypt.GenerateFromPassword(s512[:],bcrypt.MinCost)
	if err != nil{
		return nil,err
	}
	encrypt, err := Encrypt(b, aesKey)
	if err != nil {
		return nil,err
	}
	return encrypt,nil
}

func PwdConfirm(s string,d []byte,k []byte) error {
	s512 := sha512.Sum512([]byte(s))
	encrypt ,err := Decrypt(d,k)
	if err != nil{
		return err
	}
	err = bcrypt.CompareHashAndPassword(encrypt,s512[:])
	return err
}

func PwdDecrypt(p string,aesKey []byte) bool {
	decrypt,_:= Decrypt([]byte(p),aesKey)
	s512 := sha512.Sum512([]byte(p))

	err := bcrypt.CompareHashAndPassword(decrypt, s512[:])
	if err != nil {
		return false
	} else {
		return true
	}

}
func Encrypt(plantText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	//选择加密算法
	if err != nil {
		return nil, err
	}
	plantText = PKCS7Padding(plantText, block.BlockSize())
	blockModel := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	ciphertext := make([]byte, len(plantText))
	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func Decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	//选择加密算法
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS7UnPadding(plantText, block.BlockSize())
	return plantText, nil
}

func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func StringTypeCheck(p string, t string) bool {
	switch t {
	case "Num":
		pattern := "\\d+"
		result,_ := regexp.MatchString(pattern,p)
		return result
	default:
		return false
	}

}