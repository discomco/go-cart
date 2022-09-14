package schema

type DocFtor[TModel ISchema] func() *TModel
type GetDocKeyFunc func() string

//IWriteSchema is an Injector that represents the Write-Model Type in a CQRS Context
type IWriteSchema interface {
	GetStatus() int
}

//ISchema is an Injector that represents the Read-Model Type in a CQRS Context
type ISchema interface {
}
