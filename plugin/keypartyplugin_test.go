package keyparty_test

import (
	. "github.com/enaml-ops/omg-cli/plugins/products/vault/plugin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("given keyparty Plugin", func() {
	var plgn *Plugin

	BeforeEach(func() {
		plgn = new(Plugin)
	})

	Context("when ", func() {

		It("then ", func() {
			Ω(func() {
				panic("die badly")
			}).Should(Panic())
		})
	})

})
