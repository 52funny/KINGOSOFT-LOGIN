package Utils

import (
	"fmt"
	"strconv"
)

func StrEnc(data, key string) string {
	leng := len(data)
	var encData = ""
	var firstKeyBt [][64]byte
	var firstLength int
	if key != "" {
		firstKeyBt = getKeyBytes(key)
		firstLength = len(firstKeyBt)
	}
	if leng > 0 {
		if leng < 4 {
			bt := strToBt(data)
			var encByte [64]byte
			if key != "" {
				tempBt := bt
				for x := 0; x < firstLength; x++ {
					tempBt = enc(tempBt, firstKeyBt[x])
				}
				encByte = tempBt
			}
			encData = bt64ToHex(encByte)
			//fmt.Println(encData)
		} else {
			var iterator, remainder = leng / 4, leng % 4
			for i := 0; i < iterator; i++ {
				tempData := data[i*4+0 : i*4+4]
				tempByte := strToBt(tempData)
				var encByte [64]byte
				if key != "" {
					var tempBt [64]byte
					tempBt = tempByte
					for x := 0; x < firstLength; x++ {
						tempBt = enc(tempBt, firstKeyBt[x])
					}
					encByte = tempBt
				}
				encData += bt64ToHex(encByte)
			}
			if remainder > 0 {
				var remainderData = data[iterator*4+0 : leng]
				var tempByte = strToBt(remainderData)
				var encByte [64]byte
				if key != "" {
					tempBt := tempByte
					for x := 0; x < firstLength; x++ {
						tempBt = enc(tempBt, firstKeyBt[x])
					}
					encByte = tempBt
				}
				encData += bt64ToHex(encByte)
			}
		}
	}
	return encData
}

func bt64ToHex(byteData [64]byte) string {
	hex := ""
	for i := 0; i < 16; i++ {
		bt := ""
		for j := 0; j < 4; j++ {
			bt += fmt.Sprintf("%v", byteData[i *4 + j])
		}
		hex += bt4ToHex(bt)
		//fmt.Println(hex)
	}
	return hex
}

func bt4ToHex(binary string) string {
	var hex = ""
	switch binary {
	case "0000":
		hex = "0"
	case "0001":
		hex = "1"
	case "0010":
		hex = "2"
	case "0011":
		hex = "3"
	case "0100":
		hex = "4"
	case "0101":
		hex = "5"
	case "0110":
		hex = "6"
	case "0111":
		hex = "7"
	case "1000":
		hex = "8"
	case "1001":
		hex = "9"
	case "1010":
		hex = "A"
	case "1011":
		hex = "B"
	case "1100":
		hex = "C"
	case "1101":
		hex = "D"
	case "1110":
		hex = "E"
	case "1111":
		hex = "F"
	}
	return hex
}

func enc(dataByte [64]byte, keyByte [64]byte) [64]byte {
	var keys = generateKeys(keyByte)
	var ipByte = initPermute(dataByte)
	var ipLeft, ipRight, tempLeft [32]byte
	for k := 0; k < 32; k++ {
		ipLeft[k] = ipByte[k]
		ipRight[k] = ipByte[32+k]
	}
	for i := 0; i < 16; i++ {
		for j := 0; j < 32; j++ {
			tempLeft[j] = ipLeft[j]
			ipLeft[j] = ipRight[j]
		}
		var key [48]byte
		for m := 0; m < 48; m++ {
			key[m] = keys[i][m]
		}
		tempRight := xor32(pPermute(sBoxPermute(xor48(expandPermute(ipRight), key))), tempLeft)
		for n := 0; n < 32; n++ {
			ipRight[n] = tempRight[n]
		}
	}
	var finalData [64]byte
	for i := 0; i < 32; i++ {
		finalData[i] = ipRight[i]
		finalData[32+i] = ipLeft[i]
	}
	return finallyPermute(finalData)
}

