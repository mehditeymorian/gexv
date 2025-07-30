package cmd

import (
	"fmt"
	"os"

	"github.com/mehditeymorian/gexv/config"
	"github.com/mehditeymorian/gexv/extractor"
	"github.com/spf13/cobra"
)

const version = "v0.1.1"

var (
	cfgFile               string
	inputFile             string
	inputText             string
	outputFile            string
	includeMatchedSection bool
)

var overrideConfig = &config.Config{}

var rootCmd = &cobra.Command{
	Use:     "regex-extractor",
	Short:   "Extract named regex groups to CSV",
	Long:    `A CLI tool that extracts named regex groups from text or files and outputs as CSV.`,
	Version: version,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig(cfgFile, overrideConfig)
		if err != nil {
			return err
		}
		src, err := extractor.GetSource(inputFile, inputText)
		if err != nil {
			return err
		}
		return extractor.ExtractToCSV(cfg, src, outputFile, includeMatchedSection)
	},
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Path to JSON config file")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "output.csv", "Output CSV file path")
	rootCmd.Flags().BoolVarP(&includeMatchedSection, "include-matched-section", "i", false, "Include matched section with whole regex")
	rootCmd.Flags().StringVarP(&overrideConfig.Pattern, "pattern", "p", "", "Pattern to match against content")
	rootCmd.Flags().StringVarP(&overrideConfig.Flags, "flags", "g", "", "list of flags to pass to regex. such as <gm> for global and multiline matching")
	rootCmd.Flags().StringVarP(&inputFile, "file", "f", "", "Input file path")
	rootCmd.Flags().StringVarP(&inputText, "text", "t", "", "Inline input text")
	rootCmd.MarkFlagsOneRequired("file", "text")
	rootCmd.MarkFlagFilename("file")
	rootCmd.MarkFlagFilename("config", "json")
	rootCmd.MarkFlagFilename("output", "csv")

}
