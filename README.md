# SSRFの脆弱性があるアプリケーションです

実際に攻撃を試すことができます


## アプリケーション概要

セキュアコーディング啓蒙のために、本アプリケーションはあります。
ローカルに立てたDocker Compose上のコンテナにSSRFの脆弱性を利用した入力をし、サーバー上の情報を取得するシナリオを仮定します。
本アプリケーション及び関連記事は、セキュアコーディングの啓蒙のために作成されたものであり、悪意を持って利用することは禁止します。
自身のローカル環境以外での試用は法律で罰せられる可能性がありますので、ご注意ください。

## 環境構築

`/frontApp/.env`に`VITE_API_URL`を設定してください。


```.env
VITE_API_ENDPOINT=http://localhost:8080/api/v1
```

### フロントエンドのアプリをビルドする

```bash
$ cd frontApp/vul-app-front
$ npm install
$ npm run build
```

## Reference

- https://redfoxsec.com/blog/ssrf/
- https://owasp.org/API-Security/editions/2023/en/0x11-t10/
- https://www.blackhat.com/docs/us-17/thursday/us-17-Tsai-A-New-Era-Of-SSRF-Exploiting-URL-Parser-In-Trending-Programming-Languages.pdf

- https://ja.vitejs.dev/