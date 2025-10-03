/*
 * Copyright (c) 2025 kompl3xpr. All rights reserved.
 * Use of this source code is governed by a MIT license
 * that can be found in the LICENSE file.
 *
 * File:               test/integration/cli_integration_test.go
 * Last modified by:   kompl3xpr <kompl3xpr@proton.me>
 * Last modified time: 2025-10-03 15:48:48 UTC
 */

package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	assert.Equal(t, 2+2, 4)
}
