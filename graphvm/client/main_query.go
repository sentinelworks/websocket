package main

import (
    "log"
    "context"
    "time"
	"google.golang.org/grpc"

    pb "graphvm/graphvm"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

    defer conn.Close()
	c := pb.NewGraphvmClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // find shortest path
	req := &pb.QueryRequest{Gid: 0, First: "b", Second: "c"}
	if resp, err := c.Query(ctx, req); err == nil {
        log.Printf("Received: %T, shortest path is %d", resp, resp.Distance)
	} else {
        log.Fatalf("could not greet: %v", err)
	}

}
