package main

import "fmt"

func balancedBracket(brackets string) string {
	var checkArr []rune

	for _, char := range brackets {
		if char == '{' || char == '[' || char == '(' {
			checkArr = append(checkArr, char)
		} else if char == '}' {
			if len(checkArr) == 0 || checkArr[len(checkArr)-1] != '{' {
				return "NO"
			}
			checkArr = checkArr[:len(checkArr)-1]
		} else if char == ']' {
			if len(checkArr) == 0 || checkArr[len(checkArr)-1] != '[' {
				return "NO"
			}
			checkArr = checkArr[:len(checkArr)-1]
		} else if char == ')' {
			if len(checkArr) == 0 || checkArr[len(checkArr)-1] != '(' {
				return "NO"
			}
			checkArr = checkArr[:len(checkArr)-1]
		}
	}

	if len(checkArr) == 0 {
		return "YES"
	}

	return "NO"
}

func main() {
	var brackets string
	var valid bool

	fmt.Scan(&brackets)
	for i := 0; !valid && i < len(brackets); i++ {
		if brackets[i] == '{' || brackets[i] == '}' || brackets[i] == '[' || brackets[i] == ']' || brackets[i] == '(' || brackets[i] == ')' {
			continue
		} else {
			fmt.Println("Invalid input")
			valid = true
		}
	}

	if !valid {
		fmt.Println(balancedBracket(brackets))
	}
}
