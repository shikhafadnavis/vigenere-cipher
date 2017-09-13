/*
Author:     Shikha Fadnavis
Program:    Encryption using Vigenere Cipher 
Date:       09/02/2017
*/


package main


import "os"
import "bytes"

func main() {

	var vigenereKey string = os.Args[1]
	ciphertext := make([]string, 0)
	var imod int
	var i int
	kLen := len(vigenereKey)
	fi, err := os.Open(os.Args[2])
	if err != nil{
		panic(err)
	}
	fo, err := os.Create("output.txt")
	if err != nil{
		panic(err)
	}

	buf := make([]byte, 100000)
	newBuf := make([]byte, 100000)
	fi.Read(buf)
	buf = bytes.ToUpper(buf)
	//fmt.Println(string(buf))
	
	for i := 0; i < len(buf); i++{
		if((int(buf[i]) >= 65 && int(buf[i]) <= 90) || (int(buf[i]) >= 97 && int(buf[i]) <= 122)){
			fo.Write(buf[i:i+1])
			
		}

	}

	fi2, err2 := os.Open("output.txt")
	if err2 != nil{
		panic(err)
	}
	fi2.Read(newBuf)

	for i = 0; i < len(newBuf); i++{
		if newBuf[i] == 0{
			break
		}
	}
	newBufLen := i
	
	fcipher, errcipher := os.Create(os.Args[3])
	if errcipher != nil{
		panic(errcipher)
	}	

	for i := 0; i < newBufLen; i++ {

		imod = i % kLen
		//println(imod)
		var plainInt int = int(newBuf[i]) - 65
		var keyInt int = int(vigenereKey[imod]) - 65
		var cipMod int = (plainInt + keyInt + 1) % 26
		if cipMod < 0{
			cipMod = cipMod + 26
		}
		var cipLetter int = cipMod + 65 
		var cipherChar string = string(cipLetter)
		ciphertext = append(ciphertext, string(cipherChar))
		fcipher.WriteString(cipherChar)
		//fmt.Printf("%c", ciphertext[i+50])
		//fmt.Println(reflect.TypeOf(ciphertext[i]))
	}

	//fmt.Println(ciphertext)
	
	
	

}

