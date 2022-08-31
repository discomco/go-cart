package initialize_app

type Payload struct {
	DocId  string
	ListId string
}

func NewPayload(docId string, listId string) *Payload {
	return &Payload{
		DocId:  docId,
		ListId: listId,
	}
}
