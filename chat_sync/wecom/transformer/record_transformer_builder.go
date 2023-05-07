package transformer

import "github.com/yrosukedev/chat_record_sync/chat_sync/wecom"

type RecordTransformerBuilder struct {
	openAPIService OpenAPIService
}

func NewRecordTransformerBuilder(openAPIService OpenAPIService) *RecordTransformerBuilder {
	return &RecordTransformerBuilder{
		openAPIService: openAPIService,
	}
}

func (b *RecordTransformerBuilder) Build() wecom.RecordTransformer {
	messageTypeToTransformer := map[string]FieldTransformer{
		wecom.MessageTypeText: b.textMessageTransformer(),
	}
	return NewRecordTransformerFactory(messageTypeToTransformer, b.defaultTransformer())
}

func (b *RecordTransformerBuilder) textMessageTransformer() *FieldTransformerCollection {
	// if openAPIService is nil, the sender and receiver field transformers will be excluded

	if b.openAPIService == nil {
		return NewFieldTransformerCollection([]FieldTransformer{
			NewBasicFieldTransformer(),
			NewTextContentFieldTransformer(),
		})
	}

	return NewFieldTransformerCollection([]FieldTransformer{
		NewBasicFieldTransformer(),
		NewSenderFieldTransformer(b.senderNameFetcher()),
		NewReceiverFieldTransformer(b.receiverNameFetcher()),
		NewTextContentFieldTransformer(),
	})
}

func (b *RecordTransformerBuilder) defaultTransformer() *FieldTransformerCollection {
	// if openAPIService is nil, the sender and receiver field transformers will be excluded

	if b.openAPIService == nil {
		return NewFieldTransformerCollection([]FieldTransformer{
			NewBasicFieldTransformer(),
			NewOtherContentFieldTransformer(),
		})
	}

	return NewFieldTransformerCollection([]FieldTransformer{
		NewBasicFieldTransformer(),
		NewSenderFieldTransformer(b.senderNameFetcher()),
		NewReceiverFieldTransformer(b.receiverNameFetcher()),
		NewOtherContentFieldTransformer(),
	})
}

func (b *RecordTransformerBuilder) senderNameFetcher() NameFetcher {
	return NewAnyCombinator([]NameFetcher{
		NewUserNameFetcher(b.openAPIService),
		NewContactNameFetcher(b.openAPIService),
	})
}

func (b *RecordTransformerBuilder) receiverNameFetcher() NameFetcher {
	return NewAnyCombinator([]NameFetcher{
		NewUserNameFetcher(b.openAPIService),
		NewContactNameFetcher(b.openAPIService),
	})
}
