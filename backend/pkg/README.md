### pkg文件夹

这里存放一些不受到go modules管理的公共库，如deer-common等

#### 安装 deer-common
先从 github.com 克隆下来，然后做软链接
```bash
$ cd /Users/lancelrq/github
$ git clone git@github.com:LanceLRQ/deer-common.git
$ ln -s /Users/lancelrq/github/deer-common /Users/lancelrq/github/wejudge-polygon/pkg
```