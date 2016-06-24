package keyparty

import (
	"github.com/codegangsta/cli"
	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/omg-cli/pluginlib/product"
	"github.com/enaml-ops/omg-cli/pluginlib/util"
	"github.com/xchapter7x/lo"
)

type Plugin struct{}

func (s *Plugin) GetFlags() (flags []cli.Flag) {
	return []cli.Flag{
		cli.StringFlag{Name: "storage-name", Usage: "the hash name where your keys are stored in vault"},
		cli.StringFlag{Name: "manifest-file", Usage: "your existing manifest where you would like to inject keys from vault"},
	}
}

func (s *Plugin) GetMeta() product.Meta {
	return product.Meta{
		Name: "keyparty",
	}
}

func (s *Plugin) GetProduct(args []string, cloudConfig []byte) (b []byte) {
	if c := pluginutil.NewContext(args, s.GetFlags()); validContext(c) {
		lo.G.Debug("context", c)
		var dm = new(enaml.DeploymentManifest)
		dm.SetName("keyparty-dosomething")
		b = dm.Bytes()
	}
	return
}

func validContext(c *cli.Context) bool {
	return c.String("storage-name") == "" && c.String("manifest-file") == ""
}
