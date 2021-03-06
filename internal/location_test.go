package internal_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/twcrone/space-mines/tc43/internal"
)

var _ = Describe("Location", func() {
	Describe("Checking radiation", func() {
		location := internal.Location{X: 1, Y: 1, Z: 1}
		Context("when location IS a mine", func() {
			It("should be a 43", func() {
				Expect(location.CalculateRadiation([]internal.Location{{X: 1, Y: 1, Z: 1}})).To(Equal(43))
			})
		})

		Context("with no mines", func() {
			It("should be a 0", func() {
				Expect(location.CalculateRadiation([]internal.Location{})).To(Equal(0))
			})
		})

		Context("with one mine adjacent", func() {
			It("should be a 1", func() {
				mines := []internal.Location{{X: 0, Y: 0, Z: 0}}
				Expect(location.CalculateRadiation(mines)).To(Equal(1))
			})
		})

		Context("with two mines adjacent", func() {
			It("should be a 2", func() {
				mines := []internal.Location{{X: 0, Y: 0, Z: 0}, {X: 2, Y: 2, Z: 2}}
				Expect(location.CalculateRadiation(mines)).To(Equal(2))
			})
		})

		Context("with three mines adjacent", func() {
			It("should be a 3", func() {
				mines := []internal.Location{{X: 0, Y: 0, Z: 0}, {X: 2, Y: 2, Z: 2}, {X: 1, Y: 1, Z: 2}}
				Expect(location.CalculateRadiation(mines)).To(Equal(3))
			})
		})

	})
})
