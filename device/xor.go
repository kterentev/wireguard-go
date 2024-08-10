/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package device

func XorBuffer(buffer []byte, xorValue uint8) {
	xorLength := MaxXorSize

	if len(buffer) < xorLength {
		xorLength = len(buffer)
	}

	for i := 0; i < xorLength; i++ {
		buffer[i] ^= xorValue
	}
}
