package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpcdemo/test"
	"log"
	"net"
	"strconv"
)

type server struct {

}

func (s *server) DoMD5(ctx context.Context,in *test.Req) (*test.Res,error) {
	fmt.Println("MD5方法请求JSON:"+in.JsonStr+"，年龄："+in.Age+"，金额："+ strconv.Itoa(int(in.Price)))

	return &test.Res{BackJson:"MD5:"+fmt.Sprintf("%x",md5.Sum([]byte(in.JsonStr))),ResPrice:float32(in.Price + 10)},nil
}


func main()  {
	//监听所有网卡8028端口的TCP连接
	lis, err := net.Listen("tcp",":8028")
	if err!=nil {
		log.Fatalf("监听失败：%v",err)
	}
	//创建gRPC服务
	s := grpc.NewServer()
	/*
	* 注册接口服务
	* 以定义proto时的service为单位注册，服务中可以有多个方法
	* （proto编译时会为每个service生成Register***Server方法
	*  包.注册服务方法（grpc服务实例，包含接口方法的结构体【指针】）
	 */
	test.RegisterWaiterServer(s,&server{})

	/*
	*  如果有可以注册多个接口服务，结构体要实现对应接口方法
	* user.RegisterLoginServer(s,&server{})
	*
	*/
	reflection.Register(s)

	err = s.Serve(lis)
	if err !=nil{
		log.Fatalf("failed to serve: %v",err)
	}
}

