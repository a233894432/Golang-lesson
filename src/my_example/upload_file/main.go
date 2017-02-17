package main

import (
	"fmt"

	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

var (
	//设置上传到的空间
	bucket = "dtukuimg"
)

//构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func main() {
	//初始化AK，SK
	conf.ACCESS_KEY = "pUr2s6VcRoB8T0MKe6dmmlhDpQcqacnbJollpdL4"
	conf.SECRET_KEY = "19DhOANFb9vAkb_9o51NPIWdCuvurwjGqDojN6Gl"

	//创建一个Client
	c := kodo.New(0, nil)

	//设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: bucket,
		//设置Token过期时间
		Expires: 3600,
	}
	//生成一个上传token
	token := c.MakeUptoken(policy)

	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)

	var ret PutRet
	//设置上传文件的路径
	filepath := "E:/mygithub/Golang-lesson/src/my_example/sample1.jpg"
	//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	res := uploader.PutFileWithoutKey(nil, &ret, token, filepath, nil)
	//打印返回的信息
	fmt.Println(ret)
	//打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return
	}

}
