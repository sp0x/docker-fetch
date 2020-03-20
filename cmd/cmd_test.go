package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cmd", func() {
	Describe("Name parsing", func() {
		Context("when tag is given", func() {
			It("should parse image names", func() {
				i := ParseImageInfo("a/b:latest")
				Expect(i.Name).To(Equal("b"))
			})
		})
	})
})
