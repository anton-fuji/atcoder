// Algorithm Lists
package main

// bufioテンプレ
in  := bufio.NewReader(os.Stdin)
out := bufio.NewWriter(os.Stdout)
defer out.Flush()

var n int
fmt.Fscan(in, &n)



// 二次元配列を作成したい時
sc := bufio.NewScanner(os.Stdin)
sc.Scan()
var n int
fmt.Sscanf(sc.Text(), "%d", &n)

s := make([]string, n)
for i := 0; i < n; i++ {
	sc.Scan()
	s[i] = sc.Text()
}

t := make([]string, n)
for i := 0; i < n; i++ {
	sc.Scan()
	t[i] = sc.Text()
}