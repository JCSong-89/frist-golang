package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
 fmt.Println("숫자를 입력하세요")
 reader := bufio.NewReader(os.Stdin)
 line, _ := reader.ReadString('\n')
 line = strings.TrimSpace(line)

 n1, _  := strconv.Atoi(line)

 fmt.Printf("입력하신 숫자는 %d 입니다.", n1)
}