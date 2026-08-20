package common

import ("testing"; "reflect"; "errors")

func TestAbs(t *testing.T){
	tests := []struct{
		name		string
		input		float64
		expected	float64
	}{
		{"Integer", 5, 5},
		{"Negative Integer", -10, 10},
		{"Zero", 0, 0},
		{"Float", 123.532,123.532},
		{"Negative Float", -43.000001, 43.000001},
	}

	for _, tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			result := Abs(tt.input)
			if result != tt.expected{
				t.Errorf("Abs(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAddition(t *testing.T){
	tests := []struct{
		name		string
		x		string
		y		string
		expectedVal	string
		expectedErr	error
	}{
		{"Small Integers", "1", "2", "3", nil},
		{"Large Integer", "999999999999999999999999999999999999999999999999999999999999999999999999999999999", "1", "1000000000000000000000000000000000000000000000000000000000000000000000000000000000", nil},
		{"Zero", "0", "0", "0", nil},
		{"One Empty String", "", "1", "1", nil},
		{"Two Empty String", "", "", "", nil},
		{"Non-Numbers", "5A", "5", "", errors.New("5A is not a valid number in String form")},
		{"Negative Numbers", "-3", "3", "", errors.New("-3 is not a valid number in String form")},
	}

	for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
			//Following the way we did it in `TestMin`...
                        result, err := Addition(tt.x, tt.y)
			if (err==nil && tt.expectedErr!=nil) || (err!=nil && tt.expectedErr==nil){
				t.Errorf("Addition(%v, %v) = %v, %v; want %v,\n %v",
				tt.x, tt.y, result, err, tt.expectedVal, tt.expectedErr)
			} else if err!=nil && err.Error()!=tt.expectedErr.Error(){
				t.Errorf("Addition(%v, %v) = %v, %v; want %v,\n %v",
				tt.x, tt.y, result, err, tt.expectedVal, tt.expectedErr)
                        }
                        if result != tt.expectedVal{
				t.Errorf("Addition(%v, %v) = %v, %v; want %v,\n %v",
				tt.x, tt.y, result, err, tt.expectedVal, tt.expectedErr)
                        }
                })
        }
}

