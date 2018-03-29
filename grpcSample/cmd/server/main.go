package main

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"net"

	"golang.org/x/net/context"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/grpcCheck/todo"

	"fmt"
	"os"

	grpc "google.golang.org/grpc"
)

func main() {

	var tasks taskServer
	server := grpc.NewServer()
	todo.RegisterTasksServer(server, tasks)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("couldn't listen to :8080 %v", err)
	}
	log.Fatal(server.Serve(l))
}

type length int64

const (
	dbPath       = "mydb.pb"
	sizeOfLength = 8
)

var endianness = binary.LittleEndian

type taskServer struct {
}

func (s taskServer) Add(ctx context.Context, t *todo.Text) (*todo.Task, error) {
	done := true
	task := &todo.Task{
		Text: t.Text,
		Done: &done,
	}

	b, err := proto.Marshal(task)
	if err != nil {
		return nil, fmt.Errorf("couldn't encode task %v", err)
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 06666)
	if err != nil {
		return nil, fmt.Errorf("couldn't open file %v", err)
	}

	if err := binary.Write(f, endianness, length(len(b))); err != nil {
		return nil, fmt.Errorf("could not encode length of message: %v", err)
	}

	_, err = f.Write(b)
	if err != nil {
		return nil, fmt.Errorf("couldn't write into the file %v", err)
	}

	if err = f.Close(); err != nil {
		return nil, fmt.Errorf("couldn't close the file %s: %v", dbPath, err)
	}

	fmt.Println(b)
	return task, nil
}

func (s taskServer) List(ctx context.Context, void *todo.Void) (*todo.TaskList, error) {
	f, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return &todo.TaskList{}, fmt.Errorf("couldn't read %s: %v", dbPath, err)
	}

	var tasks todo.TaskList
	for {

		if len(f) == 0 {
			return &tasks, nil
		} else if len(f) < sizeOfLength {
			return &todo.TaskList{}, fmt.Errorf("remaining odd %d bytes", len(f))
		}

		var l length
		if err := binary.Read(bytes.NewReader(f[:sizeOfLength]), endianness, &l); err != nil {
			return &todo.TaskList{}, fmt.Errorf("could not decode message length: %v", err)
		}

		f = f[sizeOfLength:]

		var task todo.Task
		if err := proto.Unmarshal(f[:l], &task); err != nil {
			return &todo.TaskList{}, fmt.Errorf("couldn't read task %v", err)
		}

		f = f[l:]

		tasks.Tasks = append(tasks.Tasks, &task)
	}
	return &tasks, nil
}
