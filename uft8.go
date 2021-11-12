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
			} else if numCh == 2 {
				second = str[i : i+numByte]
			}

			i += numByte
		}
	}
	var buf bytes.Buffer
	buf.Write(first)
	if numCh == 1 || numCh == 2 {
		buf.WriteString("*")
	} else if numCh == 3 {
		buf.Write(second)
		buf.WriteString("*")
	} else { //  numCh >= 4
		buf.Write(second)
		buf.WriteString("**")
	}
	return buf.Bytes()
}
