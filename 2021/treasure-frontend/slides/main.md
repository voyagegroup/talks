----
marp: true
title: Treasure JavaScript 2021
description: Treasure 2021 Frontend
size: 16:9
theme: gaia
paginate: true
footer: Treasure 2021 Frontend
----

<!-- EmojiCheatSheet: https://www.webfx.com/tools/emoji-cheat-sheet/ -->

<!-- _paginate: false -->

# Treasure 2021 Frontend

----

# DO IT

:star: Join `#frontend`

----

# この講義のゴール

- 簡単なWebアプリケーションをブラウザ上で動かせるようになる
- ちょっとReact書けるようになる
- TypeScriptを使った型のある開発に慣れる

---

# 講義の進め方

- 私: 色んなコト説明していきます
  - みんな: 分からなかったら質問する
- みんな: 困ったら助けてって言う
  - TAと私: いい感じに一緒に解決していく
- これに限らずいい感じにやっていく
- 適度に休憩

Enjoy Treasure!

<!-- 当日メモ: この辺りでTAに相談しやすいようにTAの紹介入れておこう -->

---

# アイスブレイク

**JavaScript**, **TypeScript** 好きな人 :raised_hand:

---

# アイスブレイク

Treasureで行われる講義の中でフロントが一番得意だと思う人 :raised_hand:

---

<!-- _class: lead -->

# 事前課題について

---

# 事前課題どうでした？

適当に振っていきます

---

# 事前課題って何のため？

- 知識のベースラインを揃える
- 講義で扱う技術にサラッと触れてもらう
- 分からないことの調べ方、質問の仕方を学んでもらう
- ベースアプリの全体像を把握してもらう
- ~~講義のネタ作りのため（~~

---

ということで…

- 簡単なReactのコンポーネントを作ることが出来る
- TypeScriptの型をちょろっと読むことが出来る

を前提として進めていきます

---

…流石にキツい？

---

# (再掲) 講義の進め方

- 私: 色んなコト説明していきます
  - みんな: 分からなかったら質問する
- みんな: 困ったら助けてって言う
  - TAと私: いい感じに一緒に解決していく
- これに限らずいい感じにやっていく
- 適度に休憩

Enjoy Treasure!

---

# レビューについて

- 最低限気にして欲しいところは積極的にレビューしていった
  - Lintとかで防げる箇所も多かったので実開発では使ってみよう
    - ESLint
    - Prettier
- ありがちな内容は途中から指摘をやめた
  - ぶっちゃけみんな他の人のPR見たっしょ？
  - そこで指摘されてることは直すよね？
    - という期待がありました

<!-- 当日メモ: 別にコピペは悪じゃなくて、劣化コピーと理解しないコピーが悪 -->

---

## 脱線: 私のOSSへのPRの出し方

- 直近に出されているPRやレビューが行われたPRをいくつか見る
- 無事マージされたPRをいくつか見る
- コード品質を守るために、コードオーナーやレビュワーがどこを重点的に見ているのかを確認

これらを踏まえて

- 他のPRで指摘されているようなことはコードに混入させない
- 悩んだ場所や重点的に見て欲しいところはあらかじめコメント

---

# ベースアプリはどんなアプリ

<!-- 当日メモ: 説明してもらおう -->

誰かシュッと答えてください

---

# ベースアプリはどんなアプリ

- ユーザはログイン出来る
- ユーザは記事を{見れる, 書ける}

Simple

---

<!-- _footer: '' -->

# ディレクトリ構成

```
├── App.tsx
├── Contract.ts 
├── api
├── components
├── contexts
├── main.tsx
├── pages
├── services
├── style.tsx
├── util
```

Next.jsのディレクトリ構成に似てる（？） https://github.com/vercel/next.js/tree/master/examples/with-typescript

---

## Tips: 他のディレクトリ構成のパターンは？

例えばReactのページにあるパターン

https://ja.reactjs.org/docs/faq-structure.html

アプリケーションの規模によってディレクトリが増えたりもします
お好きなのをどうぞ:smile:

---

# 事前課題について

---

# おさらい: 事前課題1について

Articleを投稿する課題

- DevToolsの使い方を知る
- Authorizationヘッダーを使った認証

辺りを学んでもらう内容

---

# おさらい: 事前課題2について

Articleを取得する課題

- Reactでコンポーネントを生やせる

が出来るようになる内容

---

<!-- _footer: '' -->

## 事前課題でレビューした点

### APIのResponseのMapping

<!-- 当日メモ: https://github.com/VG-Tech-Dojo/treasure-app-2021/pull/97 -->

