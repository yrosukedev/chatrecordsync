package http_controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yrosukedev/chat_record_sync/use_case"
	"net/http"
)

type ChatSyncHTTPController struct {
	ctx     context.Context
	useCase use_case.UseCase
}

func NewChatSyncHTTPController(ctx context.Context, useCase use_case.UseCase) http.Handler {
	return &ChatSyncHTTPController{
		ctx:     ctx,
		useCase: useCase,
	}
}

func (c *ChatSyncHTTPController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	errs := c.useCase.Run(c.ctx)

	if len(errs) != 0 {
		c.writeFailureResponse(writer, errs)
		return
	}

	c.writeSuccessResponse(writer)
}

func (c *ChatSyncHTTPController) writeFailureResponse(writer http.ResponseWriter, errs []*use_case.SyncError) {
	writer.WriteHeader(http.StatusInternalServerError)

	response := ChatSyncResponse{
		Code: ResponseCodeFailure,
		Msg:  fmt.Sprintf("%v\n%v", ResponseCodeFailure.Msg(), combineErrors(errs)),
	}
	c.writeResponseBody(writer, response)
}

func (c *ChatSyncHTTPController) writeSuccessResponse(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusOK)

	response := ChatSyncResponse{
		Code: ResponseCodeOK,
		Msg:  ResponseCodeOK.Msg(),
	}
	c.writeResponseBody(writer, response)
}

func (c *ChatSyncHTTPController) writeResponseBody(writer http.ResponseWriter, response ChatSyncResponse) {
	responseBody, err := json.Marshal(response)
	if err != nil { // shouldn't happen
		writer.Write(c.responseForMarshalJsonError(err))
		return
	}
	writer.Write(responseBody)
}

func (c *ChatSyncHTTPController) responseForMarshalJsonError(err error) []byte {
	str := fmt.Sprintf("{\"code\":%v,\"msg\":\"%v:%v\"}", ResponseCodeMarshalJsonFailed, ResponseCodeMarshalJsonFailed.Msg(), err.Error())
	return []byte(str)
}
