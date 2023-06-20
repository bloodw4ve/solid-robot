package main

import (
	"bufio"
	"log"
	"net"
	"testing"
	"time"
)

func Test_randProverbIndex(t *testing.T) {
	type args struct {
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Case 1",
			args: args{
				max: 19,
			},
		}, {
			name: "Case 2",
			args: args{
				max: 0,
			},
		}, {
			name: "Case 3",
			args: args{
				max: 9,
			},
		}, {
			name: "Case 4",
			args: args{
				max: -10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomProverbIndex(tt.args.max); got > tt.args.max && tt.args.max >= 0 {
				t.Errorf("Random number randInt() = %v is bigger than field max: %v", got, tt.args.max)
			}
		})
	}
}

func Test_handleConn(t *testing.T) {
	srv, cl := net.Pipe()

	go func() {
		handleConn(srv)
	}()

	ticker := time.Tick(20 * time.Second)
	for {
		select {
		case <-ticker:
			cl.Close()
			srv.Close()
			return
		default:
			reader := bufio.NewReader(cl)
			b, err := reader.ReadBytes('\n')
			if err != nil {
				log.Fatal(err)
			}

			if len(b) == 0 {
				t.Error("Didn't get proverb from server")
			}
			t.Logf("Got proverb from server: %v", string(b))
		}
	}
}
