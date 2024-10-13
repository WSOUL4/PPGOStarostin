package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"strings"

	//"strconv"
	"bufio"
	"os"

	_ "github.com/lib/pq"
)

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
func GetSha256(str string) string {
	reader := strings.NewReader(str)
	/*data := make([]byte, len(str))
	  n, err := reader.Read(data)
	  if err != nil {
	      fmt.Println("Error reading from reader:", err)
	      return
	  }*/
	h := sha256.New()
	if _, err := io.Copy(h, reader); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("sha256: %x\n", h.Sum(nil))
	return hex.EncodeToString(h.Sum(nil)[:])
}
func GetSha512(str string) string {

	sha_512 := sha512.New()
	sha_512.Write([]byte(str))

	fmt.Printf("sha512: %x\n", sha_512.Sum(nil))
	return hex.EncodeToString(sha_512.Sum(nil)[:])
}
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	fmt.Printf("md5: %x\n", hash[:])
	return hex.EncodeToString(hash[:])
}
func main() {
	fmt.Println("Введите строку:")
	str1 := ReadUser()

	fmt.Println("1.sha512")
	fmt.Println("2.sha256")
	fmt.Println("3.md5")
	fmt.Println("4.Проверить цел-сть")

	str2 := ReadUser()
	switch str2 {
	case "1":
		GetSha512(str1)
	case "2":
		GetSha256(str1)
	case "3":
		GetMD5Hash(str1)
	case "4":
		fmt.Println("Введите хэш той строки:")
		str1h := ReadUser()
		fmt.Println("1.Это sha512")
		fmt.Println("2.Это sha256")
		fmt.Println("3.Это md5")
		str1hw := ReadUser()

		switch str1hw {
		case "1":
			if str1h == GetSha512(str1) {
				fmt.Println("OK")
			} else {
				fmt.Println("BAD")
			}

		case "2":
			if str1h == GetSha256(str1) {
				fmt.Println("OK")
			} else {
				fmt.Println("BAD")
			}
		case "3":
			if str1h == GetMD5Hash(str1) {
				fmt.Println("OK")
			} else {
				fmt.Println("BAD")
			}
		}
	case "exit":
		os.Exit(1)
	}

}
