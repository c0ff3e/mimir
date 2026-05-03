// SPDX-FileCopyrightText: Copyright 2015-2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package strfmt

import (
	"testing"
	"time"
)

func TestParseDurationTrailingWhitespace(t *testing.T) {
	got, err := ParseDuration("3s ")
	if err != nil {
		t.Fatalf("ParseDuration() returned an unexpected error: %v", err)
	}

	if got != 3*time.Second {
		t.Fatalf("ParseDuration() = %v, want %v", got, 3*time.Second)
	}
}
