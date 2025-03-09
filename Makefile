// Goを利用する
go:
	echo "go" > .mode

// C++を利用する
cpp:
	echo "cpp" > .mode

// 実行する
run:
	./shell/run.sh

// inputから入力を取らず標準入力で実行する。主にインタラクティブな問題に利用
runwi:
	./shell/runwi.sh

// ベースファイルの取得
base:
	./shell/base.sh

// ファイルの移動
m:
	./shell/movefile.sh

// コンテスト名を指定してファイルを移動
mc:
	./shell/movefilecontest.sh

// entrでのホットリロード
entr:
	./shell/entr.sh

// contestのファイルにあるコンテストからダウンロード
d:
	./shell/download.sh

// URLからダウンロード
du:
	./shell/downloadurl.sh

// テスト
t:
	./shell/test.sh

// 提出
s:
	./shell/submit.sh

// URLで提出
su:
	./shell/submiturl.sh

// AtCoderへログイン
login:
	oj login https://atcoder.jp