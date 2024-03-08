package hrank

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'rotateLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER d
 *  2. INTEGER_ARRAY arr
 */

func rotateLeft(d int32, arr []int32) (result []int32) {

	result = append(result, arr[d:]...)
	result = append(result, arr[:d]...)

	return
}

func ExecRotate() {
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create("/Users/koinworks/go/src/test/hrank/output")
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	// firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	// nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	// checkError(err)
	// n := int32(nTemp)

	// dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	// checkError(err)
	// d := int32(dTemp)

	// arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	// var arr []int32

	// for i := 0; i < int(n); i++ {
	// 	arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
	// 	checkError(err)
	// 	arrItem := int32(arrItemTemp)
	// 	arr = append(arr, arrItem)
	// }

	result := rotateLeft(4, []int32{1, 2, 3, 4, 5})

	// for i, resultItem := range result {
	// 	fmt.Fprintf(writer, "%d", resultItem)

	// 	if i != len(result)-1 {
	// 		fmt.Fprintf(writer, " ")
	// 	}
	// }

	// fmt.Fprintf(writer, "\n")

	// writer.Flush()
	fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
