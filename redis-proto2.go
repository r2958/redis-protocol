package main

import (
	"bufio"
	"fmt"
	"net"

	"redis-protocol/utils"
)


func main(){
	listen , err := net.Listen("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("listen failed", err)
	}
	for {
		conn , err := listen.Accept()
		if err !=nil {
			fmt.Println("Error happend when listen",err)
			continue
		}
		go process(conn)
	}

}

func process(conn net.Conn){
	//read data
	br := bufio.NewReaderSize(conn,1024*64)
	p, err := br.ReadBytes('\n')
	i := len(p) - 2
	content := p[:i] // *3
	content_length , _ := utils.ParseLen(content[1:])  //3

	var (
		bs []byte
		params [][]byte
		cmd string
	)
	for i:=0;i<content_length;i++{
		len , err := utils.ReadLen(br)  // $3  delete \r chracater
		if err!=nil{
			fmt.Println(err,len)
		}

		if bs, err = br.ReadBytes('\n'); err != nil {
		}
		if i == 0 {
			cmd = string(bs)
			continue
		}
		params = append(params,bs)
		//fmt.Println("content is ", string(bs))

	}

	fmt.Println("command is ", cmd)
	fmt.Println("paramaters is ",params)
	//send response

	bw := bufio.NewWriterSize(conn, 1024*64)
	if err := bw.WriteByte('+'); err!=nil{
		fmt.Println(err)
	}
	if ret,err := bw.WriteString("OK"); err!=nil{
		fmt.Println(ret,err)
	}
	_, err = bw.WriteString("\r\n")

	if err !=nil{
		fmt.Println(err)
		conn.Close()
		return
	}
	bw.Flush()
	conn.Close()
	return

}