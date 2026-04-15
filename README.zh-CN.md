# MCOS CLI（`ha`）

面向 **MCOS** 监管栈的 **Supervisor HTTP API** 命令行客户端（P0：GitHub 组织 **`muthur-command`**，Go 模块 **`github.com/muthur-command/cli`**）。

## 用法

- `ha help`
- `ha <subcommand> <action> [<options>]`

示例：`ha core info --raw-json`

全局参数与 `SUPERVISOR_*` 环境变量说明见英文 **README.md** 中的摘要块；完整子命令列表以本仓库编译出的 **`ha help`** 为准。

## 安装与制品名

本仓库 Release 仍上传 **`ha_amd64`** / **`ha_aarch64`** 资源名，便于 **`plugin-cli`** 等镜像在构建时按固定 URL 拉取。

## 开发与构建

参见 **README.md** 中 Remote Supervisor、`go build` 与 **`.github/workflows/build.yml`**（分支 **`mc`**）。

## 许可证

见 **LICENSE**（Apache-2.0；保留上游版权 / NOTICE 要求）。
