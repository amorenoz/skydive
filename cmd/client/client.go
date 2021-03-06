/*
 * Copyright (C) 2016 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package client

import (
	"os"

	"github.com/skydive-project/skydive/config"
	"github.com/skydive-project/skydive/graffiti/logging"
	"github.com/skydive-project/skydive/graffiti/service"
	"github.com/spf13/cobra"
)

const (
	CLIService service.Type = "cli"
)

var (
	host         string
	analyzerAddr string
)

// ClientCmd describe the skydive client root command
var ClientCmd = &cobra.Command{
	Use:          "client",
	Short:        "Skydive client",
	Long:         "Skydive client",
	SilenceUsage: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmd.Root().PersistentPreRun(cmd.Root(), args)
		if analyzerAddr != "" {
			config.Set("analyzers", analyzerAddr)
		} else {
			config.SetDefault("analyzers", []string{"localhost:8082"})
		}
	},
}

// RegisterClientCommands registers the 'client' CLI subcommands
func RegisterClientCommands(cmd *cobra.Command) {
	cmd.AddCommand(AlertCmd)
	cmd.AddCommand(CaptureCmd)
	cmd.AddCommand(PacketInjectorCmd)
	cmd.AddCommand(PcapCmd)
	cmd.AddCommand(QueryCmd)
	cmd.AddCommand(ShellCmd)
	cmd.AddCommand(StatusCmd)
	cmd.AddCommand(TopologyCmd)
	cmd.AddCommand(WorkflowCmd)
	cmd.AddCommand(NodeCmd)
	cmd.AddCommand(NodeRuleCmd)
	cmd.AddCommand(EdgeCmd)
	cmd.AddCommand(EdgeRuleCmd)
}

// Operation describes a JSONPatch operation
type Operation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

func newPatchOperation(op, path string, value ...interface{}) Operation {
	patchOperation := Operation{
		Op:   op,
		Path: path,
	}
	if len(value) > 0 {
		patchOperation.Value = value[0]
	}
	return patchOperation
}

// JSONPatch describes a JSON patch
type JSONPatch []Operation

func exitOnError(err error) {
	logging.GetLogger().Error(err)
	os.Exit(1)
}

func init() {
	host, _ = os.Hostname()

	ClientCmd.PersistentFlags().StringVarP(&AuthenticationOpts.Username, "username", "", os.Getenv("SKYDIVE_USERNAME"), "username auth parameter")
	ClientCmd.PersistentFlags().StringVarP(&AuthenticationOpts.Password, "password", "", os.Getenv("SKYDIVE_PASSWORD"), "password auth parameter")
	ClientCmd.PersistentFlags().StringVarP(&analyzerAddr, "analyzer", "", os.Getenv("SKYDIVE_ANALYZER"), "analyzer address")

	ClientCmd.PersistentFlags().String("cacert", "", "certificate file to verify the peer")
	config.BindPFlag("tls.ca_cert", ClientCmd.PersistentFlags().Lookup("cacert"))
	ClientCmd.PersistentFlags().String("cert", "", "client certificate file")
	config.BindPFlag("tls.client_cert", ClientCmd.PersistentFlags().Lookup("cert"))
	ClientCmd.PersistentFlags().String("key", "", "private key file name")
	config.BindPFlag("tls.client_key", ClientCmd.PersistentFlags().Lookup("key"))

	RegisterClientCommands(ClientCmd)
}
