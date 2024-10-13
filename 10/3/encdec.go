package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func publenc() []byte {
	publicKeyPEM, err := ioutil.ReadFile("./gen2/public.pem") //ШИФРОВКА(Загразука ключа и его загрузка в переменную)
	if err != nil {
		panic(err)
	}
	publicKeyBlock, _ := pem.Decode(publicKeyPEM)
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	//ШИФРОВКА
	plaintext := []byte("FIRST SENDS HIS REGARDS ABOUT YOUR MOTHER HAHAHA")                    //ПЕРВЫЙ ПОСЫЛАЕТ
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), plaintext) //ОПЯТЬ ЖЕ ПЕРВЫЙ ПАРАМЕТР МОЖНО СВОЙ РИДЕР ВЗЯТЬ, не безопасно
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encrypted: %x\n", ciphertext)
	return ciphertext
}
func privdec(ct []byte) { //СЕРВЕР БЕРЁТ ЗАКРЫТЫЙ КЛЮЧ
	privateKeyPEM, err := ioutil.ReadFile("./gen2/private.pem")
	if err != nil {
		panic(err)
	}
	privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	//ДЕШИФРОВКА
	//ciphertext := []byte{0x88, 0xaa, 0x63, 0x24, 0x2d, 0x48, 0xfd, 0xb1, 0x63, 0x71, 0x33, 0x17, 0x2a, 0x01, 0xce, 0x15, 0x1b, 0x25, 0xac, 0xcd, 0x35, 0xc1, 0x7c, 0x2a, 0x48, 0x58, 0x79, 0xae, 0x73, 0xf3, 0x5e, 0xc9, 0x89, 0xa7, 0x8a, 0x92, 0xa4, 0x3f, 0x3d, 0xb3, 0x43, 0x1d, 0x01, 0x74, 0xee, 0xd1, 0x1e, 0x95, 0x2b, 0x4f, 0x42, 0x46, 0x0b}
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ct) //СЕРВЕР ПРИНЯЛ ЗАШИФРОВАННЫЙ КЛЮЧОМ ШИФР И ИСПОЛЬЗУЕТ СВОЙ УЛЬТИМАТИВНЫЙ ЗАКР КЛЮЧ
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decrypted: %s\n", plaintext)
}
func main() {
	fmt.Println()
	et := publenc()
	privdec(et)

}
