/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package device

func xorBuf(buf []byte, value uint8) {
	if value == 'R' {
		goto packetType
	}

	buf[23] ^= value
	buf[22] ^= value
	buf[21] ^= value
	buf[20] ^= value
	buf[19] ^= value
	buf[18] ^= value
	buf[17] ^= value
	buf[16] ^= value
	buf[15] ^= value
	buf[14] ^= value
	buf[13] ^= value
	buf[12] ^= value
	buf[11] ^= value
	buf[10] ^= value
	buf[9] ^= value
	buf[8] ^= value
	buf[7] ^= value
	buf[6] ^= value
	buf[5] ^= value
	buf[4] ^= value
packetType:
	buf[3] ^= value
	buf[2] ^= value
	buf[1] ^= value
	buf[0] ^= value
}
