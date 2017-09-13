/*
Author:     Shikha Fadnavis
Program:    Decryption using Vigenere Cipher 
Date:       09/03/2017
*/


package main


import "os"
import "bytes"

func main() {

	var vigenereKey string = os.Args[1]
	plaintext := make([]string, 0)
	var imod int
	var i int
	kLen := len(vigenereKey)
	fi, err := os.Open(os.Args[2])
	if err != nil{
		panic(err)
	}
	

	buf := make([]byte, 100000)
	fi.Read(buf)
	buf = bytes.ToUpper(buf)
	//fmt.Println(string(buf))
	

	for i = 0; i < len(buf); i++{
		if buf[i] == 0{
			break
		}
	}
	bufLen := i
	//println(bufLen)
	
	fplain, errplain := os.Create(os.Args[3])
	if errplain != nil{
		panic(errplain)
	}	

	for i := 0; i < bufLen; i++ {

		imod = i % kLen
		//println(imod)
		var plainInt int = int(buf[i]) - 65
		var keyInt int = int(vigenereKey[imod]) - 65
		var plainMod int = (plainInt - keyInt -1) % 26
		if plainMod < 0{
			plainMod = plainMod + 26
		}
		var plainLetter int = plainMod + 65 
		var plainChar string = string(plainLetter)
		plaintext = append(plaintext, string(plainChar))
		fplain.WriteString(plainChar)
		//fmt.Printf("%c", ciphertext[i+50])
		//fmt.Println(reflect.TypeOf(ciphertext[i]))
	}

	//fmt.Println(plaintext)
	
	
	

}

