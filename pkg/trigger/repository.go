package trigger

import (
	"context"
	"fmt"

	pb "github.com/tiksup/tiksup-kafka-worker/internal/proto"
)

type GRPCRepository struct {
	Client pb.EventTriggerServiceClient
	CTX    context.Context
}

func ThrowTrigger(client pb.EventTriggerServiceClient, ctx context.Context, userID string) error {
	request := &pb.EventRequest{
		EventName: "next",
		UserId:    userID,
	}

	res, err := client.TriggerEvent(ctx, request)
	if err != nil {
		return fmt.Errorf("Error to sending request: %w", err)
	}
	if !res.Received {
		return fmt.Errorf("Event not allowed")
	}
	return nil
}
