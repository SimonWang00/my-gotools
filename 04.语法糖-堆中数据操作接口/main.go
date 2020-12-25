package main

import (
	"container/heap"
	"log"
)

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/25

// Item队列
type Queue []*Item

type Item struct {
	data  interface{}
	ref   int //优先级
	index int //在堆里面的索引
}

// 队列长度
func (m Queue) Len() int {
	return len(m)
}

// i大于j 则true 否则false
func (m Queue) Less(i, j int) bool {
	return m[i].ref > m[j].ref
}

// i和j位置的数据交换
func (m Queue) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
	m[j].index = j
	m[i].index = i
}


// 往容器最后一位插入, 强转d的类型为Item
func (m *Queue) Push(d interface{}) {
	d.(*Item).index = len(*m)
	*m = append(*m, d.(*Item))
}

// 抛出后面的一个数据
func (m *Queue) Pop() interface{} {
	l := len(*m)
	s := (*m)[l-1]
	s.index = -1
	*m = (*m)[:l-1]
	return s
}


func main() {
	queue := make(Queue, 10)
	for i := 0; i < 10; i++ {
		item:=&Item{
			data:  i + 1,
			ref:   i + 1,
			index: i,
		}
		queue[i] = item
	}
	// 堆初始化
	heap.Init(&queue)
	item := Item{
		data: 8,
		ref:  1,
	}
	heap.Push(&queue, &item)
	// 更改index为2之处的值后,重新排序
	heap.Fix(&queue, 2)
	for queue.Len() > 0 {
		// 删除最后一位
		item := heap.Pop(&queue).(*Item)
		log.Println("index", item.index, "ref", item.ref, "val", item.data)
	}
}
