package bitable_storage

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"github.com/yrosukedev/chat_record_sync/business"
	"github.com/yrosukedev/chat_record_sync/config"
	"testing"
	"time"
)

func TestWriteSucceeds(t *testing.T) {
	// Given
	ctx := context.Background()
	larkConfig := config.NewLarkConfig()
	larkClient := lark.NewClient(larkConfig.AppId, larkConfig.AppSecret)
	storageAdapter := NewChatRecordStorageAdapter(ctx, larkClient, "QCBrbzgx4aKRAis9eewcV731n7d", "tblIk692K5LXte8x")

	// When
	record := &business.ChatRecord{
		MsgId:  "CAQQluDa4QUY0On2rYSAgAMgzPrShAE=",
		Action: "send",
		From: &business.User{
			UserId: "xxyzzwksksk=",
			Name:   "Harry Wang",
		},
		To: []*business.User{
			{
				UserId: "poijnfhdwp=",
				Name:   "小明",
			},
			{
				UserId: "mngwscfgyttt=",
				Name:   "小黄",
			},
		},
		RoomId:  "wmErxtDgAA9AW32YyyuYRimKr7D1KWlw",
		MsgTime: time.Date(2023, time.March, 16, 8, 0, 8, 0, time.Local),
		MsgType: "text",
		Content: "Let's go take a dinner.",
	}

	// Then
	if err := storageAdapter.Write(record, "26df9a5c-55d8-4c52-b6ce-203325568178"); err != nil {
		t.Errorf("error should not happen here, error: %+v", err)
	}
}

func TestGenerateUUID(t *testing.T) {

	fmt.Printf("UUID: %v\n", uuid.New().String())

}
