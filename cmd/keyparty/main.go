package keyparty

import (
	"github.com/enaml-ops/keyparty/plugin"
	"github.com/enaml-ops/omg-cli/pluginlib/product"
)

func main() {
	product.Run(new(keyparty.Plugin))
}
