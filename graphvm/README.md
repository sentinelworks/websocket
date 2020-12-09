
Prequiste: You must install go, grpc, protobuf3.  To verify, make helloworld example work.  Move graphvm under $GOPATH/src/google.golang.org/grpc/examples

The design point:
    1. graph ID is allocated by using microservice
    2. Create graph data structure to do different algorithm, such as shortest path
    3. Use one graph repo to manage all posted graphs.
    4. Provide API to add/query(shortest path)/delete graph
    5. package are managed in local folder, not good but easy to run

Three APIs are define in graphvm.proto
    use protoc to generate stub: Add, Delete, Query
    
How to test?
    1. under server/ folder, run go run server/main.go
       then under client/ folder run three programs: go run client/main_add.go
    2. under test/ folder, you can test it without web

The test result:
    on client side
victor@johnson:~/go/src/google.golang.org/grpc/examples/graphvm$ go run client/main_delete.go 
2020/08/15 16:48:41 Received: *graphvm.DeleteResponse, GraphID delete state is -1
victor@johnson:~/go/src/google.golang.org/grpc/examples/graphvm$ go run client/main_query.go 
2020/08/15 16:48:45 Received: *graphvm.QueryResponse, shortest path is -1
victor@johnson:~/go/src/google.golang.org/grpc/examples/graphvm$ go run client/main_add.go 
2020/08/15 16:48:50 Received: *graphvm.AddResponse, GraphID is 0
victor@johnson:~/go/src/google.golang.org/grpc/examples/graphvm$ go run client/main_query.go 
2020/08/15 16:48:55 Received: *graphvm.QueryResponse, shortest path is 2
victor@johnson:~/go/src/google.golang.org/grpc/examples/graphvm$ go run client/main_add.go 
2020/08/15 16:49:09 Received: *graphvm.AddResponse, GraphID is 1
victor@johnson:~/go/src/google.golang.org/grpc/examples/graphvm$ go run client/main_delete.go 
2020/08/15 16:49:14 Received: *graphvm.DeleteResponse, GraphID delete state is 0
victor@johnson:~/go/src/google.golang.org/grpc/examples/graphvm$ go run client/main_delete.go 
2020/08/15 16:49:18 Received: *graphvm.DeleteResponse, GraphID delete state is -1

    on server side
victor@johnson:~/go/src/google.golang.org/grpc/examples/graphvm$ go run server/main.go 
2020/08/15 16:48:41 Received: int32, 1
2020/08/15 16:48:45 Received: int32, 0, fist b, second c
2020/08/15 16:48:50 Received: []string, [a b c d b d a c]
2020/08/15 16:48:50 0 1
2020/08/15 16:48:50 2 3
2020/08/15 16:48:50 1 3
2020/08/15 16:48:50 0 2
    Graph size=4, ref=0
    0: len=2 cap=2 [1 2]
    1: len=2 cap=2 [0 3]
    2: len=2 cap=2 [3 0]
    3: len=2 cap=2 [2 1]
2020/08/15 16:48:55 Received: int32, 0, fist b, second c
    Find path from vertex 1 -- 2
    Find path from vertex 1 -- 2
2020/08/15 16:49:09 Received: []string, [a b c d b d a c]
2020/08/15 16:49:09 0 1
2020/08/15 16:49:09 2 3
2020/08/15 16:49:09 1 3
2020/08/15 16:49:09 0 2
    Graph size=4, ref=0
    0: len=2 cap=2 [1 2]
    1: len=2 cap=2 [0 3]
    2: len=2 cap=2 [3 0]
    3: len=2 cap=2 [2 1]
2020/08/15 16:49:14 Received: int32, 1
2020/08/15 16:49:18 Received: int32, 1



victor@johnson:~/go/src/google.golang.org/grpc/examples/graphvm/test$ go run graphvm.go 
2020/08/15 16:46:26 0 1
2020/08/15 16:46:26 2 3
2020/08/15 16:46:26 1 3
2020/08/15 16:46:26 0 2
    Graph size=4, ref=0
    0: len=2 cap=2 [1 2]
    1: len=2 cap=2 [0 3]
    2: len=2 cap=2 [3 0]
    3: len=2 cap=2 [2 1]
assign id  0
    Graph size=4, ref=1
    0: len=2 cap=2 [1 2]
    1: len=2 cap=2 [0 3]
    2: len=2 cap=2 [3 0]
    3: len=2 cap=2 [2 1]
    Find path from vertex 1 -- 2
    Find path from vertex 1 -- 2
Shortest is:  2
Shortest is:  -1
Shortest is:  -1
2020/08/15 16:46:26 0 1
2020/08/15 16:46:26 2 3
2020/08/15 16:46:26 1 3
2020/08/15 16:46:26 0 2
    Graph size=4, ref=0
    0: len=2 cap=2 [1 2]
    1: len=2 cap=2 [0 3]
    2: len=2 cap=2 [3 0]
    3: len=2 cap=2 [2 1]
assign id  1
2020/08/15 16:46:26 0 1
2020/08/15 16:46:26 2 3
2020/08/15 16:46:26 1 3
2020/08/15 16:46:26 0 2
    Graph size=4, ref=0
    0: len=2 cap=2 [1 2]
    1: len=2 cap=2 [0 3]
    2: len=2 cap=2 [3 0]
    3: len=2 cap=2 [2 1]
assign id  2



code structure
victor@johnson:~/go/src/google.golang.org/grpc/examples/graphvm$ tree .
.
├── client
│   ├── main_add.go
│   ├── main_delete.go
│   └── main_query.go
├── go.mod
├── go.sum
├── graphvm
│   ├── graphvm.go
│   ├── graphvm.pb.go
│   └── graphvm.proto
├── README
├── server
│   └── main.go
└── test
    └── graphvm.go

4 directories, 11 files

