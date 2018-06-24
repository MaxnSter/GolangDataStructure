package queue

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue {
	return &Queue{data: make([]interface{}, 0)}
}

func (q *Queue) PushBack(x interface{}) {
	q.data = append(q.data, x)
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) Empty() bool {
	return q.Size() <= 0
}

func (q *Queue) Peek() interface{} {
	if q.Empty() {
		panic("peek at empty queue")
	}

	return q.data[0]
}

func (q *Queue) PopFront() {
	if q.Empty() {
		return
	}

	q.data = q.data[1:]
	return
}
