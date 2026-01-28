# common-go-methods
A library of common Golang functions which either are often used or are likely to be useful. A good source of inspiration is from projects already completed like Project Euler and also with the success Python has had to look at their docs for funcions and methods built in that would likely be useful. e.g.

- https://docs.python.org/3/library/functions.html
- https://www.w3schools.com/python/python_ref_functions.asp



## Contents

There is a `common.go` file which contains the list of functions thus far.

### Abs
This takes a `float64` and returns a `float64`. It determines the absolute value of the input which means either returning the positive version of the input, if it isn't already positive.

### All
This takes an array of `bool`s and returns a `bool`. It determines if the array is all `true`, and returns false if any value is `false`.

### Any
This takes an array of `bool`s and returns a `bool`. It determines if any part of the array is `true`, and returns false if the whole array is `false`.

### IsPalendrome

This takes an `int` and retrns a `bool`. It determines if writing the number backwards is the same as forwards. _Consequently this uses the `Reverse()` function defined also in this file._

### IsPrime

This takes an `int` and returns a `bool`. It determines if the only factors of the number are the number itself and 1.

### Max

This takes an array of `int`s and returns an `int` value from this list. It determines which of these `int`s are most positive.

### NumOfDivs

This takes an `int` and returns an `int`. It determines how many values between 1 and the given value divide into the give value.

### Reverse

This takes a `string` and returns a `string`. It flips the string to reverse order.




