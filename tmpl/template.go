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

// 行列の作成
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var h, w int
	fmt.Fscan(in, &h, &w)

	a := make([][]int64, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int64, w)
		for j := 0; j < w; j++ {
			fmt.Fscan(in, &a[i][j])
		}
	}

	rowSum := make([]int64, h)
	colSum := make([]int64, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			v := a[i][j]
			rowSum[i] += v
			colSum[j] += v
		}
	}
