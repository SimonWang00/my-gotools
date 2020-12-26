package main

import (
	"log"
	"sync"
)

func main() {
	// 建立对象重用
	//对于很多需要重复分配、回收内存的地方，sync.Pool 是一个很好的选择。频繁地分配、回收内存会给 GC 带来一定的负担，
	//严重的时候会引起 CPU 的毛刺，而 sync.Pool 可以将暂时不用的对象缓存起来，待下次需要的时候直接使用，不用再次经过内存分配，
	//复用对象的内存，减轻 GC 的压力，提升系统的性能。
	var pipe = &sync.Pool{New:func()interface{}{return "new sync pool"}}
	// 准备放入的字符串
	val := "Hello, sync pool!"
	// 放入
	pipe.Put(val)
	// 取出
	log.Println(pipe.Get())
	// 再取就没有了,会自动调用NEW
	log.Println(pipe.Get())
}
