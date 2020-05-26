package ipvalid

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//
//func main() {
//	var typeM = map[string]int{}
//	var errC, priC int
//
//	reader := bufio.NewReader(os.Stdin)
//
//	for {
//		line, _, _ := reader.ReadLine()
//		if len(line) == 0 {
//			break
//		}
//
//		lines := strings.Split(string(line), "~")
//		ipStr := lines[0]
//		ip2Str := lines[1]
//
//		ip, ok := ipStr2Uint32(ipStr)
//		if !ok {
//			errC++
//			continue
//		}
//		ip2, ok := ipStr2Uint32(ip2Str)
//		if !ok || !ip2Ok(ip2) {
//			errC++
//			continue
//		}
//		tt, priF := ipType(ip)
//		typeM[tt]++
//		if priF {
//			priC++
//		}
//	}
//
//	fmt.Println(typeM["A"], typeM["B"], typeM["C"], typeM["D"], typeM["E"], errC, priC)
//}
//
//func ipStr2Uint32(in string) (uint32, bool) {
//	ins := strings.Split(in, ".")
//	var ret uint32
//	//if len(ins) != 4 {
//	//  return 0, false
//	//}
//	for _, v := range ins {
//		vInt, _ := strconv.Atoi(v)
//		if vInt > 255 || vInt < 0 {
//			return 0, false
//		}
//		ret = ret<<8 + uint32(vInt)
//	}
//	return ret, true
//}
//
//func ipType(ip uint32) (string, bool) {
//	sF := ip>>24 == 10 ||
//		ip>>20 == 172<<8+1 ||
//		ip>>16 == 192<<8+168
//
//	//tmp2 := ip >> 16 & 255
//	//if ip>>31 == 0 {
//	if tmp := ip >> 24; tmp >= 1 && tmp <= 126 {
//		return "A", sF
//	}
//	if ip>>30 == 2 {
//		//if tmp >= 128 && tmp <= 191 {
//		return "B", sF
//	}
//	if ip>>29 == 6 {
//		//if tmp >= 192 && tmp <= 223 {
//		return "C", sF
//	}
//	if ip>>28 == 14 {
//		//if tmp >= 224 && tmp <= 239 {
//		return "D", sF
//	}
//	if ip>>28 == 15 {
//		//if tmp >= 240 {
//		return "E", sF
//	}
//	return "", false
//}
//
//func ip2Ok(ip2 uint32) bool {
//	lg := math.Log2(float64(^ip2 + 1))
//	lgInt := int(lg)
//	if lg != float64(lgInt) || lgInt == 1 || lgInt == 0 {
//		return false
//	}
//	return true
//}

func main() {
	r := bufio.NewReader(os.Stdin)
	var arr = make([]string, 0, 10)
	for {
		l, _, _ := r.ReadLine()
		if len(l) == 0 {
			break
		}
		arr = append(arr, string(l))
	}
	var a, b, c, d, e, f, g int
	for i := 0; i < len(arr); i++ {
		ss := strings.Split(arr[i], "~")
		if len(ss) != 2 {
			break
		}
		ip := ss[0]
		ym := ss[1]
		ips := strings.Split(ip, ".")
		yms := strings.Split(ym, ".")

		// 统计F
		if !isValidIp(ips) {
			f++
			continue
		}
		if !isValidYm(yms) {
			f++
			continue
		}

		h, _ := strconv.Atoi(ips[0])
		h2, _ := strconv.Atoi(ips[1])

		// 统计a b c d e
		if h >= 1 && h <= 126 {
			a++

		}
		if h >= 128 && h <= 191 {
			b++
		}
		if h >= 192 && h <= 223 {
			c++
		}
		if h >= 224 && h <= 239 {
			d++
		}
		if h >= 240 && h <= 255 {
			e++
		}

		//统计G
		if h == 10 || (h == 172 && (h2 >= 16 && h2 <= 31)) || (h == 192 && h2 == 168) {
			g++
		}

	}
	fmt.Printf("%d %d %d %d %d %d %d", a, b, c, d, e, f, g)

}

func isValidIp(ips []string) (rsl bool) {
	if len(ips) != 4 {
		return
	}

	for _, v := range ips {
		if len(v) == 0 {
			return
		}
		h, e := strconv.Atoi(v)
		if e != nil {
			return
		}

		if h < 0 || h > 255 {
			return
		}
	}
	rsl = true
	return
}

func isValidYm(yms []string) (rsl bool) {
	if len(yms) != 4 {
		return
	}
	var mask = make([]int, 4)
	for k, v := range yms {
		if len(v) == 0 {
			return
		}
		h, e := strconv.Atoi(v)
		if e != nil {
			return
		}

		if h < 0 || h > 255 {
			return
		}
		mask[k] = h
	}

	// 判断子网掩码
	rsl = validMask(mask[:])

	return
}

func validMaskOne(mask int, zero bool) bool {
	judge := mask == 254 || mask == 252 || mask == 248 || mask == 240 || mask == 224 || mask == 192 || mask == 128
	// 是否包含 0
	if zero {
		judge = judge || mask == 0
	}
	return judge
}

func validMask(mask []int) bool {
	if mask[0] == 255 {
		if mask[1] == 255 {
			if mask[2] == 255 {
				return validMaskOne(mask[3], true)
			} else {
				return validMaskOne(mask[2], true) && mask[3] == 0
			}
		} else {
			return validMaskOne(mask[1], true) && mask[2] == 0 && mask[3] == 0
		}
	}
	return validMaskOne(mask[0], false) && mask[1] == 0 && mask[2] == 0 && mask[3] == 0
}
