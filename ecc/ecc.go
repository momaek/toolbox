// Package ecc ...
// ECC 椭圆曲线，支持非对称加解密
// 抄 以太坊 的代码
package ecc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
	"os"
	"runtime"
)

const (
	eccPrivateKeyPrefix = "MOMAEK TOOLBOX PRIVATE KEY"
	eccPublicKeyPrefix  = "MOMAEK TOOLBOX PUBLIC KEY"
	eccPrivateFileName  = "momaek_toolbox_private.pem"
	eccPublicFileName   = "momaek_toolbox_public.pem"
)

// GenECCKeys 生成一个公钥和一个私钥
func GenECCKeys() error {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}

	x509PrivateKey, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}

	block := pem.Block{
		Type:  eccPrivateKeyPrefix,
		Bytes: x509PrivateKey,
	}
	file, err := os.Create(eccPrivateFileName)
	if err != nil {
		return err
	}
	defer file.Close()
	if err = pem.Encode(file, &block); err != nil {
		return err
	}

	x509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}

	publicBlock := pem.Block{
		Type:  eccPublicKeyPrefix,
		Bytes: x509PublicKey,
	}
	publicFile, err := os.Create(eccPublicFileName)
	if err != nil {
		return err
	}
	defer publicFile.Close()
	if err = pem.Encode(publicFile, &publicBlock); err != nil {
		return err
	}
	return nil
}

// Encrypt ecc 加密
func Encrypt(plainText, key []byte) (cryptText []byte, err error) {
	block, _ := pem.Decode(key)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	tempPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// Decode to get the private key in the ecdsa package
	publicKey1 := tempPublicKey.(*ecdsa.PublicKey)
	// Convert to the public key in the ecies package in the ethereum package
	publicKey := ImportECDSAPublic(publicKey1)
	crypttext, err := DoEncrypt(rand.Reader, publicKey, plainText, nil, nil)

	return crypttext, err

}

// Decrypt The private key and plaintext are passed in for decryption
func Decrypt(cryptText, key []byte) (msg []byte, err error) {
	block, _ := pem.Decode(key)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	tempPrivateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// Decode to get the private key in the ecdsa package
	// Convert to the private key in the ecies package in the ethereum package
	privateKey := ImportECDSA(tempPrivateKey)

	plainText, err := privateKey.Decrypt(cryptText, nil, nil)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

// EncryptToString encrypt to base64 encode string
func EncryptToString(plainText, key string) (encrypted string, err error) {
	msg, err := Encrypt([]byte(plainText), []byte(key))
	if err != nil {
		return
	}

	encrypted = base64.StdEncoding.EncodeToString(msg)
	return
}

// DecryptString ..
func DecryptString(cryptText, key string) (decrypted string, err error) {
	b, err := base64.StdEncoding.DecodeString(cryptText)
	if err != nil {
		return
	}

	msg, err := Decrypt(b, []byte(key))
	if err != nil {
		return
	}

	decrypted = string(msg)
	return
}
