package ecc

import (
	"testing"
)

var (
	privateKey = `-----BEGIN CHANNEL PRIVATE KEY-----
MHcCAQEEIEVqmkt+cP4DgX6qH6Po63W+y1n6vqpyUcclSfdYKgDToAoGCCqGSM49
AwEHoUQDQgAEtU3nNWY//L2bv26kuV/cBF9E9s832oO38E8qZyRhAdOums8XL4SQ
Qe063dmj5VjMXQg5cRekDjzMTVvHzGMdBA==
-----END CHANNEL PRIVATE KEY-----`
	publicKey = `-----BEGIN CHANNEL PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEtU3nNWY//L2bv26kuV/cBF9E9s83
2oO38E8qZyRhAdOums8XL4SQQe063dmj5VjMXQg5cRekDjzMTVvHzGMdBA==
-----END CHANNEL PUBLIC KEY-----`
)

func TestEncryptAndDecrypt(t *testing.T) {
	plainTest := "今天礼拜五"
	msg, err := Encrypt([]byte(plainTest), []byte(publicKey))
	if err != nil {
		t.Fatal("encrypt failed ", err)
	}

	pl, err := Decrypt(msg, []byte(privateKey))
	if err != nil {
		t.Fatal("decrypt failed ", err)
	}

	if string(pl) != plainTest {
		t.Fatal("enc, dec failed")
	}

	base64S, err := EncryptToString(plainTest, publicKey)
	if err != nil {
		t.Fatal("EncryptToString failed ", err)
	}

	pp, err := DecryptString(base64S, privateKey)
	if err != nil {
		t.Fatal("DecryptString failed ", err)
	}

	if pp != plainTest {
		t.Fatal("enc,dec string failed")
	}
}
