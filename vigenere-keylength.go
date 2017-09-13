/*
Author:     Shikha Fadnavis
Program:    Finding Key length of a Vigenere Cipher 
Date:       09/06/2017
*/

package main


import ("os"
	"sort"
	"fmt"

)

func gcd(var1 int, var2 int) int{
	var gcdvar int
	for i := 1; i <= var1 && i <= var2; i++{
		if var1%i == 0 && var2%i == 0{
			gcdvar = i
		}
	}

	return gcdvar

}

func main(){

	var ciphertextLen int 
	var coincArray []int
	var coincArrayCopy []int
	var keyLengthArray [4]int
	coincArray = make([]int, 100, 100)
	coincArrayCopy = make([]int, 100, 100)
	var min int = 9999 
	var keyLength, keyBit, factor int
	key := make([]string, 0)

	// Array with frequencies of english characters
	engFrequency := [26]float64{.082, .015, .028, .043, .127, .022, .020, .061, .070, .002, .008, .040, .024, .067, .075, .019, .001, .060, .063, .091, .028, .010, .023, .001, .020, .001}
	
	var index int = 0
	var coincArrayLen int  = 100
	var i,j, k int
	var numShifts int
	var modulus int

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
		min = 9999
		for j = 0; j < coincArrayLen; j++{
			if coincArray[i] == coincArrayCopy[j]{
				if min > j{
					min = j
					keyLengthArray[keyBit]  = j + 1
					keyBit += 1
					break
				}
				
			}
		}


	}

	factor = gcd(keyLengthArray[0], keyLengthArray[1])
	factor = gcd(keyLengthArray[2], factor)
	factor = gcd(keyLengthArray[3], factor)

	
	if factor == 1{
		keyLength = min + 1
	}else{
		keyLength = factor
	}

	
	println("\nKey Length: ", keyLength)



	/*================== Part 3 Solution ==================*/

	for i = 0; i < keyLength; i++{
		var freqPickArray [26]float64
		var prodSumArray = make([]float64, 26, 26)
		var prodSumArrayCopy = make([]float64, 26, 26)
		modulus = i 
		for j = 0; j < ciphertextLen; j++{
			if (j % keyLength == modulus) || (j % keyLength + 26 == modulus){
				//pickArray = append(pickArray, string(ciphertext[j]))
				num := int(ciphertext[j]) - 65
				freqPickArray[num] += 1
			}
		}
		
		for j = 0; j < 26; j++{
			var sum float64 = 0
			for k = 0; k < 26; k++{
				var modulo int = (k+j) % 26
				if modulo < 0{
					modulo = modulo + 26
				}
				sum = sum + (engFrequency[k] * freqPickArray[modulo])
			}
			
			prodSumArray[j] = sum
			prodSumArrayCopy[j] = sum
		}
		
		sort.Float64s(prodSumArray)
		for j = 0; j < 26; j++{
			if prodSumArray[25] == prodSumArrayCopy[j]{
				key = append(key, string(j + 65 -1))
			}
		}
	
			
 
	}

	
	fmt.Println(key)


}


