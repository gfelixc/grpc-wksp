package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/gfelixc/grpc-wksp/server"

	"github.com/icrowley/fake"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	flight.RegisterFlightOperatorServer(server, Controller{})

	err = server.Serve(lis)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

type Controller struct {
	flight.UnsafeFlightOperatorServer
}

func (t Controller) FlightDetails(context.Context, *flight.FlightDetailsRequest) (*flight.FlightDetailsResponse, error) {
	return &flight.FlightDetailsResponse{
		Id:         "IB5011",
		Terminal:   1,
		LastStatus: flight.Status_SCHEDULED,
	}, nil
}

func (t Controller) Departures(req *flight.DeparturesRequest, stream flight.FlightOperator_DeparturesServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()

		default:
			err := stream.Send(&flight.DeparturesResponse{Id: fake.Country()})
			if err != nil {
				return err
			}

			time.Sleep(5 * time.Second)
		}
	}
}

func (t Controller) TravelUpdates(stream flight.FlightOperator_TravelUpdatesServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
			messageReceived, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				return stream.SendAndClose(&flight.TravelUpdatesResponse{})
			}

			if err != nil {
				return err
			}

			fmt.Printf("Flight: %s status %s\n", messageReceived.Id, messageReceived.LastStatus)
		}
	}
}

func (t Controller) SupportChat(stream flight.FlightOperator_SupportChatServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
			req, err := stream.Recv()
			if err != nil {
				return err
			}

			fmt.Printf("Traveler: %s says %s\n", req.TravelerId, req.Message)

			customerSupport := fake.FemaleFirstName()
			message := fake.Sentence()

			err = stream.Send(&flight.SupportChatResponse{
				CustomerSupportId: customerSupport,
				Message:           message,
			})
			if err != nil {
				return err
			}

			fmt.Printf("CustomerSupport: %s says %s\n", customerSupport, message)
		}
	}
}
