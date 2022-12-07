# picture-appload

## 使い方

1. localサーバを起動
```go run main.go```
2. ブラウザで 127.0.0.1:8080 にアクセス
3. ファイルを選択を押してアップロードしたい画像を選択(一部対応していない拡張子があります。)
4. 送信を押してアップロード!
5. (ファイル名)uploaded!!と表示されるので 戻る から初めの画面に戻る
6. image list から一覧を見れる

## 注意事項

* アップロードした画像は data フォルダに保存されています。
* 削除機能はまだ未実装なので、削除する場合はdataフォルダから直接削除してください。
* 一覧表示はファイル名順で表示されるっぽいです。
* ファイル名がそのままlabelとして表示されるので、ファイル名を事前に変えておくことをお勧めします。(名づけ機能実装予定)

## 開発環境

* macbook air M1
* Visual Studio Code
* go: version go1.19.3 darwin/arm64
* gin: github.com/gin-gonic/gin v1.8.1
* ginのインストール
go mod の初期化
```go mod init アプリケーションのフォルダ名```
gin パッケージの取得
```go get -u github.com/gin-gonic/gin```
