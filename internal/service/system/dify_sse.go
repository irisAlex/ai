package system

import (
	"context"
	"fmt"
	"log"
	"strings"

	dify "github.com/irisAlex/ai/pkg/sdk"
)

type DifySseService struct{}

func NewDifySse() *DifySseService {
	return &DifySseService{}
}

func (d *DifySseService) Dify_Sse(query string) string {
	var (
		ctx = context.Background()
		c   = dify.NewClient("http://124.222.235.244", "app-HBpFUDFNza8VYczkUx3fw23S")

		req = &dify.ChatMessageRequest{
			Query:        query,
			User:         "user",
			ResponseMode: "streaming",
		}

		ch  chan dify.ChatMessageStreamChannelResponse
		err error
	)
	LLMErr := fmt.Sprintf("%s", "大模型创建失败")

	if ch, err = c.Api().ChatMessagesStream(ctx, req); err != nil {
		log.Panicln(err)
		return LLMErr
	}

	var strBuilder strings.Builder

	for {
		select {
		case <-ctx.Done():
			return LLMErr
		case streamData, isOpen := <-ch:
			if err = streamData.Err; err != nil {
				log.Println(err.Error())
				return LLMErr
			}
			if !isOpen {
				return strBuilder.String()
			}

			strBuilder.WriteString(streamData.Answer)
		}
	}
}
