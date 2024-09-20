package main

import (
	"bufio"
	"fmt"
	doc "github.com/jensilo/trustdoc"
	"log/slog"
	"net"
	"os"
	"time"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	client(log)
}

func client(log doc.Log) {
	for {
		time.Sleep(time.Second)

		conn, err := net.Dial("tcp", ":42105")
		if err != nil {
			log.Info(fmt.Sprintf("error dialing port 42105: %s", err))
			return
		}

		_, err = fmt.Fprint(conn, "Hello from the Client\n")
		if err != nil {
			log.Info(fmt.Sprintf("error writing bytes: %s", err))
			return
		}

		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Info(fmt.Sprintf("error reading response: %s", err))
			return
		}

		log.Info(status)
	}
}
