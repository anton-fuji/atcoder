# 📘 Algorithm - Searching

## **探索 (Searching)**
###  `sort.SearchInts()` (二分探索)
- **O(log n)** でソート済み配列から素早く値を検索

### 全探索（ブルートフォース）
- すべての可能な組み合わせを**片っ端から試す**方法
- シンプルかつ確実で、**小さな入力サイズの問題に強い**
- 探索範囲が広くなると処理時間が増大するため、注意が必要

*プログラムを書く前にデータの初期化を忘れないようにする*
```golang
// 二次元配列の初期化のex.
    matrix := make([][]int, r)
	for i := 0; i < r; i++ {
		matrix[i] = make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
```

```go:golang
        count := 0
		for i := 1; i <= n-2; i++ {
			for j := i + 1; j <= n-1; j++ {
				for k := j + 1; k <= n; k++ {
					if i+j+k == x {
						count += 1
					}
				}
			}
		}
```
