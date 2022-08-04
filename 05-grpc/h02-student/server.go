package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
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

func main() {

	// 启动监听
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//  创建服务端
	s := grpc.NewServer()

	// 注册GreeterServer服务
	pb.RegisterStudentServiceServer(s, &StudentServiceServerImpl{})

	// 启动服务
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
