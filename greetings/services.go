package greetings

import context "context"

//GServer implements the greeting services
type GServer struct{}

//SayHello implements SayHello service
func (g *GServer) SayHello(c context.Context, req *HelloRequest) (*HelloReply, error) {
	reply := HelloReply{}
	reply.Message = "Yuhuuuuuuuuuuuu"
	return &reply, nil
}

// type GClient struct{}

// func (g * GClient) SayHello(c context.Context, req *HelloReply ) (*HelloReply, error){
// 	reply := HelloReply{}
// 	reply.Message
// }
