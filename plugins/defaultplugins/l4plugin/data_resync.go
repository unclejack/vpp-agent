// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package l4plugin

import (
	"github.com/ligato/vpp-agent/plugins/defaultplugins/common/model/l4"
)

// ResyncAppNs configures app namespaces to the empty VPP
func (plugin *L4Configurator) ResyncAppNs(appNamespaces []*l4.AppNamespaces_AppNamespace) error {
	plugin.Log.WithField("cfg", plugin).Debug("RESYNC application namespaces begin. ")
	// Calculate and log application namespaces resync
	defer func() {
		if plugin.Stopwatch != nil {
			plugin.Stopwatch.PrintLog()
		}
	}()

	var wasError error
	if len(appNamespaces) > 0 {
		for _, appNs := range appNamespaces {
			err := plugin.ConfigureAppNamespace(appNs)
			if err != nil {
				plugin.Log.Error(err)
				wasError = err
			}
		}
	}
	plugin.Log.WithField("cfg", plugin).Debug("RESYNC application namespaces end.")

	return wasError
}

// ResyncFeatures sets initital L4Features flag
func (plugin *L4Configurator) ResyncFeatures(l4Features *l4.L4Features) error {
	plugin.Log.WithField("cfg", plugin).Debug("RESYNC L4Features begin. ")
	// Calculate and log L4Features resync
	defer func() {
		if plugin.Stopwatch != nil {
			plugin.Stopwatch.PrintLog()
		}
	}()

	var wasError error
	if l4Features != nil {
		err := plugin.ConfigureL4FeatureFlag(l4Features)
		if err != nil {
			plugin.Log.Error(err)
			wasError = err
		}
	}
	plugin.Log.WithField("cfg", plugin).Debug("RESYNC L4Features end.")

	return wasError
}
