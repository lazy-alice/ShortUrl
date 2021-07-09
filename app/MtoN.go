package app

import "math"

var decToB64Map = map[int]string{0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F", 6: "G", 7: "H", 8: "I", 9: "J", 10: "K", 11: "L", 12: "M", 13: "N", 14: "O", 15: "P", 16: "Q", 17: "R", 18: "S", 19: "T", 20: "U", 21: "V", 22: "W", 23: "X", 24: "Y", 25: "Z", 26: "a", 27: "b", 28: "c", 29: "d", 30: "e", 31: "f", 32: "g", 33: "h", 34: "i", 35: "j", 36: "k", 37: "l", 38: "m", 39: "n", 40: "o", 41: "p", 42: "q", 43: "r", 44: "s", 45: "t", 46: "u", 47: "v", 48: "w", 49: "x", 50: "y", 51: "z", 52: "0", 53: "1", 54: "2", 55: "3", 56: "4", 57: "5", 58: "6", 59: "7", 60: "8", 61: "9", 62: "+", 63: "/"}

// 64进制转10进制
func B64ToDec(b64 string) int {
	b64ToDecMap := mapKeyValueChange(decToB64Map)
	dec := 0
	b64Length := len(b64)
	for i := 0; i < b64Length; i++ {
		dec = dec + b64ToDecMap[string(b64[i])]*int(math.Pow(64, float64(b64Length-i-1)))
	}
	return dec
}

// 十进制转64进制
func DecToB64(dec int) string {
	b64 := ""
	for dec != 0 {
		b64 = decToB64Map[dec%64] + b64
		dec = dec / 64
	}
	return b64
}

// map结构键值互换
func mapKeyValueChange(data map[int]string) map[string]int {
	newData := make(map[string]int)
	for k, v := range data {
		newData[v] = k
	}
	return newData
}

