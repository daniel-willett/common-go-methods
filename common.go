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

func BubbleSort(arr []int) []int{
	var length int = len(arr)
	var changesMade bool = true
	for changesMade==true{
		changesMade = false
		for index:=0; index<length-1; index++{
			if arr[index+1]<arr[index]{
				changesMade = true
				temp := arr[index]
				arr[index] = arr[index+1]
				arr[index+1] = temp
			}
		}
	}
	return arr
}

func Extend(result []int, add []int) []int{
	for i:=0; i<len(add); i++{
		result = append(result, add[i])
	}
	return result
}

//DEPENDS ON: Abs
func Gcd(a int, b int) int{
	var larger, smaller int = 0, 0
	if a>b {
		larger = a
		smaller = b
	} else {
		larger = b
		smaller = a
	}
	var temp int = 0
	for smaller!=0 {
		temp = smaller
		smaller = larger % smaller
		larger = temp
	}
	return int(Abs(float64(larger)))
}

func GetIndexOf(text string, pattern string) []int{
	var length int = len(pattern)
	var positions = []int{}

	if length==0{
		return positions
	}

	for counter:=0; counter<=len(text)-length; counter++{
		if string(text[counter:counter+length])==pattern{
			positions = append(positions, counter)
			counter += length-1
		}
	}
	return positions
}

func InsertionSort(arr []int) []int{
	var i int = 1
	var length int = len(arr)
	for i < length{
		var j int = i
		for j>0 && arr[j-1]>arr[j]{
			var temp int = arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = temp

			j -= 1
		}
		i += 1
	}
	return arr
}

//DEPENDS ON: Reverse
func IsPalendrome(n uint) bool{
	var nstr string = strconv.Itoa(int(n))
	if Reverse(nstr)==nstr{
		return true
	}
	return false
}

func IsPrime(n int) bool{
	if n<2{
		return false
	}
	var upper int = int(math.Sqrt(float64(n)))+1
	for factor:=2; factor<upper; factor++{
		if n%factor==0{
			return false
		}
	}
	return true
}

//DEPENDS ON: Split, Replace
func IsValidIPv4(ip string) bool{
	isValidOctet := func(x string) bool{
		_, err := strconv.ParseUint(x,10,8) //Note the use of 8 here because we want values 0-255
		return err==nil
	}
	
	ip = Replace(ip, " ", "")
	octets := Split(ip, ".")
	if len(octets)!=4{
		return false
	}
	for _, val := range octets{
		if isValidOctet(val)==false{
			return false
		}
	}
	return true
}

//DEPENDS ON: Split, GetIndexOf, Replace
func IsValidIPv6(ip string) bool{
	testBlocks := func(blocks []string, n int) bool{
		edgeCase := func(slice []string) bool{
			for _, val := range slice{
				if val != ""{
					return false
				}
			}
			return true
		}

		is16BitNumber := func(x string) bool{
			_, err := strconv.ParseUint(x,16,16) //Note the use of 8 here because we want values 0-FFFF
			return err==nil
		}

		if len(blocks)==3 && edgeCase(blocks){//if blocks==['','','']
			return true
		}
		var numberOfEmpty int = 0
		var index int = 0
		for _, segment := range blocks{
			if segment==""{
				if ((numberOfEmpty==1 && blocks[index-1]!="") || numberOfEmpty>1){
					return false
				}
				numberOfEmpty += 1
			} else {
				if is16BitNumber(segment)==false{
					return false
				}
			}
			index +=1
		}
		if numberOfEmpty==0 && len(blocks)!=n{
			return false
		} else {
			return true
		}
	}
	
	
	
	ip = Replace(ip, " ", "")
	blocks := Split(ip, ":")
	if len(blocks)<3 || len(blocks)>9{ 
		//Smallest possible "::"
		//Largest possible "1:2:3:4:5:6:7::"
		return false
	}
	lastBlock := blocks[len(blocks)-1]
	if len(GetIndexOf(lastBlock, ".")) != 0 { //IPv4 embedded
		ipv6Start := blocks[0:len(blocks)-1]
		if len(ipv6Start)>6{
			return false
		} else {
			return IsValidIPv4(lastBlock) && testBlocks(ipv6Start,6)
		}
	} else { //IPv4 not embedded
		return testBlocks(blocks,8)
	}
}

//DEPENDS ON: Gcd, Abs
func Lcm(a int, b int) int{
	if a==b && a==0 {
		return 0
	}
	return int(Abs(float64(a*b/Gcd(a,b))))
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

//DEPENDS ON: Extend
func MergeSort(arr []int) []int{
	merge := func(left []int, right []int) []int{
		result := []int{}
		var i, j int = 0, 0

		for i<len(left) && j<len(right){
			if left[i]<right[j]{
				result = append(result, left[i])
				i += 1
			} else {
				result = append(result, right[j])
				j += 1
			}
		}

		result = Extend(result, left[i:len(left)])
		result = Extend(result, right[j:len(right)])

		return result
	}



	if len(arr)<=1{
		return arr
	}

	var mid int = len(arr)/2 //Integer division
	leftHalf := arr[0:mid]
	rightHalf := arr[mid:len(arr)]

	sortedLeft := MergeSort(leftHalf)
	sortedRight := MergeSort(rightHalf)

	return merge(sortedLeft, sortedRight)
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
	if x==0 {
		return 1
	}
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

//DEPENDS ON: GetIndexOf
func Split(text string, pattern string) []string{
	positions := GetIndexOf(text, pattern)
	result := []string{}
	var counter int = 0
	for _, val := range positions{
		result = append(result, string(text[counter:val]))
		counter = val +len(pattern)
	}
	if counter!=len(text){
		result = append(result, string(text[counter:len(text)]))
	} else {
		result = append(result, "")
	}
	return result
}
