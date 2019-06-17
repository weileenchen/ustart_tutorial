package recordapi

import (
	"net"

	"github.com/weileenchen/ustart_tutorial/record/recordpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Run runs the api server
func (pAPI *GPRCAPI) Run() {
	listener, err := net.Listen("tcp", ":"+pAPI.port)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	recordpb.RegisterRecordServer(srv, pAPI.record)
	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		panic(err)
	}
}
