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
	var coincArray []int
	var coincArrayCopy []int
	coincArray = make([]int, 100, 100)
	coincArrayCopy = make([]int, 100, 100)
	var min int = 9999 
	var keyLength int

	// Array with frequencies of english characters
	//var engFrequency [26]int = {.082, .015, .028, .043, .127, .022, .020, .061, .070, .002, .008, .040, .024, .067, .075, .019, .001, .060, .063, .091, .028, .010, .023, .001, .020, .001}
	
	var index int = 0
	var coincArrayLen int  = 100
	var i int
	var numShifts int

	fi, err := os.Open(os.Args[1])
	if err != nil{
		panic(err)
	}
	
	ciphertext := make([]byte, 20000)
	
	fi.Read(ciphertext)

	for i = 0; i < len(ciphertext); i++{
		if ciphertext[i] == 0{
			break
		}
	}
	ciphertextLen = i

	// Finding coincidence for upto 100 shifts : Sufficient tries

	for numShifts = 1; numShifts <= coincArrayLen; numShifts++{
		count := 0
		
		for i = numShifts; i < (ciphertextLen - numShifts -1); i++{
				
			if int(ciphertext[i]) == int(ciphertext[i-numShifts]){
				count += 1
			}

		}

		coincArray[index] = count
		coincArrayCopy[index] = count
		index += 1

	}

	sort.Ints(coincArray)
	
	// Unsorted Array of coincidences
	for i = 0; i < coincArrayLen; i++{
		//println(coincArrayCopy[i])
	}

	

	for i = (coincArrayLen-1); i > (coincArrayLen-5); i--{
		//println(coincArray[i])
		for j := 0; j < coincArrayLen; j++{
			if coincArray[i] == coincArrayCopy[j]{
				//println("\nIndex: ")
				//println(j)
				if min > j{
					min = j
					break
				}
				
			}
		}
	}


	keyLength = min + 1
	println("\nKey Length: ", keyLength)
	
	


}


