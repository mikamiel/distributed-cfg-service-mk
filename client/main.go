package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "distributed-cfg-service-mk/proto"
)

func main() {

	// Some simple testing here:

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDistributedCfgServiceMKClient(conn)

	service := "Test service 1"

	params := []*pb.Parameter{
		{Key: "Tcp port", Value: "80"},
		{Key: "Memory limit", Value: "2gb"},
		{Key: "Root dir", Value: "/root"},
	}

	resp, err := c.CreateConfig(context.Background(), &pb.Config{Service: service, Parameters: params})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.Service)
		fmt.Println(resp.Timestamp.AsTime())
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
		fmt.Println(resp3.Service)
		fmt.Println(resp3.Timestamp.AsTime())
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

	resp5, err := c.DeleteConfig(
		context.Background(),
		&pb.Service{Name: service},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted config ", resp5.Service, " at ", resp5.Timestamp.AsTime())
	}

	// clientApp := "client app 1"
	// resp, err := c.SubscribeClientApp(
	// 	context.Background(),
	// 	&pb.SubscriptionRequest{Service: service, ClientApp: clientApp},
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(resp)
	// }

	// GOOD:
	// clientApp := "client app 1"
	// resp, err := c.UnSubscribeClientApp(
	// 	context.Background(),
	// 	&pb.SubscriptionRequest{Service: service, ClientApp: clientApp},
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(resp)
	// }

	// GOOD:
	// resp, err := c.ListConfigSubscribers(
	// 	context.Background(),
	// 	&pb.Service{Name: service},
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(resp)
	// }

}
