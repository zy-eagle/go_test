package writetest

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	data := []byte("C语言中文网, Go语言入门教程\r\n我都喜欢,但是,我不喜欢C#语言入门教程")
	fmt.Println("写入前未使用和已使用的缓冲区大小为：", w.Available(), w.Buffered())
	w.Write(data)
	fmt.Println("写入后未使用和已使用的缓冲区大小为：", w.Available(), w.Buffered())
	//将缓存区中的数据写入底层的IO.Writer
	w.Flush()
	fmt.Println("Flush后未使用和已使用的缓冲区大小为：", w.Available(), w.Buffered())
	fmt.Printf("Flush后缓冲区中的输出为：%q\n", string(wr.Bytes()))

	w.WriteByte('G')
	w.Flush()
	fmt.Printf("WriteByte后缓冲区中的输出为：%q\n", string(wr.Bytes()))

	var str = "我爱中华人民共和国"
	n, err := w.WriteString(str)
	w.Flush()
	fmt.Println("writeString:", string(wr.Bytes()), n, err)
}
