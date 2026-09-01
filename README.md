# common-go-methods
A library of common Golang functions which either are often used or are likely to be useful. A good source of inspiration is from projects already completed like Project Euler and also with the success Python has had to look at their docs for funcions and methods built in that would likely be useful. e.g.

- https://docs.python.org/3/library/functions.html
- https://www.w3schools.com/python/python_ref_functions.asp

## Go version

This project is using Go 1.22. The CI will use 1.22, the `go.mod` should say 1.22.

## Contents

There is a `common.go` file which contains the list of functions thus far. This contains various type constraints to be used for generics in many of the functions. The generics we have are `SignedNum`, `UnsignedNum`, `Number`, `Generic`, and `Integer`. These can all be found at the start of the `common.go` file.



## Big Numbers

Some numbers are very big. Bigger than Go can store using the default number types. For this we can represent numbers as strings and there is the basic operations: `Addition`, +; `Subtraction`, -; `Multiplication`, x
These take arbitrarily long strings which represent numbers and perform the associated operation on them.
If the input is not a valid number then the method will return an error.

Note: Current support is for positive integers only. `Subtraction` can result in a negative integer but does not yet support it back as an input.

### CompareAMoreThanB

This evaluates the mathematical logical computation of `A>B` for inputs `A` and `B` as numerical representation of the `string` inputs.

### IsNumber

This takes in a `string` representation of a Big Number and returns whether the `string` is valid to represent a number i.e. is it only composed of digits.

### Padding

This takes in a `smaller` and `larger` `string`s and then pads the front of `smaller` with `"0"`s until the result has the same length as the `larger` string resulting in a `string`.

### RemoveLeadingZeros

This takes an input `string` which should be a number and removes the leading `0`s it has. This is like an opposite to `Padding()`.


## Maths

### Abs

This functions as the absolute value, | |, taking in a `Number` and returning a `Number`.

### Gcd

This finds the Greatest Common Divisor of two `Integer`s returning the result in the same type (`Integer`)

### IsPalendrome

This determines if an input `Number` reversed gives the same value resulting in either `true` or `false`.

### IsPrime

This determines if an input `Integer` is a prime number resulting in either `true` or `false`.

### Lcm

This finds the Lowest Common Multiple of two `Integer`s returning the result in the same type (`Integer`)

### NumOfDivs

This determines the number of proper divisors of an input `Integer` resulting in an `int` value. This will include `1` and the input number itself.


## Arrays/Sclices

### All
This takes an array/slice of `bool`s and returns a `bool`. It determines if the array is all `true`, and returns false if any value is `false`.

### Any

This takes an array/slice of `bool`s and returns a `bool`. It determines if any part of the array is `true`, and returns false if the whole array is `false`.

### BubbleSort

This takes an array/slice of `Number`s and returns an array of `Numbers`s. It determines the numerically ascending order of the elements with the largest at the end of the array and the smallest at the front. This is achieved using Bubble Sort.

### Count
This takes an array/slice of `Generic`s & a `Generic` and returns an `int`. It counts how many times the single `Generic` value appears in the `Generic` array.

### Insert

This takes an array/slice of `Generic`s, a positions/index `int`, & a `Generic` value and results in an array/slice of `Generic`s. It finds the index value, regardless of value or array/slice size, and inserts the `Generic` value. If the index is larger than the array/slice length then the default value for `Generic` will be padding the result. If the index is less than 0 then the same occurs but in reverse.

### InsertionSort

This takes an array/slice of `Number`s and returns an array of `Numbers`s. It determines the numerically ascending order of the elements with the largest at the end of the array and the smallest at the front. This is achieved using Insertion Sort.

### Max

This takes an array/slice of `Number`s and returns the largest `Number` with an `error` for if the array is empty.

### MergeSort

This takes an array/slice of `Number`s and returns an array of `Numbers`s. It determines the numerically ascending order of the elements with the largest at the end of the array and the smallest at the front. This is achieved using Merge Sort.

### Min

This takes an array/slice of `Number`s and returns the smallest `Number` with an `error` for if the array is empty.

### Pop

This takes an array/slice of `Generic`s and an index value `int`. It then removes the `Generic` and position `index` and returns both the 'popped' value along with the modified array/slice. If the index is outside of the range of valid index values then the result is the original array/slice and the defaut value for the given `Generic`.

### QuickSort

This takes an array/slice of `Number`s and returns an array of `Numbers`s. It determines the numerically ascending order of the elements with the largest at the end of the array and the smallest at the front. This is achieved using Quick Sort.


## String Manipulation

### GetIndexOf

This takes an input text of `string` and a pattern to search for within the text as `string`. It will result in an `int` slice of all the non-overlapping indexes where the pattern occurs within the text.

### Join

This takes an array of `string`s and a `string` and returns a `string`. Elements of the array are connected together with the second `string` used as a delimiter between them. This is in essence the opposite of `Split()`.

### Lower

This takes a `string` and turns the alphabet characters A-Z to their lower case value a-z, resutling in the modified version of the input `string`.

### Replace

This takes an input text `string`, a pattern `string` and a result `string`. It will search through the text for the occurances of the pattern and then replace those patterns with the result. The resulting output is the modified text `string`.

### Reverse

This takes a n input text `string` and results with the reversed `string`.

### Split

This takes an input text `string` and a pattern `string`. It will search for instances of the pattern and delimit the input `string` into a slice with each entry being the values between the patterns.

### SwapCase

This takes a `string` and turns the alphabet characters A-Z to their lower case value a-z whilst also turning the alphabet characters a-z to their upper case A-Z, resutling in the modified version of the input `string`.

