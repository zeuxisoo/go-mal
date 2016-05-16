package server

import (
    "fmt"
    "net"
    "log"
    "bufio"
    "time"
)

type Config struct {
    Address string
    Port    int
}

type Client struct {
    connection      net.Conn
    remoteAddress   string
    time            int64
    bufferIn        *bufio.Reader
    bufferOut       *bufio.Writer
}

type Server struct {
    Address string
    Port    int
}

func NewServer(config *Config) *Server {
    return &Server{
        Address: config.Address,
        Port   : config.Port,
    }
}

func (s *Server) Run() {
    addressWithPort := fmt.Sprintf("%s:%d", s.Address, s.Port)
    listener, err   := net.Listen("tcp", addressWithPort)

    if err != nil {
        log.Fatal("[Error] ", err.Error())
    }

    log.Printf("listening on %s", addressWithPort)

    for {
        conn, err := listener.Accept()

        if err != nil {
            log.Fatal("[Error] ", err.Error())
        }

        go handleClient(&Client{
            connection   : conn,
            remoteAddress: conn.RemoteAddr().String(),
            time         : time.Now().Unix(),
            bufferIn     : bufio.NewReader(conn),
            bufferOut    : bufio.NewWriter(conn),
        })
    }
}

func handleClient(client *Client) {

}
