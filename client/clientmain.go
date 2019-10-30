package main

import (
	"context"
	"google.golang.org/grpc"
	"grpcdemo/test"
	"log"
	"os"
	"time"
)

func main() {

	//建立连接grpc服务

	conn,err := grpc.Dial("127.0.0.1:8028",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect : %v",err)
	}
	// 函数结束时关闭链接
	defer conn.Close()

	//创建waiter服务的客户端
	t := test.NewWaiterClient(conn)

	//模拟请求数据
	res := "test123"

	if len(os.Args)>1{
		res = os.Args[1]
	}

	// 调用grps接口
	rt ,err := t.DoMD5(context.Background(),&test.Req{JsonStr:res,Age:"20",Price:2.46})
	//测试定时循环
	for i:=0;i<1000 ; i++ {

		t.SayHello(context.Background(),&test.Req{JsonStr:res})
		time.Sleep(1*time.Second)
	}
	//rt1,err := t.SayHello(context.Background(),&test.Req{JsonStr:res})
	if err !=nil{
		log.Fatalf("colud not greet :%v",err)
	}

	log.Printf("服务端响应：%s,金额：%v",rt.BackJson,rt.ResPrice)

	//log.Println("sayhello 返回:",rt1.BackJson)
}
