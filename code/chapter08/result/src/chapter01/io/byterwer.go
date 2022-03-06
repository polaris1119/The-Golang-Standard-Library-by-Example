package main

import (
	"bytes"
	"fmt"
	"os"
)

func ByteRWerExample() {
FOREND:
	for {
		fmt.Println("请输入要通过WriteByte写入的一个ASCII字符（b：返回上级；q：退出）：")
		var ch byte
		fmt.Scanf("%c\n", &ch)
		switch ch {
		case 'b':
			fmt.Println("返回上级菜单！")
			break FOREND
		case 'q':
			fmt.Println("程序退出！")
			os.Exit(0)
		default:
			buffer := new(bytes.Buffer)
			err := buffer.WriteByte(ch)
			if err == nil {
				fmt.Println("写入一个字节成功！准备读取该字节……")
				newCh, _ := buffer.ReadByte()
				fmt.Printf("读取的字节：%c\n", newCh)
			} else {
				fmt.Println("写入错误")
			}
		}

	}
}
