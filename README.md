# yotta-adapter

## 项目依赖库

- [openwallet](https://github.com/blocktree/openwallet.git)
- [eosio-adapter](https://github.com/blocktree/eosio-adapter.git)

注意：
要在最终引用此适配包的项目中，replace eos-go引用，即在go.mod中：
`replace github.com/eoscanada/eos-go => github.com/blocktree/eos-go v0.8.13-blocktree`


## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf目录，新建YTA.ini文件，编辑如下内容：

```ini

#wallet api url
ServerAPI = "http://localhost:8888"
# Cache data file directory, default = "", current directory: ./data
dataDir = ""

```
