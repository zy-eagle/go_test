package readtest

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

// Reader 从缓冲区读数据
func TestRead(t *testing.T) {

	data := []byte("C语言中文网, Go语言入门教程\r\n我都喜欢,但是,我不喜欢C#语言入门教程")
	fmt.Printf("%T\n", data)
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)

	var buf [10]byte
	//Read
	n, err := r.Read(buf[:])
	fmt.Println("Read:", string(buf[:n]), n, err)
	//ReadByte
	c, err := r.ReadByte()
	fmt.Println("ReadByte:", string(c), err)
	//ReadBytes
	var delim byte = ','
	line, err := r.ReadBytes(delim)
	fmt.Println("ReadBytes:", string(line), err)
	//Readline
	l, prefix, err := r.ReadLine()
	fmt.Println("Readline:", string(l), prefix, err)
	//ReadRune
	ch, size, err := r.ReadRune()
	fmt.Println("ReadRune:", string(ch), size, err)
	//ReadSlice
	sl, err := r.ReadSlice(delim)
	fmt.Println("ReadSlice:", string(sl), err)
	s, err := r.ReadString(delim)
	fmt.Println("ReadString:", s, err)
	//Buffered 读取当前缓存大小
	fmt.Println("Buffered:", r.Buffered())
	//Peek, 不会清除缓存数据
	p1, err := r.Peek(12)
	fmt.Println("Peek:", string(p1), err)
	p2, err := r.Peek(14)
	fmt.Println("Peek:", string(p2), err)
}
