package initialize_app

type IEvt interface {
}

type Evt struct {
	payload *Payload
}

func NewEvt(docId string, listId string) IEvt {
	return &Evt{
		payload: NewPayload(docId, listId),
	}
}
