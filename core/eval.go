package core

import(
	"errors"
	// "log"
	// "net"
	"io"
)

// func evalPING(args []string, c net.Conn) error {
func evalPING(args []string, c io.ReadWriter) error {
	var b []byte

	if len(args) >= 2 {
		return errors.New("err wrong no. of args for 'ping'")
	}

	if len(args) == 0 {
		b = Encode("PONG", true)
	} else {
		b = Encode(args[0], false)
	}

	_, err := c.Write(b)
	return err 
}

// func EvalAndRespond(cmd *RedisCmd, c net.Conn) error {
	func EvalAndRespond(cmd *RedisCmd, c io.ReadWriter) error {
	// log.Println("cmd: ", cmd)
	switch cmd.Cmd {
	case "PING":
		return evalPING(cmd.Args, c)
	default:
		return evalPING(cmd.Args, c)
	}
}