func TestAll(t *testing.T){
        tests := []struct{
                name            string
                input           []bool
                expected        bool
        }{
                {"3 Trues", []bool{true, true, true}, true},
                {"3 Falses", []bool{false, false, false}, false},
                {"Empty", []bool{}, true}, //The empty intersection is the whole space so thus true here
                {"One Of Each", []bool{true, false}, false},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := All(tt.input)
                        if result != tt.expected{
                                t.Errorf("All(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestAny(t *testing.T){
        tests := []struct{
                name            string
                input           []bool
                expected        bool
        }{
                {"3 Trues", []bool{true, true, true}, true},
                {"3 Falses", []bool{false, false, false}, false},
                {"Empty", []bool{}, false}, //The empty union is the empty set so thus false here
                {"One Of Each", []bool{true, false}, true},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Any(tt.input)
                        if result != tt.expected{
                                t.Errorf("Any(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestCompareAMoreThanB(t *testing.T){
        tests := []struct{
                name            string
                A	        string
		B		string
                expectedVal     bool
		expectedErr	error
        }{
		{"Small Integers", "5", "60", false, nil},
		{"Large Integers - Same Length", "1000000000000000000000000000000000000000000000000000000000000000000000000000000001", "1000000000000000000000000000000000000000000000000000000000000000000000000000000000", true, nil},
		{"One Empty", "", "300", false, nil},
		{"Both Empty", "", "", false, errors.New("CompareAMoreThanB: Cannot compare two empty strings")},
		{"Non-Numbers", "5A", "5", false, errors.New("5A is not a valid number in String form")},
		{"Negative Numbers", "-3", "3", false, errors.New("-3 is not a valid number in String form")},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
			//Following the way we did it in `TestMin`...
                        result, err := CompareAMoreThanB(tt.A, tt.B)
			if (err==nil && tt.expectedErr!=nil) || (err!=nil && tt.expectedErr==nil){
				t.Errorf("CompareAMoreThanB(%v, %v) = %v, %v; want %v,\n %v",
				tt.A, tt.B, result, err, tt.expectedVal, tt.expectedErr)
			} else if err!=nil && err.Error()!=tt.expectedErr.Error(){
				t.Errorf("CompareAMoreThanB(%v, %v) = %v, %v; want %v,\n %v",
				tt.A, tt.B, result, err, tt.expectedVal, tt.expectedErr)
                        }
                        if result != tt.expectedVal{
				t.Errorf("CompareAMoreThanB(%v, %v) = %v, %v; want %v,\n %v",
				tt.A, tt.B, result, err, tt.expectedVal, tt.expectedErr)
                        }
                })
        }
}

func TestCount(t *testing.T){
        tests := []struct{
                name            string
                arr	        []int
		countValue	int
                expected        int
        }{
		{"Standard", []int{1,1,1,1,2,2,3,3,4}, 1, 4},
		{"Empty", []int{}, 10, 0},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Count(tt.arr, tt.countValue)
                        if result != tt.expected{
                                t.Errorf("Count(%v, %v) = %v; want %v", tt.arr, tt.countValue, result, tt.expected)
                        }
                })
        }
}


func TestGcd(t *testing.T){
        tests := []struct{
                name            string
                a        	int
		b		int
                expected        int
        }{
		//Normal cases
                {"Two Positive", 20, 15, 5},
                {"One Negative", -20, 15, 5},
		{"Two Negative", -20, -15, 5},
		{"Coprime", 7, 13, 1},

		//Edge cases
		{"One 0", 3, 0, 3},
		{"Two 0", 0, 0, 0},
		{"One 1", 3, 1, 1},
		{"Two 1", 1, 1, 1},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Gcd(tt.a, tt.b)
                        if result != tt.expected{
                                t.Errorf("Gcd(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
                        }
                })
        }
}

func TestGetIndexOf(t *testing.T){
        tests := []struct{
                name            string
                text            string
                pattern        	string
                expected        []int
        }{
                //Normal cases
		{"One Character", "Hello World", "o", []int{4, 7}},
		{"Two Character", "Hello World", "ll", []int{2}},
		{"No Match", "Hello World", "cat", []int{}},
		{"Space", "Hello World", " ", []int{5}},

		//Edge cases
		{"Empty Text", "", "pattern", []int{}},
		{"Empty Pattern", "Hello World", "", []int{}},
		{"Regex-Style Response", "aaaaa", "aa", []int{0,2}},
		{"Emojis", "😊", "😊", []int{0}},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := GetIndexOf(tt.text, tt.pattern)
                        if !reflect.DeepEqual(result, tt.expected){
                                t.Errorf("GetIndexOf(%v, %v) = %v; want %v", tt.text, tt.pattern, result, tt.expected)
                        }
                })
        }
}

func TestInsert(t *testing.T){
	tests := []struct{
                name            string
                inputArr	[]string
		inputPos	int
		inputVal	string
                expected        []string
        }{
                //Normal cases
                {"Standard", []string{"a", "b", "c", "e", "f", "g"}, 3, "d", []string{"a", "b", "c", "d", "e", "f", "g"}},
		{"Empty Array", []string{}, 5, "Hi", []string{"", "", "", "", "", "Hi"}},
		{"Empty Array & 0 Position", []string{}, 0, "a", []string{"a"}},
		{"Empty Value", []string{"a", "b", "c"}, 2, "", []string{"a", "b", "", "c"}},
		{"Position too large", []string{"a", "b", "c"}, 5, "f", []string{"a", "b", "c", "", "", "f"}}, //abcdef
		{"Position too small", []string{"a", "b", "c"}, -4, "w", []string{"w", "", "", "", "a", "b", "c"}}, //wxyzabc
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Insert(tt.inputArr, tt.inputPos, tt.inputVal)
			if !reflect.DeepEqual(result, tt.expected){
                                t.Errorf("Split(%v, %v, %v) = %v; want %v", tt.inputArr, tt.inputPos, tt.inputVal, result, tt.expected)
                        }
                })
        }
}

