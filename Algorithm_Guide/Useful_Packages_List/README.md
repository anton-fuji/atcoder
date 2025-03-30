# Goの便利パッケージ一覧

このドキュメントでは、AtCoder（およびその他の競技プログラミングサイト）で問題を解く際に役立つGoの標準ライブラリをまとめています。

---

## 標準ライブラリ一覧

### 入力処理関連

- [`bufio`](https://pkg.go.dev/bufio)  
  バッファ付き高速入力。`os.Stdin`と組み合わせて使う。

- [`os`](https://pkg.go.dev/os)  
  入力元（`os.Stdin`）やファイル処理で使用。

- [`fmt`](https://pkg.go.dev/fmt)  
  フォーマット付き入出力（小規模な問題ならこれだけでもOK）。

- [`strconv`](https://pkg.go.dev/strconv)  
  文字列と数値の相互変換に使う（例：`Atoi`, `Itoa` など）。

- [`strings`](https://pkg.go.dev/strings)  
  文字列の分割、置換、トリミングなどで便利。

---

### 数学系

- [`math`](https://pkg.go.dev/math)  
  基本的な数学関数（`Abs`, `Pow`, `Sqrt`など）。

- [`math/big`](https://pkg.go.dev/math/big)  
  多倍長整数演算。桁数の非常に大きい数を扱うときに使用（レアケース）。

---

### データ構造・アルゴリズム

- [`sort`](https://pkg.go.dev/sort)  
  スライスのソート、カスタム比較関数によるソートも可。

- [`container/heap`](https://pkg.go.dev/container/heap)  
  優先度付きキューを作るために使う（例：ダイクストラ法など）。

- [`container/list`](https://pkg.go.dev/container/list)  
  双方向リンクリスト。BFSでdeque的に使いたいときに便利。

- [`container/ring`](https://pkg.go.dev/container/ring)  
  循環バッファ（使用頻度は低いが知っておくと良い）。

---

### ⏱ 時間計測（デバッグ用）

- [`time`](https://pkg.go.dev/time)  
  処理時間の計測などに使用（提出前には削除推奨）。

---

## 🔧 便利なTips & テンプレート

- 入力は `bufio.NewScanner()` や `bufio.NewReader()` + `os.Stdin` で高速化。
- `strings.Split()` や `strconv.Atoi()` を活用して整数変換。
- `gcd`, `lcm`, `modpow` などのよく使う関数は事前に用意しておくと便利。
- `nextInt()`, `nextInts()` のようなカスタム入力関数を用意しておくと楽。

---

## 📁 よく使うスニペット例（関数テンプレ）

- 高速入力リーダー
- GCD / LCM
- 累乗（modあり/なし）
- BFS / DFS テンプレ
- Union-Find（素集合データ構造）
- ダイクストラ法 / 優先度付きキュー
- セグメントツリー（必要なら）

---

## 📝 注意点

- AtCoderでは**標準ライブラリ以外使用不可**
- 提出前に **デバッグ出力や未使用パッケージの削除** を忘れずに

