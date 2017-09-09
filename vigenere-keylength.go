/*
Author:     Shikha Fadnavis
Program:    Finding Key length of a Vigenere Cipher 
Date:       09/06/2017
*/

package main

//import "fmt"
import ("os"
	"sort"

)

func main(){

	var ciphertextLen int
	coincArray := []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	coincArrayCopy := []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
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
	ciphertextCopy = ciphertext
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
				
				if int(ciphertext[j]) == int(ciphertextCopy[k]){
					count += 1
				}
			}

		}

		coincArray[index] = count
		coincArrayCopy[index] = count
		index += 1

	}

	sort.Ints(coincArray)
	
	// Unsorted Array of coincidences
	for i = 0; i < 20; i++{
		println(coincArrayCopy[i])
	}
	println("Hello")
	println(coincArrayCopy[7])

	for i = 19; i > 15; i--{
		println(coincArray[i])
		for j := 0; j < 20; j++{
			if coincArray[i] == coincArrayCopy[j]{
				println("\nIndex: ")
				println(j)
			}
		}
	}


}