func finallyPermute(endByte [64]byte) [64]byte {
	var fpByte [64]byte
	fpByte[0] = endByte[39]
	fpByte[1] = endByte[7]
	fpByte[2] = endByte[47]
	fpByte[3] = endByte[15]
	fpByte[4] = endByte[55]
	fpByte[5] = endByte[23]
	fpByte[6] = endByte[63]
	fpByte[7] = endByte[31]
	fpByte[8] = endByte[38]
	fpByte[9] = endByte[6]
	fpByte[10] = endByte[46]
	fpByte[11] = endByte[14]
	fpByte[12] = endByte[54]
	fpByte[13] = endByte[22]
	fpByte[14] = endByte[62]
	fpByte[15] = endByte[30]
	fpByte[16] = endByte[37]
	fpByte[17] = endByte[5]
	fpByte[18] = endByte[45]
	fpByte[19] = endByte[13]
	fpByte[20] = endByte[53]
	fpByte[21] = endByte[21]
	fpByte[22] = endByte[61]
	fpByte[23] = endByte[29]
	fpByte[24] = endByte[36]
	fpByte[25] = endByte[4]
	fpByte[26] = endByte[44]
	fpByte[27] = endByte[12]
	fpByte[28] = endByte[52]
	fpByte[29] = endByte[20]
	fpByte[30] = endByte[60]
	fpByte[31] = endByte[28]
	fpByte[32] = endByte[35]
	fpByte[33] = endByte[3]
	fpByte[34] = endByte[43]
	fpByte[35] = endByte[11]
	fpByte[36] = endByte[51]
	fpByte[37] = endByte[19]
	fpByte[38] = endByte[59]
	fpByte[39] = endByte[27]
	fpByte[40] = endByte[34]
	fpByte[41] = endByte[2]
	fpByte[42] = endByte[42]
	fpByte[43] = endByte[10]
	fpByte[44] = endByte[50]
	fpByte[45] = endByte[18]
	fpByte[46] = endByte[58]
	fpByte[47] = endByte[26]
	fpByte[48] = endByte[33]
	fpByte[49] = endByte[1]
	fpByte[50] = endByte[41]
	fpByte[51] = endByte[9]
	fpByte[52] = endByte[49]
	fpByte[53] = endByte[17]
	fpByte[54] = endByte[57]
	fpByte[55] = endByte[25]
	fpByte[56] = endByte[32]
	fpByte[57] = endByte[0]
	fpByte[58] = endByte[40]
	fpByte[59] = endByte[8]
	fpByte[60] = endByte[48]
	fpByte[61] = endByte[16]
	fpByte[62] = endByte[56]
	fpByte[63] = endByte[24]
	return fpByte
}

func pPermute(sBoxByte [32]byte) [32]byte {
	var pBoxPermute [32]byte
	pBoxPermute[0] = sBoxByte[15]
	pBoxPermute[1] = sBoxByte[6]
	pBoxPermute[2] = sBoxByte[19]
	pBoxPermute[3] = sBoxByte[20]
	pBoxPermute[4] = sBoxByte[28]
	pBoxPermute[5] = sBoxByte[11]
	pBoxPermute[6] = sBoxByte[27]
	pBoxPermute[7] = sBoxByte[16]
	pBoxPermute[8] = sBoxByte[0]
	pBoxPermute[9] = sBoxByte[14]
	pBoxPermute[10] = sBoxByte[22]
	pBoxPermute[11] = sBoxByte[25]
	pBoxPermute[12] = sBoxByte[4]
	pBoxPermute[13] = sBoxByte[17]
	pBoxPermute[14] = sBoxByte[30]
	pBoxPermute[15] = sBoxByte[9]
	pBoxPermute[16] = sBoxByte[1]
	pBoxPermute[17] = sBoxByte[7]
	pBoxPermute[18] = sBoxByte[23]
	pBoxPermute[19] = sBoxByte[13]
	pBoxPermute[20] = sBoxByte[31]
	pBoxPermute[21] = sBoxByte[26]
	pBoxPermute[22] = sBoxByte[2]
	pBoxPermute[23] = sBoxByte[8]
	pBoxPermute[24] = sBoxByte[18]
	pBoxPermute[25] = sBoxByte[12]
	pBoxPermute[26] = sBoxByte[29]
	pBoxPermute[27] = sBoxByte[5]
	pBoxPermute[28] = sBoxByte[21]
	pBoxPermute[29] = sBoxByte[10]
	pBoxPermute[30] = sBoxByte[3]
	pBoxPermute[31] = sBoxByte[24]
	return pBoxPermute
}

