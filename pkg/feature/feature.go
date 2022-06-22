/*
 Copyright 2021 The Hybridnet Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package feature

import (
	"os"

	"github.com/spf13/pflag"
	"k8s.io/apiserver/pkg/util/feature"
	"k8s.io/component-base/featuregate"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func init() {
	logger := log.Log.WithName("feature")

	if err := feature.DefaultMutableFeatureGate.Add(DefaultHybridnetFeatureGates); err != nil {
		logger.Error(err, "failed to init feature gate")
		os.Exit(1)
	}

	feature.DefaultMutableFeatureGate.AddFlag(pflag.CommandLine)
}

const (
	// owner: @bruce.mwj
	// alpha: v0.1
	//
	// Enable multi-cluster network connection.

	MultiCluster featuregate.Feature = "MultiCluster"
)

var DefaultHybridnetFeatureGates = map[featuregate.Feature]featuregate.FeatureSpec{
	MultiCluster: {
		Default:    false,
		PreRelease: featuregate.Alpha,
	},
}

func MultiClusterEnabled() bool {
	return feature.DefaultMutableFeatureGate.Enabled(MultiCluster)
}

func KnownFeatures() []string {
	return feature.DefaultMutableFeatureGate.KnownFeatures()
}
