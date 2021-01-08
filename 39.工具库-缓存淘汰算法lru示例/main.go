package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/7

import (
	"container/list"
	"fmt"
)

type LRU struct {
	maxByte  int64
	useByte  int64
	ll       *list.List
	cache    map[string]*list.Element
	CallBack func(key string, val []byte)
}

type Node struct {
	Key string
	Val []byte
}

func NewLRU(maxByte int64, callBack func(key string, val []byte)) *LRU {
	return &LRU{
		maxByte:  maxByte,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
		CallBack: callBack,
	}
}

func (g *LRU) Add(key string, data []byte) {
	if val, ok := g.cache[key]; ok {
		g.ll.MoveToFront(val)
	} else {
		g.useByte += int64(len(key)) + int64(len(data))
		val := g.ll.PushFront(&Node{
			Key: key,
			Val: data,
		})
		g.cache[key] = val

	}
	for g.maxByte != 0 && g.maxByte < g.useByte {
		g.Remove()
	}
}

func (g *LRU) Del(key string) {
	g.Remove(key)
}

func (g *LRU) Remove(k ...string) {
	var (
		val *list.Element
		ok  bool
	)
	if len(k) == 0 {
		val = g.ll.Back()
	} else {
		if val, ok = g.cache[k[0]]; !ok {
			return
		}
	}
	if val != nil {
		g.ll.Remove(val)
		node := val.Value.(*Node)
		delete(g.cache, node.Key)
		g.useByte -= int64(len(node.Key)) + int64(len(node.Val))
		if g.CallBack != nil {
			g.CallBack(node.Key, node.Val)
		}
	}
}

func (g *LRU) Get(key string) (v []byte, ok bool) {
	var val *list.Element
	if val, ok = g.cache[key]; ok {
		g.ll.MoveToFront(val)
		node := val.Value.(*Node)
		v = node.Val
		return
	}
	return
}

func (g *LRU) Len() int {
	return g.ll.Len()
}


/*
LRU算法的设计原则是：如果一个数据在最近一段时间没有被访问到，那么在将来它被访问的可能性也很小。
也就是说，当限定的空间已存满数据时，应当把最久没有被访问到的数据淘汰。


利用一个链表来实现，每次新插入数据的时候将新数据插到链表的头部；每次缓存命中（即数据被访问），
则将数据移到链表头部；那么当链表满的时候，就将链表尾部的数据丢弃。

*/
func main() {
	lru := NewLRU(2, nil)
	lru.Add("1", []byte("1"))
	val, ok := lru.Get("1")
	if !ok {
		fmt.Println("lru not get")
	}
	if string(val) != "1" {
		fmt.Println("lru  get val err")
	}
	fmt.Println("val : ", string(val))
	lru.Del("1")
	_, ok = lru.Get("1")
	if ok {
		fmt.Println("lru del err")
	}
}
