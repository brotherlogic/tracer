package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/brotherlogic/goserver"
	pbg "github.com/brotherlogic/goserver/proto"
	"github.com/brotherlogic/keystore/client"
	pb "github.com/brotherlogic/tracer/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//Server main server type
type Server struct {
	*goserver.GoServer
	calls      map[string]*pb.Trace
	callsMutex *sync.Mutex
}

// Init builds the server
func Init() *Server {
	s := &Server{
		&goserver.GoServer{},
		make(map[string]*pb.Trace),
		&sync.Mutex{},
	}
	return s
}

// DoRegister does RPC registration
func (s *Server) DoRegister(server *grpc.Server) {
	pb.RegisterTracerServiceServer(server, s)
}

// ReportHealth alerts if we're not healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	s.callsMutex.Lock()
	defer s.callsMutex.Unlock()

	count := int64(0)
	key := ""
	for k, c := range s.calls {
		count += int64(len(c.Events))
		key = k
	}

	return []*pbg.State{
		&pbg.State{Key: "calls", Value: int64(len(s.calls))},
		&pbg.State{Key: "size", Value: count},
		&pbg.State{Key: "sample_key", Text: key},
	}
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()

	//Turn off logging
	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	server := Init()
	server.GoServer.KSclient = *keystoreclient.GetClient(server.GetIP)
	server.PrepServer()
	server.Register = server

	server.RegisterServer("tracer", false)

	server.SendTrace = false

	fmt.Printf("%v\n", server.Serve())
}
