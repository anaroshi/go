package main

import (
	"fmt"
	//"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/anaroshi/learngo/decorator/cipher"
	"github.com/anaroshi/learngo/decorator/lzw"
)

func checkErr(err error) {
	if err!= nil {
		panic(err)
	}
}

type Component interface {
	Operator(string)
}

var sentData string
var ReceiveData string

type SendComponent struct {}

func (selfy *SendComponent) Operator(data string) {
	sentData = data
}

type ReceiveComponent struct {}

func (selfy *ReceiveComponent) Operator(data string) {
	ReceiveData = data
}

// 압축
type ZipComponent struct {
	com Component
}

func (selfy *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	checkErr(err)
	selfy.com.Operator(string(zipData))
}

// 암호화
type EncryptComponent struct {
	key string
	com Component
}

func (selfy *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), selfy.key)
	checkErr(err)
	selfy.com.Operator(string(encryptData))
}	


// ----------------------------------------------------

type DecryptComponent struct {
	key string
	com Component
}

func (selfy *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), selfy.key)
	checkErr(err)
	selfy.com.Operator(string(decryptData))
}

type UnzipComponent struct {
	com Component
}

func (selfy *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	checkErr(err)
	selfy.com.Operator(string(unzipData))
}


func main() {
	sender := &EncryptComponent{ key:"abcde", 
	com: &ZipComponent{
		com: &SendComponent{},
	}}

	sender.Operator("Hello World")
	fmt.Println(sentData)

	receiver := &UnzipComponent{		
		com: &DecryptComponent{
			key:"abcde", 
			com: &ReceiveComponent{},
	}}

	receiver.Operator(sentData)
	fmt.Println(ReceiveData)

}	

//9번 보는중(decorator패턴)