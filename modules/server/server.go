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
            bufferedIn   : bufio.NewReader(conn),
            bufferedOut  : bufio.NewWriter(conn),
        })
    }
}

func handleClient(client *Client) {
    client.Write("220 Welcome to the Mal fake server")

    //
    hello, _ := client.Read()
    log.Print(hello)

    client.Write("250 Hello Domain")

    //
    mailFrom, _ := client.Read()
    log.Print(mailFrom)

    client.Write("250 Ok")

    //
    mailTo, _ := client.Read()
    log.Print(mailTo)

    client.Write("250 Ok")

    //
    data, _ := client.Read()
    log.Print(data)

    client.Write("354 End data with <CR><LF>.<CR><LF>")

    for {
        body, _ := client.Read()
        log.Print(body)

        // Submit
        bytes := []byte(body)

        if bytes[0] == 46 && bytes[1] == 13 && bytes[2] == 10 {
            client.Write("250 Ok: queued as 12345")
            break
        }
    }

    client.connection.Close()
}
