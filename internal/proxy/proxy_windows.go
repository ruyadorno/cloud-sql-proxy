// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proxy

import (
	"context"
	"errors"
	"path/filepath"
	"strings"
)

var errFUSENotSupported = errors.New("FUSE is not supported on Windows")

// UnixAddress returns the Unix socket for a given instance in the provided
// directory, by replacing all colons in the instance's name with periods.
func UnixAddress(dir, inst string) string {
	inst2 := strings.ReplaceAll(inst, ":", ".")
	return filepath.Join(dir, inst2)
}

type fuseMount struct {
	// fuseDir is always an empty string on Windows.
	fuseDir string
}

func configureFUSE(c *Client, conf *Config) (*Client, error)         { return nil, errFUSENotSupported }
func (c *Client) fuseMounts() []*socketMount                         { return nil }
func (c *Client) serveFuse(ctx context.Context, notify func()) error { return errFUSENotSupported }
func (c *Client) unmountFUSE() error                                 { return nil }
func (c *Client) waitForFUSEMounts()                                 {}
