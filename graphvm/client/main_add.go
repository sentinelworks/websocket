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

    // create graph
	req := &pb.AddRequest{Data: []string{"a", "b", "c", "d", "b", "d", "a", "c"}}
	if resp, err := c.Add(ctx, req); err == nil {
        log.Printf("Received: %T, GraphID is %d", resp, resp.Gid)
	} else {
        log.Fatalf("could not greet: %v", err)
	}

}
