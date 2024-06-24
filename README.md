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
