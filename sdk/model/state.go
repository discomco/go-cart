package model

//IWriteModel is an Injector that represents the Write-Model Type in a CPQRS Context
type IWriteModel interface {
	//	GetStatus() int
}

//IReadModel is an Injector that represents the Read-Model Type in a CPQRS Context
type IReadModel interface {
}
