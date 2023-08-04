package live

import (
	"context"
	"fmt"
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
	c.Header("Content-Type", "multipart/x-mixed-replace; boundary=frame")

	ws, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		log.Fatalf("Failed to upgrade WebSocket connection: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	stream, err := s.grpcClient.GetLiveStream(ctx)
	if err != nil {
		log.Fatalf("Failed to call GetLiveStream: %v", err)
	}

	go func() {
		defer cancel()

		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Println("Error receiving streaming response:", err)
				return
			}
			if err := ws.WriteMessage(websocket.BinaryMessage, resp.Data); err != nil {
				fmt.Println("Failed to send video bytes over WebSocket:", err)
				return
			}
		}
	}()

	_, _, err = ws.ReadMessage()
	if err != nil && !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		fmt.Println("Failed to read WebSocket message:", err)
	}
}
