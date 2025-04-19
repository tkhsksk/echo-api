## 初回起動

# まずはgolangが入っていることを確認
go version

# dbを作成
mysql -u root -p
# パスワードの入力 Enter
CREATE DATABASE IF NOT EXISTS [データベース名] DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;


# モジュールを一括導入
cd /your/project/directory
go mod tidy

# .env.sampleファイルを元に.envファイルの作成
cp .env.sample .envZ

# echoの実行
go run main.go

## db操作

# dbのデータ確認
mysql -u root -p
# パスワードの入力 Enter
show databases;
show tables from [データベース名];
use [データベース名];
select * from [テーブル名];

# テーブル内のデータを削除
truncate table [テーブル名];

## dbテスト
# テストデータの登録
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
	"email": "testuser@example.com",
	"password": "password123"
}'

# ログイン確認
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "email": "testuser@example.com",
    "password": "password123"
}'

# ポストの登録
curl -X POST http://localhost:8080/posts/ \
-H "Content-Type: application/json" \
-H "Session-ID: da3105ad-1fd7-46d4-92ec-ca8a99d2ab6a" \
-d '{
	"title": "テスト投稿",
	"content": "これはテスト投稿の内容です"
}'