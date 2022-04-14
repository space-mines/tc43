package internal_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/twcrone/space-mines/tc43/internal"
)

var _ = Describe("./Internal/Sector", func() {
	Describe("Removing sectors with no radiation", func() {
		Context("when there are no sectors", func() {
			sectors := internal.GenerateBlankSectors(0)
			Context("when there are no sectors", func() {
				It("should result in empty list of sectors", func() {
					shouldBeEmpty := internal.RemoveSectorsWithNoRadiation(sectors)
					Expect(len(shouldBeEmpty)).To(Equal(0))
				})
			})
			It("should result in empty list of sectors", func() {
				shouldBeEmpty := internal.RemoveSectorsWithNoRadiation(sectors)
				Expect(len(shouldBeEmpty)).To(Equal(0))
			})
		})

		Context("when there is one sector with no radiation", func() {
			sectors := internal.GenerateBlankSectors(2)
			sectors[0].Radiation = 0
			originalSize := len(sectors)
			It("should result in empty list of with one less sector", func() {
				shouldBeOneLess := internal.RemoveSectorsWithNoRadiation(sectors)
				Expect(len(shouldBeOneLess)).To(Equal(originalSize - 1))
			})
		})
	})
})
