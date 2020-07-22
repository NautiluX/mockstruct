package mypackage_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"

	mypkgmock "github.com/NautiluX/mockstruct/mockinpkg/pkg/mypackage/mocks"
	mypackage "github.com/NautiluX/mockstruct/mockinpkg/pkg/otherpackage"
	otherpkgmock "github.com/NautiluX/mockstruct/mockinpkg/pkg/otherpackage/mocks"
)

var _ = Describe("Pkg/Mypackage/Myinterface", func() {
	Context("When things happen", func() {
		It("Does stuff", func() {
			mockCtrl := gomock.NewController(GinkgoT())
			myMock := mypkgmock.NewMockMyInterface(mockCtrl)
			otherMock := otherpkgmock.NewMockMyInterface(mockCtrl)
			myMock.EXPECT().DoStuff()
			otherMock.EXPECT().DoStuff()
			mypackage.AllTheStuff(myMock, otherMock)
		})
	})
})
