package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/11

import (
	"context"
	"fmt"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
	"io"
	"log"
	"math"
	"strconv"
	"time"
)

var (
	HOST   = "127.0.0.1"
	TABLE  = "my_test"
	client 		gohbase.Client
	adminClient gohbase.AdminClient
)

//每一秒产生一个
func GetRowKey(uid uint32) string {
	return fmt.Sprintf("%s%d", strconv.Itoa(int(uid)), math.MaxInt64-time.Now().Unix())
}

func reverse(str string) string {
	var result string
	length := len(str)
	for i := 0; i < length; i++ {
		result += fmt.Sprintf("%c", str[length-i-1])
	}
	return result
}

func initHbase()  {
	client = gohbase.NewClient(HOST)
	adminClient = gohbase.NewAdminClient(HOST)
}

func init()  {
	initHbase()
}

// CreateTable 创建表
func CreateTable()  error{
	fam := map[string]map[string]string{
		"info":{
			// 是否使用布隆过虑及使用何种方式
			// ROW: 对 ROW，行键的哈希在每次插入行时将被添加到布隆。
			// ROWCOL: 行键 + 列族 + 列族修饰的哈希将在每次插入行时添加到布隆
			// NONE:
			"BLOOMFILTER":         "ROW",
			// 数据保留的版本数, 如果不需要保留老版本数据, 可以调低此参数来节约磁盘空间
			"VERSIONS":            "3",
			// 是否常驻内存
			"IN_MEMORY":           "false",
			//
			"KEEP_DELETED_CELLS":  "false",
			//
			"DATA_BLOCK_ENCODING": "FAST_DIFF",
			// 默认是 2147483647 即:Integer.MAX_VALUE 值大概是68年
			// 这个参数是说明该列族数据的存活时间，单位是s
			// 这个参数可以根据具体的需求对数据设定存活时间，超过存过时间的数据将在表中不在显示，
			// 待下次major compact的时候再彻底删除数据.
			// 注意的是TTL设定之后 MIN_VERSIONS=>’0’ 这样设置之后，TTL时间戳过期后，
			// 将全部彻底删除该family下所有的数据，如果MIN_VERSIONS 不等于0那将保留
			// 最新的MIN_VERSIONS个版本的数据，其它的全部删除，比如MIN_VERSIONS=>’1’ 届时将保留一个最新版本的数据，
			// 其它版本的数据将不再保存。
			"TTL":                 "2147483647",
			// 指定列族是否采用压缩算法以及使用什么压缩算法NONE: 不使用压缩
			// SNAPPY/ZIPPY: 压缩比22.2%, 压缩172MB/s, 解压409MB/s
			// LZO: 压缩比20.5%, 压缩135MB/s, 解压410MB/s
			// GZIP: 压缩比13.4%, 压缩21MB/s, 解压118MB/s
			"COMPRESSION":         "NONE",
			// major compact的后保留的最少版本数
			"MIN_VERSIONS":        "0",
			//
			"BLOCKCACHE":          "true",
			//
			"BLOCKSIZE":           "65536",
			//
			"REPLICATION_SCOPE":   "0",
		},
	}
	table := hrpc.NewCreateTable(context.Background(), []byte(TABLE), fam)
	err := adminClient.CreateTable(table)
	if err != nil{
		log.Printf("Create hbase table failed! err:%v", err.Error())
		return err
	}
	log.Println("create table success!")
	return nil
}

// GetRow 查询单条数据
func GetRow(key string)  {
	newGetStr, err := hrpc.NewGetStr(context.Background(), TABLE, key)
	result, err := client.Get(newGetStr)
	if err != nil{
		panic(err)
	}
	for _, v := range result.Cells{
		log.Printf("%v : %v \n", v.Qualifier, string((v.Value)))
	}
}

// AddRow 插入一条数据
func AddRow()  {
	rowKey := "88888888888888888" // RowKey
	value := map[string]map[string][]byte {
			// 列族名, 与创建表时指定的名字保持一致
			"info": {
				"name": []byte("SimonWang00"),  // 列与值, 列名可自由定义
				"age":  []byte("18"),
		},
	}
	putRequest, err := hrpc.NewPutStr(context.Background(), TABLE, rowKey, value)
	_, err = client.Put(putRequest)
	if err != nil{
		panic(err)
	}
	log.Println("add one data success !")
}

// DelRow 删除一条数据
func DelRow(key string)  {
	newDel, err := hrpc.NewDel(context.Background(), []byte(TABLE), []byte(key), nil)
	result, err := client.Delete(newDel)
	if err != nil{
		panic(err)
	}
	log.Println("delete row success! id = ", result.String())
}

// GetFamily查询指定字段
func GetFamily()  {
	family := map[string][]string{"info": []string{"name"}}
	newGetStr, err := hrpc.NewGetStr(context.Background(), TABLE, "88888888888888888", hrpc.Families(family))
	result, err := client.Get(newGetStr)
	if err != nil {
		panic(err)
	}
	for _, v := range result.Cells {
		val := v
		log.Println(string(val.Qualifier), string(val.Value))
	}
}

// ScanStr前缀匹配查询
func ScanStr()  {
	//rowKey 前缀匹配
	pFilter := filter.NewPrefixFilter([]byte("88"))
	q, err := hrpc.NewScanStr(context.Background(), TABLE, hrpc.Filters(pFilter))
	if err != nil {
		panic(err)
		return
	}
	scanRsp := client.Scan(q)
	for {
		res, err := scanRsp.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		for _, v := range res.Cells {
			val := v
			log.Println(string(val.Qualifier), string(val.Value))
		}
	}
}

// ScanRangeStr分页查询
func ScanRangeStr(starts , ends string)  {
	f := filter.NewPageFilter(10)
	//q, err := hrpc.NewScanRangeStr(context.Background(), "user", startKey, endKey)
	q, err := hrpc.NewScanRangeStr(context.Background(), TABLE, starts, ends, hrpc.Filters(f))
	if err != nil {
		panic(err)
		return
	}
	scanRsp := client.Scan(q)
	for {
		res, err := scanRsp.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		for _, v := range res.Cells {
			val := v
			log.Println(string(val.Qualifier), string(val.Value), string(val.Row))
		}
	}
}

func main() {
	//CreateTable()
	//AddRow()
	//GetRow("88888888888888888")
	//DelRow("88888888888888888")
	//AddRow()
	//GetFamily()
	ScanStr()
}
