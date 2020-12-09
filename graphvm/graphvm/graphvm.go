//package main
package graphvm

import (
    "fmt"
    "sync"
    "log"
)

const (
    BadId   = -1
    NoPath  = -2
    InUse   = -3
    BadArg  = -4
)

var repo *Grepo

// v should contain pairs of connected network nodes
func PostGraph(v []string) int32 {
    n := len(v)
    if(n == 0 || ((n&1) != 0)) {
        return BadArg
    }

    g := new(Graph)
    return g.InitGraph(v)
}

func DeleteGraph(id int32) int32 {
    return repo.DeleteId(id)
}

func GetShortest(id int32, x, y string) int32 {
    s := repo.GetId(id)
    if s == nil {
        return BadId
    }
    defer repo.PutId(id)
    return s.Shortest(x, y)
}

func PrintGraph(id int32) {
    s := repo.GetId(id)
    if s != nil {
        defer repo.PutId(id)
        s.Print()
    }
}

// This is repository of all graphs
type Grepo struct {
    reqid chan bool
    respid chan int32
    m map[int32](*Graph)
    mux sync.Mutex
}

// This must be called at system startup
func SystemStart() {
    repo = new(Grepo)
    repo.InitGrepo()
}

func(r *Grepo) InitGrepo() {
    r.m = make(map[int32](*Graph))
    r.reqid = make(chan bool)
    r.respid = make(chan int32)
    go r.IdManager()
}

// As id must be unique only, we do not
// track which thread request id
func(r *Grepo) IdManager() {
    for gid, more := int32(0), true; more; gid++ {
        more = <-r.reqid
        r.respid<- gid
    }
}

func(r *Grepo) AddId(g *Graph) int32 {
    r.reqid<- true
    id := <-r.respid
    r.mux.Lock()
    defer r.mux.Unlock()
    r.m[id] = g
    return id
}

func(r *Grepo) GetId(id int32) *Graph {
    r.mux.Lock()
    defer r.mux.Unlock()
    s, ok := r.m[id];
    if ok {
        s.ref++
        return s
    }
    return nil
}

func(r *Grepo) PutId(id int32) {
    r.mux.Lock()
    defer r.mux.Unlock()
    r.m[id].ref--
}

func(r *Grepo) DeleteId(id int32) int32 {
    r.mux.Lock()
    defer r.mux.Unlock()
    s, ok := r.m[id];
    if !ok {
        return BadId
    }
    if(s.ref != 0) {
        return InUse
    }
    delete(r.m, id)
    return 0
}

type Graph struct {
    ref, size int32
    v [][]int32
    m map[string]int32
}

// convert network node (string) to an ID (int32)
func(s* Graph) GetId(node string, add bool) int32 {
    id, ok := s.m[node];
    if !ok {
        if (!add) {
            return BadId
        }
        id, s.m[node] = s.size, s.size
        s.size++;
    }
    return id;
}

// construct edge list to a adjacent list graph
// which is convenient for some graph algorithm
func(s* Graph) InitGraph(node []string) int32 {
    n := len(node)
    v := make([]int32, n)
    s.m = make(map[string]int32)

    for i, me := range node {
        v[i] = s.GetId(me, true)
    }

    s.v = make([][]int32, s.size)
    for i:=0; i<n; i = i+2 {
        a, b := v[i], v[i+1]
        log.Println(a, b)
        s.v[a] = append(s.v[a], b)
        s.v[b] = append(s.v[b], a)
    }

    s.Print()
    return repo.AddId(s)
}

// Only BFS is needed to find shortest path
// For weighted graph, dijkstra (fastest) or 
// bellman-ford (DP) must be used.  
func(s* Graph) Shortest(x, y string) int32 {
    a := s.GetId(x, false)
    b := s.GetId(y, false)
    fmt.Printf("\tFind path from vertex %d -- %d\n", a, b);
    if a == BadId || b == BadId {
        return BadId
    }

    if a == b {
        return 0
    }

    fmt.Printf("\tFind path from vertex %d -- %d\n", a, b);
    dist := make([]int32, s.size)
    q := make([]int32, s.size)
    q[0], dist[a] = a, 1
    for i,j:=0,1; i<j ; i++ {
        d := dist[q[i]]
        for _, k := range s.v[q[i]] {
            if (k == b) {
                return d
            } else if(dist[k] == 0) {
                q[j], dist[k], j = k, d+1, j+1
            }
        }
    }
    return NoPath // no path from a to b
}

// For formal design, Stringer int32erface should be used here
// Change all Printf to Sprint32f and return string
func (s *Graph) Print() {
    fmt.Printf("\tGraph size=%d, ref=%d\n", s.size, s.ref);
    for i:= int32(0); i<s.size; i++ {
        fmt.Printf("\t%d: len=%d cap=%d %v\n", i, len(s.v[i]), cap(s.v[i]), s.v[i])
    }
}

/*
func main() {
    SystemStart()
    v := []string{"a", "b", "c", "d", "b", "d", "a", "c"}
    s := PostGraph(v)
    fmt.Println("assign id ", s)
    PrintGraph(s)
    a := GetShortest(s, "b", "c")
    fmt.Println("Shortest is: ", a)
    a = GetShortest(-1, "b", "c")
    fmt.Println("Shortest is: ", a)
    DeleteGraph(100)
    DeleteGraph(s)
    a = GetShortest(s, "b", "c")
    fmt.Println("Shortest is: ", a)

    s = PostGraph(v)
    fmt.Println("assign id ", s)

    s = PostGraph(v)
    fmt.Println("assign id ", s)
}
*/
