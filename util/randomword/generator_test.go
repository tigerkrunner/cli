package randomword_test

import (
	"strings"
	"time"

	. "code.cloudfoundry.org/cli/util/randomword"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generator", func() {
	var gen Generator

	BeforeEach(func() {
		gen = Generator{}
	})

	Describe("RandomAdjective", func() {
		It("generates a random adjective each time it is called", func() {
			setOne := []string{}
			setTwo := []string{}

			for i := 0; i < 3; i++ {
				setOne = append(setOne, gen.RandomAdjective())
				// We wait for 3 millisecond because the seed we use to generate the
				// randomness has a unit of 1 nanosecond plus random test flakiness
				time.Sleep(3)
				setTwo = append(setTwo, gen.RandomAdjective())
			}
			Expect(setOne).ToNot(ConsistOf(setTwo))
		})
	})

	Describe("RandomNoun", func() {
		It("generates a random noun each time it is called", func() {
			setOne := []string{}
			setTwo := []string{}

			for i := 0; i < 3; i++ {
				setOne = append(setOne, gen.RandomNoun())
				// We wait for 3 millisecond because the seed we use to generate the
				// randomness has a unit of 1 nanosecond plus random test flakiness
				time.Sleep(3)
				setTwo = append(setTwo, gen.RandomNoun())
			}
			Expect(setOne).ToNot(ConsistOf(setTwo))
		})
	})

	Describe("Babble", func() {
		It("generates a random adjective noun pair each time it is called", func() {
			wordPair := gen.Babble()
			Expect(wordPair).To(MatchRegexp(`^\w+-\w+-\w+$`))
		})
	})

	Describe("Non-repeating adjectives", func() {
		It("generates a string of three words where the first two adjectives are distinct", func() {
			for i := 0; i < 50; i++ {
				wordPair := gen.Babble()
				result := strings.Split(wordPair, "-")
				Expect(result[0]).ToNot(Equal(result[1]))
			}
		})
	})
})
