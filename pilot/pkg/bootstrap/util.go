// Copyright Istio Authors
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

package bootstrap

import (
	"github.com/apache/dubbo-go-pixiu/pilot/pkg/model"
	"github.com/apache/dubbo-go-pixiu/pilot/pkg/serviceregistry/provider"
	"istio.io/pkg/ledger"
)

func hasKubeRegistry(registries []string) bool {
	for _, r := range registries {
		if provider.ID(r) == provider.Kubernetes {
			return true
		}
	}
	return false
}

func buildLedger(ca RegistryOptions) ledger.Ledger {
	var result ledger.Ledger
	if ca.DistributionTrackingEnabled {
		result = ledger.Make(ca.DistributionCacheRetention)
	} else {
		result = &model.DisabledLedger{}
	}
	return result
}
