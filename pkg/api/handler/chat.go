package handler

import (
	"fmt"
	"healy-apigateway/pkg/api/response"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/helper"
	"healy-apigateway/pkg/logging"
	models "healy-apigateway/pkg/utils"
	"net/http"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type ChatHandler struct {
	Grpc_Client interfaces.ChatClient
}

func NewChatHandler(ChatingClient interfaces.ChatClient) *ChatHandler {
	return &ChatHandler{
		Grpc_Client: ChatingClient,
	}
}

var User = make(map[string]*websocket.Conn)

// WebSocket
func (ch *ChatHandler) FriendMessage(c *websocket.Conn) {
	var mu sync.Mutex
	userID := strconv.Itoa(c.Locals("user_id").(int))
	mu.Lock()
	defer mu.Unlock()
	defer delete(User, userID)
	defer c.Close()

	User[userID] = c

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			response := response.ClientResponse("Details not in correct format", nil, err.Error())
			c.WriteJSON(response)
			return
		}
		helper.SendMessageToUser(User, msg, userID)
	}
}
func (ch *ChatHandler) GetChat(c *fiber.Ctx) error {
	logEntry := logging.Logger().WithField("context", "GetChatHandler")
	logEntry.Info("Processing GetChat request")

	// Retrieve query parameters from the request
	friendID := c.Query("FriendID")
	offset := c.Query("Offset")
	limit := c.Query("Limit")

	// Validate parameters
	if friendID == "" || offset == "" || limit == "" {
		errMsg := "FriendID, Offset, or Limit missing or invalid"
		logEntry.Error(errMsg)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": errMsg,
			"data":    nil,
			"error":   "Bad Request",
		})
	}

	// Convert offset and limit to appropriate types if needed
	// ...

	userID := strconv.Itoa(c.Locals("user_id").(int))

	// Prepare the chat request object
	chatRequest := models.ChatRequest{
		FriendID: friendID,
		Offset:   offset,
		Limit:    limit,
	}

	// Call your gRPC client method with the retrieved parameters
	result, err := ch.Grpc_Client.GetChat(userID, chatRequest)
	fmt.Println(result)
	if err != nil {
		logEntry.WithError(err).Error("Error in getting chat details")
		errs := response.ClientResponse("Failed to get chat details", nil, err.Error())
		return c.Status(http.StatusBadRequest).JSON(errs)
	}

	logEntry.Info("get chat successful")
	successRes := response.ClientResponse("successfully got all chat details", result, nil)
	return c.Status(http.StatusOK).JSON(successRes)
}
