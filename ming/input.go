package ming

import (
	"bufio"
	"fmt"
	"os"
)
//Input 键盘输入
func Input() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	i := 0
	for input.Scan() {
		//控制循环退出
		if input.Text() == "end" {
			break
		}
		i++
		counts[input.Text()] = i
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}