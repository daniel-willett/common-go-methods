package main

import ("testing";"reflect")

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

func TestBubbleSort(t *testing.T){
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
                t.Run(tt.name, func(t *testing.T){
                        result := BubbleSort(tt.input)
                        if !reflect.DeepEqual(result, tt.expected){
                                t.Errorf("BubbleSort(%v) = %v; want %v", tt.input, result, tt.expected)
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

func TestIsPalendrome(t *testing.T){
        tests := []struct{
                name            string
                input		uint	
                expected        bool
        }{
                //Normal cases
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
		{"Too Few Blocks", "0", false},
		{"Random Characters", "dasjkhdj", false},
		{"Too Many Blocks", "0:0:0:0:0:0:0:0:0", false},
		{"Too Many Blocks IPv4 Embedded", "0:0:0:0:0:0:0:1.2.3.4", false},
		{"Too Many Octets IPv4 Embedded", "0:0:0:0:0:0:1.2.3.4.5", false},
		{"Too Few Blocks IPv4 Embedded", "12:12.12.32.4", false},
		{"Too Large Block", "10000:0:0:0:0:0:0:0", false},
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

func TestMax(t *testing.T){
        tests := []struct{
                name            string
                input           []int
                expected        int
        }{
		{"Signleton", []int{5}, 5},
		{"Positive Values", []int{1,2,3,4,5}, 5},
		{"Negative Values", []int{-1,-2,-3,-4,-5}, -1},
		{"Positive & Negative Values", []int{1,-1,2,-2,3,-3}, 3},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Max(tt.input)
                        if result != tt.expected{
                                t.Errorf("Max(%v) = %v; want %v", tt.input, result, tt.expected)
                        }
                })
        }
}

func TestMin(t *testing.T){
        tests := []struct{
                name            string
                input           []int
                expected        int
        }{
                {"Signleton", []int{5}, 5},
                {"Positive Values", []int{1,2,3,4,5}, 1},
                {"Negative Values", []int{-1,-2,-3,-4,-5}, -5},
                {"Positive & Negative Values", []int{1,-1,2,-2,3,-3}, -3},
        }

        for _, tt := range tests{
                t.Run(tt.name, func(t *testing.T){
                        result := Min(tt.input)
                        if result != tt.expected{
                                t.Errorf("Min(%v) = %v; want %v", tt.input, result, tt.expected)
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
