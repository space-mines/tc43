package internal_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/twcrone/space-mines/tc43/internal"
)

var _ = Describe("Game", func() {
	Describe("Reveal Sector", func() {
		Context("when radiation > 1", func() {
			It("only that sector is revealed", func() {
				mines := []internal.Location{{X: 1, Y: 1, Z: 1}}
				sectors := internal.GenerateBlankSectors(3)
				game := internal.NewGame("test", mines, sectors, 3)
				game.Reveal(0)
				Expect(game.Sectors[0].Radiation).To(Equal(1))
				for i := 1; i < len(game.Sectors); i++ {
					Expect(game.Sectors[i].Radiation).To(Equal(-1))
				}
			})
		})
	})
})
