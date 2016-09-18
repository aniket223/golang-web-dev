package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"io"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		// now handles multiple connections
		go serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	var i int
	var rMethod, rURI string
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		tx := scanner.Text()
		fmt.Println(tx)
		if i == 0 {
			// we're in REQUEST LINE
			xs := strings.Fields(tx)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("****METHOD:",rMethod)
			fmt.Println("****URI:",rURI)
		}
		if tx == "" {
			// when tx is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}

	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	body += "\n"
	body += rMethod
	body += "\n"
	body += rURI
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
