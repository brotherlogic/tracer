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

const (
	//KEY is where we store config
	KEY = "github.com/brotherlogic/tracer/config"
)

//Server main server type
type Server struct {
	*goserver.GoServer
	config     *pb.Config
	calls      []*pb.Trace
	counts     map[string]int
	mostCalled string
	allCalls   int64
	markedIds  []*pb.MarkRequest
	goodList   []string
}

// Init builds the server
func Init() *Server {
	s := &Server{
		&goserver.GoServer{},
		&pb.Config{},
		nil,
		make(map[string]int),
		"",
		int64(0),
		make([]*pb.MarkRequest, 0),
		[]string{"recordmatcher", "recordmover"},
	}
	return s
}

func (s *Server) save(ctx context.Context) {
	s.KSclient.Save(ctx, KEY, s.config)
}

func (s *Server) load(ctx context.Context) error {
	config := &pb.Config{}
	data, _, err := s.KSclient.Read(ctx, KEY, config)
	if err != nil {
		return err
	}
	s.config = data.(*pb.Config)
	return nil
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
	if master {
		err := s.load(ctx)
		return err
	}

	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	return []*pbg.State{
		&pbg.State{Key: "good_list", Text: fmt.Sprintf("%v", s.goodList)},
		&pbg.State{Key: "last_mark", TimeValue: s.config.LastMarkSent},
		&pbg.State{Key: "calls", Value: int64(len(s.calls))},
		&pbg.State{Key: "most_calls", Text: s.mostCalled},
		&pbg.State{Key: "all_calls", Value: s.allCalls},
		&pbg.State{Key: "marked", Value: int64(len(s.markedIds))},
	}
}

func (s *Server) staleAlert(ctx context.Context) error {
	if time.Now().Sub(time.Unix(s.config.LastMarkSent, 0)) > time.Hour*24*7 {
		s.RaiseIssue(ctx, "Adjust alert settings", fmt.Sprintf("Last mark alert was sent at %v", time.Unix(s.config.LastMarkSent, 0)), false)
	}

	return nil
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
	server.RegisterRepeatingTask(server.staleAlert, "stale_alert", time.Hour)

	server.SendTrace = false

	fmt.Printf("%v\n", server.Serve())
}
