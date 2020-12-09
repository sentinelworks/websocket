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

    // delete graph
    var i int32 = 1;
	req := &pb.DeleteRequest{Gid: i}
	if resp, err := c.Delete(ctx, req); err == nil {
        log.Printf("Received: %T, GraphID delete state is %d", resp, resp.Result)
	} else {
        log.Fatalf("could not greet: %v", err)
	}
}
