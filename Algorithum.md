# 📘 Algorithm

## **ソート (Sorting)**
### **1. `sort.Ints()` (昇順ソート)**
- Go の標準ライブラリ `sort` を使用
- 計算量: **O(n log n)**
- 配列を簡単に昇順に並び替え可能

### **2. 手動ソート (条件分岐を使う)**
- `if-else` を使って 3 つの値を並べ替え
- 計算量: **O(1)**
- 値の個数が決まっている場合に有効

## **探索 (Searching)**
### **1. `sort.SearchInts()` (二分探索)**
- **O(log n)** でソート済み配列から素早く値を検索

## **データ構造 (Data Structures)**
### **1. スライス (Slice)**
- `make([]int, size)` で作成
- `append()` で動的に追加可能

### **2. マップ (Map)**
- `map[string]int{}` の形で宣言
- キーと値のペアを保持

## **数学 (Math)**
### **1. `math.Max()` & `math.Min()`**
- `math.Max(a, b)` → 2 つの値の最大値
- `math.Min(a, b)` → 2 つの値の最小値

## **文字列操作 (String Manipulation)**
### **1. `strings.Split()` (文字列の分割)**
- 文字列を特定の区切りで分割

### **2. `strings.Join()` (配列の結合)**
- スライスを文字列に結合

