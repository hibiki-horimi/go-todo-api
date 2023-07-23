# 基本的な Todo 用の CRUD RESTful API

これは Todo タスクを管理するための基本的な CRUD（Create、Read、Update、Delete）RESTful API です。この API を使用すると、新しいタスクの作成、タスクの詳細の取得、タスク情報の更新、タスクの削除など、よく使われる操作を行うことができます。

## 目次

- [機能](#機能)
- [使用されている技術](#使用されている技術)
- [インストール](#インストール)
- [使用方法](#使用方法)
- [API エンドポイント](#APIエンドポイント)

## 機能

- 新しい Todo タスクの作成
- すべての Todo タスクのリストの取得
- 特定の Todo タスクの詳細の取得（ID で指定）
- Todo タスクの名前とステータスの更新（ID で指定）
- Todo タスクの削除（ID で指定）

## 使用されている技術

この API の作成には以下の技術が使用されています：

- Go: サーバーサイドの開発に使用されるプログラミング言語
- Echo: Go 用の高速でミニマルなウェブフレームワーク
- PostgreSQL: Todo タスクのデータを格納する強力なオープンソースのリレーショナルデータベース
- Docker: アプリケーションとその依存関係をパッケージ化するために使用されるコンテナ化プラットフォーム

## インストール

1. システムに Docker がインストールされていることを確認してください。

2. このリポジトリをローカルマシンにクローンします。

```bash
git clone https://github.com/hibiki-horimi/go-todo-api

```

3. プロジェクトディレクトリに移動します。

```bash
cd go-todo-api
```

4. Docker コンテナを実行します。

```bash
docker compose up
```

5. 別のシェルを開き、データベースマイグレーションを実行して必要なテーブルをセットアップします。

```bash
go run main.go migrate
```

## 使用方法

curl や Postman などの REST クライアントツールを使用して API とやり取りすることができます。以下は利用可能なエンドポイントとその機能の概要です。

## API Endpoints

`GET /api/todos`

すべての Todo タスクのリストを取得します。

```bash
curl localhost:8080/api/todos
```

`POST /api/todos`

新しい Todo タスクを作成します。リクエストボディに以下のフィールドを持つ JSON ペイロードを必要とします：

task（必須）：Todo タスクの名前

```bash
curl -X POST -H "Content-Type: application/json" -d '{"task":"買い物に行く"}' localhost:8080/api/todos
```

`GET /api/todos/:id`

特定の Todo タスクの詳細を ID で指定して取得します。

```bash
curl localhost:8080/api/todos/12345678-90ab-cdef-ghij-klmnopqrstuv
```

`PUT /api/todos/:id`

Todo タスクの名前とステータスを ID で指定して更新します。更新するフィールドを含む JSON ペイロードが必要です。

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"task":"航空券を買う", "done": true}' localhost:8080/api/tasks/cae0bb14-d854-4387-b27f-28b1df880d02
```

`DELETE /api/todos/:id`

ID で指定した Todo タスクを削除します。

```bash
curl -X DELETE localhost:8080/api/todos/12345678-90ab-cdef-ghij-klmnopqrstuv
```
