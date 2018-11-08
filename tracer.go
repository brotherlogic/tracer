package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"sync"
	"time"

	pbgh "github.com/brotherlogic/githubcard/proto"
	"github.com/brotherlogic/goserver"
	pbg "github.com/brotherlogic/goserver/proto"
	"github.com/brotherlogic/goserver/utils"
	"github.com/brotherlogic/keystore/client"
	pb "github.com/brotherlogic/tracer/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//Server main server type
type Server struct {
	*goserver.GoServer
	calls      map[string]*pb.ContextCall
	callsMutex *sync.Mutex
}

// Init builds the server
func Init() *Server {
	s := &Server{
		&goserver.GoServer{},
		make(map[string]*pb.ContextCall),
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

// Mote promotes/demotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	return nil
}

// GetState gets the state of the server
func (s *Server) GetState() []*pbg.State {
	count := 0
	for _, c := range s.calls {
		if c.Properties.Died == 0 && c.Properties.Created == 0 {
			count++
		}
	}
	return []*pbg.State{
		&pbg.State{Key: "calls", Value: int64(len(s.calls))},
		&pbg.State{Key: "unfinished_calls", Value: int64(count)},
	}
}

func (s *Server) buildLong(call *pb.ContextCall) string {
	retstring := fmt.Sprintf("%v - %v\n", call.GetProperties().Id, (call.Properties.Died-call.Properties.Created)/1000000)
	for _, m := range call.GetMilestones() {
		retstring += fmt.Sprintf("[%v] - %v (%v)\n", (m.GetTimestamp()-call.Properties.Created)/1000000, m.Label, m.GetType())
	}

	return retstring
}

func (s *Server) findLongest(ctx context.Context) {
	longest := s.getLongContextCall()
	if longest != nil && (longest.Properties.Died-longest.Properties.Created)/1000000 > 500 {
		ip, port, _ := utils.Resolve("githubcard")
		if port > 0 {
			conn, err := grpc.Dial(ip+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
			if err == nil {
				defer conn.Close()
				client := pbgh.NewGithubClient(conn)
				client.AddIssue(ctx, &pbgh.Issue{Service: longest.Properties.Origin, Title: "Long", Body: fmt.Sprintf("%v", s.buildLong(longest))}, grpc.FailFast(false))
			}
		}
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

	server.RegisterRepeatingTask(server.findLongest, "find_longest", time.Hour)
	server.Log("Tracer is starting!")
	fmt.Printf("%v\n", server.Serve())
}
