package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func splitPacket(str string) (string, string) {
	var split = strings.Split(str, "|")
	return split[0], split[1]
}

func splitMouseMovement(str string) (int, int) {
	var split = strings.Split(str, ",")
	xpos, xerr := strconv.Atoi(split[0])
	ypos, yerr := strconv.Atoi(split[0])

	if xerr != nil {
		fmt.Println("Error", xerr)
	}
	if yerr != nil {
		fmt.Println("Error", yerr)
	}

	return xpos, ypos
}

func main() {

	s, err := net.ResolveUDPAddr("udp4", "127.0.0.1:44906")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Starting Server on 44906")
	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		fmt.Println("Reading data")
		n, addr, err := connection.ReadFromUDP(buffer)
		_ = addr

		msgType, msg := splitPacket(string(buffer))

		if msgType == "0" {
			//Mouse Movement
			mousePosX, mousePosY := robotgo.GetMousePos()
			mouseX, mouseY := splitMouseMovement(msg)
			robotgo.MoveMouse(mousePosX+mouseX, mousePosY+mouseY)
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		/*data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}*/
	}
}
