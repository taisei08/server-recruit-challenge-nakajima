# サーバーエンジニア向け 2026新卒採用事前課題

あなたは歌手とアルバムを管理するAPIの機能開発にたずさわることになりました。

次の課題に順に取り組んでください。

できない課題があっても構いません。

面接中に課題に関して質問をしますので、分かる範囲で説明してください。

## 課題1
プログラムのコードを読み、中身を把握しましょう。

## 課題2
Docker と Go をインストールし(各自で調べてください)、歌手を管理するAPIの動作を確認しましょう。

```
# (ターミナルを開いて)
# Docker コンテナを起動する
docker compose up -d
# HTTP サーバーを起動する
go run main.go
```

```
# (別のターミナルを開いて)
# 歌手の一覧を取得する
curl http://localhost:8888/singers

# 指定したIDの歌手を取得する
curl http://localhost:8888/singers/1

# 歌手を追加する
curl -X POST -d '{"id":10,"name":"John"}' http://localhost:8888/singers

# 歌手を削除する
curl -X DELETE http://localhost:8888/singers/10
```

## 課題3
アルバムを管理するAPIを新規作成しましょう。

### 3-1
アルバムの一覧を取得するAPI
```
curl http://localhost:8888/albums

# このようなレスポンスを期待しています
[{"id":1,"title":"Alice's 1st Album","singer_id":1},{"id":2,"title":"Alice's 2nd Album","singer_id":1},{"id":3,"title":"Bella's 1st Album","singer_id":2}]
```

### 3-2
指定したIDのアルバムを取得するAPI
```
curl http://localhost:8888/albums/1

# このようなレスポンスを期待しています
{"id":1,"title":"Alice's 1st Album","singer_id":1}
```

### 3-3
アルバムを追加するAPI
```
curl -X POST -d '{"id":10,"title":"Chris 1st","singer_id":3}' http://localhost:8888/albums

# このようなレスポンスを期待しています
{"id":10,"title":"Chris 1st","singer_id":3}

# そして、アルバムを取得するAPIでは、追加したものが存在するように
curl http://localhost:8888/albums/10
```

### 3-4
アルバムを削除するAPI
```
curl -X DELETE http://localhost:8888/albums/10
```

## 課題4
アルバムを取得するAPIでは、歌手の情報も付加するように改修しましょう。

### 4-1
指定したIDのアルバムを取得するAPI
```
curl http://localhost:8888/albums/1

# このようなレスポンスを期待しています
{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}}
```

### 4-2
アルバムの一覧を取得するAPI
```
curl http://localhost:8888/albums

# このようなレスポンスを期待しています
[{"id":1,"title":"Alice's 1st Album","singer":{"id":1,"name":"Alice"}},{"id":2,"title":"Alice's 2nd Album","singer":{"id":1,"name":"Alice"}},{"id":3,"title":"Bella's 1st Album","singer":{"id":2,"name":"Bella"}}]
```

## 課題5
歌手とそのアルバムを管理するという点で、現状の実装の改善点を検討し思いつく限り書き出してください。

実装をする必要はありません。

### 機能面
1. GetALL関数が全てのデータを取得してしまうためLIMIT・OFFSETで区分けで取得する
2. 管理者機能・ユーザー機能をつけて`POST`や`DELETE`ができる人を制限する
3. 曲名や歌手名を検索できる機能の追加
4. 歌手・アルバムデータの更新処理も可能にする
5. 一つのアルバムに対して複数人の歌手を登録できるようにする
6. 5のために`albums`と`singers`の中間テーブルを作成し、作曲者が複数いても1対多リレーションに落とし込める設計に変更する

### 開発面
1. `infra/mysqldb`と`service`と`controller`それぞれへのテストの追加
2. アプリケーションもdockerで動かす
3. 基本的なCRUDのルーティングの処理を共通化したい
    ``` go
      // この辺りの処理をエンドポイント毎に共通化したい
      mux.HandleFunc("GET /singers", singerController.GetSingerListHandler)
      mux.HandleFunc("GET /singers/{id}", singerController.GetSingerDetailHandler)
      mux.HandleFunc("POST /singers", singerController.PostSingerHandler)
      mux.HandleFunc("DELETE /singers/{id}", singerController.DeleteSingerHandler)

      mux.HandleFunc("GET /albums", albumController.GetAlbumListHandler)
      mux.HandleFunc("GET /albums/{id}", albumController.GetAlbumDetailHandler)
      mux.HandleFunc("POST /albums", albumController.PostAlbumHandler)
      mux.HandleFunc("DELETE /albums/{id}", albumController.DeleteAlbumHandler)
    ```
4. AUTO_INCREMENTを使用しているためINSERT時のID指定は無くしたい（課題を解く上では今のままの方がいいが）
5. AlbumのモデルからSingerIDフィールドを消したい
6. 5のためにリクエスト形式をモデルに沿ったものにしたい
    ``` go
    type Album struct {
      ID       AlbumID  `json:"id,omitempty"`
      Title    string   `json:"title,omitempty"`
      SingerID SingerID `json:"singer_id,omitempty"` // リクエストの形式を変えてこのフィールドを削除したい
      Singer   Singer   `json:"singer,omitempty"`
    }
    ```
    ``` plaintext
    以下の形のリクエストにしたい
    curl -X POST -d '{"id":1,"title":"Alice's 1st Album","singer":{"name":"Alice"}}'
    ```
7. POST時のバリデーションエラーのステータスコードが`500`なので`400`にしたい
8. GET時に存在しないIDを指定したときエラーのステータスコードが`500`なので`404`にしたい
9. OpenAPI仕様書の導入
