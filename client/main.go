package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "distributed-cfg-service-mk/proto"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDistributedCfgServiceMKClient(conn)

	service := "PostGres10"

	// params := []*pb.Parameter{
	// 	{Key: "Tcp port", Value: "111"},
	// 	// {Key: "Tcp port", Value: "33334"},
	// 	// {Key: "Memory limit", Value: "2gb"},
	// 	// {Key: "Memory limit", Value: "4gb"},
	// 	// {Key: "Root dir", Value: "/my_root/"},
	// 	// {Key: "Root dir", Value: "/my_root2/"},
	// 	{Key: "TTL", Value: "111"},
	// }

	// resp, err := c.ListConfigTimestamps(
	// 	// resp, err := c.UpdateConfig(

	// 	context.Background(),
	// 	&pb.Service{Name: service},
	// 	// &pb.ConfigTimestamp{Service: service, Timestamp: nil},
	// 	// &pb.Config{Service: service, Parameters: params},

	// )

	// st, ok := status.FromError(err)
	// if !ok {
	// 	log.Println("Error was not a status error")
	// }

	// // st2 := status.Convert(err)
	// // st3 := status.

	// if err != nil {
	// 	log.Println(st)
	// } else {
	// 	// log.Println(st2.Details()...)
	// 	// log.Println(st2.Message())
	// 	// fmt.Println(resp.GetService(), resp.GetTimestamp())
	// 	fmt.Println(resp.GetService())
	// 	fmt.Println(resp.GetTimestamps())
	// }

	// resp2, err := c.GetConfig(
	// 	// resp, err := c.UpdateConfig(

	// 	context.Background(),
	// 	&pb.Service{Name: service},
	// 	// &pb.ConfigTimestamp{Service: service, Timestamp: nil},
	// 	// &pb.Config{Service: service, Parameters: params},

	// )
	// fmt.Println(resp2.GetService())
	// fmt.Println(resp2.GetParameters())

	// resp3, err := c.GetArchivedConfig(
	// 	// resp, err := c.UpdateConfig(

	// 	context.Background(),
	// 	&pb.Timestamp{Service: "PostGres11",
	// 		Timestamp: resp.GetTimestamps()[1],
	// 	},
	// 	// &pb.ConfigTimestamp{Service: service, Timestamp: nil},
	// 	// &pb.Config{Service: service, Parameters: params},

	// )
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(resp3.GetService())
	// 	fmt.Println(resp3.GetTimestamp())
	// 	fmt.Println(resp3.GetParameters())
	// } else {
	// }

	// GOOD:
	resp, err := c.DeleteConfig(

		context.Background(),
		&pb.Service{Name: service},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp, resp.Timestamp.AsTime())
	}

	// GOOD:
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

	//GOOD
	// resp, err := c.GetConfig(
	// 	context.Background(),
	// 	&pb.Service{Name: service},
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(resp.Parameters)
	// }

	// // fmt.Println(resp.Service)
	// // fmt.Println(resp.Parameters)

}
