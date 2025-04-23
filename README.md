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
# 管理者の登録
curl -X POST http://localhost:4207/admin/register \
-H "Content-Type: application/json" \
-d '{
	"email": "user@example.com",
	"password": "password123"
}'

# ユーザーの登録
curl -X POST http://localhost:4207/user/register \
-H "Content-Type: application/json" \
-d '{
	"email": "user@example.com",
	"password": "password123"
}'

# ユーザーの取得
curl -X GET http://localhost:4207/users \
-H "Content-Type: application/json" \
-H "Session-ID: [ログイン確認時に返ってきた管理者セッションID]"

# セッションの取得
curl -X GET http://localhost:4207/users/sessions \
-H "Content-Type: application/json" \
-H "Session-ID: [ログイン確認時に返ってきた管理者セッションID]"

# ログイン
curl -X POST http://localhost:4207/user/login \
-H "Content-Type: application/json" \
-d '{
    "email": "user@example.com",
    "password": "password123"
}'

# ポストの登録
curl -X POST http://localhost:4207/posts \
-H "Content-Type: application/json" \
-H "Session-ID: [ログイン確認時に返ってきたユーザーセッションID]" \
-d '{
	"title": "テスト投稿",
	"content": "これはテスト投稿の内容です"
}'
```