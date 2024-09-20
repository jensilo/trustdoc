package main

import (
	"bufio"
	"fmt"
	doc "github.com/jensilo/trustdoc"
	"log/slog"
	"net"
	"os"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	server(log)
}

func server(log doc.Log) {
	ln, err := net.Listen("tcp", ":42105")
	if err != nil {
		log.Info(fmt.Sprintf("error listening on port 42105: %s", err))
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Info(fmt.Sprintf("error accepting connection: %s", err))
		}

		go handleConnection(log, conn)
	}
}

func handleConnection(log doc.Log, conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Info(fmt.Sprintf("error closing connection: %s", err))
		}
	}(conn)

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Info(fmt.Sprintf("error reading from buffer: %s", err))
			return
		}

		log.Info(fmt.Sprintf("received line: %s", data))

		_, err = conn.Write([]byte(fmt.Sprintf("Hello TCP Client, you said: %s\n", data)))
		if err != nil {
			log.Info(fmt.Sprintf("error writing back response: %s", err))
			return
		}
	}
}
