package test

import (
	"example.com/backend/entity"
	"testing"
	. "github.com/onsi/gomega"
)
// test
func TestValid(t *testing.T){
	g := NewGomegaWithT(t)
	books := entity.Book {
		Title: "harry potter",
		Price: 3500,
		Code: "BK123456",
	}
	t.Run("Test Valid Input",func(t *testing.T) {
		obj := books

		err := obj.ValidateBook()

		g.Expect(err).To(BeNil())
	})
}