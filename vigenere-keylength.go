/*
Author:     Shikha Fadnavis
Program:    Finding Key length of a Vigenere Cipher 
Date:       09/06/2017
*/

package main

//import "fmt"
import "os"

func main(){

	var ciphertextLen int
	var coincArray [30]int
	var index int = 0
	var i int
	var numShifts int
	fi, err := os.Open(os.Args[1])
	if err != nil{
		panic(err)
	}
	
	ciphertext := make([]byte, 4096)
	ciphertextCopy := make([]byte, 4096)
	
	fi.Read(ciphertext)
	//fi.close()

	for i = 0; i < len(ciphertext); i++{
		if ciphertext[i] == 0{
			break
		}
	}
	ciphertextLen = i

	// Finding coincidence for upto 20 shifts : Sufficient tries

	for numShifts = 1; numShifts <= 20; numShifts++{
		count := 0
		
		for j, k := numShifts, 0; j < ciphertextLen; j, k = j+1, k+1{

			if k < (ciphertextLen - numShifts){
				print("hello")

				if int(ciphertext[j]) == int(ciphertextCopy[k]){
					print("hello")
					count += 1
				}
			}

		}
		println(count)

		coincArray[index] = count
		index += 1

	}

	for i = 0; i < 20; i++{

		println(coincArray[i])
	}	




}


