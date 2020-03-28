package ircserver 
import (
	"net"
	"fmt"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

type Server struct {

}

func (s *Server) GetTodo(ctx context.Context, request *TodoRequest) (*TodoResponse, error) {
	return &TodosResponse {
		Todo: &Todo {
			ID: request.ID,
			Name: request.Name,
			Completed: request.Completed,
			CreatedAt: request.CreatedAt,	
		}	
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", "localhost", "8000"))
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	grpcServer := grpc.NewServer();

}
