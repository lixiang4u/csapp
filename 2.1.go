package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	hexToBinMap = map[string][]byte{
		"0": []byte("0000"),
		"1": []byte("0001"),
		"2": []byte("0010"),
		"3": []byte("0011"),
		"4": []byte("0100"),
		"5": []byte("0101"),
		"6": []byte("0110"),
		"7": []byte("0111"),
		"8": []byte("1000"),
		"9": []byte("1001"),
		"A": []byte("1010"),
		"B": []byte("1011"),
		"C": []byte("1100"),
		"D": []byte("1101"),
		"E": []byte("1110"),
		"F": []byte("1111"),
	}

	binToHexMap = map[string]byte{
		"0000": byte('0'),
		"0001": byte('1'),
		"0010": byte('2'),
		"0011": byte('3'),
		"0100": byte('4'),
		"0101": byte('5'),
		"0110": byte('6'),
		"0111": byte('7'),
		"1000": byte('8'),
		"1001": byte('9'),
		"1010": byte('A'),
		"1011": byte('B'),
		"1100": byte('C'),
		"1101": byte('D'),
		"1110": byte('E'),
		"1111": byte('F'),
	}
)

func main() {
	// 4个题目
	log.Println(hexToBin("0x8F7A93"))
	log.Println(binToHex("1011011110011100"))
	log.Println(hexToBin("0xC4E5D"))
	log.Println(binToHex("1101011011011111100110"))

	// 测试
	var hex = "0x8F7A93"
	if hex == binToHex(hexToBin(hex)) {
		log.Println("测试OK")
	} else {
		log.Println("测试Error")
	}

	// 测试
	var bin = "1011011110011100"
	if bin == hexToBin(binToHex(bin)) {
		log.Println("测试OK")
	} else {
		log.Println("测试Error")
	}

}

func hexToBin(hex string) string {
	//将0x8F7A93转换为二进制
	// 16转2 直接查表
	// 16转2 直接格式化？？？

	var result []byte
	hex = strings.ToUpper(hex)

	if !strings.HasPrefix(hex, "0X") {
		log.Println("不是十六进制格式")
		return ""
	}
	hex = strings.TrimPrefix(hex, "0X")
	for _, c := range hex {
		if v, ok := hexToBinMap[string(c)]; !ok {
			log.Println("十六进制内容异常")
		} else {
			result = append(result, v...)
		}
	}

	return string(result)
}

func binToHex(bin string) string {
	// 2转16 从右往左四个一组，最左侧用零补齐四位，查表
	var l = 4
	var result []byte
	var binGroup []string
	for {
		if len(bin) <= 0 {
			break
		}
		if len(bin) < l {
			for i := 0; i <= l-len(bin); i++ {
				bin = "0" + bin
			}
			binGroup = append(binGroup, bin)
			break
		}
		binGroup = append(binGroup, bin[len(bin)-l:])
		bin = bin[:len(bin)-l]
	}
	for i := len(binGroup) - 1; i >= 0; i-- {
		if v, ok := binToHexMap[binGroup[i]]; !ok {
			log.Println("二进制数据内容异常")
		} else {
			result = append(result, v)
		}
	}
	return fmt.Sprintf("0x%s", string(result))
}
