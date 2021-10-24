
/**
	加密、解密相关方法

	生成随机数
	aes-ecb
	RSA
	Sha256
	HMAC
 */
package secret

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
)

/**
	生成随机数后转为16进制
 */
func RandToken(num int) string {
	b := make([]byte, num)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

/**
	RSA加密
 */
func RsaEncrypt(originalData, publicKey []byte) ([]byte, error) {
	pubKey, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey),originalData)
	if err != nil {
		return nil, err
	}
	return encryptedData,err
}

/**
	RSA解密
 */
func RsaDecrypt(encryptedData, privateKey []byte) ([]byte, error) {
	prvKey, err := x509.ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	originalData, err := rsa.DecryptPKCS1v15(rand.Reader, prvKey, encryptedData)
	if err != nil {
		return nil, err
	}
	return originalData,err
}

