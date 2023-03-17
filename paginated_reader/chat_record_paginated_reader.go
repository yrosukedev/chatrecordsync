package paginated_reader

import (
	"github.com/yrosukedev/chat_record_sync/business"
)

type ChatRecordPaginatedReader struct {
	paginatedBufferedReader ChatRecordPaginatedBufferedReader
	paginationStorage       ChatRecordPaginationStorage
}

func NewChatRecordPaginatedReader(paginatedBufferedReader ChatRecordPaginatedBufferedReader, paginationStorage ChatRecordPaginationStorage) *ChatRecordPaginatedReader {
	return &ChatRecordPaginatedReader{
		paginatedBufferedReader: paginatedBufferedReader,
		paginationStorage:       paginationStorage,
	}
}

func (r *ChatRecordPaginatedReader) Read() (records []*business.ChatRecord, err error) {
	pageToken, err := r.paginationStorage.Get()
	if err != nil {
		return nil, err
	}

	return r.paginatedBufferedReader.Read(pageToken)
}
