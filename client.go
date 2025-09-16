// package main

// import (
// 	"fmt"
// 	"io"
// 	"net"
// 	"os"
// )

// func main() {
// 	connection, err := net.Dial("tcp", "localhost:8075")
// 	if err != nil {
// 		fmt.Println("There is an error during connection: ", err)
// 	}
// 	defer connection.Close()

// 	_, err = io.Copy(os.Stdout, connection)
// 	if err != nil {
// 		fmt.Println("There is an error during copy: ", err)
// 	}

// }

package main

import (
	"io"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8070"
	TYPE = "tcp"
)

func main() {

	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	_, err = conn.Write([]byte("A message from client 2\n")) //داده‌ای که کلاینت به سرور می‌فرستد
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	//new line of code
	conn.CloseWrite()

	received := make([]byte, 4096)
	for {
		println("Reading data...")
		temp := make([]byte, 4096)
		_, err = conn.Read(temp) //داده‌های ارسالی از سمت سرور خوانده شده و در تمپ ریخته می‌شوند
		if err != nil {
			if err == io.EOF {
				break
			}
			println("Read data failed:", err.Error())
			os.Exit(1)
		}
		  received = append(received, temp...)
		//   fmt.Println("Using io Package: ")
		// io.Copy(os.Stdout, conn)
	}

	println("Received message from Server:", string(received))

}
