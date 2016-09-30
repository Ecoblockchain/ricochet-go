package main

import (
	"fmt"
	rpc "github.com/special/notricochet/rpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gopkg.in/readline.v1"
	"log"
	"os"
	"strings"
)

const (
	defaultAddress = "127.0.0.1:58281"
)

func main() {
	conn, err := grpc.Dial(defaultAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("connection failed: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	input, err := readline.NewEx(&readline.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer input.Close()
	log.SetOutput(input.Stdout())

	c := &Client{
		Backend: rpc.NewRicochetCoreClient(conn),
		Input:   input,
	}

	if err := c.Initialize(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input.SetPrompt("> ")
	for {
		line := input.Line()
		if line.CanContinue() {
			continue
		} else if line.CanBreak() {
			break
		}

		words := strings.SplitN(line.Line, " ", 1)
		switch words[0] {
		case "clear":
			readline.ClearScreen(readline.Stdout)

		case "quit":
			os.Exit(0)

		case "status":
			fmt.Printf("server: %v\n", c.ServerStatus)
			fmt.Printf("identity: %v\n", c.Identity)

		case "connect":
			status, err := c.Backend.StartNetwork(context.Background(), &rpc.StartNetworkRequest{})
			if err != nil {
				fmt.Printf("start network error: %v\n", err)
			} else {
				fmt.Printf("network started: %v\n", status)
			}

		case "disconnect":
			status, err := c.Backend.StopNetwork(context.Background(), &rpc.StopNetworkRequest{})
			if err != nil {
				fmt.Printf("stop network error: %v\n", err)
			} else {
				fmt.Printf("network stopped: %v\n", status)
			}

		case "contacts":
			byStatus := make(map[rpc.Contact_Status][]*rpc.Contact)
			for _, contact := range c.Contacts {
				byStatus[contact.Status] = append(byStatus[contact.Status], contact)
			}

			order := []rpc.Contact_Status{rpc.Contact_ONLINE, rpc.Contact_UNKNOWN, rpc.Contact_OFFLINE, rpc.Contact_REQUEST, rpc.Contact_REJECTED}
			for _, status := range order {
				contacts := byStatus[status]
				if len(contacts) == 0 {
					continue
				}
				fmt.Printf(". %s\n", strings.ToLower(status.String()))
				for _, contact := range contacts {
					fmt.Printf("... %s\n", contact.Nickname)
				}
			}

		case "help":
			fallthrough

		default:
			fmt.Println("Commands: clear, quit, status, connect, disconnect, contacts, help")
		}
	}
}
