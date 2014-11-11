package main

import "fmt"

import "../go-dockerclient"

func main() {
	client, _ := docker.NewClient("tcp://192.168.1.104:2376")

	fmt.Printf("%T", client)
}
