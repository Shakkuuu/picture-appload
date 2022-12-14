# picture-appload

## 使い方

1. localサーバを起動  
``` go run main.go ```
2. ブラウザで 127.0.0.1:8080 にアクセス
3. ファイルを選択を押してアップロードしたい画像を選択(一部対応していない拡張子があります。)
4. 付けたい名前を入力
5. 送信を押してアップロード!!
6. 付けた名前(元のファイル名)uploaded!!と表示されるので 戻る から初めの画面に戻る
7. image list から一覧を見れる
8. image list内の削除したい画像の削除を押すと、確認画面を経由して削除ができる
9. 検索の枠に探す画像のファイル名を入力すると、その画像のみ表示される。(例:sample.png)

## 注意事項

* アップロードした画像は data フォルダに保存されています。
* 一度に複数のファイルをアップロードすることはできません。
* 一覧表示はファイル名順で表示されるっぽいです。
* 付けたい名前を入力せずに送信すると、そのままのファイル名で保存されます。
* 何かしらエラーが発生したらエラーページに移動するので、戻るから戻ってください。

## 開発環境

* macbook air M1
* Visual Studio Code
* go: version go1.19.3 darwin/arm64
* 使用パッケージ  
```fmt,log,os,gin```  
* gin: github.com/gin-gonic/gin v1.8.1
* ginのインストール
go mod の初期化  
``` go mod init アプリケーションのフォルダ名 ```  
gin パッケージの取得  
``` go get -u github.com/gin-gonic/gin ```
