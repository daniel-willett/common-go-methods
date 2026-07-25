package common 

import("math"; "strconv"; "errors")

type SignedNum interface{
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type UnsignedNum interface{
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Number interface{
	SignedNum | UnsignedNum
}

type Generic interface{
	~string | ~bool | Number
}

type Integer interface{
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func Abs[Num Number](n Num) Num{
	if n<0{
		return -n
	}
	return n
}

//DEPENDS ON: Reverse
func Addition(num1 string, num2 string) string{
	padding := func(smaller string, larger string) string{
		var smallLen int = len(smaller)
		var largeLen int = len(larger)
		var result string = ""
		for i:=1; i<=largeLen-smallLen; i++{
			result += "0" //Padding of zeros
		}
		result += smaller //Append the orignal part to the pad
		return result
	}

	if len(num1)>len(num2){
		num2 = padding(num2,num1)
	} else {
		num1 = padding(num1,num2)
	}
	var result string = ""
	var (
		firstDigit = 0 
		secondDigit = 0
		carry = 0
		units = 0
		total = 0
	)
	for index := len(num1)-1; index>=0; index--{
		firstDigit, _ = strconv.Atoi(string(num1[index]))
		secondDigit, _ = strconv.Atoi(string(num2[index]))

		total = firstDigit + secondDigit + carry

		carry = total/10
		units = total%10

		result += strconv.Itoa(units)
	}
	if carry==1{
		result += "1"
	}

	//At this point, `result` has been appended to in reverse order so we need to reverse this string

	result = Reverse(result)

	return result
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

func BubbleSort[N Number](arr []N) []N{
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

func Count[G Generic](arr []G, countValue G) int{
	var n int = 0
	for _, val := range arr{
		if val==countValue{
			n+=1
		}
	}
	return n
}

//DEPENDS ON: Abs
func Gcd[I Integer](a I, b I) I{
	var larger, smaller I = 0, 0
	if a>b {
		larger = a
		smaller = b
	} else {
		larger = b
		smaller = a
	}
	var temp I = 0
	for smaller!=0 {
		temp = smaller
		smaller = larger % smaller
		larger = temp
	}
	return Abs(larger)
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

func Insert[G Generic](arr []G, pos int, val G) []G{
	result := []G{}
	var defaultVal G //As we're using generics we need some way to have a 'default' value
	var n int = len(arr)
	if pos<0{
		result = append(result, val)
		for counter:=pos+1; counter<0; counter++{
                        result = append(result, defaultVal)
                }
		result = append(result, arr...)
	} else if n-1<pos{ //pos longer than arr length
		result = append(result, arr...)
		for counter:=n; counter<pos; counter++{
			result = append(result, defaultVal)
		}
		result = append(result, val)
	} else { //pos inside array length
		result = append(result, arr[:pos]...)
		result = append(result, val)
		for counter:=pos+1; counter<=n; counter++{
			result = append(result, arr[counter-1])
		}
	}
	return result
}

func InsertionSort[N Number](arr []N) []N{
	var i int = 1
	var length int = len(arr)
	for i < length{
		var j int = i
		for j>0 && arr[j-1]>arr[j]{
			temp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = temp

			j -= 1
		}
		i += 1
	}
	return arr
}

//DEPENDS ON: Reverse
func IsPalendrome[N Number](n N) bool{
	var nstr string = strconv.Itoa(int(n))
	if Reverse(nstr)==nstr{
		return true
	}
	return false
}

func IsPrime[I Integer](n I) bool{
	if n<2{
		return false
	}
	var upper I = I(math.Sqrt(float64(n)))+1
	var factor I
	for factor = 2; factor<upper; factor++{
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

//DEPENDS ON: Split, GetIndexOf, Replace, IsValidIPv4
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
		for index, segment := range blocks{
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
		}
		if numberOfEmpty==0 && len(blocks)!=n{
			return false
		} else {
			return true
		}
	}
	isValidIPv4Inv6 := func(block string) bool{
		isValidOctet := func(x string) bool{
			_, err := strconv.ParseUint(x,10,8) //Note the use of 8 here because we want values 0-255
			return err==nil
		}
		
		octets := Split(block, ".")
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
			return isValidIPv4Inv6(lastBlock) && testBlocks(ipv6Start,6)
		}
	} else { //IPv4 not embedded
		return testBlocks(blocks,8)
	}
}

func Join(arr []string, delim string) string{
	var result string = ""
	var n int = len(arr)
	if n==0{
		return ""
	}
	for index:=0; index<n-1; index++{
		result += arr[index]
		result += delim
	}
	result += arr[n-1]
	return result
}

//DEPENDS ON: Gcd, Abs
func Lcm[I Integer](a I, b I) I{
	if a==b && a==0 {
		return 0
	}
	return Abs(a*b/Gcd(a,b))
}

func Lower(text string) string{
        var result string = ""
        for _, char := range text{
                if char<=90 && char>=65{
                        result += string(char+32)
                } else {
                        result += string(char)
                }
        }
        return result
}

func Max[N Number](arr []N) (N, error){
	if len(arr)==0{
                err := errors.New("Max: Cannot find maximum value of empty array/slice")
                return 0, err
	}
	var largest N = arr[0]
	for _, val := range arr{
		if val>largest{
			largest=val
		}
	}
	return largest, nil
}

func MergeSort[N Number](arr []N) []N{
	merge := func(left []N, right []N) []N{
		result := []N{}
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

		result = append(result, left[i:len(left)]...)
		result = append(result, right[j:len(right)]...)

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

func Min[N Number](arr []N) (N, error){
	if len(arr)==0{
		err := errors.New("Min: Cannot find minimum value of empty array/slice")
		return 0, err
	}
	var smallest N = arr[0]
	for _, val := range arr{
		if val<smallest{
			smallest=val
		}
	}
	return smallest, nil
}

func Multiplication(num1 string, num2 string) string{
	var result string = ""
	if len(num1)<len(num2){
		return Multiplication(num2,num1)
	}
	//Ensures we have Multiplication(bigger,smaller)
	var partial = make([]string, len(num2))
	for i:=len(num2)-1; i>=0; i--{
		var (
			firstDigit = 0 
			secondDigit = 0
			carry = 0
			units = 0
			total = 0
		)
		firstDigit, _ = strconv.Atoi(string(num2[i]))
		for k:=len(num2)-1-i; k>0; k--{
			partial[i] += "0"
		}
		for j:=len(num1)-1; j>=0; j--{
			secondDigit, _ = strconv.Atoi(string(num1[j]))
			total = (firstDigit*secondDigit)+carry
			carry = total/10
			units = total%10
			partial[i] += strconv.Itoa(units)

		}
		if carry!=0{
			partial[i] += strconv.Itoa(carry)

		}
		partial[i] = Reverse(partial[i])
	}
	for counter:=0; counter<len(num2); counter++{
		result = Addition(result, partial[counter])
	}
	return result
	
}




func NumOfDivs[I Integer](x I) int{
	if x==0 {
		return 1
	}
	var counter int = 0
	var upper I = I(math.Sqrt(float64(x)))+1
	var factor I
	for factor=1; factor<upper; factor++{
		if x%factor==0{
			counter += 1
		}
	}
	counter *= 2
	if I((upper-1)*(upper-1))==x{
		counter -= 1
	}
	return counter
}

func Pop[G Generic](arr []G, pos int) ([]G, G){
	result := []G{}
	var n int = len(arr)
	var value G
	if pos<0 || pos>=n{
		result = arr
		//value takes default value
	} else {
		result = append(result, arr[:pos]...)
		value = arr[pos]
		for counter:=pos+1; counter<n; counter++{
			result = append(result, arr[counter])
		}
	}
	return result, value
}

func QuickSort[N Number](arr []N) []N{
	partition := func(arr []N, low int, high int) int{
		pivot := arr[high]
		var i int = low-1
		
		for j:=low; j<high; j++{
			if arr[j]<=pivot{
				i+=1
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
		temp := arr[i+1]
		arr[i+1] = arr[high]
		arr[high] = temp
		return i+1
	}

	var quicksort func(arr []N, low int, high int)
	quicksort = func(arr []N, low int, high int){
		if low<high{
			pivotIndex := partition(arr, low, high)
			quicksort(arr, low, pivotIndex-1)
			quicksort(arr, pivotIndex+1, high)
		}
	}

	quicksort(arr, 0, len(arr)-1)

	return arr
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

func Reverse(text string) string{
	var result string = ""
	runes := []rune(text)
	for counter := len(runes)-1; counter>=0; counter--{
		result += string(runes[counter])
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

//DEPENDS ON: Upper, Lower
func SwapCase(text string) string{
	//normal XOR upper XOR lower = swapcase
	upper := []rune(Upper(text))
	lower := []rune(Lower(text))
	start := []rune(text)
	var result string = ""
	for index, _ := range start{
		result += string(upper[index] ^ lower[index] ^ start[index])
	}
	return result
}

func Upper(text string) string{
	var result string = ""
	for _, char := range text{
		if char<=122 && char>=97{
			result += string(char-32)
		} else {
			result += string(char)
		}
	}
	return result
}
