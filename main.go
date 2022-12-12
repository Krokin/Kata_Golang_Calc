package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	numOne int
	numTwo int
	operand string
}

const (
	RIM int8 = 1
	ARAB int8 = 0
	ERR int8 = -1
)



func main() {
	str := input()
	operation, types, err := validationAndParse(str)
	if err == false {
		fmt.Println("Err input")
		return
	}
	calc(*operation, types)
}

func input() string {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	return s.Text()
}

func validationAndParse(str string) (*Operation, int8, bool) {
	o := strings.Split(str, " ")

	if len(o) != 3 || !checkOperand(o[1]) {
		return nil, ERR, false
	}
	oper := &Operation{operand: o[1]}

	if checkArab(oper, o) {
		return oper, ARAB, true
	}

	if checkRim(oper, o) {
		return oper, RIM, true
	}
	return nil, ERR, false
}

func checkOperand(s string) bool {
	if strings.Contains("+-/*", s) && len(s) == 1 {
		return true
	}
	return false
}

func checkArab(oper *Operation, o []string) (ok bool) {
	a, err := strconv.Atoi(o[0])
	if err != nil || a > 10 || a < 1 {
		ok = false
		return
	}
	b, err := strconv.Atoi(o[2])
	if err != nil || b > 10 || b < 1 {
		ok = false
		return
	}
	oper.numOne, oper.numTwo = a, b
	ok = true
	return
}	

func checkRim(oper *Operation, o []string) bool {
	rim := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	a , ok := rim[o[0]]
	if !ok {
		return false
	}
	oper.numOne = a

	a , ok = rim[o[2]]
	if !ok {
		return false
	}
	oper.numTwo = a
	
	return true
}

func calc(oper Operation, t int8) {
	res := 0
	switch oper.operand {
	case "+":
		res = oper.numOne + oper.numTwo
	case "-":
		res = oper.numOne - oper.numTwo
	case "/":
		res = oper.numOne / oper.numTwo
	case "*":
		res = oper.numOne * oper.numTwo
	}

	if t == RIM {
		if res < 1 {
			fmt.Printf("Результат нельзя отобразить римскими цифрами: %d", res)
			return
		}
		fmt.Println(outputRim(res))
	} else {
		fmt.Println(res)
	}
}

func outputRim(res int) string {
	fmt.Println(res)
	if res == 0 {
		return ""
	} else if res < 4 {
		return "I" + outputRim(res - 1)
	} else if res < 9 {
		if res == 4 {
			return "IV" + outputRim(res - 4)
		}
		return "V" + outputRim(res - 5)
	} else if res < 40 {
		if res == 9 {
			return "IX" + outputRim(res - 9)
		}
		return "X" + outputRim(res - 10)
	} else if res < 50 {
		return "XL" + outputRim(res - 40)
	} else if res < 90 {
		return "L" + outputRim(res - 50)
	} else if res < 100 {
		return "XC" + outputRim(res - 90)
	} else if res == 100 {
		return "C" + outputRim(res - 100)
	}
	return ""
}