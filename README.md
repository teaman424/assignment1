### Mac環境建置
---------------
Mac環境，可以直接使用Homebrew安裝
```
brew install go
```
設定環境變數
```
export PATH=$PATH:/usr/local/opt/go/libexec/bin
export GOROOT=$HOME/go1.X
```

檢查一下是否安裝成功
```
go version
```
### Ubuntu環境建置
---------------
更新套件包
```
sudo apt-get update  
sudo apt-get -y upgrade  
```

download go binary
```
wget https://dl.google.com/go/go1.17.7.linux-amd64.tar.gz 
```

設定環境變數
```
sudo tar -xvf go1.17.7.linux-amd64.tar.gz
sudo mv go /usr/local 
export GOROOT=/usr/local/go 
export GOPATH=$HOME
export PATH=$PATH:/usr/local/opt/go/libexec/bin
```

檢查一下是否安裝成功
```
go version
```

### 執行程式
---------------

```
cd assignment1
go run main.go

```
