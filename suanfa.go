package main

/*import (
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
*/

/*type SeqQueue struct {
	data     []int // 存储元素的切片
	capacity int   // 队列容量
	front    int   // 队头指针
	rear     int   // 队尾指针
}

func initQueue(cap int) *SeqQueue {
	return &SeqQueue{
		data:     make([]int, 0, cap),
		capacity: cap,
		front:    0,
		rear:     0,
	}
}

func (s *SeqQueue) insertQueue(val int) bool {
	if s.rear >= s.capacity {
		return false
	}
	s.data = append(s.data, val)
	s.rear++
	return true
}
func (s *SeqQueue) popQueue() (int, bool) {
	if s.front >= s.rear {
		return 0, false
	}
	val := s.data[s.front]
	s.front++
	return val, true
}
func (s *SeqQueue) isEmpty() bool {
	return s.front >= s.rear
}
func (s *SeqQueue) Peek() int {
	if s.front >= s.rear {
		return -1
	}
	return s.data[s.front]
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	nums := strings.Fields(scanner.Text())

	queue := initQueue(n)

	for _, numStr := range nums {
		num, _ := strconv.Atoi(numStr)
		if !queue.insertQueue(num) {
			fmt.Println("错误：队列已满。")
		}
	}
	for i := 0; i <= n; i++ {
		if queue.isEmpty() {
			fmt.Println("错误：队列为空。")
			fmt.Println(queue.Peek()) // 强制返回-1
			continue
		}
		fmt.Println(queue.Peek())
		queue.popQueue()
	}
	fmt.Println("错误：队列为空。")
}*/

func main() {

}
