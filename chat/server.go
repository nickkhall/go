package server

type ChatServer interface {
	Listen(address string) error
	Broadcast(command interface{}) error
	Start()
	Close()
}

type TcpChatServer struct {
	listener net.Listener
	clients  []*clients
	mutex    *sync.Mutex
}

type client struct {
	conn   net.Conn
	name   string
	writer *protocol.CommandWriter
}

func (s *TcpChatServer) Listen(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		s.listener = l
	}

	fmt.Println("Server listening on port 8000...\n")

	return err
}

func (s *TcpChatServer) Close() {
	s.listener.Close()
}

func (s *TcpChatServer) Start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("ERROR: ", err)
		} else {
			client := s.accept(conn)
			go s.serve(client)
		}
	}
}

func (s *TcpChatServer) accept(con net.Conn) *client {
	fmt.Println("Accepting connection from %v, total clients: %v", conn.RemoteAddr().String(), len(s.clients)+1)
	
	s.mutex.Lock()
	defer s.mutex.Unlock()
}
