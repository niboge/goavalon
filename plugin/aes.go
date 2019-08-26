package plugin

import (
	"avalon/app/model"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"strings"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(data string) ([]byte, error) {
	origData := []byte(data)
	key := key()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(data string) (string, error) {
	crypted := []byte(data)
	key := key()
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	retStr := string(origData)

	return retStr, nil
}

func key() []byte {
	return []byte("0123456789abcdefghijklmn")
	// result, err := AesEncrypt([]byte("hello world"), key)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(base64.StdEncoding.EncodeToString(result))
	// origData, err := AesDecrypt(result, key)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(origData))
}

/*
 go 的变量不能重复定义坑，处理格式要搞n个变量中间变量(却成了整个函数的变量)，带来了代码的混乱性。

 值的比较也很坑，我知道你是强类型，强个毛啊，我就想比较一个值而已，不想比较类型,学学php的===和c++的 opreator ==(int b),
 老告诉类型不一样，我特么在乎你是人是狗？我只在乎你们的出生日期是否一样！！！！为什么人的出生日期是int，
 狗的出生日期就得是string???可以不同，但是比较的时候要可以比较才行啊
 试问，啥时候给加一个~= 判断??

 多参数返回值的处理也略坑 用起来不爽

 其他还好
*/
func UserAuth(ticket string) (user *model.UserSt) {
	session, _ := NewRedis("")

	// aes decode
	split := strings.Split(ticket, ",")
	split[1] = split[1][:len(split[1])]
	ticket, _ = AesDecrypt(split[1])

	// get user from session
	split = strings.Split(ticket, "-")
	res, _ := session.Get("UserAuth:" + string(split[0]))

	user = new(model.UserSt)
	if jsonErr := json.Unmarshal(res.([]byte), user); jsonErr != nil {
		fmt.Printf("[!ERROR!] %v %v \n", user, jsonErr)
		return nil
	}
	// *user = res.(model.UserSt)

	// verify login_time
	if strings.Compare(string(user.LoginTime), split[1]) == 0 {
		fmt.Println("[!ERROR!] cmp fail \n")
		return nil
	}

	return
}
