package cmd

import (
	"fmt"
	. "github.com/inherd/sqling/parser"
	. "github.com/inherd/sqling/render"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var (
	path        string
	output_type string
	rootCmd     = &cobra.Command{
		Use:   "sqling",
		Short: "Sqling is a modeling tool to build from SQL file",
		Long:  `Sqling is a modeling tool to build from SQL file.`,
		Run: func(cmd *cobra.Command, args []string) {
			dat, err := ioutil.ReadFile(path)
			Check(err)
			sql := string(dat)
			parseSql, refs := ParseSql(sql)

			Write(parseSql, refs)
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

	rootCmd.PersistentFlags().StringVarP(&path, "input", "i", "", "input file")
	rootCmd.PersistentFlags().StringVarP(&output_type, "output_type", "t", "puml", "output file type, support for puml, json")
}

func initConfig() {

}
