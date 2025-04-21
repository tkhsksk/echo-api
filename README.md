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
go run main.go
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
# ユーザーの登録
curl -X POST http://localhost:4207/register \
-H "Content-Type: application/json" \
-d '{
	"email": "user@example.com",
	"password": "password123"
}'

# ログイン
curl -X POST http://localhost:4207/login \
-H "Content-Type: application/json" \
-d '{
    "email": "user@example.com",
    "password": "password123"
}'

# ポストの登録
curl -X POST http://localhost:4207/posts/ \
-H "Content-Type: application/json" \
-H "Session-ID: [ログイン確認時に返ってきたセッションID]" \
-d '{
	"title": "テスト投稿",
	"content": "これはテスト投稿の内容です"
}'
```
# ユーザーの登録
curl -X POST http://localhost:4207/register \
-H "Content-Type: application/json" \
-d '{
	"email": "user@example.com",
	"password": "password123"
}'

# ログイン
curl -X POST http://localhost:4207/login \
-H "Content-Type: application/json" \
-d '{
    "email": "user@example.com",
    "password": "password123"
}'

# ポストの登録
curl -X POST http://localhost:4207/posts/ \
-H "Content-Type: application/json" \
-H "Session-ID: [ログイン確認時に返ってきたセッションID]" \
-d '{
	"title": "テスト投稿",
	"content": "これはテスト投稿の内容です"
}'