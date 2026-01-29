package main

import("math"; "strconv")

func Abs(n float64) float64{
	if n<0{
		return (-1)*n
	}
	return n
}

func All(arr []bool) bool{
	for _, val := range arr{
		if val==false{
			return false
		}
	}
	return true
}

func Any(arr []bool) bool{
	for _, val := range arr{
		if val==true{
			return true
		}
	}
	return false
}

func GetIndexOf(text string, pattern string) []int{
	var length int = len(pattern)
	var positions = []int{}
	for counter:=0; counter<=len(text)-length; counter++{
		if string(text[counter:counter+length])==pattern{
			positions = append(positions, counter)
		}
	}
	return positions
}

//DEPENDS ON: Reverse
func IsPalendrome(n int) bool{
	var nstr string = strconv.Itoa(n)
	if Reverse(nstr)==nstr{
		return true
	}
	return false
}

func IsPrime(n int) bool{
	var upper int = int(math.Sqrt(float64(n)))+1
	for factor:=2; factor<upper; factor++{
		if n%factor==0{
			return false
		}
	}
	return true
}

func Max(arr []int) int{
	var largest int = arr[0]
	for _, val := range arr{
		if val>largest{
			largest=val
		}
	}
	return largest
}

func Min(arr []int) int{
	var smallest int = arr[0]
	for _, val := range arr{
		if val<smallest{
			smallest=val
		}
	}
	return smallest
}

func NumOfDivs(x int) int{
	var counter int = 0
	var upper int = int(math.Sqrt(float64(x)))+1
	for factor:=1; factor<upper; factor++{
		if x%factor==0{
			counter += 1
		}
	}
	counter *= 2
	if int((upper-1)*(upper-1))==x{
		counter -= 1
	}
	return counter
}

//DEPENDS ON: GetIndexOf
func Replace(text string, pattern string, newPattern string) string{
	positions := GetIndexOf(text, pattern)
	var result string = ""
	var counter int = 0
	for _, val := range positions{
		result += string(text[counter:val])
		result += newPattern
		counter = val + len(pattern)
	}
	if counter!=len(text){
		result += string(text[counter:len(text)])
	}
	return result

}

func Reverse(nstr string) string{
	var result string = ""
	for counter := len(nstr)-1; counter>=0; counter--{
		result += string(nstr[counter])
	}
	return result
}
