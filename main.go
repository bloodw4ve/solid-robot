package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

const addr = "0.0.0.0:12345"

const proto = "tcp4"

var interval time.Duration = 3 * time.Second

var proverbs = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func main() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		if len(proverbs) == 0 {
			continue
		}
		i := randomProverbIndex(len(proverbs) - 1)
		conn.Write([]byte(proverbs[i] + "\n"))
		time.Sleep(interval)
	}
}

func randomProverbIndex(max int) int {
	if max <= 0 {
		return 0
	}
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return r.Intn(max)
}