func TestIsNumber(t *testing.T){
        tests := []struct{
                name            string
                input		string	
                expected        bool
        }{
		{"Success", "1221", true},
		{"Failure", "122A", false},
		{"Negative", "-10", false},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := IsNumber(tt.input)
                        if result != tt.expected{
                                t.Errorf("IsNumber(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestIsPalendrome(t *testing.T){
        tests := []struct{
                name            string
                input		uint	
                expected        bool
        }{
		{"Success", 1221, true},
		{"Failure", 1245, false},
		{"One Char", 1, true},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := IsPalendrome(tt.input)
                        if result != tt.expected{
                                t.Errorf("IsPalendrome(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestIsPrime(t *testing.T){
        tests := []struct{
                name            string
                input		int	
                expected        bool
        }{
                //Normal cases
		{"Success", 13, true},
		{"Failure", 12, false},
		{"One", 1, false},
		{"Zero", 0, false},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := IsPrime(tt.input)
                        if result != tt.expected{
                                t.Errorf("IsPrime(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestIsValidIPv4(t *testing.T){
        tests := []struct{
                name            string
                input		string	
                expected        bool
        }{
                //True cases
		{"Smallest", "0.0.0.0", true},
		{"Standard", "19.24.62.12", true},
		{"Largest", "255.255.255.255", true},
		//False cases
		{"Too Large Octet", "1000.0.0.1", false},
		{"Too Few Octets", "1.1.1", false},
		{"Invalid Octet", "1.1.A.1", false},
		{"Too Many Octets", "1.1.1.1.1", false},
		{"Random Chatacters", "Hello", false},
		{"Emoji", "1.1.1.😊", false},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := IsValidIPv4(tt.input)
                        if result != tt.expected{
                                t.Errorf("IsValidIPv4(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestIsValidIPv6(t *testing.T){
        tests := []struct{
                name            string
                input           string
                expected        bool
        }{
		//True cases
		{"Small", "0:0:0:0:0:0:0:1", true},
		{"Small Compact", "::1", true},
		{"Reverse Small Compact", "1::", true},
		{"Zeros Compact", "::", true},
		{"Largest", "FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF:FFFF", true},
		{"Standard Compact", "98a:329:adae::23:0:0:1", true},
		{"Standard", "1f5e:8fbe:e550:7a07:4679:14fc:1.2.3.4", true},
		{"IPv4 Embedded Compact Zeros", "::1.2.3.4", true},
		{"IPv4 Embedded Compact", "1::1:1.2.3.4", true},
		{"IPv4 Zeros Embedded Compact","132::0.0.0.0", true},

		//False cases
		{"Random Characters", "dasjkhdj", false},
		{"Too Few Blocks", "0", false},
		{"Too Many Blocks", "0:0:0:0:0:0:0:0:0", false},
		{"Too Few Blocks IPv4 Embedded", "0:0:0:0:1.2.3.4", false},
		{"Too Many Blocks IPv4 Embedded", "0:0:0:0:0:0:0:1.2.3.4", false},
		{"Too Few Octets IPv4 Embedded", "0:0:0:0:0:0:1.2.3", false},
		{"Too Many Octets IPv4 Embedded", "0:0:0:0:0:0:1.2.3.4.5", false},
		{"Too Large Block", "10000:0:0:0:0:0:0:0", false},
		{"Emoji", "1::1:😊", false},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := IsValidIPv6(tt.input)
                        if result != tt.expected{
                                t.Errorf("IsValidIPv6(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestJoin(t *testing.T){
        tests := []struct{
                name            string
                arr		[]string
		delim		string
                expected        string
        }{
		{"Standard", []string{"Hello", "World"}, " ", "Hello World"},
		{"Empty Array", []string{}, "A", ""},
		{"Empty Pattern", []string{"Hello", "World"}, "", "HelloWorld"},
		{"Empty Array Values", []string{"","","",""}, "Hi", "HiHiHi"},
		{"Emojis", []string{"😊","😊","😊"}, "👻", "😊👻😊👻😊"},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Join(tt.arr, tt.delim)
                        if result != tt.expected{
                                t.Errorf("Join(%v, %v) = %v; want %v", tt.arr, tt.delim, result, tt.expected)
                        }
                })
        }
}

func TestLcm(t *testing.T){
        tests := []struct{
                name            string
                a               int
                b               int
                expected        int
        }{
                //Normal cases
                {"Two Positive", 20, 15, 60},
                {"One Negative", -20, 15, 60},
                {"Two Negative", -20, -15, 60},
                {"Coprime", 7, 13, 91},

                //Edge cases
                {"One 0", 3, 0, 0},
                {"Two 0", 0, 0, 0},
                {"One 1", 3, 1, 3},
                {"Two 1", 1, 1, 1},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Lcm(tt.a, tt.b)
                        if result != tt.expected{
                                t.Errorf("Lcm(%v, %v) = %v; want %v", tt.a, tt.b, result, tt.expected)
                        }
                })
        }
}

func TestLower(t *testing.T){
        tests := []struct{
                name            string
                input           string
                expected        string
        }{
                {"Standard", "HELLO WORLD", "hello world"},
                {"Digit", "1", "1"},
                {"Empty", "", ""},
		{"Emoji", "Hello😊World", "hello😊world"},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Lower(tt.input)
                        if result != tt.expected{
                                t.Errorf("Lower(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestMax(t *testing.T){
        tests := []struct{
                name            string
                input           []int
                expectedVal     int
		expectedErr	error
        }{
		{"Signleton", []int{5}, 5, nil},
		{"Positive Values", []int{1,2,3,4,5}, 5, nil},
		{"Negative Values", []int{-1,-2,-3,-4,-5}, -1, nil},
		{"Positive & Negative Values", []int{1,-1,2,-2,3,-3}, 3, nil},
		{"Empty", []int{}, 0, errors.New("Max: Cannot find maximum value of empty array/slice")},
	}

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result, err := Max(tt.input)
			//We have a slight problem with nil vs non-nil errors.
			//So first we check explicitly if they are equal in the nil sense
			//Should that pass by then we check if they are equal in a non-nil sense
			if (err==nil && tt.expectedErr!=nil) || (err!=nil && tt.expectedErr==nil){
				t.Errorf("Max(%v) = %v, %v; want %v, %v",
				tt.input, result, err, tt.expectedVal, tt.expectedErr)
			} else if err!=nil && err.Error()!=tt.expectedErr.Error(){
                                t.Errorf("Max(%v) = %v, %v; want %v, %v",
				tt.input, result, err, tt.expectedVal, tt.expectedErr)
                        }
			//Finally we check if the expected value is what we wanted
			if result!=tt.expectedVal {
				t.Errorf("Max(%v) = %v, %v; want %v, %v",
				tt.input, result, err, tt.expectedVal, tt.expectedErr)
			}
                })
        }
}

func TestMin(t *testing.T){
        tests := []struct{
                name            string
                input           []int
                expectedVal     int
		expectedErr	error
        }{
                {"Signleton", []int{5}, 5, nil},
                {"Positive Values", []int{1,2,3,4,5}, 1, nil},
                {"Negative Values", []int{-1,-2,-3,-4,-5}, -5, nil},
                {"Positive & Negative Values", []int{1,-1,2,-2,3,-3}, -3, nil},
		{"Empty", []int{}, 0, errors.New("Min: Cannot find minimum value of empty array/slice")},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result, err := Min(tt.input)
			//We have a slight problem with nil vs non-nil errors.
                        //So first we check explicitly if they are equal in the nil sense
                        //Should that pass by then we check if they are equal in a non-nil sense
                        if (err==nil && tt.expectedErr!=nil) || (err!=nil && tt.expectedErr==nil){
                                t.Errorf("Min(%v) = %v, %v; want %v, %v",
                                tt.input, result, err, tt.expectedVal, tt.expectedErr)
                        } else if err!=nil && err.Error()!=tt.expectedErr.Error(){
                                t.Errorf("Min(%v) = %v, %v; want %v, %v",
                                tt.input, result, err, tt.expectedVal, tt.expectedErr)
                        }
                        //Finally we check if the expected value is what we wanted
                        if result!=tt.expectedVal {
                                t.Errorf("Min(%v) = %v, %v; want %v, %v",
                                tt.input, result, err, tt.expectedVal, tt.expectedErr)
                        }
                })
        }
}

func TestMultiplication(t *testing.T){
	tests := []struct{
                name            string
                num1           	string
		num2		string
                expectedVal	string
		expectedErr	error
        }{
		{"Small Integers", "2", "3", "6", nil},
		{"Large Integer", "999999999999999999999999999999999999999999999999999999999999999999999999999999999", "1000000000000000000000000000000000000000000000000000000000000000000000000000000000", "999999999999999999999999999999999999999999999999999999999999999999999999999999999000000000000000000000000000000000000000000000000000000000000000000000000000000000", nil},
		{"Zero", "150", "0", "0", nil},
		{"One Empty", "", "5", "0", nil},
		{"Both Empty", "", "", "0", nil},
		{"Non-Numbers", "5A", "5", "", errors.New("5A is not a valid number in String form")},
		{"Negative Numbers", "-3", "3", "", errors.New("-3 is not a valid number in String form")},
        }

	for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
			//Following the way we did it in `TestMin`...
                        result, err := Multiplication(tt.num1, tt.num2)
			if (err==nil && tt.expectedErr!=nil) || (err!=nil && tt.expectedErr==nil){
				t.Errorf("Multiplication(%v, %v) = %v, %v; want %v,\n %v",
				tt.num1, tt.num2, result, err, tt.expectedVal, tt.expectedErr)
			} else if err!=nil && err.Error()!=tt.expectedErr.Error(){
				t.Errorf("Multiplication(%v, %v) = %v, %v; want %v,\n %v",
				tt.num1, tt.num2, result, err, tt.expectedVal, tt.expectedErr)
                        }
                        if result != tt.expectedVal{
				t.Errorf("Multiplication(%v, %v) = %v, %v; want %v,\n %v",
				tt.num1, tt.num2, result, err, tt.expectedVal, tt.expectedErr)
                        }
                })
        }
}

func TestNumOfDivs(t *testing.T){
        tests := []struct{
                name            string
                input           int
                expected        int
        }{
		{"Composite", 20, 6}, //1,2,4,5,10,20
		{"Prime", 13, 2},
		{"One", 1, 1},
		{"Zero", 0, 1},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := NumOfDivs(tt.input)
                        if result != tt.expected{
                                t.Errorf("NumOfDivs(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestPop(t *testing.T){
	tests := []struct{
                name            string
		arr		[]string
		pos		int
                expectedArr     []string
		expectedVal	string
        }{
		{"Standard", []string{"Apple", "Pear", "Banana", "Orange"}, 2, []string{"Apple", "Pear", "Orange"}, "Banana"},
		{"Empty", []string{}, 0, []string{}, ""},
		{"Position Too Large", []string{"Hello", "World"}, 5, []string{"Hello", "World"}, ""},
		{"Negative Position", []string{"Hello", "World"}, -5, []string{"Hello", "World"}, ""},
	}

	for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        resultArr, resultVal := Pop(tt.arr, tt.pos)
			if !reflect.DeepEqual(resultArr, tt.expectedArr) || resultVal!=tt.expectedVal{
                                t.Errorf("Pop(%v, %v) = %v, %v; want %v, %v", 
				tt.arr, tt.pos, resultArr, resultVal, tt.expectedArr, tt.expectedVal)
                        }
                })
        }
}

func TestReplace(t *testing.T){
        tests := []struct{
                name            string
		text		string
		pattern		string
		newPattern	string
                expected        string
        }{
		//Normal cases
		{"Standard", "Hello World", "World", "Dogs", "Hello Dogs"},
		{"Same length", "Hello Cats", "Cats", "Dogs", "Hello Dogs"},
		{"Spaces", "Hello World", " ", "", "HelloWorld"},
		{"No match", "Hello World", "Cats", "Dogs", "Hello World"},
		//Edge Cases
		{"Empty Text", "", "Cats", "Dogs", ""},
		{"Empty Pattern", "Hello World", "", "Dogs", "Hello World"},
		{"Regex-Style Matching", "aaaaa", "aa", "b", "bba"},
		{"Emoji Text", "😊😊👻👻😭😭😂😂🤣🤣❤️❤️😍😍😒😒👌👌😘😘💕💕😁😁", "128513", "a", "😊😊👻👻😭😭😂😂🤣🤣❤️❤️😍😍😒😒👌👌😘😘💕💕😁😁"},
		{"Emoji Text & Pattern", "😊😊👻👻😭😭", "😊😊", "emoji", "emoji👻👻😭😭"},
		{"Emoji Text & Pattern & New Pattern", "😊😊👻👻😭😭", "😊😊", "👻👻", "👻👻👻👻😭😭"},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Replace(tt.text, tt.pattern, tt.newPattern)
                        if result != tt.expected{
                                t.Errorf("Replace(%v, %v, %v) = %v; want %v", tt.text, tt.pattern, tt.newPattern, result, tt.expected)
                        }
                })
        }
}

func TestReverse(t *testing.T){
        tests := []struct{
                name            string
                input           string
                expected        string
        }{
        	{"Standard", "Hello World", "dlroW olleH"},
		{"Singleton", "A", "A"},
		{"Empty", "", ""},
		{"Emoji", "😁", "😁"},
		{"Emoji & Text", "Hello😁World", "dlroW😁olleH"},
	}

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Reverse(tt.input)
                        if result != tt.expected{
                                t.Errorf("Reverse(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestSorting(t *testing.T){
        tests := []struct{
                name            string
                input           []int
                expected        []int
        }{
                {"Normal", []int{9,2,5,4,3,6,8,7,1}, []int{1,2,3,4,5,6,7,8,9}},
                {"Singleton", []int{1}, []int{1}},
                {"Empty", []int{}, []int{}},
                {"Negatives", []int{-9,-2,-5,-4,-3,-6,-8,-7,-1}, []int{-9,-8,-7,-6,-5,-4,-3,-2,-1}},
                {"Positives and Negatives", []int{1,2,3,0,-1,-2,-3}, []int{-3,-2,-1,0,1,2,3}},
        }

        for _, tt := range tests{
		t.Run("BubbelSort: " + tt.name, func(t *testing.T){
                        result := BubbleSort(tt.input)
                        if !reflect.DeepEqual(result, tt.expected){
                                t.Errorf("BubbleSort(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
		t.Run("InsertionSort: " + tt.name, func(t *testing.T){
                        result := InsertionSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected){
                                t.Errorf("InsertionSort(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
		t.Run("MergeSort: " + tt.name, func(t *testing.T){
                        result := MergeSort(tt.input)
                        if !reflect.DeepEqual(result, tt.expected){
                                t.Errorf("MergeSort(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
		t.Run("QuickSort: " + tt.name, func(t *testing.T){
                        result := QuickSort(tt.input)
                        if !reflect.DeepEqual(result, tt.expected){
                                t.Errorf("QuickSort(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestSplit(t *testing.T){
        tests := []struct{
                name            string
                text            string
                pattern         string
                expected        []string
        }{
		{"Standard", "Hello World", "o W", []string{"Hell", "orld"}},
		{"No Match", "Hello World", "Cats", []string{"Hello World"}},
		{"Empty Text", "", "Cats", []string{""}},
		{"Empty Pattern", "Hello World", "", []string{"Hello World"}},
		{"Match at the end", "Hello World", "d", []string{"Hello Worl", ""}},
		{"Emojis", "😁😁👻😁😁", "👻", []string{"😁😁","😁😁"}},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Split(tt.text, tt.pattern)
                        if !reflect.DeepEqual(result, tt.expected){
                                t.Errorf("Split(%v, %v) = %v; want %v", tt.text, tt.pattern, result, tt.expected)
                        }
                })
        }
}

func TestSwapCase(t *testing.T){
        tests := []struct{
                name            string
                input           string
                expected        string
        }{
                {"Standard", "Hello World", "hELLO wORLD"},
                {"Digit", "1", "1"},
                {"Empty", "", ""},
		{"Emoji", "Hello😁World", "hELLO😁wORLD"},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := SwapCase(tt.input)
                        if result != tt.expected{
                                t.Errorf("SwapCase(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestUpper(t *testing.T){
        tests := []struct{
                name            string
                input           string
                expected        string
        }{
                {"Standard", "hello world", "HELLO WORLD"},
                {"Digit", "1", "1"},
                {"Empty", "", ""},
		{"Emoji", "Hello😊World", "HELLO😊WORLD"},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Upper(tt.input)
                        if result != tt.expected{
                                t.Errorf("Upper(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}
