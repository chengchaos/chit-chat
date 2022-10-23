package data

import "testing"

func TestCreateUUID(t *testing.T) {
	uuid := CreateUUID()
	println("uuid =>", uuid)
}

func TestEncrypt(t *testing.T) {
	plainText := CreateUUID()
	t.Logf("plainText => %s\n", plainText)
	encryptText := Encrypt(plainText)
	t.Logf("encryptText => %s\n", encryptText)
}
