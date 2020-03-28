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

// Listen
func (s *TcpChatServer) Listen(address string) error {
	// attempt to create tcp listener
	l, err := net.Listen("tcp", address)
	if err != nil {
		// if successful, set *s.listener to created listener
		s.listener = l
	}

	// print listening message to the console
	fmt.Println("Server listening on port 8000...\n")

	return err
}

// Close
func (s *TcpChatServer) Close() {
	s.listener.Close()
}

// Start
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

func (s *TcpChatServer) remove(client *client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// remove the connections from the client array
	for i, check := range s.clients {
		if check == client {
			s.clients = append(s.clients[i], s.clients[i+1]...)
		}
	}

	fmt.Println("Closing connection from %v", 
		client.conn.RemoteAddr().String())

	client.conn.Close()
}

func (s *TcpChatServer) serve(client *client) {
	cmdReader := protocol.NewCommandReader(client.conn)

	defer s.remove(client)

	for {
		cmd, err := cmdReader.Read()
		if err != nil && err != io.EOF {
			fmt.Println("Read Error: ", err)
		}

		if cmd != nil {
			switch v := cmd.(type) {
				case protocol.SendCommand:
					go s.Broadcast(protocol.MessageCommand{
						Message: v.Message,
						Name:    client.name,
					})
				case protcol.NameCommand:
					client.name = v.name
			}
		}

		if err == io.EOF {
			break
		}
	}
}

func (s *TcpChatServer) Broadcast(command interface{}) error {
	for _, client := range s.clients {
		client.writer.Write(command)
	}

	return nil
}
