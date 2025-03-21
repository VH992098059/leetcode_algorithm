package sequence_table

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	data     []int // 用切片模拟栈存储
	capacity int   // 最大容量
}

func NewStack(capacity int) *Stack {
	return &Stack{
		data:     make([]int, 0, capacity),
		capacity: capacity,
	}
}

func (s *Stack) Push(val int) bool {
	if len(s.data) >= s.capacity {
		return false
	}
	s.data = append(s.data, val)
	return true
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, true
}

func (s *Stack) Peek() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	nums := strings.Fields(scanner.Text())

	stack := NewStack(n)

	// 压栈操作
	for _, numStr := range nums {
		num, _ := strconv.Atoi(numStr)
		if !stack.Push(num) {
			fmt.Println("错误：栈已满。")
		}
	}

	// 执行 n+1 次取顶并出栈
	for i := 0; i <= n; i++ {
		if stack.IsEmpty() {
			fmt.Println("错误：栈为空。")
			fmt.Println(stack.Peek())
			continue
		}

		// 取顶
		top := stack.Peek()
		fmt.Println(top)

		// 出栈
		stack.Pop()
	}
}
