package news_manager_client

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/zhangminghui6106/delinkcious-0.5/pb/news_service/pb"
	om "github.com/zhangminghui6106/delinkcious-0.5/pkg/object_model"
	"google.golang.org/grpc"
)

type DisconnectFunc func() error

func NewClient(grpcAddr string) (cli om.NewsManager, disconnectFunc DisconnectFunc, err error) {
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	disconnectFunc = func() (err error) {
		if conn == nil {
			return
		}

		err = conn.Close()
		return
	}

	if err != nil {
		return
	}
	var getNewsEndpoint = grpctransport.NewClient(
		conn, "pb.News", "GetNews",
		encodeGetNewsRequest,
		decodeGetNewsResponse,
		pb.GetNewsResponse{},
	).Endpoint()

	cli = EndpointSet{
		GetNewsEndpoint: getNewsEndpoint,
	}
	return
}
