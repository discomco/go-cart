package schema

type DocFtor[TModel ISchema] func() *TModel
type GetDocKeyFunc func() string

//IModel is an Injector that represents the Write-Model Type in a CQRS Context
type IModel interface {
	ISchema
	GetStatus() int
}

//ISchema is an Injector that represents the Read-Model Type in a CQRS Context
type ISchema interface {
}
