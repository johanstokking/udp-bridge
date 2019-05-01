package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	localAddr := os.Getenv("LOCAL_ADDR")
	if localAddr == "" {
		localAddr = ":1701"
	}
	remoteAddr := os.Getenv("REMOTE_ADDR")
	if remoteAddr == "" {
		remoteAddr = "127.0.0.1:1700"
	}

	listenAddr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		panic(err)
	}
	ln, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		panic(err)
	}

	conns := make(map[string]net.Conn)
	buf := make([]byte, 65507)
	for {
		n, addr, err := ln.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}
		conn, ok := conns[addr.String()]
		if !ok {
			var err error
			conn, err = net.Dial("udp", remoteAddr)
			if err != nil {
				panic(fmt.Sprintf("dial %v failed: %v", addr, err))
			}
			conns[addr.String()] = conn
			go func() {
				buf := make([]byte, 65507)
				for {
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Printf("< read failed: %v\n", err)
						return
					}
					_, err = ln.WriteToUDP(buf[:n], addr)
					if err != nil {
						fmt.Printf("< write failed: %v\n", err)
					} else {
						fmt.Printf("< %v\n", strconv.QuoteToASCII(string(buf[:n])))
					}
				}
			}()
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Printf("> write failed: %v\n", err)
		} else {
			fmt.Printf("> %v\n", strconv.QuoteToASCII(string(buf[:n])))
		}
	}
}
