package controllers

import (
	"context"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sashabaranov/go-openai"
)

// type bodyMessage struct {
// 	Model string `json:"model" default0:"gpt-3.5-turbo"`

// }

func HaveConversation(c *fiber.Ctx) error {
	// Get params off of request body.
	var body struct {
		UserId  int `json:"user_id"`
		ConvoId int `json:"convo_id"`
		// Messages []map[string]string
		Request string `json:"request"`
	}

	if err := c.BodyParser(&body); err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	// Check if body.Request contains `dad joke`, else append.
	if !strings.Contains(strings.ToLower(body.Request), "dad joke") {
		body.Request = strings.Join([]string{body.Request, "dad", "joke"}, " ")
	}

	// Get basic call to OpenAi and return response.
	// Create client.
	client := openai.NewClient(os.Getenv("OPENAIKEY"))
	chatCompletionMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: body.Request,
	}
	var chatCompletionMessageList []openai.ChatCompletionMessage
	chatCompleteionRequest := openai.ChatCompletionRequest{
		Model:    "gpt-3.5-turbo",
		Messages: append(chatCompletionMessageList, chatCompletionMessage),
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		chatCompleteionRequest,
	)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"Error": err.Error(),
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": resp.Choices[0].Message.Content,
	})

	// TODO full experience:
	// Check if Convo exists in DB.
	// Exists: Get all Messages associated with Convo
	// Else: Create Convo record.

	// Create newMessage.

	// Create chatCompleteionRequest object.

	// Send chatCompletionRequst to openAI.

	// Add response to newMessage and push to DB.
	// Add newMessage.Id to Convo.Messages list and push to DB.

	// Return response.
	// return c.JSON(fiber.Map{
	// 	"message": "Reached HaveConversation",
	// })

}
