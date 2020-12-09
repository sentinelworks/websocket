package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
    "log"

    //pb "google.golang.org/grpc/examples/graphvm/graphvm"
    pb "graphvm/graphvm"

)

type server struct{
    pb.UnimplementedGraphvmServer
}

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

    pb.SystemStart()
	srv := grpc.NewServer()
	pb.RegisterGraphvmServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
    req := in.GetData()
    log.Printf("Received: %T, %v", req, req)
    // we can use multithreaded here to improve the potential performance
    ret := pb.PostGraph(req)
    return &pb.AddResponse{Gid: ret}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
    req := in.GetGid()
    log.Printf("Received: %T, %v", req, req)
    res := pb.DeleteGraph(req)
    return &pb.DeleteResponse{Result: res}, nil
}

func (s *server) Query(ctx context.Context, in *pb.QueryRequest) (*pb.QueryResponse, error) {
    req := in.GetGid()
    first := in.GetFirst()
    second := in.GetSecond()
    log.Printf("Received: %T, %v, fist %s, second %s", req, req, first, second)
    dist := pb.GetShortest(req, first, second);
    return &pb.QueryResponse{Distance: dist}, nil
}
