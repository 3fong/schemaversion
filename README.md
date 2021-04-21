
## go modules用法

https://github.com/golang/go/wiki/Modules

Your typical day-to-day workflow can be:

Add import statements to your .go code as needed.
Standard commands like go build or go test will automatically add new dependencies as needed to satisfy imports (updating go.mod and downloading the new dependencies).
When needed, more specific versions of dependencies can be chosen with commands such as go get foo@v1.2.3, go get foo@master (foo@default with mercurial), go get foo@e3702bed2, or by editing go.mod directly.

## 库,模块,包关系

Summarizing the relationship between repositories, modules, and packages:

A repository contains one or more Go modules.
Each module contains one or more Go packages.
Each package consists of one or more Go source files in a single directory.



