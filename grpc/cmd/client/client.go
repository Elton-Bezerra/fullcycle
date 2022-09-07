package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/Elton-Bezerra/fullcycle/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUsersStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "João",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not connect make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "João",
		Email: "j@j.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not connect make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}

		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}

}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "Elton",
			Email: "test@hotmail.com",
		},
		{
			Id:    "2",
			Name:  "Wesley",
			Email: "test@hotmail.com",
		},
		{
			Id:    "3",
			Name:  "Weasley",
			Email: "test@hotmail.com",
		},
		{
			Id:    "4",
			Name:  "Test",
			Email: "test@hotmail.com",
		},
		{
			Id:    "5",
			Name:  "Another Test",
			Email: "test@hotmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response: %v", err)
	}

	fmt.Println("Response: ", res)
}

func AddUsersStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUsersStreamBoth(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "Elton",
			Email: "test@hotmail.com",
		},
		{
			Id:    "2",
			Name:  "Wesley",
			Email: "test@hotmail.com",
		},
		{
			Id:    "3",
			Name:  "Weasley",
			Email: "test@hotmail.com",
		},
		{
			Id:    "4",
			Name:  "Test",
			Email: "test@hotmail.com",
		},
		{
			Id:    "5",
			Name:  "Another Test",
			Email: "test@hotmail.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
			}

			fmt.Printf("Receiving user %v with status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
