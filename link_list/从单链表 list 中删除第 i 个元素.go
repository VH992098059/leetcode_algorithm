package link_list

import "fmt"

type FifthNode struct { //链表中的数据
	Data int
	Next *FifthNode
}
type FifthLinkedList struct { //链表总体
	Head   *FifthNode
	Length int
}

func (l *FifthLinkedList) FifthInsertHead(val int) {
	newNode := &FifthNode{Data: val}
	newNode.Next = l.Head
	l.Head = newNode
	l.Length++
}
func (l *FifthLinkedList) FifthDeleteHead(pos int) error {
	if pos < 1 || pos > l.Length+1 {
		return fmt.Errorf("invalid position")
	}
	if pos == 1 {
		l.Head = l.Head.Next
	} else {
		current := l.Head
		for i := 1; i < pos-1; i++ {
			current = current.Next
		}
		current.Next = current.Next.Next
	}
	l.Length--
	return nil
}

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	n, _ := strconv.Atoi(scanner.Text())
// 	scanner.Scan()
// 	numsList := strings.Fields(scanner.Text())
// 	nums := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		nums[i], _ = strconv.Atoi(numsList[i])
// 	}
// 	list := &link_list.FifthLinkedList{}
// 	for _, num := range nums {
// 		list.FifthInsertHead(num)
// 	}
// 	scanner.Scan()
// 	i, _ := strconv.Atoi(scanner.Text())
// 	err := list.FifthDeleteHead(i)
// 	if err != nil {
// 		fmt.Println("错误：删除位置不合法。")
// 	}
// 	fmt.Printf("%d:", list.Length)
// 	current := list.Head
// 	for current != nil {
// 		fmt.Printf(" %d", current.Data)
// 		current = current.Next
// 	}
// 	fmt.Println()
// }
