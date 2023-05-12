# go-testify-assert-test

## 知りたいこと

- testify の Equal は何を比べているか
  - [reflect.DeepEqual](https://pkg.go.dev/reflect#DeepEqual)
- pb の何が影響してるか
  - 内部で `MessageState` という unexport のフィールドを持っており、その値が参照されることで変わってしまう
    - `reflect.DeepEqual` では unexport の struct のフィールドも比較対象にするため
    - unexport は json 出力には関係してない
  - 一般に、どういう条件の時影響しうるか
- 要素出力以外で影響を与えうるコード
