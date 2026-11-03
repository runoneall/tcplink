package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	local_addr := net.JoinHostPort(local_host, local_port)
	local_listener, err := net.Listen("tcp", local_addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		local_conn, err := local_listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go func(local_conn net.Conn) {
			defer local_conn.Close()

			remote_addr := net.JoinHostPort(remote_host, remote_port)
			remote_conn, err := net.Dial("tcp", remote_addr)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer remote_conn.Close()

			done := make(chan struct{}, 2)
			go func() { io.Copy(remote_conn, local_conn); done <- struct{}{} }()
			go func() { io.Copy(local_conn, remote_conn); done <- struct{}{} }()
			<-done
		}(local_conn)
	}
}
