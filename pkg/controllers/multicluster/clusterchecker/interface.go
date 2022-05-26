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

package clusterchecker

import (
	"context"
	"time"

	controllerruntime "sigs.k8s.io/controller-runtime"
)

type Checker interface {
	Register(name string, check Check) error
	Unregister(name string) error
	CheckAll(ctx context.Context, clusterManager controllerruntime.Manager, opts ...Option) (map[string]CheckResult, error)
	Check(ctx context.Context, name string, clusterManager controllerruntime.Manager, opts ...Option) (CheckResult, error)
}

type CheckResult interface {
	Succeed() bool
	Error() error
	TimeStamp() time.Time
}

type Check interface {
	Check(ctx context.Context, clusterManager controllerruntime.Manager, opts ...Option) CheckResult
}
