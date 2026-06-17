# to-do-application-api
Udemy講座「Echo/Go + Reactで始めるモダンWebアプリケーション開発」の内容をベースに、REST API開発、クリーンアーキテクチャ、テスト、CI/CD、デプロイまでを学習することを目的としています。

# 使用している主な技術
### Backend
- Go
- Echo

### Frontend
- React
- TypeScript

### Testing
- Unit Test
- E2E Test
- Playwright

### CI/CD
- GitHub Actions

### AI Tools
- Claude

### Design
- API設計
- データベース設計
- クリーンアーキテクチャ

## 開発環境

### 必要ソフトウェア

- Node.js v24.16.0 以上
- Go
- MySQL

Node.js ダウンロード

https://nodejs.org/ja/download

## セットアップ
### リポジトリのクローン

```bash
git clone <repository-url>
cd to-do-application-api
```

### Goモジュールのインストール

```bash
go mod tidy
```

### データベース構築

```bash
GO_ENV=dev go run migrate/migrate.go
```

---

## アプリケーション起動

### Backend

```bash
go run main.go
```

# 参考資料
https://www.udemy.com/course/echo-go-react-restapi/learn/lecture/37108418#overview
