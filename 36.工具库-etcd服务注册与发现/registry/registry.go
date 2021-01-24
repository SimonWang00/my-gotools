package registry

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"hash/crc32"
	"log"
	"time"
)

var prefix = "/registry/server/"

// 注册接口
type Registry interface {
	RegistryNode(node PutNode) error
	UnRegistry()
}


type registryServer struct {
	cli        *clientv3.Client
	stop       chan bool
	isRegistry bool
	options    Options
	leaseID    clientv3.LeaseID
}

type PutNode struct {
	Addr string `json:"addr"`
}

type Node struct {
	Id   uint32 `json:"id"`
	Addr string `json:"addr"`
}

type Options struct {
	Name   string
	Ttl    int64
	Config clientv3.Config
}

func NewRegistry(options Options) (Registry, error) {
	cli, err := clientv3.New(options.Config)
	if err != nil {
		return nil, err
	}
	return &registryServer{
		stop:       make(chan bool),
		options:    options,
		isRegistry: false,
		cli:        cli,
	}, nil
}

func (s *registryServer) RegistryNode(put PutNode) error {
	if s.isRegistry {
		return errors.New("only one node can be registered")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.options.Ttl)*time.Second)
	defer cancel()
	// 创建租约时间ttl
	grant, err := s.cli.Grant(context.Background(), s.options.Ttl)
	fmt.Println("grant:",grant)
	if err != nil {
		return err
	}
	var node = Node{
		Id:   s.HashKey(put.Addr),
		Addr: put.Addr,
	}
	nodeVal, err := s.GetVal(node)
	if err != nil {
		return err
	}
	// 新增加key
	_, err = s.cli.Put(ctx, s.GetKey(node), nodeVal, clientv3.WithLease(grant.ID))
	if err != nil {
		return err
	}
	s.leaseID = grant.ID
	s.isRegistry = true
	go s.KeepAlive()
	return nil
}

func (s *registryServer) UnRegistry() {
	s.stop <- true
}

// 撤销租约
func (s *registryServer) Revoke() error {
	_, err := s.cli.Revoke(context.TODO(), s.leaseID)
	if err != nil {
		log.Printf("[Revoke] err : %s", err.Error())
	}
	s.isRegistry=false
	return err
}

// 保持续租的逻辑
func (s *registryServer) KeepAlive() error {
	keepAliveCh, err := s.cli.KeepAlive(context.TODO(), s.leaseID)
	if err != nil {
		log.Printf("[KeepAlive] err : %s", err.Error())
		return err
	}
	for {
		select {
		case <-s.stop:
			_ = s.Revoke()
			return nil
		case _, ok := <-keepAliveCh:
			if !ok {
				_ = s.Revoke()
				return nil
			}
		}
	}
}


func (s *registryServer) GetKey(node Node) string {
	key := fmt.Sprintf("%s%s/%d", prefix, s.options.Name, s.HashKey(node.Addr))
	fmt.Println(key)
	return key
}

func (s *registryServer) GetVal(node Node) (string, error) {
	data, err := json.Marshal(&node)
	return string(data), err
}

func (e *registryServer) HashKey(addr string) uint32 {
	hid := crc32.ChecksumIEEE([]byte(addr))
	fmt.Println("crc32.ChecksumIEEE([]byte(addr)):",crc32.ChecksumIEEE([]byte(addr)))
	return hid
}