### Upper

This takes a `string` and turns the alphabet characters a-z to their upper case value A-Z, resutling in the modified version of the input `string`.


## Networking

### IsValidIPv4

This takes a `string` input and determines if it is a valid IPv4 address resulting in `true` or `false`.

### IsValidIPv6

This takes a `string` input and determines if it is a valid IPv6 address resulting in `true` or `false`. This includes IPv4 embeddings.


## Examples

### Big Numbers:

```
result, error := Addition("1234", "5678")
// error is nil and result is "6912"
result, error := Addition("123four", "1234")
// result is "" and error is "123four is not a valid number in String form"
```
```
result, error := CompareAMoreThanB("100", "20")
// error is nil and result is true
result, error := CompareAMoreThanB("100", "100")
// error is nil and result is false
```
```
result := IsNumber("100")
// result is true
result := IsNumber("five")
// result is false
```
```
result, error := Multiplication("4", "12")
// error is nil and result is "48"
result, error := Multiplication("four", "12")
// result is "" and error is "four is not a valid number in String form"
```
```
str1 := "100"
str2 := "7654321"
str3 := ""

result := Padding(str1, str2)
// result is "0000100"
result := Padding(str3, str2)
// result is "0000000"
```
```
result := RemoveLeadingZeros("000100")
// result is "100"
```
```
result, error := Subtraction("50","15")
// error is nil result is "35"
result, error := Subtraction("15","50")
// error is nil result is "-35"
result, error := Subtraction("-50", "15")
// result is "" and error is "-50 is not a valid number in String form"
```

### Math:

```
result = Gcd(15,10)
// result is 5
```
```
result := IsPalendrome(1000)
// result is false
result := IsPalendrome(1001)
// result is true
result := IsPalendrome(1)
// result is true
```
```
result := IsPrime(13)
// result is true
result := IsPrime(10)
// result is false
```
```
result := Lcm(3,7)
// result is 21
```
```
result := NumOfDivs(13)
// result is 2
result := NumOfDivs(20)
// result is 6 (as 1,2,4,5,10,20 are the divisors)
```

### Arrays/Sclices:

```
arr1 := []bool{false, true, true}
arr2 := []bool{true, true, true}
arr3 := []bool{false, false , false}

result := All(arr1)
// result is false
result := All(arr2)
// result is true

result := Any(arr1)
// result is true
result := Any(arr3)
// result is false
```
```
arr1 := []string{"Hi", "hi", "hI", "HI"}
arr2 := []int{1,1,1,1,1,2,3,4}

result := count(arr1, "hi")
// result is 1
result := count(arr2, 1)
// result is 5
```
```
arr1 := []string{"a", "b", "c", "e", "f", "g"}
arr2 := []string{}

result := Insert(arr1, 3, "d")
// result is []string{"a", "b", "c", "d", "e", "f", "g"}
result := Insert(arr2, 5, "Hello World")
// result is []string{"", "", "", "", "", "Hello World"}
result := Insert(arr1, -4, "w")
// result is []string{"w","", "", "", "a", "b", "c", "e", "f", "g"}
```
```
arr1 := []string{"Hello", "World"}
arr2 := []string{"","","",""}

result := Join(arr1, " ")
// result is "Hello World"
result := Join(arr2, "Hi")
// result is "HiHiHi"
```
```
arr1 := []int{1,2,3,4,5,6}
arr2 := []int{}

result, error := Max(arr1)
// error is nil and result is 6
result, error := Max(arr2)
// result is 0 and error is Max: "Cannot find maximum value of empty array/slice"

result, error := Min(arr1)
// error is nil and result is 1
result, error := Min(arr2)
// result is 0 and error is Max: "Cannot find minimum value of empty array/slice"
```
```
arr1 := []int{20, 30, 40, 100, 200, 300}
arr2 := []string{"Hello", "World"}

result, val := Pop(arr1, 2)
// result is []int{20, 30, 100, 200, 300} and val is 40
result, val := Pop(arr2, -5)
// result is []string{"Hello", "World"} and val is ""
```

For all the sorting algorithms:
```
arr := []int{9,8,7,6,5,4,3,2,1}
result := Algorithm(arr)
// result is []int{1,2,3,4,5,6,7,8,9}
```

### String Manipulation:

```
str1 := "Hello World"
str2 := "aaaaa"

result := GetIndexOf(str1, "l")
// result is []int{2,3,9}
result := GetIndexOf(str2, "aa")
// result is []int{0,2}
```
```
result := Lower("Hello World")
// result is "hello world"
```
```
str1 := "Hello World"
str2 := "aaaaa"

result := Replace(str1, "World", "There")
// result is "Hello There"
result := Replace(str2, "aa", "b")
// result is "bba"
```
```
result := Reverse("Hello")
// result is "olleH"
```
```
result := Split("Hello World", " ")
// result is []string{"Hello", "World"}
```
```
result := SwapCase("Hello World")
// result is "hELLO wORLD"
```
```
result := Upper("Hello World")
// result is "HELLO WORLD"
```

### Networking:

```
result := IsValidIPv4("120.10.22.10")
// result is true
result := IsValidIPv4("1.1.A.1")
// result is false

result := IsValidIPv6("::1.2.3.4")
// result is true
result := IsValidIPv6("1f5e:8fbe:e550:7a07:4679:14fc:1.2.3.4")
// result is true
result := IsValidIPv6("10000:0:0:0:0:0:0:0")
// result is false
```
