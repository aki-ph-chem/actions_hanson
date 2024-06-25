# GitHub Actions のハンズオンをやる

Software Design 7月号のGitHub Actions実践講座

## 基本

とりあえず、簡単なgoのプロジェクトをビルドする例

.github/workflows/go.yaml
```YAML
name: Go Test

on:
  push:

jobs:

  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: '1.22'

    - name: Build
      run: go build -v .
```

## CIを実現する

このCIではフォーマッタをかけた後で、静的解析、テスト、ビルドを実行する。

CIのトリガーはプルリクエストが作成とする。

```YAML
name: ci

on:
  pull_request:
    branches:
      - main

jobs:

  ci:
    name: ci
    runs-on: ubuntu-latest
    steps:
      - name: checkout repositry
        uses: actions/checkout@v4

      - name: setup environment of Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.22.2'

      - name: format src
        run: test -z $(gofmt -l .)

      - name: static analysis
        run: go vet ./...

      - name: test
        run: go test

      - name: build
        run: go build .
```

## CDを実現する
