package mypackage

//go:generate mockgen -destination=../mocks/mypackage/myinterface.go -package=mypackage -source=myinterface.go
type MyInterface interface {
	DoStuff()
	DoOtherStuf()
}

func AllTheStuff(m MyInterface) {
	m.DoStuff()
}
