# common-go-methods
A library of common Golang functions which either are often used or are likely to be useful. A good source of inspiration is from projects already completed like Project Euler and also with the success Python has had to look at their docs for funcions and methods built in that would likely be useful. e.g.

- https://docs.python.org/3/library/functions.html
- https://www.w3schools.com/python/python_ref_functions.asp

## Go version

This project is using Go 1.22. The CI will use 1.22, the `go.mod` should say 1.22.

## Contents

There is a `common.go` file which contains the list of functions thus far. This contains various type constraints to be used for generics in many of the functions. The generics we have are `SignedNum`, `UnsignedNum`, `Number`, `Generic`, and `Integer`.

### Abs

This takes a `Number` and returns a `Number`. It determines the absolute value of the input which means returning the positive version of the input, if it isn't already positive.

### All

This takes an array of `bool`s and returns a `bool`. It determines if the array is all `true`, and returns false if any value is `false`.

### Any

This takes an array of `bool`s and returns a `bool`. It determines if any part of the array is `true`, and returns false if the whole array is `false`.

### BubbleSort

This takes an array of `Number`s and returns an array of `Numbers`s. It determines the numerically ascending order of the elements with the largest at the end of the array and the smallest at the front.

### Count
This takes an array of `Genric`s & a `Generic` and returns an `int`. It counts how many times the single `Generic` value appears in the `Generic` array.

### Gcd

This takes two `Integer`s and returns an `Integer`. It determines the the largest positive integer which will divide into both given `Integer`s. It does this by Euclid's Algorithm.

### GetIndexOf

This takes a `string` and returns an `[]int`. It determines all the occurances of a pattern within a string and gives an empty slice if no such pattern exists.

### Insert

This takes an array of `Generic`s, an `int` & a `Generic` and returns an array of `Generic`s. It takes the `int` and this is the position in the `Generic` array it will insert the given `Generic` value based on indexing. If the index is negative, it works out the position going backwards and uses the default `Generic`'s value to backfill the empty positions.

### InsertionSort

This takes an array of `Number`s and returns an array of `Number`s. It determines the numerically ascending order of the elements with the largest at the end of the array and the smallest at the front.

### IsPalendrome

This takes a `Number` and returns a `Number`. It determines if writing the number backwards is the same as forwards.

### IsPrime

This takes an `Integer` and returns a `bool`. It determines if the only factors of the number are the number itself and 1.

### IsValidIPv4

This takes a `string` and returns a `bool`. It determines if the given string is a valid IPv4 address. Currently this is just a crude checker and doesn't currently do more complex cases of IPv4 formatting.

### IsValidIPv6

This takes a `string` and returns a `bool`. It determines if the given string is a valid IPv4 address. This does allow for IPv4 embedding and thus relies on the `IsValidIPv4()` function.

### Join

This takes an array of `string`s and a `string` and returns a `string`. Elements of the array are connected together with the second `string` used as a delimiter between them. This is in essence the opposite of `Split()`.

### Lcm

This takes two `Integers`s and returns an `Integer`. It determines the largest positive integer which the two given `Integers`s divide into. It does this through using `Gcd()`.

### Lower

This takes a `string` and returns a `string`. It determines the ASCII value each character of the string and shifts it appropriately when needed to make the result lower case.

### Max

This takes an array of `Number`s and returns a `Number` value from this list and an `error`. It determines which of these `Number`s are most positive. Should the input be irrelevant or produce an undefined result, then the `error` value will switch from `nil` to an error value.

### MergeSort
This takes an array of `Number`s and returns an array of `Number`s. It determines the numerically ascending order of the elements with the largest at the end of the array and the smallest at the front.

### Min

This takes an array of `Number`s and returns a `Number` value from this list and an `error`. It determines which of these `Number`s are least positive. Should the input be irrelevant or produce an undefined result, then the `error` value will switch from `nil` to an error value.

### NumOfDivs

This takes an `Integer` and returns an `int`. It determines how many values between 1 and the given value divide into the give value.

### Pop
This takes an array of `Generic`s & an `int` and returns an array of `Generic`s & a `Generic. It determines the index position of the `Generic` array to remove and return as the `Generic` value along side with the `Generic` array now reduced by 1 value.

### QuickSort

This takes an array of `Number`s and returns an array of `Number`s. It determines the numerically ascending order of the elements with the largest at the end of the array and the smallest at the front.

### Replace

This takes three `string`s and returns a `string`. It determines the occurances of a pattern within a given text and replaces those patterns with a new pattern.

### Reverse

This takes a `string` and returns a `string`. It flips the string to reverse order.

### Split

This takes two `string`s and returns an `[]string`. It determines the positions of a pattern within a given text and splits the text accross a slice along those pattern matches.

### SwapCase

This takes a `string` and returns a `string`. It determines if characters within a string are upper case or lower case and swaps them to the other casing.

### Upper

This takes a `string` and returns a `string`. It determines the ASCII value each character of the string and shifts it appro
priately when needed to make the result upper case.
