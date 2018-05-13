/*
Copyright (c) 2018 Red Hat, Inc.

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

// This file contains the implementation of the job template type.

package awx

type JobTemplate struct {
	id               int
	name             string
	askLimitOnLaunch bool
	askVarsOnLaunch  bool
}

func (t *JobTemplate) Id() int {
	return t.id
}

func (t *JobTemplate) Name() string {
	return t.name
}

func (t *JobTemplate) AskLimitOnLaunch() bool {
	return t.askLimitOnLaunch
}

func (t *JobTemplate) AskVarsOnLaunch() bool {
	return t.askVarsOnLaunch
}
