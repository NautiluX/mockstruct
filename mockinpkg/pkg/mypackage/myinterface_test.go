package mypackage_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"

	"github.com/NautiluX/mockstruct/mockinpkg/pkg/mypackage"
	"github.com/NautiluX/mockstruct/mockinpkg/pkg/mypackage/mocks"
)

var _ = Describe("Pkg/Mypackage/Myinterface", func() {
	Context("When things happen", func() {
		It("Does stuff", func() {
			mockCtrl := gomock.NewController(GinkgoT())
			mock := mocks.NewMockMyInterface(mockCtrl)
			mock.EXPECT().DoStuff()
			mypackage.AllTheStuff(mock)
		})
	})
})
