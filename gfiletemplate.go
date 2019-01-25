package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func cleanString(stream string, seperator string) []int {
	// Trims the stream and then splits
	trimmed_stream := strings.TrimSpace(stream)
	split_arr := strings.Split(trimmed_stream, seperator)
	// convert strings to integers and store them in a slice
	clean_arr := make([]int, len(split_arr))
	for index, val := range split_arr {
		clean_arr[index], _ = strconv.Atoi(val)
	}
	return clean_arr
}
func main() {
	inputreader := bufio.NewReader(os.Stdin)
	input, _ := inputreader.ReadString('\n')
	noOfTestCases, _ := strconv.Atoi(strings.TrimSpace(input))
	for i := 0; i < noOfTestCases; i++ {
		lineInput, _ := inputreader.ReadString('\n')
		fmt.Println(cleanString(lineInput, " "))
	}
}
