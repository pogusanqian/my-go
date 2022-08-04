package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	pb "pogusanqin.com/student/stub"
)

var stu = &pb.Student{
	Name: "张三",
	Age:  23,
	Sex:  "男",
}

type StudentServiceServerImpl struct {
	pb.UnimplementedStudentServiceServer
}

func (s *StudentServiceServerImpl) GetStudentById(ctx context.Context, req *pb.GetStudentByIdReq) (*pb.Student, error) {
	return stu, nil
}

func (s *StudentServiceServerImpl) GetStudents(ctx context.Context, req *emptypb.Empty) (*pb.GetStudentsRsp, error) {
	stus := *&pb.GetStudentsRsp{
		Students: []*pb.Student{stu, stu, stu},
	}
	return &stus, nil
}

func Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var user string
	var password string
	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}

	if user != "admin" || password != "admin" {
		return status.Errorf(codes.Unauthenticated, "token不合法")
	}
	return nil
}

func main() {
	// 启动监听
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 证书认证-双向认证
	cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	if err != nil {
		log.Fatal("证书读取错误", err)
	}
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert/ca.crt")
	if err != nil {
		log.Fatal("ca证书读取错误", err)
	}
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ClientAuth: tls.RequireAndVerifyClientCert,
		// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})

	// 实现token验证, 需要使用拦截器
	var authInterceptor grpc.UnaryServerInterceptor = func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		// 自定义验证token方法
		err = Auth(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	//  创建服务器
	s := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(authInterceptor))

	// 注册GreeterServer服务
	pb.RegisterStudentServiceServer(s, &StudentServiceServerImpl{})

	// 启动服务
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
