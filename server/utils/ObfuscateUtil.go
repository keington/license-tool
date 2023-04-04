/*
 * package utils 工具包，提供了一些常用的工具函数
 */
package utils

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"unicode"
)

/*
 * ObfuscationUtil 是一个字节数组混淆工具函数，用于将输入的字节数组与签名代码进行异或运算并进行加密
 * @params: input []byte - 需要加密的字节数组
 * 			signatureCode string - 签名代码
 * @returns: string - 加密后的字符串
 * 			 error - 如果签名代码为空，则返回一个错误
 */
func ObfuscationUtil(input []byte, signatureCode string) (string, error) {

	// targetLen 目标字符常量 - 4096
	const targetLen = 4096

	// 判断特征码是否为空
	if signatureCode == "" {
		return "", errors.New("signature code is empty")
	}

	scBytes := []byte(signatureCode)
	ciphertext := make([]byte, targetLen)
	for i := 0; i < targetLen; i++ {
		if i < len(input) {
			ciphertext[i] = input[i] ^ scBytes[i%len(scBytes)] // 将输入的字节数组与签名代码进行异或运算，生成密文
		} else {
			randByte := byte(rand.Intn(256))
			for randByte != '0' && randByte != '1' && !unicode.IsLetter(rune(randByte)) {
				randByte = byte(rand.Intn(256))
			}
		}
	}

	return base64.URLEncoding.EncodeToString(ciphertext), nil // 将密文进行 URL 编码并返回
}
