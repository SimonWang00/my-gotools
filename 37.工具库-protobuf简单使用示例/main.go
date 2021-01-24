package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/7

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"my-gotools/37.工具库-protobuf简单使用示例/school"
)

func main() {
	s1:=&school.Student{} //第一个学生信息
	s1.Name="SimonWang00"
	s1.Age=18
	s1.Address="wuhan"
	s1.Cn=school.ClassName_class2 //枚举类型赋值
	ss:=&school.Students{}
	ss.Person=append(ss.Person,s1) //将第一个学生信息添加到Students对应的切片中
	s2:=&school.Student{}  //第二个学生信息
	s2.Name="jack ma"
	s2.Age=60
	s2.Address="hangzhou"
	s2.Cn=school.ClassName_class3
	ss.Person=append(ss.Person,s2)//将第二个学生信息添加到Students对应的切片中
	ss.School="杭州师范"
	fmt.Println("Students信息为：",ss)

	// Marshal takes a protocol buffer message
	// and encodes it into the wire format, returning the data.
	buffer, _ := proto.Marshal(ss)
	fmt.Println("序列化之后的信息为：",buffer)
	// 	Use UnmarshalMerge to preserve and append to existing data.
	data:=&school.Students{}
	proto.Unmarshal(buffer,data)
	fmt.Println("反序列化之后的信息为：",data)
}
