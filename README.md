# seqa sequence図作成CLI ツール

## はじめに、これは何なのか

テキストで記載したデータをもとに、テキストでシーケンス図を出力するCLIツール
gnuplot あるいは msggenというツールにインスパイヤされたもの。

元データはmarkdown形式を意識した単純なフォーマットとする。
出力は、あまり手をかけないが、業務に必要十分な出力を可能とする。

goの把握を兼ねて作成する。結果、binary形式にてある程度汎用的に使えることが期待できると思う。

MITライセンスとする（あとでGitHub登録時に設定すること）
非コピーレフト。

# 使い方の概要

```
#sample seq

- a: Aさん
- b: Bさん
- a->b:AからBへ
- b->c
- c->b:戻る方向の書き方
- a<-c:逆方向の矢印も使うよな
```

## 解説

＃にて、Title他、シーケンスの中での文を記載可能とする

以降、＊にて１行のシーケンスを記載する。（あるいは、Nodeの宣言）

各Node（a,b,c）は記載された順に横に展開する。

先頭に記載することで、それぞれのNodeの宣言をすればよいし、説明も不要で順序もきにしないならいきなりシーケンスを記載すればよい。

シーケンスは、-> あるいは逆方向を記載可能とする。　− は表記上複数繰り返しもOK（無視される）。>（あるいは<）の複数はエラー扱いとする。

今後よりバリエーションを増やすケースにおいては、この仕様は変更する。
ただし、>(<)を記載せず、-のみの記載は許容とする。これは---のみの記載。

それぞれのStatementにて、：にて説明を記載可能とする。
Node Statementにおいては、Nodeの説明、Line Statementではその矢印の説明

よって上記の出力イメージはこんな感じか
```
Sample seq

A:Asan               B:Bsan             C
|---- AからBへ ----->|                  |
|                    |----------------->|
|                    |<-戻る方向の書き方-|
|<-- 逆方向の矢印も使うよな -------------|
```

幅は、コメント長を調整し計算するようにしたい。

```
content = <statement>*
statement = "-" <<nodes> ":" <description> | <comment>
comment = "#" <descrption>
nodes = <node> | <node> <arrow> <node>
node = ident
arrow = "<"{0,1} + "-"{1..} + ">"{0,1}
```
