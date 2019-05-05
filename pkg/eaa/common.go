// Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eaa

import (
	"errors"
	"strings"
)

func CommonNameStringToURN(commonName string) (URN, error) {
	splittedCN := strings.Split(commonName, ":")

	if len(splittedCN) != 2 {
		return URN{}, errors.New("Cannot translate Common Name to URN")
	}

	return URN{
		Namespace: splittedCN[0],
		ID:        splittedCN[1],
	}, nil
}