## 標準入力
func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
}

## ホットリロード＋他の言語使用
- entr.sh, run.sh
- test.sh, submit.sh, submiturl.sh

-------
[C - X: Yet Another Die Game :解答](https://atcoder.jp/contests/abc053/tasks/arc068_a)
```golang
package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	result := x/11*2 + min(2, (x%11+5)/6)

	fmt.Println(result)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

