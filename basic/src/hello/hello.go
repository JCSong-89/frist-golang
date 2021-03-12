package main

import "fmt"

func main() {
	var a int
	var b int
	a = 3
	b = 4

	c := 4
	
	var d bool
	d = a == 3 &&  b > 2

	if c < 2 {
		fmt.Println("C출력", c)
	} else {
		fmt.Println("c미출력")
	}

	fmt.Println("Hello 월드")
	fmt.Println(a + b)
	fmt.Println(c)
	fmt.Printf("%v\n", c) // 타입추론하여 알아서 벨류 출력 
	fmt.Printf("%v\n", a|b) 
	fmt.Printf("%v\n", a&b) 
	fmt.Println("결과=", a^b)
	fmt.Println(d) 

}

/*
int  4/8 바이트
int32 4바이트
int64 8바이트
int8 1바이트 -128 ~ 127 uint8 0~255
int16 2바이트 -23768 ~ 32767 uint16 0~65535
*/

/*
float32 4바이트
float64 8바이트
*/

/*
string = rune[]
bool
rune 1~3 바이트
*/
