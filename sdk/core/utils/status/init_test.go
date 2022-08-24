package status

// MyStatus is an enumeration that indicates my liveliness
type MyStatus int

const (
	Unknown MyStatus = 1 << iota
	Born    MyStatus = 2
	Living  MyStatus = 4
	Dead    MyStatus = 8
)

var me MyStatus

func init() {
	me = Unknown
}
