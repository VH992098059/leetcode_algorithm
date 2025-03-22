package sequence_table

type SeqQueue struct {
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

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	n, _ := strconv.Atoi(scanner.Text())
// 	scanner.Scan()
// 	nums := strings.Fields(scanner.Text())

// 	queue := initQueue(n)

// 	for _, numStr := range nums {
// 		num, _ := strconv.Atoi(numStr)
// 		if !queue.insertQueue(num) {
// 			fmt.Println("错误：队列已满。")
// 		}
// 	}
// 	for i := 0; i <= n; i++ {
// 		if queue.isEmpty() {
// 			fmt.Println("错误：队列为空。")
// 			fmt.Println(queue.Peek()) // 强制返回-1
// 			continue
// 		}
// 		fmt.Println(queue.Peek())
// 		queue.popQueue()
// 	}
// 	fmt.Println("错误：队列为空。")
// }
