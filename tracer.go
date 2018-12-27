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
	calls            map[string]*pb.ContextCall
	callsMutex       *sync.Mutex
	silencedAlerts   int
	whitelist        []string
	longestDelivered int64
	timeOfLongest    time.Duration
}

// Init builds the server
func Init() *Server {
	s := &Server{
		&goserver.GoServer{},
		make(map[string]*pb.ContextCall),
		&sync.Mutex{},
		0,
		[]string{
			"dropboxsync",
		},
		int64(0),
		0,
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
	s.callsMutex.Lock()
	defer s.callsMutex.Unlock()

	count := 0
	for _, c := range s.calls {
		if c.Properties.Died == 0 && c.Properties.Created == 0 {
			count++
		}
	}
	return []*pbg.State{
		&pbg.State{Key: "calls", Value: int64(len(s.calls))},
		&pbg.State{Key: "unfinished_calls", Value: int64(count)},
		&pbg.State{Key: "silenced_alerts", Value: int64(s.silencedAlerts)},
		&pbg.State{Key: "num_whitelisted", Value: int64(len(s.whitelist))},
		&pbg.State{Key: "longest_delivered", Value: s.longestDelivered},
		&pbg.State{Key: "time_of_longest", TimeDuration: s.timeOfLongest.Nanoseconds()},
	}
}

func (s *Server) buildLong(call *pb.ContextCall) string {
	retstring := fmt.Sprintf("%v - %v\n", call.GetProperties().Id, (call.Properties.Died-call.Properties.Created)/1000000)
	for _, m := range call.GetMilestones() {
		retstring += fmt.Sprintf("[%v] - %v (%v)\n", (m.GetTimestamp()-call.Properties.Created)/1000000, m.GetLabel(), m.GetType())
	}

	return retstring
}

func (s *Server) findLongest(ctx context.Context) {
	longest := s.getLongContextCall(ctx)
	if longest != nil && (longest.Properties.Length)/1000000 >= 500 {
		ip, port, _ := utils.Resolve("githubcard")
		if port > 0 {
			conn, err := grpc.Dial(ip+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
			if err == nil {
				defer conn.Close()
				client := pbgh.NewGithubClient(conn)
				client.AddIssue(ctx, &pbgh.Issue{Service: longest.Properties.Origin, Title: "Long", Body: fmt.Sprintf("%v", s.buildLong(longest))}, grpc.FailFast(false))
				longest.Properties.Delivered = true
				s.longestDelivered++
				s.timeOfLongest = time.Duration(longest.Properties.Length)
			}
		}
	} else if longest != nil {
		s.Log(fmt.Sprintf("Rejecting %v because of length %v", longest.Properties.Origin, longest.Properties.Length))
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

	server.RegisterRepeatingTask(server.findLongest, "find_longest", time.Minute*2)

	fmt.Printf("%v\n", server.Serve())
}
