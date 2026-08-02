package server

import (
	"net"
	"strconv"
	"log"
	"io"
	"fmt"
	"strings"

	"github.com/Moksh-10/redis/config"
	"github.com/Moksh-10/redis/core"
)

// func readCommand(c net.Conn) (*core.RedisCmd, error) {
func readCommand(c io.ReadWriter) (*core.RedisCmd, error) {
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return nil, err
	}

	tokens, err := core.DecodeArrayString(buf[:n])
	if err != nil {
		return nil, err
	}

	return &core.RedisCmd{
		Cmd: strings.ToUpper(tokens[0]),
		Args: tokens[1:],
	}, nil
}

// func respondError(err error, c net.Conn) {
	func respondError(err error, c io.ReadWriter) {
	c.Write([]byte(fmt.Sprintf("-%s\r\n", err)))
}

// func respond(cmd *core.RedisCmd, c net.Conn) {
func respond(cmd *core.RedisCmd, c io.ReadWriter) {
	err := core.EvalAndRespond(cmd, c)
	if err != nil {
		respondError(err, c)
	}
}

func RunSyncTCPServer() {
	log.Println("starting a sync tcp server on ", config.Host, config.Port)

	var con_clients int = 0

	lsnr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		// panic(err)
		log.Println("err", err)
		return
	}

	for {
		c, err := lsnr.Accept()
		if err != nil {
			// panic(err)
			log.Println("err", err)
		}

		con_clients += 1
		// log.Println("clients connected with add: ", c.RemoteAddr(), "concurrent clients: ", con_clients)

		for {
			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				con_clients -= 1
				// log.Println("client disconnected ", c.RemoteAddr(), "concurrent clients: ", con_clients)
				if err == io.EOF {
					break
				}
				// log.Println("err", err)
			}
			respond(cmd, c)
		}
	}
}