// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2021 Canonical Ltd
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 * SPDX-License-Identifier: Apache-2.0'
 */

package main

import (
	"os"
	"path/filepath"
	"testing"

	hooks "github.com/canonical/edgex-snap-hooks/v2"
)

func TestInstallFiles(t *testing.T) {
	// override snap paths to temp directories,
	// 	which get cleaned up automatically after the tests
	hooks.Snap = t.TempDir()
	hooks.SnapData = t.TempDir()

	t.Run("install-one", func(t *testing.T) {
		path := "/config/device-camera/res/configuration.toml"
		createSourceFile(hooks.Snap+path, t)

		err := installFiles(path)
		if err != nil {
			t.Fatalf("error installing files: %s", err)
		}

		checkDestFile(hooks.SnapData+path, t)
	})
}

// utility functions

const (
	testFileContent = "file data"
)

func createSourceFile(path string, t *testing.T) {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		t.Fatalf("error creating directory hierarchy: %s", err)
	}

	err = os.WriteFile(path, []byte(testFileContent), 0644)
	if err != nil {
		t.Fatalf("error writing file: %s", err)
	}
}

func checkDestFile(path string, t *testing.T) {
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("error reading file: %s", err)
	}

	if string(content) != testFileContent {
		t.Fatalf("expected file content: %s\ngot: %s", content, testFileContent)
	}
}
