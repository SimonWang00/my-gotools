package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/curve25519"
	"io"
	"os"
)

func main() {

	var Aprivate, Apublic [32]byte
	//产生随机数
	if _, err := io.ReadFull(rand.Reader, Aprivate[:]); err != nil {
		os.Exit(0)
	}
	curve25519.ScalarBaseMult(&Apublic, &Aprivate)
	fmt.Println("A私钥", base64.StdEncoding.EncodeToString(Aprivate[:]))
	fmt.Println("A公钥", base64.StdEncoding.EncodeToString(Apublic[:])) //作为椭圆起点

	var Bprivate, Bpublic [32]byte
	//产生随机数
	if _, err := io.ReadFull(rand.Reader, Bprivate[:]); err != nil {
		os.Exit(0)
	}
	curve25519.ScalarBaseMult(&Bpublic, &Bprivate)
	fmt.Println("B私钥",  base64.StdEncoding.EncodeToString(Bprivate[:]))
	fmt.Println("B公钥",  base64.StdEncoding.EncodeToString(Bpublic[:])) //作为椭圆起点

	var Akey, Bkey [32]byte

	//A的私钥加上Ｂ的公钥计算A的key
	curve25519.ScalarMult(&Akey, &Aprivate, &Bpublic)

	//B的私钥加上A的公钥计算B的key
	curve25519.ScalarMult(&Bkey, &Bprivate, &Apublic)

	fmt.Println("A交互的KEY",  base64.StdEncoding.EncodeToString(Akey[:]))
	fmt.Println("B交互的KEY",  base64.StdEncoding.EncodeToString(Bkey[:]))
}
