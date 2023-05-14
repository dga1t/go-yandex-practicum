package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"yandex-go-intro/exercise/hashbyte"
	"yandex-go-intro/exercise/logger"
	"yandex-go-intro/exercise/mathslice"
	"yandex-go-intro/exercise/randbyte"
	"yandex-go-intro/exercise/stopwatch"
)

func main() {

	// testLimitReader()
	// testHashbyte()
	// testLogger()
	// testStopWatch()
	// testMathslice()

	// fmt.Println("Hello, world!1")

	// ver := "v0.0.1"
	// id := 0
	// pi := 3.1415

	// fmt.Println("ver =", ver, "id =", id, "pi =", pi)

	// fizzBuzz()
	// testSlice()
	// testMap()

	// exercise #4 in maps
	//aaa := [4]int{1,2,3,5}
	//result := findApair(aaa, 4)
	//fmt.Println(result)

	// input := []string{
	// 	"cat",
	// 	"dog",
	// 	"bird",
	// 	"dog",
	// 	"parrot",
	// 	"cat",
	// }
	// fmt.Println(removeDuplicates(input))

	// serializeJson()

	// fmt.Println(Global)
	// UseGlobal()
	// fmt.Println(Global)
}

func getGeneration(year int) {
	switch {
	case year >= 1946 && year <= 1964:
		fmt.Println("Hi, Boomer!")
	case year >= 1965 && year <= 1980:
		fmt.Println("Hi, representer X!")
	case year >= 1981 && year <= 1996:
		fmt.Println("Hi, millenial!")
	case year >= 1997 && year <= 2012:
		fmt.Println("Hi, zoomer!")
	case year >= 2013:
		fmt.Println("Hi, alfa!")
	}
}

func fizzBuzz() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func testSlice() {
	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
	workDaysSlice := weekTempArr[:5]
	weekendSlice := weekTempArr[5:]
	fromTuesdayToThursDaySlice := weekTempArr[1:4]
	weekTempSlice := weekTempArr[:]

	workDaysSlice[0] = 9

	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))                                        // [1 2 3 4 5] 5 7
	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))                                           // [6 7] 2 2
	fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))                                        // [1 2 3 4 5 6 7] 7 7
	fmt.Println("------------------------------")

	dim := 100
	s := make([]int, dim)

	for i := range s {
		s[i] = i + 1
	}
	fmt.Println("s = ", s)

	s = append(s[:10], s[90:]...)
	dim = len(s)
	fmt.Println("s = ", s)
	fmt.Println("dim = ", dim)

	for i := range s[:dim/2] {
		s[i], s[dim-i-1] = s[dim-i-1], s[i]
	}
	fmt.Println("s = ", s)
}

func testMap() {
	m := make(map[string]int)
	m["хлеб"] = 50
	m["молоко"] = 100
	m["колбаса"] = 200
	m["сыр"] = 500

	fmt.Println("m = ", m)

	for k, v := range m {
		if v >= 100 {
			fmt.Printf("Product %s costs %d \n", k, v)
		}
	}

	order := []string{"молоко", "колбаса"}
	sum := 0

	for _, v := range order {
		sum += m[v]
	}

	fmt.Println("Order sum = ", sum)
}

// exercise #4 in maps
func findApair(arr []int, k int) []int {
	m := make(map[int]int)
	for i, a := range arr {
		if j, ok := m[k-a]; ok {
			return []int{i, j}
		}
		m[a] = i
	}
	return nil
}

// exercise #5 in maps
func removeDuplicates(input []string) []string {
	output := make([]string, 0)
	inputSet := make(map[string]struct{}, len(input))

	for _, v := range input {
		if _, ok := inputSet[v]; !ok {
			output = append(output, v)
		}
		inputSet[v] = struct{}{}
	}
	return output
}

// exercise #2 in structures
type Person struct {
	Name  string
	Email string
	Age   int
}

func serializeJson() {
	p := Person{"Joe", "j@k.io", 69}
	res, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("unable marshal to json")
	}
	fmt.Println(string(res))
}

// exercise #3 in structures
const rawResp = `
{
	"header": {
		"code": 0,
		"message": ""
	},
	"data": [{
		"type": "user",
		"id": 100,
		"attributes": {
			"email": "bob@yandex.ru",
			"article_ids": [10, 11, 12]
		}
	}]
}
`

