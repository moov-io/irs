// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/moov-io/base/log"
	"github.com/moov-io/irs/pkg/service"

	"github.com/moov-io/irs/pkg/config"
	"github.com/moov-io/irs/pkg/file"
	"github.com/spf13/cobra"
)

var (
	inputFile string
	rawData   []byte
)

var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "Launches web server",
	Long:  "Launches web server",
	RunE: func(cmd *cobra.Command, args []string) error {
		test, _ := cmd.Flags().GetBool("test")
		if test {
			return nil
		}

		env := &service.Environment{
			Logger: log.NewDefaultLogger().Set("app", log.String("irs")),
		}

		env, err := service.NewEnvironment(env)
		if err != nil {
			env.Logger.Fatal().LogErrorf("error loading up environment", err)
			os.Exit(1)
		}
		defer env.Shutdown()

		env.Logger.Info().Log("starting services")
		shutdown := env.RunServers(true)
		defer shutdown()
		return nil
	},
}

var Validate = &cobra.Command{
	Use:   "validator",
	Short: "Validate irs file",
	Long:  "Validate an incoming irs file",
	RunE: func(cmd *cobra.Command, args []string) error {
		f, err := file.CreateFile(rawData)
		if err != nil {
			return err
		}
		return f.Validate()
	},
}

var Print = &cobra.Command{
	Use:   "print",
	Short: "Print irs file",
	Long:  "Print an incoming irs file with special format (options: irs, json)",
	RunE: func(cmd *cobra.Command, args []string) error {
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}
		if format != config.OutputJsonFormat && format != config.OutputIrsFormat {
			return errors.New("format not supported")
		}

		f, err := file.CreateFile(rawData)
		if err != nil {
			return err
		}

		output := f.Ascii()
		if format == config.OutputJsonFormat {
			buf, err := json.Marshal(f)
			if err != nil {
				return err
			}
			var pretty bytes.Buffer
			err = json.Indent(&pretty, buf, "", "  ")
			if err != nil {
				return err
			}
			output = pretty.Bytes()
		}

		fmt.Println(string(output))
		return nil
	},
}

var Convert = &cobra.Command{
	Use:   "convert [output]",
	Short: "Convert irs file format",
	Long:  "Convert an incoming irs file into another format (options: irs, json)",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires output argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			return err
		}
		if format != config.OutputJsonFormat && format != config.OutputIrsFormat {
			return errors.New("format not supported")
		}

		f, err := file.CreateFile(rawData)
		if err != nil {
			return err
		}

		output := f.Ascii()
		if format == config.OutputJsonFormat {
			buf, err := json.Marshal(f)
			if err != nil {
				return err
			}
			var pretty bytes.Buffer
			err = json.Indent(&pretty, buf, "", "  ")
			if err != nil {
				return err
			}
			output = pretty.Bytes()
		}

		wFile, err := os.Create(args[0])
		if err != nil {
			return err
		}
		_, err = wFile.Write(output)
		wFile.Close()

		return err
	},
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		isWeb := false
		cmdNames := make([]string, 0)
		getName := func(c *cobra.Command) {}
		getName = func(c *cobra.Command) {
			if c == nil {
				return
			}
			cmdNames = append([]string{c.Name()}, cmdNames...)
			if c.Name() == "web" {
				isWeb = true
			}
			getName(c.Parent())
		}
		getName(cmd)

		if !isWeb {
			if inputFile == "" {
				path, err := os.Getwd()
				if err != nil {
					log.NewDefaultLogger().LogErrorf("error getting current working directory: %v", err)
				}
				inputFile = filepath.Join(path, "irs.json")
			}
			_, err := os.Stat(inputFile)
			if os.IsNotExist(err) {
				return errors.New("invalid input file")
			}
			rawData, err = os.ReadFile(inputFile)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func initRootCmd() {
	WebCmd.Flags().BoolP("test", "t", false, "test server")
	Convert.Flags().String("format", "json", "format of irs file(required)")
	Convert.MarkFlagRequired("format")
	Print.Flags().String("format", "json", "print format")

	rootCmd.SilenceUsage = true
	rootCmd.PersistentFlags().StringVar(&inputFile, "input", "", "input file (default is $PWD/irs.json)")
	rootCmd.AddCommand(WebCmd)
	rootCmd.AddCommand(Convert)
	rootCmd.AddCommand(Print)
	rootCmd.AddCommand(Validate)
}

func main() {
	initRootCmd()

	rootCmd.Execute()
}
