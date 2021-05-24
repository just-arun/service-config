package database

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

type GetDGraphDBParam struct {
	Host     string
	Port     int
	UserId   string
	Password string
}

func (s *GetDGraphDBParam) conStr() string {
	return fmt.Sprintf("%v:%v", s.Host, s.Port)
}

func getDGraphDB(param GetDGraphDBParam) (dGraphClient *dgo.Dgraph, cancelFunction CancelFunc) {
	conn, err := grpc.Dial(param.conStr(), grpc.WithInsecure())
	if err != nil {
		panic("While trying to dial gRPC")
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)
	ctx := context.Background()

	// Perform login call. If the Dgraph cluster does not have ACL and
	// enterprise features enabled, this call should be skipped.
	for {
		// Keep retrying until we succeed or receive a non-retriable error.
		err = dg.Login(ctx, param.UserId, param.Password)
		if err == nil || !strings.Contains(err.Error(), "Please retry") {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		panic(fmt.Sprintf("While trying to login %v", err.Error()))
	}

	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}
