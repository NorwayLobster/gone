package main

import (
	"bytes"
	"fmt"
)

// utf-8解码
func getText(str []byte) []byte {
	var numCh int = 0
	var first []byte
	var second []byte
	for i := 0; i < len(str) && numCh <= 4; {
		if str[i] < 128 { //a single byte
			numCh++
			if numCh == 1 {
				first = str[i : i+1]
			} else if numCh == 2 {
				second = str[i : i+1]
			}
			fmt.Printf("first: %s\n", first)
			fmt.Printf("second: %s\n", second)
			i++
		} else { //multibyte
			var numByte int = 0
			for j := 7; j >= 0; j-- {
				fmt.Printf("%b\n", str[i])
				fmt.Printf("%b\n", str[i]>>j)
				if ((str[i] >> j) & 1) == 1 {
					numByte++
				} else {
					break
				}
			}
			numCh++
			if numCh == 1 {
				first = str[i : i+numByte]
			} else if numCh == 2 {
				second = str[i : i+numByte]
			}

			fmt.Printf("first: %s\n", first)
			fmt.Printf("second: %s\n", second)
			// for k := i; k < len(str) && k < i+numByte; k++ {
			// result = append(result, str[k])
			// }
			// fmt.Printf("result: %v", result)
			// result = append(result, str[i:i+numByte])
			// fmt.Printf("%s\n", str[i:i+numByte])
			i += numByte
		}
	}
	var buf bytes.Buffer
	buf.Write(first)
	if numCh == 1 || numCh == 2 {
		buf.WriteString("*")
		// result = append(result, '*')
	} else if numCh == 3 {
		buf.Write(second)
		buf.WriteString("*")
		// result = first + second + '*'
	} else { //  numCh >= 4
		buf.Write(second)
		buf.WriteString("**")
		// result = first + second + "**"
	}
	fmt.Printf("result: %s\n", buf.Bytes())
	return buf.Bytes()
}

// utf-8解码
func getText1(str []byte) []byte {
	if len(str) == 0 {
		return nil
	}
	var numCh int = 0
	var first []byte
	var last []byte
	for i := 0; i < len(str) && numCh <= 4; {
		if str[i] < 128 { //a single byte
			numCh++
			if numCh == 1 {
				first = str[i : i+1]
			}
			last = str[i : i+1]
			i++
		} else { //multibyte
			var numByte int = 0
			for j := 7; j >= 0; j-- {
				if ((str[i] >> j) & 1) == 1 {
					numByte++
				} else {
					break
				}
			}
			numCh++
			if numCh == 1 {
				first = str[i : i+numByte]
			}
			last = str[i : i+numByte]
			i += numByte
		}
	}
	var buf bytes.Buffer
	buf.Write(first)
	if numCh == 1 || numCh == 2 {
		buf.WriteString("*")
	} else if numCh == 3 {
		buf.WriteString("*")
		buf.Write(last)
	} else { //  numCh >= 4
		buf.WriteString("**")
		buf.Write(last)
	}
	return buf.Bytes()
}

func utf8_demo() {
	// var str []byte = []byte("くうじょう じょうたろう夜桜")
	// var str []byte = []byte("中")
	// var str []byte = []byte("中华")
	// var str []byte = []byte("中华人")
	// var str []byte = []byte("中华人民")
	var str []byte = []byte("中华人民共和国")
	fmt.Printf("str: %s\n", str)
	result := getText1(str)
	fmt.Printf("result: %s\n", result)
}
