// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"time"
// )

// func main() {

// 	listener, err := net.Listen("tcp", "localhost:8075")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer listener.Close()

// 	for {
// 		connection, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("There is an error during connection: ", err)
// 		}
// 		fmt.Println("Connection accepted")
// 		go func() {
// 			fmt.Println("Starting a new goroutine...")
// 			handleConnection(connection)
// 		}()
// 	}

// }

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()
// 	for {
// 		_, err := io.WriteString(conn, time.Now().Format(time.Stamp+"\n"))
// 		if err != nil {
// 			fmt.Println("There is an error during writing: ", err)
// 		}
// 		time.Sleep(time.Second)
// 	}

// }

package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "8070"
	TYPE = "tcp"
)

func main() {

	// arguments := os.Args
	// if len(arguments) == 1 {
	// 	fmt.Println("Please provide a port number!")
	// 	return
	// }

	// PORT := ":" + arguments[1]

	// تابع لیسن به سرور اجازه می‌دهد تا درخواست‌های ورودی از کلاینت‌ها را بپذیرد
	listener, err := net.Listen(TYPE, HOST+":"+PORT) 
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		// سرور منتظر اتصالی از سمت یک کلاینت می‌ماند. زمانی که یک اتصال برقرار شد، تابع اکسپت آن را می‌پذیرد و یک کانکشن برمی‌گرداند.
		c, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	
	packet := make([]byte, 4096)
	tmp := make([]byte, 4096)

	testData := []byte("A test message from Server, Hahahaha\n")
	
	for {
		// داده‌ها از اتصال بین کلاینت و سرور خوانده شده و به تمپ ریخته می‌شوند
		_, err := c.Read(tmp)
		fmt.Println("Data recieved from client: ", string(tmp))
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		packet = append(packet, testData...)
		packet = append(packet, tmp...)
	}
	c.Write(packet) //داده‌ها به کلاینت ارسال می‌شوند. مشابه همان خط io.WriteString
	fmt.Println("All messages transferred in the connection: ", string(packet))
}
