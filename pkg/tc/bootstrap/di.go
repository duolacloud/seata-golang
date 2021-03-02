package bootstrap

import (
	"github.com/micro/cli/v2"
	xconfig "github.com/xxxmicro/base/config"
	xsource "github.com/xxxmicro/base/config/source"
	"go.uber.org/fx"
)

func App(c *cli.Context) *fx.App {
	return fx.New(fx.Options(
		fx.Provide(
			func() *cli.Context { return c },
			xsource.NewSourceProvider,
			xconfig.NewConfigProvider,
		),
		fx.Invoke(StartServer),
	))
}
