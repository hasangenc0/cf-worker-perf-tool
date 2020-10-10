package cmd

import (
	"github.com/hasangenc0/cf-worker-perf-tool/pkg/profiler"
	"os"

	"github.com/hasangenc0/cf-worker-perf-tool/pkg/reporter"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	url     string
	profile int
)

var rootCmd = &cobra.Command{
	Use:   "cf-worker-perf-tool",
	Short: "Performance analysis tool for Cloudflare workers",
	Run: func(cmd *cobra.Command, args []string) {
		if profile > 0 && url != "" {
			profiler.GetMetrics(profile, url)
			return
		}

		if url != "" {
			reporter.ReportHttpResponse(url)
			return
		}

		cmd.Help()
		os.Exit(0)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&url, "url", "", "Full url of the application to analyse")
	rootCmd.PersistentFlags().IntVar(&profile, "profile", 0, "Number of times to send request for profiling (must be bigger than zero)")
}
