// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2019 Intel Corporation

package main

import (
	"os"

	// Imports required to run agent
	"github.com/otcshare/edgenode/pkg/interfaceservice"
	"github.com/otcshare/edgenode/pkg/service"
)

// EdgeServices array contains function pointers to services start functions
var EdgeServices = []service.StartFunction{interfaceservice.Run}

func main() {

	if !service.RunServices(EdgeServices) {
		os.Exit(1)
	}

	service.Log.Infof("Service stopped gracefully")
}