type (
	Response struct {
		Header ResponseHeader `json:"header"`
		Data   ResponseData   `json:"data,omitempty"`
	}

	ResponseHeader struct {
		Code    int    `json:"code"`
		Message string `json:"message,omitempty"`
	}

	ResponseData []ResponseDataItem

	ResponseDataItem struct {
		Type       string                `json:"type"`
		Id         int                   `json:"id"`
		Attributes ResponseDataItemAttrs `json:"attributes"`
	}

	ResponseDataItemAttrs struct {
		Email      string `json:"email"`
		ArticleIds []int  `json:"article_ids"`
	}
)

func ReadResponse(rawResp string) (Response, error) {
	resp := Response{}
	if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
		return Response{}, fmt.Errorf("JSON unmarshal: %w", err)
	}

	return resp, nil
}

// exercise #1 in functions
func PrintAllFilesWithFilter(path string, filter string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}

	for _, f := range files {
		filename := filepath.Join(path, f.Name())
		if strings.Contains(filename, filter) {
			fmt.Println(filename)
		}
		if f.IsDir() {
			PrintAllFilesWithFilter(filename, filter)
		}
	}
}

// exercise #6 in functions
type figures int

const (
	square figures = iota
	circle
	triangle
)

func area(f figures) (func(float64) float64, bool) {
	switch f {
	case square:
		return func(x float64) float64 { return x * x }, true
	case circle:
		return func(x float64) float64 { return 3.142 * x * x }, true
	case triangle:
		return func(x float64) float64 { return 0.433 * x * x }, true
	default:
		return nil, false
	}
}

// exercise #7 in functions
func PrintFilesWithFuncFilter(path string, predicate func(string) bool) {
	var walk func(string)
	walk = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}

		for _, f := range files {
			filename := filepath.Join(path, f.Name())
			if predicate(filename) {
				fmt.Println(filename)
			}
			if f.IsDir() {
				walk(filename)
			}
		}
	}

	walk(path)
}

// exercise #4 in operator defer
var Global = 5

func UseGlobal() {
	defer func(checkout int) {
		Global = checkout
	}(Global)
	Global = 69
	fmt.Println(Global)
}

func testMathslice() {
	s := mathslice.Slice{1, 2, 3}
	fmt.Println(s)
	fmt.Println("Сумма слайса: ", mathslice.SumSlice(s))

	mathslice.MapSlice(s, func(i mathslice.Element) mathslice.Element {
		return i * 2
	})

	fmt.Println("Слайс, умноженный на два: ", s)

	fmt.Println("Сумма слайса: ", mathslice.SumSlice(s))

	fmt.Println("Свёртка слайса умножением ",
		mathslice.FoldSlice(s,
			func(x mathslice.Element, y mathslice.Element) mathslice.Element {
				return x * y
			},
			1))

	fmt.Println("Свёртка слайса сложением ",
		mathslice.FoldSlice(s,
			func(x mathslice.Element, y mathslice.Element) mathslice.Element {
				return x + y
			},
			0))
}

// exercise #4 in OOP
func testStopWatch() {
	sw := stopwatch.Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())
}

// exercise #5 in embeding
func testLogger() {
	logger := logger.NewLogExtended()
	logger.SetLogLevel(1)
	logger.Infoln("Не должно напечататься")
	logger.Warnln("Hello")
	logger.Errorln("World")
	logger.Println("Debug")
}

// exercise #1 in interfaces
func testHashbyte() {
	// создаём генератор случайных чисел
	generator := randbyte.New(time.Now().UnixNano()) // в качестве затравки передаём ему текущее время — при каждом запуске оно будет разным

	buf := make([]byte, 16)

	for i := 0; i < 5; i++ {
		n, _ := generator.Read(buf)
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}

	hasher := hashbyte.New(0)
	hasher.Write(buf)
	fmt.Printf("Hash: %v \n", hasher.Hash())
}

// exercise #2 in interfaces
// write a func that limits the amount of bytes that are read from io.Reader
type LimitedReader struct {
	reader io.Reader
	left   int // запоминаем количество считанных байт
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &LimitedReader{reader: r, left: n}
}

func (r *LimitedReader) Read(p []byte) (int, error) {
	if r.left == 0 {
		return 0, io.EOF
	}
	if r.left < len(p) {
		p = p[0:r.left]
	}
	n, err := r.reader.Read(p)
	r.left -= n
	return n, err
}

func testLimitReader() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
