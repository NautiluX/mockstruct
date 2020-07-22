package mypackage

import "github.com/NautiluX/mockstruct/mockinpkg/pkg/mypackage"

//go:generate mockgen -destination=mocks/myinterface.go -package=mocks -source=myinterface.go
type MyInterface interface {
	DoStuff()
	DoOtherStuf()
}

func AllTheStuff(m MyInterface, m2 mypackage.MyInterface) {
	m.DoStuff()
}
