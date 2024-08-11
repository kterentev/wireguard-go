/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package conn

const (
	MaxXorSize = 24  // maximum size of packet to be xored (less than device.MinMessageSize 32)
	XorValue   = 81  // Q uppercase q
)

func XorBuffer(buffer []byte) {
		xorLength := MaxXorSize

		if len(buffer) < xorLength {
				xorLength = len(buffer)
		}

		for i := 0; i < xorLength; i++ {
				buffer[i] ^= XorValue
		}
}
