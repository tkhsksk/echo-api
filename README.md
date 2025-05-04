## 初回起動

```bash
# まずはgolangが入っていることを確認
go version

# dbを作成
mysql -u root -p
# パスワードの入力 Enter
CREATE DATABASE IF NOT EXISTS [データベース名] DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;
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
show databases;
show tables from [データベース名];
use [データベース名];
select * from [テーブル名];

# テーブル内のデータを削除
truncate table [テーブル名];
```

## apiテスト

```bash
# 管理者関連
# 管理者の登録
curl -X POST http://localhost:4207/auth/admin/register \
-H "Content-Type: application/json" \
-d '{
	"name": "takahshi",
	"email": "user+001@example.com",
	"password": "Password123"
}'

# 管理者ログイン
curl -X POST http://localhost:4207/auth/admin/login \
-H "Content-Type: application/json" \
-d '{
    "email": "user@example.com",
    "password": "Password123"
}'

# ユーザーの取得
curl -X GET http://localhost:4207/users \
-H "Content-Type: application/json" \
-H "Session-ID: 538e54d0-0e21-4222-9d12-55845e573f2f"

# 個別ユーザーの取得
curl -X GET http://localhost:4207/users/5 \
-H "Content-Type: application/json" \
-H "Session-ID: [管理者セッション]"

# セッションの取得
curl -X GET http://localhost:4207/users/sessions \
-H "Content-Type: application/json" \
-H "Session-ID: [管理者セッション]"

# ユーザー関連
# ユーザーの登録
curl -X POST http://localhost:4207/auth/user/register \
-H "Content-Type: application/json" \
-d '{
	"name": "テスト太郎",
	"email": "user+001@example.com",
	"password": "Password123"
}'

# ユーザーログイン
curl -X POST http://localhost:4207/auth/user/login \
-H "Content-Type: application/json" \
-d '{
    "email": "user@example.com",
    "password": "Password123"
}'

# ポストの登録
curl -X POST http://localhost:4207/auth/user/posts \
-H "Content-Type: application/json" \
-H "Session-ID: 6ee62750-6276-44f7-b3ee-14f66632027a" \
-d '{
	"title": "テスト投稿",
	"content": "これはテスト投稿の内容です"
}'

# ポストの更新
curl -X PUT http://localhost:4207/auth/user/posts/1 \
-H "Content-Type: application/json" \
-H "Session-ID: [ユーザーセッション]" \
-d '{
	"title": "テスト投稿edit",
	"content": "これはテスト投稿の内容ですedit"
}'

# ポストの取得
curl -X GET http://localhost:4207/auth/user/posts \
-H "Content-Type: application/json" \
-H "Session-ID: 6ee62750-6276-44f7-b3ee-14f66632027a"

# 個別ポストの取得
curl -X GET http://localhost:4207/auth/user/posts/1 \
-H "Content-Type: application/json" \
-H "Session-ID: a55fab5b-7e51-4f87-8a94-a87f9d5f671b"

# プロフィール取得
curl -X GET http://localhost:4207/auth/user/profiles \
-H "Content-Type: application/json" \
-H "Session-ID: 6ee62750-6276-44f7-b3ee-14f66632027a"

# データベースのリセット
curl -X POST http://localhost:4207/delete \
-H "Content-Type: application/json" \
-d '{
    "password": "パスワード"
}'
```