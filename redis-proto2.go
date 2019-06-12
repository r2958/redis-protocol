package main

import (
	"bufio"
	"fmt"
	"net"
	"errors"
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
	//fmt.Println("content is ",string(p), "raw data is", p, "length is ", len(p))
	i := len(p) - 2
	//fmt.Println(p[:i] , string(p[:i])) // delete "\n" byte  == *3
	content := p[:i] // *3

	content_length , _ := parseLen(content[1:])  //3

	//fmt.Println("get content length is",content_length, content[1:])


	var bs []byte

	var (
		params [][]byte
		cmd string
	)

	for i:=0;i<content_length;i++{

		len , err := readLen(br)  // $3  delete \r chracater
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

	fmt.Println("get Data is ", cmd, params)


	//send response

	bw := bufio.NewWriterSize(conn, 1024*64)
	bw.WriteByte('+')
	bw.WriteString("OOKKK-Andy")
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



func parseLen(p []byte) (int, error) {
	if len(p) == 0 {
		return -1, errors.New("malformed length")
	}
	if p[0] == '-' && len(p) == 2 && p[1] == '1' {
		// handle $-1 and $-1 null replies.
		return -1, nil
	}
	var n int

	for _, b := range p {
		n *= 10

		if b < '0' || b > '9' {
			return -1, errors.New("illegal bytes here in length")
		}
		n += int(b - '0')
	}
	return n, nil
}



func readLen(br *bufio.Reader) (int, error) {
	//prefix :=byte('$')
	ls, err := br.ReadBytes('\n')

	ls = ls[:len(ls)-2]  // delete \n chracaters

	if err != nil {
		return 0, err
	}
	if len(ls) < 2 {

		return 0, errors.New("illegal bytes ddd in length")
	}
	if ls[0] != '$' { // start flag

		return 0, errors.New("illegal bytes bbb  in length")
	}


	return parseLen(ls[1:])
}