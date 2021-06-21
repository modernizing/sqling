package cmd

import (
	"fmt"
	"github.com/inherd/sqling/modeling/parser"
	"github.com/inherd/sqling/modeling/render"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var (
	path       string
	outputType string
	rootCmd    = &cobra.Command{
		Use:   "sqling",
		Short: "Sqling is a modeling tool to build from SQL file",
		Long:  `Sqling is a modeling tool to build from SQL file.`,
		Run: func(cmd *cobra.Command, args []string) {
			dat, err := ioutil.ReadFile(path)
			render.Check(err)

			sql := string(dat)

			structs, refs := parser.ParseSql(sql)

			if outputType == "json" {
				render.OutputJson(structs, refs)
				return
			}

			render.OutputPuml(structs, refs)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&path, "input", "i", "", "input file (required)")
	rootCmd.Flags().StringVarP(&outputType, "output_type", "t", "puml", "output file type, support for puml, json")

	rootCmd.MarkFlagRequired("input")
}

func initConfig() {

}
