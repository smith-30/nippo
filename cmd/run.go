package cmd

import (
	"fmt"
	"net/http"
	"time"

	stats_api "github.com/fukata/golang-stats-api-handler"
	"github.com/google/gops/agent"
	"github.com/smith-30/ootd/config"
	"github.com/smith-30/ootd/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	configPath = ""
	conf       = config.Config{}
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	Long:  `you can use "gops" and /api/stats endpoint to see app performance.`,
	Run: func(cmd *cobra.Command, args []string) {
		//
		// create logger
		//
		zl, err := logger.NewLogger(logger.Config{
			Development:         conf.Logger.Development,
			Level:               conf.Logger.Level,
			Encoding:            conf.Logger.Encoding,
			OutputPaths:         conf.Logger.OutputPaths,
			AppErrorOutputPaths: conf.Logger.AppErrorOutputPaths,
			ErrorOutputPaths:    conf.Logger.ErrorOutputPaths,
			EncoderConfig: logger.EncoderConfig{
				MessageKey:    conf.Logger.MessageKey,
				LevelKey:      conf.Logger.LevelKey,
				TimeKey:       conf.Logger.TimeKey,
				NameKey:       conf.Logger.NameKey,
				CallerKey:     conf.Logger.CallerKey,
				StacktraceKey: conf.Logger.StacktraceKey,
				LevelEncoder:  conf.Logger.LevelEncoder,
				CallerEncoder: conf.Logger.CallerEncoder,
			},
		})
		if err != nil {
			panic(err)
		}
		zl.Info("start process", zap.String("version_info", verInfo()))

		// for logging panic
		defer func() {
			if err := recover(); err != nil {
				zl.Panic(err)
			}
		}()

		// gops
		if conf.Gops.Enable {
			if err := agent.Listen(agent.Options{}); err != nil {
				panic(err)
			}
		}

		// for monitoring
		if conf.HealthCheck.Enable {
			go func() {
				http.HandleFunc("/api/stats", stats_api.Handler)
				if err := http.ListenAndServe(conf.HealthCheck.Host+":"+conf.HealthCheck.Port, nil); err != nil {
					panic(err)
				}
			}()
		}

		fmt.Printf("%#v\n", conf)
		time.Sleep(1 * time.Minute)
	},
}

func init() {
	// merge command
	rootCmd.AddCommand(runCmd)

	//
	// command args
	//
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "example.toml", "config file name")

	// Here you will define your flags and configuration settings.
	viper.SetConfigFile(configPath)

	// read in environment variables that match
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}
}
