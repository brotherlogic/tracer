package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

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
	calls []*pb.Trace
}

// Init builds the server
func Init() *Server {
	s := &Server{
		&goserver.GoServer{},
		nil,
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
	return []*pbg.State{
		&pbg.State{Key: "calls", Value: int64(len(s.calls))},
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

	server.RegisterRepeatingTask(server.clean, "clean", time.Minute*5)

	server.SendTrace = false

	fmt.Printf("%v\n", server.Serve())
}
