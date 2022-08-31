# Changelog for the go-scream SDK

## 2022.07.22
- Improved SDK layout to better show purpose
- Removed underscores from some test names

## 2022.07.21
- Added support for generic status flags - checkout utils/status
- Added KnowsCmd and KnowsEvt to IAggregate
- Some cleanup work in domain/interfaces
- Refactored ASCtor to ASFtor
- Refactored ESDBCtor to EventStoreDBFtor
- added godocs 

## 2022.07.20
- Improved README.MD
- domain.Fbk now also has Warnings and Infos
- Got rid of the "Base" constructors for Fbk,Cmd,Event
- privatized NewEvent, only NewEvt is public


## 2022.07.19
- Added CHANGES.MD
- Refactored IMediator injector to make a better distinction with IBus
- Added godoc for IMediator
- Updated README.MD
- Refactored "Add" to "Inject" in DI
- Refactored functor type AggCtor to AggFtor
- refactored functor CacheCtor to CacheFtor
- added test Injectors to domain/aggregate_test
- Renamed IAggActor to IAggPlugin because an "Actor" is a well-defined pattern
