package structs

type Node struct {
	value int
	next  *Node
}
type CyclicList struct {
	head *Node
	tail *Node
	size int
}

func CreateList(size int) *CyclicList {
	list := CyclicList{}
	for i := 0; i < size; i++ {
		list.Add(i)
	}
	return &list
}

func (list *CyclicList) Length() int {
	return list.size
}

func (list *CyclicList) Add(element int) Node {
	if list.head == nil {
		list.head = &Node{element, list.tail}
		list.head.next = list.head
		list.tail = list.head
	} else {
		newNode := &Node{element, list.head}
		list.tail.next = newNode
		list.tail = newNode
	}
	list.size++
	return *list.tail
}

func (list *CyclicList) ToArray() []int {
	if list.head == nil {
		return []int{}
	}
	elements := make([]int, 0, list.size)
	temp := list.head
	for i := 0; i < list.size; i++ {
		elements = append(elements, temp.value)
		temp = temp.next
	}
	return elements
}

func (list *CyclicList) Get(index int) Node {
	if list.head == nil {
		panic("cannot get from empty list")
	}

	index = ((index % list.size) + list.size) % list.size

	temp := list.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return *temp
}
