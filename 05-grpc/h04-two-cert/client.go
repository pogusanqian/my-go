package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pb "pogusanqin.com/student/stub"
)

func main() {
	// 证书认证-双向认证
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.crt")
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ServerName: "*.pogu.com",
		RootCAs:    certPool,
	})
	
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
