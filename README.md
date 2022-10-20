# chit chat

《Go Web 编程》里面的论坛示例的 Copy。

## Init

```bash
$ go mod init
go: creating new go.mod: module github.com/chengchaos/chit-chat

$ git init
提示：使用 'master' 作为初始分支的名称。这个默认分支名称可能会更改。要在新仓库中
提示：配置使用初始分支名，并消除这条警告，请执行：
提示：
提示：	git config --global init.defaultBranch <名称>
提示：
提示：除了 'master' 之外，通常选定的名字有 'main'、'trunk' 和 'development'。
提示：可以通过以下命令重命名刚创建的分支：
提示：
提示：	git branch -m <name>
已初始化空的 Git 仓库于 /Users/chengchao/go-path/src/github.com/chengchaos/chit-chat/.git/
$ git status
位于分支 master

尚无提交

未跟踪的文件:
  （使用 "git add <文件>..." 以包含要提交的内容）
	.idea/
	README.md
	chitchat.log
	config.json
	go.mod
	main.go
	public/
	utils.go

提交为空，但是存在尚未跟踪的文件（使用 "git add" 建立跟踪）
$ git push
fatal: 没有配置推送目标。
或通过命令行指定 URL，或用下面命令配置一个远程仓库

    git remote add <名称> <地址>

然后使用该远程仓库名执行推送

    git push <名称>

$ git remote add git@github.com:chengchaos/chit-chat.git
用法：git remote add [<选项>] <名称> <地址>

    -f, --fetch           抓取远程的分支
    --tags                抓取时导入所有的标签和关联对象
                          或不抓取任何标签（--no-tags）
    -t, --track <分支>    跟踪的分支
    -m, --master <分支>   主线分支
    --mirror[=(push|fetch)]
                          把远程设置为用以推送或抓取的镜像

$ git remote add origin git@github.com:chengchaos/chit-chat.git
$ git status
位于分支 master

尚无提交

未跟踪的文件:
  （使用 "git add <文件>..." 以包含要提交的内容）
	.idea/
	README.md
	chitchat.log
	config.json
	go.mod
	main.go
	public/
	utils.go

提交为空，但是存在尚未跟踪的文件（使用 "git add" 建立跟踪）
$ touch .gitignore
$ git status
位于分支 master

尚无提交

未跟踪的文件:
  （使用 "git add <文件>..." 以包含要提交的内容）
	.gitignore
	README.md
	config.json
	go.mod
	main.go
	public/
	utils.go

提交为空，但是存在尚未跟踪的文件（使用 "git add" 建立跟踪）
$ git add .
$ git commit -m 'init '
[master（根提交） 5c9ab58] init
 7 files changed, 139 insertions(+)
 create mode 100644 .gitignore
 create mode 100644 README.md
 create mode 100644 config.json
 create mode 100644 go.mod
 create mode 100644 main.go
 create mode 100644 public/a.html
 create mode 100644 utils.go
$ git status
位于分支 master
无文件要提交，干净的工作区
$ git pull
当前分支没有跟踪信息。
请指定您要合并哪一个分支。
详见 git-pull(1)。

    git pull <远程> <分支>

如果您想要为此分支创建跟踪信息，您可以执行：

    git branch --set-upstream-to=origin/<分支> master

$ git branch --set-upstream-to=origin/master master
fatal: 请求的上游分支 'origin/master' 不存在
提示：
提示：如果您正计划基于远程一个现存的上游分支开始您的工作，
提示：您可能需要执行 "git fetch" 来获取分支。
提示：
提示：如果您正计划推送一个能与对应远程分支建立跟踪的新的本地分支，
提示：您可能需要使用 "git push -u" 推送分支并配置和上游的关联。
提示：Disable this message with "git config advice.setUpstreamFailure false"

$ git branch --set-upstream-to=origin/main master
分支 'master' 设置为跟踪 'origin/main'。
$ git pull
fatal: 拒绝合并无关的历史
$ git fetch
$ git status
位于分支 master
您的分支和 'origin/main' 出现了偏离，
并且分别有 1 和 1 处不同的提交。
  （使用 "git pull" 来合并远程分支）

无文件要提交，干净的工作区
$ git pull
fatal: 拒绝合并无关的历史

请确认您有正确的访问权限并且仓库存在。
$ git branch -a
* master
  remotes/origin/main
  
$ git pull --allow-unrelated-histories
Merge made by the 'ort' strategy.
 LICENSE | 25 +++++++++++++++++++++++++
 1 file changed, 25 insertions(+)
 create mode 100644 LICENSE

$ git status
位于分支 master
您的分支领先 'origin/main' 共 3 个提交。
  （使用 "git push" 来发布您的本地提交）

无文件要提交，干净的工作区
$ git push
fatal: 您当前分支的上游分支和您当前分支名不匹配，为推送到远程的上游分支，
使用

    git push origin HEAD:main

为推送至远程同名分支，使用

    git push origin HEAD

为了永久地选择任一选项，参见 'git help config' 中的 push.default。

为了避免在与本地分支名字不匹配时自动设置上游分支，参见 'git help config'
中 branch.autoSetupMerge 的 'simple' 选项。

$ git push origin HEAD:main
枚举对象中: 16, 完成.
对象计数中: 100% (16/16), 完成.
使用 8 个线程进行压缩
压缩对象中: 100% (11/11), 完成.
写入对象中: 100% (15/15), 3.97 KiB | 3.97 MiB/s, 完成.
总共 15（差异 2），复用 0（差异 0），包复用 0
remote: Resolving deltas: 100% (2/2), done.
To github.com:chengchaos/chit-chat.git
   d7574ba..f973382  HEAD -> main
```