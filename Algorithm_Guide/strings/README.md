# 📚 Go標準 `strings` パッケージ便利関数まとめ

## 🔍 検索系

| 関数                            | 説明                                | 例                                   |
|---------------------------------|-------------------------------------|--------------------------------------|
| `strings.Contains(s, substr)`   | `s` に `substr` が含まれているか    | `"go"` in `"golang"` → ✅            |
| `strings.HasPrefix(s, prefix)`  | `s` が `prefix` で始まるか          | `"http"` in `"https://"`             |
| `strings.HasSuffix(s, suffix)`  | `s` が `suffix` で終わるか          | `"jpg"` in `"cat.jpg"`               |
| `strings.Index(s, substr)`      | 最初の出現位置（なければ -1）       | `"l"` in `"hello"` → `2`             |
| `strings.LastIndex(s, substr)`  | 最後の出現位置（なければ -1）       | `"l"` in `"hello"` → `3`             |

## 🔁 変換・置換系

| 関数                                 | 説明                              | 例                                     |
|--------------------------------------|-----------------------------------|----------------------------------------|
| `strings.ToUpper(s)`                 | 大文字に変換                      | `"Go"` → `"GO"`                        |
| `strings.ToLower(s)`                 | 小文字に変換                      | `"Go"` → `"go"`                        |
| `strings.ReplaceAll(s, old, new)`    | 全部置換                          | `"aabb"` → `"xxxx"`                    |
| `strings.Repeat(s, count)`           | 文字列を繰り返す                  | `"na"` * 4 → `"nananana"`              |

## ✂️ 分割・結合系

| 関数                       | 説明                                | 例                                             |
|----------------------------|-------------------------------------|------------------------------------------------|
| `strings.Split(s, sep)`    | 区切り文字で分割                    | `"a,b,c"` → `["a", "b", "c"]`                 |
| `strings.Join(slice, sep)` | スライスを区切り文字で結合         | `["a", "b"]` → `"a,b"`                         |
| `strings.Fields(s)`        | 空白で分割（複数スペースもOK）      | `"Go is fun"` → `["Go", "is", "fun"]`         |

## 🧼 空白・トリム系

| 関数                                 | 説明                                   | 例                                               |
|--------------------------------------|----------------------------------------|--------------------------------------------------|
| `strings.TrimSpace(s)`               | 前後の空白・改行を削除                 | `"  hello\n"` → `"hello"`                       |
| `strings.Trim(s, cutset)`            | 指定文字を前後から除去                 | `"--hi--"` → `Trim("--hi--", "-")` → `"hi"`     |
| `strings.TrimPrefix(s, prefix)`      | 指定接頭文字列を削除                   | `"abcde"` → `TrimPrefix("abcde", "ab")` → `"cde"` |
| `strings.TrimSuffix(s, suffix)`      | 指定接尾文字列を削除                   | `"file.txt"` → `TrimSuffix("file.txt", ".txt")` → `"file"` |

## 💡 その他便利系

| 関数                            | 説明                                | 例                                       |
|---------------------------------|-------------------------------------|------------------------------------------|
| `strings.Count(s, substr)`      | `substr` の出現回数                 | `"banana"` → `"a"` → `3`                 |
| `strings.Compare(a, b)`         | 辞書順比較（-1, 0, 1 を返す）       | `"apple" < "banana"` → `-1`              |
| `strings.EqualFold(a, b)`       | 大文字小文字を無視して比較         | `"Go"` と `"go"` → `true`                |

## 🧪 使用例（スニペット）

```go
s := "   hello,world,go!   "
s = strings.TrimSpace(s)                  // "hello,world,go!"
parts := strings.Split(s, ",")            // ["hello", "world", "go!"]
joined := strings.Join(parts, "-")        // "hello-world-go!"
fmt.Println(strings.ToUpper(joined))      // "HELLO-WORLD-GO!"
```