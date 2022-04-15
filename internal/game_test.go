package internal_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/twcrone/space-mines/tc43/internal"
)

var _ = Describe("Game", func() {
	Describe("State is 'LOSE'", func() {
		Context("when game state is 'LOSE'", func() {
			mines := []internal.Location{{X: 1, Y: 1, Z: 1}}
			sectors := internal.GenerateBlankSectors(3)
			game := internal.NewGame("test", mines, sectors, 3)
			game.State = "LOSE"

			It("cannot reveal sectors", func() {
				game.Reveal(0)
				for _, sector := range game.Sectors {
					Expect(sector.Radiation).To(Equal(-1))
					Expect(sector.Marked).To(Equal(false))
				}
			})

			It("cannot mark sectors", func() {
				game.Mark(0)
				for _, sector := range game.Sectors {
					Expect(sector.Radiation).To(Equal(-1))
					Expect(sector.Marked).To(Equal(false))
				}
			})
		})
	})

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

		Context("when radiation = 0", func() {
			mines := []internal.Location{{X: 0, Y: 0, Z: 0}}
			sectors := internal.GenerateBlankSectors(3)
			game := internal.NewGame("test", mines, sectors, 3)
			game.Reveal(2) // click edge
			sectorIdsAdjacentToMine := internal.GetAdjacentSectorIdsFor(0, 0, 0, 3)

			It("reveals all sectors with 0 radiation", func() {
				for sectorId := range game.Sectors {
					if !internal.Contains(sectorIdsAdjacentToMine, sectorId) && sectorId != 0 {
						sector := game.Sectors[sectorId]
						Expect(sector.Radiation).To(Equal(0))
					}
				}
			})

			It("does not reveal the mine", func() {
				Expect(game.Sectors[0].Radiation).To(Equal(-1))
			})

			It("reveals sectors around mine", func() {
				for sectorId := range game.Sectors {
					if internal.Contains(sectorIdsAdjacentToMine, sectorId) {
						sector := game.Sectors[sectorId]
						Expect(sector.Radiation).To(Equal(1))
					}
				}
			})

		})

		Context("when sector is already marked", func() {
			mines := []internal.Location{{X: 0, Y: 0, Z: 0}}
			sectors := internal.GenerateBlankSectors(3)
			game := internal.NewGame("test", mines, sectors, 3)
			game.Sectors[2].Marked = true
			game.Reveal(2) // click edge
			It("should not reveal any sectors", func() {
				for sectorId := range game.Sectors {
					sector := game.Sectors[sectorId]
					Expect(sector.Radiation).To(Equal(-1))
				}
			})
		})

		Context("when sector is already revealed", func() {
			mines := []internal.Location{{X: 0, Y: 0, Z: 0}}
			sectors := internal.GenerateBlankSectors(3)
			game := internal.NewGame("test", mines, sectors, 3)
			game.Sectors[2].Radiation = 0
			game.Reveal(2) // click edge
			It("should not reveal any more sectors", func() {
				for sectorId := range game.Sectors {
					sector := game.Sectors[sectorId]
					if sectorId == 2 {
						Expect(sector.Radiation).To(Equal(0))
					} else {
						Expect(sector.Radiation).To(Equal(-1))
					}
				}
			})
		})

		Context("when sector ID is invalid", func() {
			It("does nothing", func() {
				mines := []internal.Location{{X: 1, Y: 1, Z: 1}}
				sectors := internal.GenerateBlankSectors(3)
				game := internal.NewGame("test", mines, sectors, 3)
				game.Reveal(100)
				for i := 0; i < len(game.Sectors); i++ {
					Expect(game.Sectors[i].Radiation).To(Equal(-1))
				}
				game.Reveal(-100)
				for i := 0; i < len(game.Sectors); i++ {
					Expect(game.Sectors[i].Radiation).To(Equal(-1))
				}
			})
		})

		Context("when sector contains mine", func() {
			mines := []internal.Location{{X: 0, Y: 0, Z: 0}}
			sectors := internal.GenerateBlankSectors(3)
			game := internal.NewGame("test", mines, sectors, 3)
			game.Reveal(0)

			It("sets game state to LOSE", func() {
				Expect(game.State).To(Equal("LOSE"))
			})

			It("removes all sectors", func() {
				Expect(len(game.Sectors)).To(Equal(0))
			})
		})
	})

	Describe("Mark Sector", func() {
		Context("when sector is not marked or revealed", func() {
			mines := []internal.Location{{X: 0, Y: 0, Z: 0}}
			sectors := internal.GenerateBlankSectors(3)
			game := internal.NewGame("test", mines, sectors, 3)
			game.Mark(2) // click edge
			It("should mark that sector", func() {
				for sectorId := range game.Sectors {
					sector := game.Sectors[sectorId]
					if sectorId == 2 {
						Expect(sector.Radiation).To(Equal(-1))
						Expect(sector.Marked).To(Equal(true))
					} else {
						Expect(sector.Radiation).To(Equal(-1))
					}
				}
			})
		})

		Context("when sector is already marked", func() {
			mines := []internal.Location{{X: 0, Y: 0, Z: 0}}
			sectors := internal.GenerateBlankSectors(3)
			game := internal.NewGame("test", mines, sectors, 3)
			game.Sectors[2].Marked = true
			game.Mark(2)
			It("should unmark that sector", func() {
				for sectorId := range game.Sectors {
					sector := game.Sectors[sectorId]
					if sectorId == 2 {
						Expect(sector.Radiation).To(Equal(-1))
						Expect(sector.Marked).To(Equal(false))
					} else {
						Expect(sector.Radiation).To(Equal(-1))
					}
				}
			})
		})

		Context("when sector is already revealed", func() {
			mines := []internal.Location{{X: 0, Y: 0, Z: 0}}
			sectors := internal.GenerateBlankSectors(3)
			game := internal.NewGame("test", mines, sectors, 3)
			game.Sectors[2].Radiation = 1
			game.Mark(2)
			It("should NOT mark that sector", func() {
				for sectorId := range game.Sectors {
					sector := game.Sectors[sectorId]
					if sectorId == 2 {
						Expect(sector.Radiation).To(Equal(1))
					} else {
						Expect(sector.Radiation).To(Equal(-1))
					}
					Expect(sector.Marked).To(Equal(false))
				}
			})
		})
	})

	Context("when sector ID is invalid", func() {
		It("does nothing", func() {
			mines := []internal.Location{{X: 1, Y: 1, Z: 1}}
			sectors := internal.GenerateBlankSectors(3)
			game := internal.NewGame("test", mines, sectors, 3)
			game.Mark(100)
			for i := 0; i < len(game.Sectors); i++ {
				Expect(game.Sectors[i].Radiation).To(Equal(-1))
			}
			game.Mark(-100)
			for i := 0; i < len(game.Sectors); i++ {
				Expect(game.Sectors[i].Radiation).To(Equal(-1))
			}
		})
	})

	Context("when sector is only mine", func() {
		mines := []internal.Location{{X: 0, Y: 0, Z: 0}}
		sectors := internal.GenerateBlankSectors(3)
		game := internal.NewGame("test", mines, sectors, 3)
		game.Mark(0)
		It("should set game state to 'WIN'", func() {
			Expect(game.State).To(Equal("WIN"))
		})

		It("should reveal all sectors", func() {
			for i := 0; i < len(game.Sectors); i++ {
				Expect(game.Sectors[i].Radiation).NotTo(Equal(-1))
			}
		})
	})

})
