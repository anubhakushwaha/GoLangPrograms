package main

import (
	"log"

	"golang.org/x/net/context"

	"github.com/golang/protobuf/protoc-gen-go/grpcCheck/todo"

	"flag"
	"fmt"
	"os"
	"strings"

	grpc "google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect to server %v", err)
	}

	client := todo.NewTasksClient(conn)

	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list(context.Background(), client)
	case "add":
		err = add(context.Background(), client, strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("Unknown subcommand %s", cmd)
	}
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

func add(ctx context.Context, client todo.TasksClient, text string) error {
	_, err := client.Add(ctx, &todo.Text{Text: &text})
	if err != nil {
		return fmt.Errorf("couldn't fetch from server: %s", err)
	}

	fmt.Println("Task added successfully")
	return nil
}

func list(ctx context.Context, client todo.TasksClient) error {
	l, err := client.List(ctx, &todo.Void{})
	if err != nil {
		return fmt.Errorf("couldn't fetch from server: %s", err)
	}
	for _, task := range l.Tasks {
		if *task.Done {
			fmt.Println("DONE !!!")
		} else {
			fmt.Println("NOT WORKING !!!")
		}
		fmt.Println(*task.Text)
	}
	return nil
}
