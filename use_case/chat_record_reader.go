package use_case

import "github.com/yrosukedev/chat_record_sync/business"

type ChatRecordReader interface {
	// Read reads one record from the data source.
	// If there is no data left to be read, Read return nil, io.EOF.
	Read() (record *business.ChatRecord, err error)
}
