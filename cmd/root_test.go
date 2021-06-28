package cmd

import (
	"bytes"
	"github.com/mattn/go-shellwords"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"testing"
)


func Test_Json(t *testing.T) {
	path := "../_fixtures/constraint.sql"

	analysis := []CmdTestCase{{
		Name:   "",
		Cmd:    " -i " + filepath.FromSlash(path),
	}}
	RunTestCmd(t, analysis)
}

func RunTestCmd(t *testing.T, tests []CmdTestCase) {
	RunTestCaseWithCmd(t, tests, NewRootCmd)
}

func NewRootCmd() *cobra.Command {
	return rootCmd
}

func RunTestCaseWithCmd(t *testing.T, tests []CmdTestCase, rootCmd func() *cobra.Command) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			defer ResetEnv()()

			t.Log("running Cmd: ", tt.Cmd)
			_, output, err := executeActionCommandC(tt.Cmd, rootCmd)
			if (err != nil) != tt.WantError {
				t.Errorf("expected error, got '%v'", err)
			}
			if tt.Golden != "" {
				abs, _ := filepath.Abs(tt.Golden)
				slash := filepath.FromSlash(abs)
				AssertGoldenString(t, output, slash)
			}
		})
	}
}

type CmdTestCase struct {
	Name      string
	Cmd       string
	Golden    string
	WantError bool
}

func executeActionCommandC(cmd string, rootCmd func() *cobra.Command) (*cobra.Command, string, error) {
	args, err := shellwords.Parse(cmd)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	//command := rootCmd(buf)
	command := rootCmd()

	command.SetArgs(args)

	c, err := command.ExecuteC()

	return c, buf.String(), err
}


func ResetEnv() func() {
	origEnv := os.Environ()
	return func() {
		os.Clearenv()
		for _, pair := range origEnv {
			kv := strings.SplitN(pair, "=", 2)
			os.Setenv(kv[0], kv[1])
		}
	}
}
