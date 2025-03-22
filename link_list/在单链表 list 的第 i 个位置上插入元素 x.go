package link_list

import (
	"fmt"
	"strconv"
	"strings"
)

type FourthNode struct { //链表中的数据
	data int
	next *FourthNode
}
type FourthLinkedList struct { //链表总体
	head   *FourthNode
	Length int
}

/*插入数据*/
func (list *FourthLinkedList) FourthInsert(pos int, value int) error {
	if pos < 1 || pos > list.Length+1 {
		return fmt.Errorf("错误：插入位置不合法。")
	}
	newFourthNode := &FourthNode{data: value}
	if pos == 1 {
		newFourthNode.next = list.head
		list.head = newFourthNode
	} else {
		current := list.head
		for i := 1; i < pos-1; i++ {
			current = current.next
		}
		newFourthNode.next = current.next
		current.next = newFourthNode
	}
	list.Length++
	return nil
}
func (list *FourthLinkedList) FourthStringList() string {
	var elements []string
	current := list.head
	for current != nil {
		elements = append(elements, strconv.Itoa(current.data))
		current = current.next
	}
	return strings.Join(elements, " ")
}

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	nums := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Scan(&nums[i])
// 	}

// 	list := &LinkedList{}
// 	for i := 0; i < n; i++ {
// 		list.Insert(i+1, nums[i])
// 	}

// 	fmt.Printf("%d: %s\n", list.Length, list.StringList())

// 	lastNum := nums[n-1]
// 	if err := list.Insert(0, lastNum); err != nil {
// 		fmt.Println(err)
// 	}
// 	if err := list.Insert(n+2, lastNum); err != nil {
// 		fmt.Println(err)
// 	}
// }