func sBoxPermute(expandByte [48]byte) [32]byte {
	var sBoxByte [32]byte
	binary := ""

	var s1 = [4][16]byte{
		{14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
		{0, 15, 7, 4, 14, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
		{4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
		{15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13}}

	/* Table - s2 */
	var s2 = [4][16]byte{
		{15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
		{3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
		{0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
		{13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9}}

	/* Table - s3 */
	var s3 = [4][16]byte{
		{10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
		{13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
		{13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
		{1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12}}
	/* Table - s4 */
	var s4 = [4][16]byte{
		{7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
		{13, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
		{10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
		{3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14}}

	/* Table - s5 */
	var s5 = [4][16]byte{
		{2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
		{14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
		{4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
		{11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3}}

	/* Table - s6 */
	var s6 = [4][16]byte{
		{12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
		{10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
		{9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
		{4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13}}

	/* Table - s7 */
	var s7 = [4][16]byte{
		{4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
		{13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
		{1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
		{6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12}}

	/* Table - s8 */
	var s8 = [4][16]byte{
		{13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
		{1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
		{7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
		{2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11}}

	for m := 0; m < 8; m++ {
		i, j := byte(0), byte(0)
		i = expandByte[m*6+0]*2 + expandByte[m*6+5]
		j = expandByte[m*6+1]*2*2*2 + expandByte[m*6+2]*2*2 + expandByte[m*6+3]*2 + expandByte[m*6+4]
		switch m {
		case 0:
			binary = getBoxBinary(s1[i][j])
		case 1:
			binary = getBoxBinary(s2[i][j])
		case 2:
			binary = getBoxBinary(s3[i][j])
		case 3:
			binary = getBoxBinary(s4[i][j])
		case 4:
			binary = getBoxBinary(s5[i][j])
		case 5:
			binary = getBoxBinary(s6[i][j])
			break
		case 6:
			binary = getBoxBinary(s7[i][j])
		case 7:
			binary = getBoxBinary(s8[i][j])
		}
		x, _ := strconv.Atoi(binary[0:1])
		sBoxByte[m*4+0] = byte(x)
		x, _ = strconv.Atoi(binary[1:2])
		sBoxByte[m*4+1] = byte(x)
		x, _ = strconv.Atoi(binary[2:3])
		sBoxByte[m*4+2] = byte(x)
		x, _ = strconv.Atoi(binary[3:4])
		sBoxByte[m*4+3] = byte(x)
	}
	return sBoxByte
}

func getBoxBinary(b byte) string {
	var binary = ""
	switch b {
	case 0:
		binary = "0000"
	case 1:
		binary = "0001"
	case 2:
		binary = "0010"
	case 3:
		binary = "0011"
	case 4:
		binary = "0100"
	case 5:
		binary = "0101"
	case 6:
		binary = "0110"
	case 7:
		binary = "0111"
	case 8:
		binary = "1000"
	case 9:
		binary = "1001"
	case 10:
		binary = "1010"
	case 11:
		binary = "1011"
	case 12:
		binary = "1100"
	case 13:
		binary = "1101"
	case 14:
		binary = "1110"
	case 15:
		binary = "1111"
	}
	return binary
}

func xor32(one, two [32]byte) [32]byte {
	var xorByte [32]byte
	for i := 0; i < 32; i++ {
		xorByte[i] = one[i] ^ two[i]
	}
	return xorByte
}

func xor48(one, two [48]byte) [48]byte {
	var xorByte [48]byte
	for i := 0; i < 48; i++ {
		xorByte[i] = one[i] ^ two[i]
	}
	return xorByte
}

func expandPermute(right [32]byte) [48]byte {
	var epByte [48]byte
	for i := 0; i < 8; i++ {
		if i == 0 {
			epByte[i*6+0] = right[31]
		} else {
			epByte[i*6+0] = right[i*4-1]
		}
		epByte[i*6+1] = right[i*4+0]
		epByte[i*6+2] = right[i*4+1]
		epByte[i*6+3] = right[i*4+2]
		epByte[i*6+4] = right[i*4+3]
		if i == 7 {
			epByte[i*6+5] = right[0]
		} else {
			epByte[i*6+5] = right[i*4+4]
		}
	}
	return epByte
}

func initPermute(dataByte [64]byte) [64]byte {
	var ipByte [64]byte
	for i, m, n := 0, 1, 0; i < 4; i, m, n = i+1, m+2, n+2 {
		for j, k := 7, 0; j >= 0; j, k = j-1, k+1 {
			ipByte[i*8+k] = dataByte[j*8+m]
			ipByte[i*8+k+32] = dataByte[j*8+n]
		}
	}
	return ipByte
}

func generateKeys(keyByte [64]byte) [16][48]byte {
	var key [56]byte
	var keys [16][48]byte
	var loop = [16]int{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}
	for i := 0; i < 7; i++ {
		for j, k := 0, 7; j < 8; j, k = j+1, k-1 {
			key[i*8+j] = keyByte[8*k+i]
		}
	}
	var i = 0
	for ; i < 16; i++ {
		tempLeft, tempRight := byte(0), byte(0)
		for j := 0; j < loop[i]; j++ {
			tempLeft = key[0]
			tempRight = key[28]
			for k := 0; k < 27; k++ {
				key[k] = key[k+1]
				key[28+k] = key[29+k]
			}
			key[27], key[55] = tempLeft, tempRight
		}
		var tempKey [48]byte
		tempKey[0] = key[13]
		tempKey[1] = key[16]
		tempKey[2] = key[10]
		tempKey[3] = key[23]
		tempKey[4] = key[0]
		tempKey[5] = key[4]
		tempKey[6] = key[2]
		tempKey[7] = key[27]
		tempKey[8] = key[14]
		tempKey[9] = key[5]
		tempKey[10] = key[20]
		tempKey[11] = key[9]
		tempKey[12] = key[22]
		tempKey[13] = key[18]
		tempKey[14] = key[11]
		tempKey[15] = key[3]
		tempKey[16] = key[25]
		tempKey[17] = key[7]
		tempKey[18] = key[15]
		tempKey[19] = key[6]
		tempKey[20] = key[26]
		tempKey[21] = key[19]
		tempKey[22] = key[12]
		tempKey[23] = key[1]
		tempKey[24] = key[40]
		tempKey[25] = key[51]
		tempKey[26] = key[30]
		tempKey[27] = key[36]
		tempKey[28] = key[46]
		tempKey[29] = key[54]
		tempKey[30] = key[29]
		tempKey[31] = key[39]
		tempKey[32] = key[50]
		tempKey[33] = key[44]
		tempKey[34] = key[32]
		tempKey[35] = key[47]
		tempKey[36] = key[43]
		tempKey[37] = key[48]
		tempKey[38] = key[38]
		tempKey[39] = key[55]
		tempKey[40] = key[33]
		tempKey[41] = key[52]
		tempKey[42] = key[45]
		tempKey[43] = key[41]
		tempKey[44] = key[49]
		tempKey[45] = key[35]
		tempKey[46] = key[28]
		tempKey[47] = key[31]
		var m int
		switch i {
		case 0:
			for m = 0; m < 48; m++ {
				keys[0][m] = tempKey[m]
			}
		case 1:
			for m = 0; m < 48; m++ {
				keys[1][m] = tempKey[m]
			}
		case 2:
			for m = 0; m < 48; m++ {
				keys[2][m] = tempKey[m]
			}
		case 3:
			for m = 0; m < 48; m++ {
				keys[3][m] = tempKey[m]
			}
		case 4:
			for m = 0; m < 48; m++ {
				keys[4][m] = tempKey[m]
			}
		case 5:
			for m = 0; m < 48; m++ {
				keys[5][m] = tempKey[m]
			}
		case 6:
			for m = 0; m < 48; m++ {
				keys[6][m] = tempKey[m]
			}
		case 7:
			for m = 0; m < 48; m++ {
				keys[7][m] = tempKey[m]
			}
		case 8:
			for m = 0; m < 48; m++ {
				keys[8][m] = tempKey[m]
			}
		case 9:
			for m = 0; m < 48; m++ {
				keys[9][m] = tempKey[m]
			}
		case 10:
			for m = 0; m < 48; m++ {
				keys[10][m] = tempKey[m]
			}
		case 11:
			for m = 0; m < 48; m++ {
				keys[11][m] = tempKey[m]
			}
		case 12:
			for m = 0; m < 48; m++ {
				keys[12][m] = tempKey[m]
			}
		case 13:
			for m = 0; m < 48; m++ {
				keys[13][m] = tempKey[m]
			}
		case 14:
			for m = 0; m < 48; m++ {
				keys[14][m] = tempKey[m]
			}
		case 15:
			for m = 0; m < 48; m++ {
				keys[15][m] = tempKey[m]
			}
		}
	}
	return keys
}

func getKeyBytes(key string) [][64]byte {
	var array [][64]byte
	var length = len(key)
	var iterator = length / 4
	var remainder = length % 4
	i := 0
	for ; i < iterator; i++ {
		array = append(array, strToBt(key[i*4:i*4+4]))
	}
	if remainder > 0 {
		array = append(array, strToBt(key[i*4:length]))
	}
	return array
}

func strToBt(str string) [64]byte {
	var leng = len(str)
	var bt [64]byte
	if leng < 4 {
		for i := 0; i < leng; i++ {
			var k = int(str[i])
			for j := 0; j < 16; j++ {
				pow := 1
				for m := 15; m > j; m-- {
					pow *= 2
				}
				bt[16*i+j] = byte((k / pow) % 2)
			}
		}

		for p := leng; p < 4; p++ {
			k := 0
			for q := 0; q < 16; q++ {
				pow := 1
				for m := 15; m > q; m-- {
					pow *= 2
				}
				bt[16*p+q] = byte((k / pow) % 2)
			}
		}
	} else {
		for i := 0; i < 4; i++ {
			k := int(str[i])
			for j := 0; j < 16; j++ {
				pow := 1
				for m := 15; m > j; m-- {
					pow *= 2
				}
				bt[16*i+j] = byte((k / pow) % 2)
			}
		}
	}
	return bt
}

