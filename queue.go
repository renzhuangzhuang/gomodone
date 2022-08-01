package queue

//定义队列数据结构
type (
	Queue struct {
		top    *node
		rear   *node
		length int
	}

	node struct {
		pre   *node
		net   *node
		value interface{}
	}
)

// 创建一个队列
func New() *Queue {
	return &Queue{nil, nil, 0}
}

//获取队列长度
func (this *Queue) Len() int {
	return this.length
}

//判断队列是否为空
func (this *Queue) Any() bool {
	return this.length > 0
}

//返回队列顶端元素
func (this *Queue) Peek() interface{} {
	if this.top == nil {
		return nil
	}
	return this.top.value
}

//入队操作
func (this *Queue) Pusk(v interface{}) {
	n := &node{nil, nil, v}
	if this.length == 0 {
		this.top = n
		this.rear = this.top
	} else {
		n.pre = this.rear
		this.rear.net = n
		this.rear = n
	}
	this.length++
}

// 出队操作
func (this *Queue) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.top
	if this.top.net == nil {
		this.top = nil
	} else {
		this.top = this.top.net
		this.top.pre.net = nil
		this.top.pre = nil
	}
	this.length--
	return n.value
}
