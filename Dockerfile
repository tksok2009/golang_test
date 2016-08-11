FROM golang:1.3.1-onbuild

# 対象goアプリ/ツールがコンテナ上では app という名前でインストールされるので、その名前を指定して実行
## 実行時引数が無いケース
# CMD ["app"]
## 実行時引数があるケース
CMD ["app", "192.168.56.24/0"]