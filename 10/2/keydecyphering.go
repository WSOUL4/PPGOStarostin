package main

import (
	"bufio"
	"crypto/aes"
	"encoding/hex"
	"fmt"

	//"strconv"

	"os"
)

func EncryptAES(key []byte, plaintext string) string {
	// создание шифра
	c, err := aes.NewCipher(key)
	if err != nil {
		os.Exit(1)
	}
	// выделение места для зашифрованных данных
	out := make([]byte, len(plaintext))
	// шифрование
	c.Encrypt(out, []byte(plaintext))
	// возврат шестнадцатеричной строки
	return hex.EncodeToString(out)
}
func aescypher(pt string, key string) string {
	// ключ шифрования

	c := EncryptAES([]byte(key), pt)
	return c
}

/*
	func DecryptAES(key []byte, plaintext string) string {
		// создание шифра
		c, err := aes.NewCipher(key)
		if err != nil {
			os.Exit(1)
		}
		// выделение места для зашифрованных данных
		out := make([]byte, len(plaintext))
		// шифрование
		c.Decrypt(out, []byte(plaintext))
		// возврат шестнадцатеричной строки
		return hex.EncodeToString(out)
	}
*/
func DecryptAES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)
	c, err := aes.NewCipher(key)
	if err != nil {
		os.Exit(1)
	}
	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
	s := string(pt)
	//fmt.Println("DECRYPTED:", s)
	return s
}

func aesdecypher(pt string, key string) string {
	// ключ шифрования

	c := DecryptAES([]byte(key), pt)
	return c
}
func ReadUser() (str string) {

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	//fmt.Printf("\n--\t%s\t--\n", in.Text())
	str = fmt.Sprintf("%s", in.Text())
	return str
	//line = strings.TrimSuffix(line, "\n")
	//fmt.Println(line)
	//fmt.Scanln(&str)
	//fmt.Println("ReadUser:", str)
	//return line
}
func main() {
	fmt.Println("Ключ:")
	key := ReadUser()
	//key := "thisis32bitlongpassphraseimusing"

	// открытый текст
	fmt.Println("Что зашифровать:")
	pt := ReadUser()
	k := aescypher(pt, key)

	fmt.Println(k)

	fmt.Println("Ключ:")
	key = ReadUser()
	//key := "thisis32bitlongpassphraseimusing"

	// открытый текст
	fmt.Println("Что расшифровать:")
	pt = ReadUser()
	k = aesdecypher(pt, key)
	fmt.Println(k)

}
