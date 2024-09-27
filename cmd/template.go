package cmd

import (
	"github.com/atomeight/distributed-event-collector/component/dec"
	"github.com/atomeight/distributed-event-collector/core"
	"github.com/golobby/container/v3"
	"github.com/latifrons/commongo/safe_viper"
	"github.com/latifrons/latigo"
	"github.com/latifrons/latigo/logging"
	"github.com/latifrons/latigo/program"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(run)
}

var run = &cobra.Command{
	Use:   "dec",
	Short: "start dec server",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		program.LoadConfigs(program.FolderConfig{
			Root: "data",
		}, "INJ")

		lvl, err := zerolog.ParseLevel(safe_viper.ViperMustGetString("debug.log_level"))
		if err != nil {
			panic(err)
		}
		logging.SetupDefaultLogger(lvl)

		err = core.BuildDependencies(dec.Singletons)
		if err != nil {
			panic(err)
		}

		var decEngineSetup *dec.DECSetup
		err = container.Resolve(&decEngineSetup)
		if err != nil {
			panic(err)
		}

		engine := latigo.NewDefaultEngineV2()
		engine.Jobs = decEngineSetup.ProvideBootSequence()
		engine.Start()
	},
}
