package server

import (
    "net"
    "bufio"
)

type Client struct {
    connection      net.Conn
    remoteAddress   string
    time            int64
    bufferedIn      *bufio.Reader
    bufferedOut     *bufio.Writer
}

func (c *Client) Write(text string) {
    c.bufferedOut.WriteString(text + "\r\n")
    c.bufferedOut.Flush()
}

func (c *Client) Read() (string, error) {
    line, err := c.bufferedIn.ReadString('\n')

    if err != nil {
        return "", err
    }

    return line, nil
}
