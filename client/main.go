package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "distributed-cfg-service-mk/proto"
)

func main() {

	// Some naive testing here:

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDistributedCfgServiceMKClient(conn)

	service := "Service number one"
	clientApp := "Client app one"
	// clientApp2 := "Client app 2"

	params := []*pb.Parameter{
		{Key: "Tcp port", Value: "80"},
		{Key: "Memory limit", Value: "2gb"},
		{Key: "Root dir", Value: "/root"},
	}

	resp, err := c.CreateConfig(context.Background(), &pb.Config{Service: service, Parameters: params})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Config for service", resp.Service, " was successfully created at ", resp.Timestamp.AsTime())
	}

	resp2, err := c.GetConfig(
		context.Background(),
		&pb.Service{Name: service},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp2.Service)
		fmt.Println(resp2.Parameters)
	}

	resp3, err := c.UpdateConfig(context.Background(),
		&pb.Config{
			Service: service,
			Parameters: []*pb.Parameter{
				{Key: "Tcp port", Value: "9000"},
				{Key: "Root dir", Value: ""},
				{Key: "TTL", Value: "15"},
			},
		},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Config for service", resp3.Service, " was successfully updated at ", resp3.Timestamp.AsTime())
	}

	resp4, err := c.GetConfig(
		context.Background(),
		&pb.Service{Name: service},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp4.Service)
		fmt.Println(resp4.Parameters)
	}

	configTSlist, err := c.ListConfigTimestamps(
		context.Background(),
		&pb.Service{Name: service},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Available timestamps for config ", configTSlist.Service, ": ")
		for _, timstmp := range configTSlist.Timestamps {

			fmt.Println(timstmp.AsTime())
		}
	}

	resp5, err := c.GetArchivedConfig(
		context.Background(),
		&pb.Timestamp{Service: service,
			Timestamp: configTSlist.Timestamps[0],
		})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp5.Service, " ", resp5.Timestamp.AsTime())
		fmt.Println(resp5.Parameters)
	}

	resp6, err := c.SubscribeClientApp(
		context.Background(),
		&pb.SubscriptionRequest{Service: service, ClientApp: clientApp},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Client app ", clientApp, "was successfully subscribed to config ", service, resp6)
	}

	resp7, err := c.DeleteConfig(
		context.Background(),
		&pb.Service{Name: service},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted config ", resp7.Service, " at ", resp7.Timestamp.AsTime())
	}

	resp8, err := c.ListConfigSubscribers(
		context.Background(),
		&pb.Service{Name: service},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Config subscribers list:", resp8)
	}

	resp9, err := c.UnSubscribeClientApp(
		context.Background(),
		&pb.SubscriptionRequest{Service: service, ClientApp: clientApp},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Client app ", clientApp, "was successfully UnSubscribed from config ", service, resp9)
	}

	resp10, err := c.DeleteConfig(
		context.Background(),
		&pb.Service{Name: service},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted config ", resp10.Service, " at ", resp10.Timestamp.AsTime())
	}

}
