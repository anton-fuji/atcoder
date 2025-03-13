# 📘 Algorithm - Sorting

## **ソート (Sorting)**
### **1. `sort.Ints()` (昇順ソート)**
- Go の標準ライブラリ `sort` を使用
- 計算量: **O(n log n)**
- 配列を簡単に昇順に並び替え可能
### ex.) intスライスの最小最大 sort.Ints()
```golang
	s := []int{4980, 7980, 6980}
	sort.Ints(s)             // => [4980 6980 7980]
	fmt.Println(s[len(s)-1]) // => 7980 max
	fmt.Println(s[0])        // => 4980 min
```
---------
### **2. 手動ソート (条件分岐を使う)**
- `if-else` を使って 3 つの値を並べ替え
- 計算量: **O(1)**
- 値の個数が決まっている場合に有効

