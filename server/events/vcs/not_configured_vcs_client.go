// Copyright 2017 HootSuite Media Inc.
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Modified hereafter by contributors to runatlantis/atlantis.

package vcs

import (
	"fmt"

	"github.com/runatlantis/atlantis/server/events/models"
)

// NotConfiguredVCSClient is used as a placeholder when Atlantis isn't configured
// on startup to support a certain VCS host. For example, if there is no GitHub
// config then this client will be used which will error if it's ever called.
type NotConfiguredVCSClient struct {
	Host models.VCSHostType
}

func (a *NotConfiguredVCSClient) GetModifiedFiles(repo models.Repo, pull models.PullRequest) ([]string, error) {
	return nil, a.err()
}
func (a *NotConfiguredVCSClient) CreateComment(repo models.Repo, pullNum int, comment string) error {
	return a.err()
}
func (a *NotConfiguredVCSClient) PullIsApproved(repo models.Repo, pull models.PullRequest) (bool, error) {
	return false, a.err()
}
func (a *NotConfiguredVCSClient) PullIsMergeable(repo models.Repo, pull models.PullRequest) (bool, error) {
	return false, a.err()
}
func (a *NotConfiguredVCSClient) UpdateStatus(repo models.Repo, pull models.PullRequest, state models.CommitStatus, description string) error {
	return a.err()
}
func (a *NotConfiguredVCSClient) err() error {
	//noinspection GoErrorStringFormat
	return fmt.Errorf("Atlantis was not configured to support repos from %s", a.Host.String())
}
func (a *NotConfiguredVCSClient) MergePull(repo models.Repo, pull models.PullRequest) (bool, error) {
	return false, a.err()
}
