// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go mod init

package modcmd

import (
	"cmd/go/internal/base"
	"cmd/go/internal/modload"
	"context"
)

var cmdInit = &base.Command{
	UsageLine: "go mod init [module-path]",
	Short:     "initialize new module in current directory",
	Long: `
Init initializes and writes a new go.mod file in the current directory, in
effect creating a new module rooted at the current directory. The go.mod file
must not already exist.

Init accepts one optional argument, the module path for the new module. If the
module path argument is omitted, init will attempt to infer the module path
using import comments in .go files, vendoring tool configuration files (like
Gopkg.lock), and the current directory (if in GOPATH).

See https://golang.org/ref/mod#go-mod-init for more about 'go mod init'.
`,
	Run: runInit,
}

func init() {
	base.AddChdirFlag(&cmdInit.Flag)
	base.AddModCommonFlags(&cmdInit.Flag)
}

// 执行go mod init 命令执行入口
func runInit(ctx context.Context, cmd *base.Command, args []string) {
	//args 是go mod init 后面带的参数 仅支持一个参数
	if len(args) > 1 {
		base.Fatalf("go: 'go mod init' accepts at most one argument")
	}
	var modPath string
	//获取参数即包的名称
	if len(args) == 1 {
		modPath = args[0]
	}
	//设置forceUseModules为true
	modload.ForceUseModules = true
	//创建mod文件
	modload.CreateModFile(ctx, modPath) // does all the hard work
}
