// socket-server project main.go
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	// "strings"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)

	if err != nil {

		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")
	connection, err := server.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("client connected")
	//go processClient(connection)
	for {
		fmt.Println("Enter command")
		var command string
		fmt.Scan(&command)
		if command == "download" {
			_, err = connection.Write([]byte(command))
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			var filename string
			fmt.Println("Enter filename")
			fmt.Scan(&filename)
			_, err = connection.Write([]byte(filename))
			//start recieving file bytes
			var fileSize string
			buf := make([]byte, 4096)
			mLen, err := connection.Read(buf)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if string(buf[:mLen]) == "File_Not_Exists" {
				fmt.Println("File not exists")
			} else {
				fileSize = string(buf[:mLen])
			}
			fmt.Println("Recieved Size :", fileSize)
			filesize, err := strconv.ParseInt(fileSize, 10, 0)
			var iterator int = 0
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
			fmt.Println("File size :", filesize)
			//time.Sleep(5 * time.Second)
			f, err := os.Create(filename)
			if err != nil {
				log.Fatalf("unable to read file: %v", err)
			}
			//defer f.Close()
			for {
				mLen, err := connection.Read(buf)
				iterator = iterator + mLen
				//fmt.Println("Bytes Recieved :", string(buf[:mLen]))
				if iterator >= int(filesize) {
					n, err := f.Write(buf[:mLen])
					if err != nil {
						fmt.Println(err)
						continue
					}
					if n > 0 {
						//fmt.Println(string(buf[:n]))
					}
					fmt.Println("File has Ended")
					break
				}
				if err == io.EOF {
					//_, err = connection.Write([]byte("File_Ended"))
					fmt.Println("File has Ended")
					//_, err = connection.Write([]byte("File_Ended"))
					break
				}
				// if strings.Contains(string(buf[:mLen]), "File_Ended") {
				// 	fmt.Println("File_Ended")
				// 	break
				// }
				n, err := f.Write(buf[:mLen])
				if err != nil {
					fmt.Println(err)
					continue
				}
				if n > 0 {
					//fmt.Println(string(buf[:n]))
				}
				fmt.Println(iterator*100/int(filesize), "% completed ")
			}
		} else if command == "dir" {
			_, err = connection.Write([]byte(command))
			buffer := make([]byte, 4096)
			mLen, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			fmt.Println(string(buffer[:mLen]))
		} else if command == "back" {
			_, err = connection.Write([]byte(command))
			buffer := make([]byte, 4096)
			mLen, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			fmt.Println(string(buffer[:mLen]))
		} else if command == "move" {
			_, err = connection.Write([]byte(command))
			buffer := make([]byte, 4096)
			mLen, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			fmt.Printf(string(buffer[:mLen]))
			//time.Sleep(2 * time.Second)
			var dir string
			fmt.Println("Enter Dir name :")
			fmt.Scan(&dir)
			_, err = connection.Write([]byte(dir))
			mLen, err1 := connection.Read(buffer)
			if err1 != nil {
				fmt.Println("Error reading:", err.Error())
			}
			fmt.Println(string(buffer[:mLen]))
		} else if command == "disks" {
			_, err = connection.Write([]byte(command))
			buffer := make([]byte, 4096)
			mLen, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			fmt.Println(string(buffer[:mLen]))
		} else {
			fmt.Printf("Invalid Input")
		}
		//fmt.Println("Received: ", string(buffer[:mLen]))

	}
}

/*
func processClient(connection net.Conn) {

	//connection.Close()
}


for {
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		var command string = string(buffer[:mLen])
		if command == "whoami" {
			fmt.Println("Received: ", string(buffer[:mLen]))
			_, err = connection.Write([]byte("command is whoami" + string(buffer[:mLen])))
		} else if string(command) == "dir" {
			fmt.Println("Received: ", string(buffer[:mLen]))
			_, err = connection.Write([]byte("command is dir" + string(buffer[:mLen])))
		} else {
			fmt.Println("Received: ", string(buffer[:mLen]))
			_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
			connection.Close()
		}
	}*/