```ts
// ArticleClientGraphQL.ts
const json = await res.json();
return await json.data.articleById; // Article型が返ってくるはず
```

returnされる値は

```ts
{ id: number; title: string; body: string;
  user_id: number; // !? userId: number; では！？
}
```

使う側にどういう影響が出そう？

---

## 事前課題でレビューした点

### APIのResponseのMapping

Mappingしてないでそのまま使うと…

- 型が嘘をついてしまうことに繋がるかも
- API側で名前変えたら、影響範囲が膨大に

<!-- 当日メモ:
  Mappingするか、Queryで userId: user_id みたいにRenameとか
  Viewでundefinedが出始める例
-->

--- 

## 事前課題でレビューした点

### fetchのErrorハンドリング

<!-- 当日メモ: https://github.com/VG-Tech-Dojo/treasure-app-2021/pull/97 -->

```ts
fetch({ ... }).then(/* ここでErrorHandling */)
```

:clap:

--- 

## 事前課題でレビューした点

### fetchのErrorハンドリング

https://developer.mozilla.org/ja/docs/Web/API/Fetch_API/Using_Fetch

> fetch() から返される Promise は レスポンスが HTTP 404 や 500 を返して HTTP エラーステータスの場合でも拒否されません
>（中略） 正常に解決し、拒否されるのはネットワークのエラーや、何かがリクエストの完了を妨げた場合のみです

--- 

## 事前課題でレビューした点

### fetchのErrorハンドリング


```ts
fetch({ ... }).then(/* ここでErrorHandling */)
```

例えばここでStatusCodeやok statusを見て

- 認証エラー
- リクエストが正常か

などをハンドリングできたりするとGood

--- 

## 事前課題でレビューした点

### CustomError

<!-- 当日メモ: https://github.com/VG-Tech-Dojo/treasure-app-2021/pull/110/files#diff-82125ddf744c97aab032e6305a8f45a685da73537e511c82471b4c1b3cb0964fR58 -->

```ts
class MyAwesomeError extends Error
```

Errorに型が付いていると意図してthrowしていることが伝わりやすい
エラーの種類で処理を分けたいとかの時に思い出してみよう

---

## 事前課題でレビューした点

### == vs ===

https://developer.mozilla.org/ja/docs/Web/JavaScript/Equality_comparisons_and_sameness

顔みたいでカワイイですね :cat:

`null` も `undefined` も別個で扱わないとかであれば `==` も許容…？
意味を持たせていたりオプショナルチェイニングを使った場合に痛い目を見るかも

---

### ここで == vs === Quiz

https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Operators/Optional_chaining

```ts
type Obj = { value: { innerValue: number | null } | null }

const nullObj: Obj = { value : null }

if (nullObj.value?.innerValue === null) { console.log(1); }
if (nullObj.value?.innerValue == null) { console.log(2); }
if (nullObj.value?.innerValue === undefined) { console.log(3); }
if (nullObj.value?.innerValue == undefined) { console.log(4); }
```

コンソールには何が表示されるでしょう

<!-- 2, 3, 4 -->

---

## 事前課題でレビューした点

### 暗黙の型変換

<!-- 当日メモ: https://github.com/VG-Tech-Dojo/treasure-app-2021/pull/137 -->

```ts
const numericStr = "100";
const implicitValue = +numericStr;
```

`implicitValue` の型は？

---

### すごいぞJavaScript

https://twitter.com/cdesio/status/1013166206877163520?s=20

```ts
+'1' // 1
1 + '1' // '11'
'1' == true // true
if (1 == '1') { console.log(1) } // 1
```

型変換用の関数（`parseInt` とか）や `===` を使おう

```ts
// これ好き
new Date(2021, 8, 10) // Fri Sep 10 2021 00:00:00 GMT+0900 (日本標準時)
```

---

<!-- _class: lead -->

# 演習

---


## Update出来るように

---

## 何のためのFormValidation?

色々意見出してもらおう

---

## Tips: QueryParameter………

https://github.com/VG-Tech-Dojo/treasure-app-2021/pull/110

```ts
parseInt(val /* ここで数値以外の文字列だったらNaN */)
```

とかのやつ
変な値をAPIに投げないためにも、実際のViewに謎の値を入れてやらないためにもやるといい

例えばRedirectさせちゃうとか

---

その上での型変換

---

## そう言えばHooksどうでした？

https://github.com/VG-Tech-Dojo/treasure-app-2021/pull/110

早期Returnなー

---

## Hooksをいい感じに見たい

React拡張入れてみるとか

---

## そう言えば、API作ってある？

Interfaceをアレしてアレする方法もある


