## 初回起動

```bash
# まずはgolangが入っていることを確認
go version

# dbを作成
mysql -u root -p
# パスワードの入力 Enter
# db作成
CREATE DATABASE IF NOT EXISTS [データベース名] DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;
# db権限付与
GRANT ALL PRIVILEGES ON [データベース名].* TO '[ユーザー名]'@'localhost';


# モジュールを一括導入
cd /path/to/your-project-directory
go mod tidy

# .env.sampleファイルを元に.envファイルの作成
cp .env.sample .env
vi .env

# echoの実行
go run /path/to/your-project-directory/main.go
# 自身の環境の場合
# go run ~/git/echo.api/main.go
# http://localhost:4207/
```

## 本番起動(myappディレクトリにビルドファイルを)

```bash
go build -o myapp
./myapp
```

## db操作

```bash
# dbのデータ確認
mysql -u root -p
# パスワードの入力 Enter
# db移動
use [データベース名];
# データ確認
select * from [テーブル名];

# テーブル内のデータを削除
truncate table [テーブル名];
# テーブルの構造確認
show full columns from [テーブル名];
```

## apiテスト

```bash
# 管理者関連
# 管理者の登録
curl -X POST http://localhost:4207/auth/admin/register \
-H "Content-Type: application/json" \
-d '{
	"name": "takahshi",
	"email": "tkhsksk0318@example.com",
	"password": "Password123"
}'

# 管理者ログイン
curl -s -X POST http://localhost:4207/auth/admin/login \
-H "Content-Type: application/json" \
-d '{
    "email": "tkhsksk0318@example.com",
    "password": "Password123"
}' | jq

# ユーザーの取得
curl -X GET "http://localhost:4207/authed/admin/users" \
-H "Content-Type: application/json" \
-H "Session-ID: 4f714410-a4df-4688-84c8-ba1fc483d2cc" | jq

# limit指定してユーザーの取得
curl -X GET "http://localhost:4207/authed/admin/users?limit=2" \
-H "Content-Type: application/json" \
-H "Session-ID: 4f714410-a4df-4688-84c8-ba1fc483d2cc" | jq

# 個別ユーザーの取得
curl -X GET http://localhost:4207/authed/admin/users/5 \
-H "Content-Type: application/json" \
-H "Session-ID: [管理者セッション]"

# セッションの取得
curl -X GET http://localhost:4207/authed/admin/users/sessions \
-H "Content-Type: application/json" \
-H "Session-ID: 5a5340c9-edcd-4ba0-8853-22c8548e73d2"

# 商品
# 作成
curl -s -X POST http://localhost:4207/authed/admin/products \
-H "Content-Type: application/json" \
-H "Session-ID: 4f714410-a4df-4688-84c8-ba1fc483d2cc" \
-d '{
    "name": "ジャケット赤",
    "price": 12800,
    "content": "",
    "status": "active",
    "category_id": 3
}' | jq

# 取得
curl -X GET http://localhost:4207/authed/admin/products/3 \
-H "Content-Type: application/json" \
-H "Session-ID: 1f8a6d75-a555-4fda-8ce6-2548e9f328b3" | jq

# 一覧取得
curl -s -X GET http://localhost:4207/authed/admin/products \
-H "Content-Type: application/json" \
-H "Session-ID: 4f714410-a4df-4688-84c8-ba1fc483d2cc" | jq

# 更新
curl -X PUT http://localhost:4207/authed/admin/products/6 \
-H "Content-Type: application/json" \
-H "Session-ID: 1f8a6d75-a555-4fda-8ce6-2548e9f328b3" \
-d '{
    "name": "ジャケット黒",
    "price": 12800,
    "content": "",
    "status": "suspended",
    "category_id": 3
}' | jq

# カテゴリー
# 作成
curl -X POST http://localhost:4207/authed/admin/categories \
-H "Content-Type: application/json" \
-H "Session-ID: 4f714410-a4df-4688-84c8-ba1fc483d2cc" \
-d '{
    "name": "ボトムス",
    "content": "",
    "status": "active",
    "parent_id": 2
}' | jq

# ツリー取得
curl -X GET http://localhost:4207/authed/admin/categories/tree \
-H "Content-Type: application/json" \
-H "Session-ID: 4f714410-a4df-4688-84c8-ba1fc483d2cc" | jq

# 全取得
curl -X GET http://localhost:4207/authed/admin/categories \
-H "Content-Type: application/json" \
-H "Session-ID: 4f714410-a4df-4688-84c8-ba1fc483d2cc" | jq

# id指定取得
curl -X GET http://localhost:4207/authed/admin/categories/4 \
-H "Content-Type: application/json" \
-H "Session-ID: 4f714410-a4df-4688-84c8-ba1fc483d2cc" | jq

# 更新
curl -X PUT http://localhost:4207/authed/admin/categories/4 \
-H "Content-Type: application/json" \
-H "Session-ID: 4f714410-a4df-4688-84c8-ba1fc483d2cc" \
-d '{
    "name": "bottoms",
    "content": "",
    "status": "active",
    "parent_id": 2
}' | jq

# ユーザー関連
# ユーザーの登録
curl -X POST http://localhost:4207/auth/user/register \
-H "Content-Type: application/json" \
-d '{
	"name": "テスト太郎",
	"email": "user+003@example.com",
	"password": "Password123"
}'

# ユーザーログイン
curl -X POST http://localhost:4207/auth/user/login \
-H "Content-Type: application/json" \
-d '{
    "email": "user@example.com",
    "password": "Password123"
}'

# 商品
# 取得
curl -X GET http://localhost:4207/authed/user/products/7 \
-H "Content-Type: application/json" \
-H "Session-ID: 6c83928f-26fc-483c-a19a-7b5d2d96feee" | jq

# 一覧取得
curl -X GET http://localhost:4207/authed/user/products \
-H "Content-Type: application/json" \
-H "Session-ID: 6c83928f-26fc-483c-a19a-7b5d2d96feee" | jq

# ポスト
# 登録
curl -X POST http://localhost:4207/authed/user/posts \
-H "Content-Type: application/json" \
-H "Session-ID: 6ee62750-6276-44f7-b3ee-14f66632027a" \
-d '{
	"title": "テスト投稿",
	"content": "これはテスト投稿の内容です"
}'

# 更新
curl -X PUT http://localhost:4207/authed/user/posts/1 \
-H "Content-Type: application/json" \
-H "Session-ID: [ユーザーセッション]" \
-d '{
	"title": "テスト投稿edit",
	"content": "これはテスト投稿の内容ですedit"
}'

# 取得
curl -X GET http://localhost:4207/authed/user/posts \
-H "Content-Type: application/json" \
-H "Session-ID: 6ee62750-6276-44f7-b3ee-14f66632027a"

# 個別取得
curl -X GET http://localhost:4207/authed/user/posts/1 \
-H "Content-Type: application/json" \
-H "Session-ID: a55fab5b-7e51-4f87-8a94-a87f9d5f671b"

# プロフィール
# 取得
curl -X GET http://localhost:4207/authed/user/profiles \
-H "Content-Type: application/json" \
-H "Session-ID: c4278f0b-f9e0-4c21-9415-7ccf665820f3"

# 更新
curl -X PUT http://localhost:4207/authed/user/profiles \
-H "Content-Type: application/json" \
-H "Session-ID: c4278f0b-f9e0-4c21-9415-7ccf665820f3" \
-d '{
    "name": "テスト投稿edit"
}'

# データベースのリセット
curl -X POST http://localhost:4207/delete \
-H "Content-Type: application/json" \
-d '{
    "password": "パスワード"
}'
```