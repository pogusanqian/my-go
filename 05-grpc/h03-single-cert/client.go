package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pb "pogusanqin.com/student/stub"
)

func main() {
	// 读取证书
	creds, _ := credentials.NewClientTLSFromFile("cert/server.pem", "*.pogu.com")
	// 建立与服务端的连接
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// 退出时关闭连接
	defer conn.Close()

	// 获取客户端
	c := pb.NewStudentServiceClient(conn)

	// 设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 发起请求
	rsp, err := c.GetStudentById(ctx, &pb.GetStudentByIdReq{Id: 1001})
	if err != nil {
		log.Fatalf("请求GetStudentById接口失败: %v", err)
	}
	log.Printf("GetStudentById: %s", rsp)

	rsp2, _ := c.GetStudents(ctx, &emptypb.Empty{})
	log.Printf("GetStudents: %s", rsp2)
}
