# blog-public

# 目的
* 本格的ではないが、どのようにブログが作られるかを通してサーバサイドについての知識と、Reactについての知識をつける

# 開発方法

```sh
cd scripts/docker
docker-compose up --build
```

とすればサーバが立ち上がります。

その後は

```
docker-compose up
```

を使います。 buildオプションはビルドで時間がかかるので、依存ライブラリを行ったりDockerfileの更新があったときだけでよいです。

サーバを落とすときは

```sh
docker-compose down
```

としてください。

これでMySQLサーバ(データベース)とバックエンドサーバ(golang)が立ち上がります。

フロントエンドはローカルから行うことにします。(node_modules周りでWSLが心配なため)

```sh
cd frontend
npm install
npm run dev
```

これでフロントエンドのサーバが立ち上がります。 `http://localhost:8080` を見てみましょう。何か表示されていれば成功です。

# 開発の進め方
1. issueを見る
2. できそうなものを探す、なかったらdiscordで相談する
3. やってみる、途中でもdiscordで実況ログを書いてみる

手を動かしてみましょう！

# docker, docker-composeの入れ方
WSL2では https://zenn.dev/sprout2000/articles/95b125e3359694 などを参考にしてください。

参考に僕のインストールログを残します
```sh
sudo apt-get update
sudo apt-get install     ca-certificates     curl     gnupg     lsb-release
echo   "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
$(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
sudo docker run hello-world
sudo curl -L "https://github.com/docker/compose/releases/download/v2.0.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version
```
docker-composeはv2を使っています。
