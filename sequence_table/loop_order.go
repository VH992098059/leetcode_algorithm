package sequence_table

const MAXSIZE = 100

type SeQueue struct {
	data  [MAXSIZE]int
	front int
	rear  int
}

func initSeQueue() *SeQueue {
	return &SeQueue{
		front: 0,
		rear:  0,
		data:  [MAXSIZE]int{},
	}
}
func (q *SeQueue) isEmptySeQueue() bool {
	return q.front == q.rear
}
func (q *SeQueue) isMaxSeQueue() bool {
	return (q.rear+1)%MAXSIZE == q.front
}
func (q *SeQueue) lengthSeQueue() int {
	return (q.rear - q.front + MAXSIZE) % MAXSIZE
}
func (q *SeQueue) inputSeQueue(e int) {
	if q.isMaxSeQueue() {
		return
	}
	q.data[q.rear] = e
	q.rear = (q.rear + 1) % MAXSIZE
	return
}
func (q *SeQueue) outputSeQueue(e int) {
	if q.isEmptySeQueue() {
		return
	}
	e = q.data[q.front]
	q.front = (q.front + 1) % MAXSIZE
	return
}
