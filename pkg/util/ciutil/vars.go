// Copyright 2016-2018, Pulumi Corporation.
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

package ciutil

import (
	"os"
)

// Vars contains a set of metadata variables about a CI system.
type Vars struct {
	// Name is a required friendly name of the CI system.
	Name System
	// BuildID is an optional unique identifier for the current build/job.
	BuildID string
	// BuildType is an optional friendly type name of the build/job type.
	BuildType string
	// BuildURL is an optional URL for this build/job's webpage.
	BuildURL string
	// SHA is the SHA hash of the code repo at which this build/job is running.
	SHA string
}

// DetectVars detects and returns the CI variables for the current environment.
func DetectVars() Vars {
	v := Vars{Name: DetectSystem()}
	// All CI systems have a name (that's how we know we're in a CI system). After detecting one, we will
	// try to detect some additional CI-specific metadata that the CLI will use. It's okay if we can't,
	// we'll just have reduced functionality for our various CI integrations.
	switch v.Name {
	case GitLab:
		// We are running in GitLab CI. See https://docs.gitlab.com/ee/ci/variables/.
		v.BuildID = os.Getenv("CI_JOB_ID")
		v.BuildURL = os.Getenv("CI_JOB_URL")
		v.SHA = os.Getenv("CI_COMMIT_SHA")
	case Travis:
		// We are running in Travis. See https://docs.travis-ci.com/user/environment-variables/. Travis doesn't
		// set a build URL in its environment -- see  https://github.com/travis-ci/travis-ci/issues/8935.
		v.BuildID = os.Getenv("TRAVIS_JOB_ID")
		v.BuildType = os.Getenv("TRAVIS_EVENT_TYPE")
		v.SHA = os.Getenv("TRAVIS_PULL_REQUEST_SHA")
	}
	return v
}
