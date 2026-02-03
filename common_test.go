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
                {"3 Trues", []bool{true,true,true}, true},
                {"3 Falses", []bool{false,false,false}, false},
                {"Empty", []bool{}, true}, //The empty intersection is the whole space so thus true here
                {"One Of Each", []bool{true, false},false},
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
                {"3 Trues", []bool{true,true,true}, true},
                {"3 Falses", []bool{false,false,false}, false},
                {"Empty", []bool{}, false}, //The empty union is the empty set so thus false here
                {"One Of Each", []bool{true, false},true},
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




















