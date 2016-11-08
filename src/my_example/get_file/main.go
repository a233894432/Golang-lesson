package main

import (
	"fmt"

	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

var (
	//设置需要操作的空间
	bucket = "dtukuimg"
	//设置需要操作的文件的key
	key = "FnGTTISMfKXpAET23Sm10nu3YQ4I"
)

func main() {

	conf.ACCESS_KEY = "pUr2s6VcRoB8T0MKe6dmmlhDpQcqacnbJollpdL4"
	conf.SECRET_KEY = "19DhOANFb9vAkb_9o51NPIWdCuvurwjGqDojN6Gl"

	//new一个Bucket管理对象
	c := kodo.New(0, nil)
	p := c.Bucket(bucket)

	//调用Stat方法获取文件的信息
	entry, err := p.Stat(nil, key)
	//打印列取的信息
	fmt.Println(entry)
	//打印出错时返回的信息
	if err != nil {
		fmt.Println(err)
	}
}
