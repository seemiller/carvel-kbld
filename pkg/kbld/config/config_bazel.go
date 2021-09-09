// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package config

type SourceBazelOpts struct {
	Run SourceBazelBuildOpts
}

type SourceBazelBuildOpts struct {
	Target     *string   `json:"target"`
	RawOptions *[]string `json:"rawOptions"`
}
