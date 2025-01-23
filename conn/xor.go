/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package conn

const xorValue byte = 'R' // R version xors only first 4 bytes of a packet (its type)

func xorBuf(buf []byte) {
	buf[3] ^= xorValue
	buf[2] ^= xorValue
	buf[1] ^= xorValue
	buf[0] ^= xorValue
}
