package mypackage_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"

	mocks "github.com/NautiluX/mockstruct/centralmock/pkg/mocks/mypackage"
	"github.com/NautiluX/mockstruct/centralmock/pkg/mypackage"
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
