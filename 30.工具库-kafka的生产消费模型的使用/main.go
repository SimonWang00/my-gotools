package main

//File  : main.go
//Author: Simon
//Describe: describle your function
//Date  : 2021/1/5

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

var wg sync.WaitGroup

const (
	KafkaHost = "127.0.0.1:9092"
	Topic	  = "mytopic"		// 主题
	Group	  = "SimonWang00"	// 消费组
	Partition = 0				// 分区
)

func main()  {
	asyncMode()
	//syncMode()
}

// 同步模型
func syncMode()  {
	SyncProducer()
	SyncConsumer()
}

// 异步模型
func asyncMode() {
	go SaramaProducer()
	go SaramaConsumer()
	select {}
}

// 同步生产模型
func SyncProducer()  {
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V2_2_0_0
	fmt.Println("start make my kafka producer")

	//使用配置,新建一个同步生产者
	producer, e := sarama.NewSyncProducer([]string{KafkaHost}, config)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer producer.Close()

	for i:=0; i<5; i++ {
		//创建消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = Topic
		msg.Value = sarama.StringEncoder("this is a good test,hello kai")
		//发送消息
		pid, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		//time.Sleep(time.Second)
	}
}


// 同步消费模型
func SyncConsumer()  {
	fmt.Println("start consume")
	config := sarama.NewConfig()
	//consumer新建的时候会新建一个client，这个client归属于这个consumer，并且这个client不能用作其他的consumer
	consumer, err := sarama.NewConsumer([]string{KafkaHost}, config)
	if err != nil {
		panic(err)
	}

	//获取 kafka 主题
	partitions, err := consumer.Partitions(Topic)
	if err != nil {
		fmt.Println("get partitions failed, err:", err)
		return
	}

	for _, p := range partitions {
		//sarama.OffsetNewest：从当前的偏移量开始消费，sarama.OffsetOldest：从最老的偏移量开始消费
		partitionConsumer, err := consumer.ConsumePartition(Topic, p, sarama.OffsetNewest)
		if err != nil {
			fmt.Println("partitionConsumer err:", err)
			continue
		}
		wg.Add(1)
		go func(){
			for m := range partitionConsumer.Messages() {
				fmt.Println("message:", m)
				fmt.Printf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}


// SaramaProducer 异步生产者
func SaramaProducer() {
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V2_2_0_0
	fmt.Println("start make my kafka producer")

	//使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer([]string{KafkaHost}, config)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer producer.AsyncClose()
	//循环判断哪个通道发送过来数据.
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case suc := <-p.Successes():
				if suc != nil{
					fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
				}
			case fail := <-p.Errors():
				if fail != nil{
					fmt.Println("err: ", fail.Err)
					return
				}
			}
		}
	}(producer)

	// 发送数据
	var value string
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		time11 := time.Now()
		value = "this is a message from SimonWang00 " + time11.Format("15:04:05")
		// 发送的消息,主题。
		// 注意：这里的msg必须得是新构建的变量，不然你会发现发送过去的消息内容都是一样的，因为批次发送消息的关系。
		msg := &sarama.ProducerMessage{
			Topic: Topic,
		}
		//将字符串转化为字节数组
		msg.Value = sarama.ByteEncoder(value)
		//fmt.Println(value)
		//使用通道发送
		producer.Input() <- msg
	}
}


// SaramaConsumer 异步消费者
func SaramaConsumer()  {

	fmt.Println("start consume")
	config := sarama.NewConfig()

	//提交offset的间隔时间，每秒提交一次给kafka
	config.Consumer.Offsets.CommitInterval = 1 * time.Second

	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V0_10_0_1

	//consumer新建的时候会新建一个client，这个client归属于这个consumer，并且这个client不能用作其他的consumer
	consumer, err := sarama.NewConsumer([]string{KafkaHost}, config)
	if err != nil {
		panic(err)
	}

	//新建一个client，为了后面offsetManager做准备
	client, err := sarama.NewClient([]string{KafkaHost}, config)
	if err != nil {
		panic("client create error")
	}
	defer client.Close()

	//新建offsetManager，为了能够手动控制offset
	offsetManager,err:=sarama.NewOffsetManagerFromClient(Group,client)
	if err != nil {
		panic("offsetManager create error")
	}
	defer offsetManager.Close()

	//创建一个第2分区的offsetManager，每个partition都维护了自己的offset
	partitionOffsetManager,err:=offsetManager.ManagePartition(Topic,0)
	if err != nil {
		panic("partitionOffsetManager create error")
	}
	defer partitionOffsetManager.Close()


	fmt.Println("consumer init success")
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	//sarama提供了一些额外的方法，以便我们获取broker那边的情况
	topics,_:=consumer.Topics()
	fmt.Println(topics)
	partitions,_:=consumer.Partitions(Topic)
	fmt.Println(partitions)

	//第一次的offset从kafka获取(发送OffsetFetchRequest)，之后从本地获取，由MarkOffset()得来
	nextOffset,_:=partitionOffsetManager.NextOffset()
	fmt.Println(nextOffset)

	//创建一个分区consumer，从上次提交的offset开始进行消费
	partitionConsumer, err := consumer.ConsumePartition(Topic, Partition, nextOffset+1)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	fmt.Println("start consume really")

ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n message:%s", msg.Offset,string(msg.Value))
			//拿到下一个offset
			nextOffset,offsetString:=partitionOffsetManager.NextOffset()
			fmt.Println(nextOffset+1,"...",offsetString)
			//提交offset，默认提交到本地缓存，每秒钟往broker提交一次（可以设置）
			partitionOffsetManager.MarkOffset(nextOffset+1,"modified metadata")

		case <-signals:
			break ConsumerLoop
		}
	}
}
