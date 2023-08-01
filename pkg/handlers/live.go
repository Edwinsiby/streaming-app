package handlers

import (
	"context"
	"io"
	"log"
	"net/http"

	"stream/pb"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

type StreamCamHandler struct {
	grpcClient pb.StreamingServiceClient
}

func NewAuthenticationHandler(cc *grpc.ClientConn) *StreamCamHandler {
	client := pb.NewStreamingServiceClient(cc)
	return &StreamCamHandler{
		grpcClient: client,
	}
}

func (s *StreamCamHandler) LivePage(c *gin.Context) {
	c.HTML(http.StatusFound, "live.html", nil)
}

func (s *StreamCamHandler) LiveStreamCam(c *gin.Context) {
	ws, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		log.Fatalf("Failed to upgrade WebSocket connection: %v", err)
	}
	stream, err := s.grpcClient.GetLiveStream(context.Background())
	if err != nil {
		log.Fatalf("Failed to call GetLiveStream: %v", err)
	}

	c.Header("Content-Type", "multipart/x-mixed-replace; boundary=frame")

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving streaming response: %v", err)
		}
		if err := ws.WriteMessage(websocket.BinaryMessage, resp.Data); err != nil {
			log.Fatalf("Failed to send video bytes over WebSocket: %v", err)
		}
	}
}
