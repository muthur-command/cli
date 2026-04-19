# Muthur Command CLI

用于与 **Muthur Command Supervisor** 交互的命令行界面。

## 用法

- `mc help`
- `mc <subcommand> <action> [<options>]`

示例：

- `mc core info --raw-json`

### 修饰符

#### 全局

```text
      --api-token string   Muthur Command Supervisor API 令牌
      --config string      可选配置文件（默认为 $HOME/.muthurcommand.yaml）
      --endpoint string    Muthur Command Supervisor 端点（默认为 'supervisor'）
  -h, --help               mc 的帮助信息
      --log-level string   日志级别（默认为 Warn）
      --no-progress        关闭进度指示动画
      --raw-json           输出 API 返回的原始 JSON
```

所有选项也可通过带 `SUPERVISOR_` 前缀的环境变量使用，例如 `SUPERVISOR_LOG_LEVEL`。

#### 子命令

可用命令：

```text
  apps           安装、更新、移除与配置 Muthur Command 应用
  audio          音频设备处理。
  authentication Muthur Command 用户认证。
  cli            获取信息、更新或配置 Muthur Command CLI 后端
  core           控制 Muthur Command Core
  dns            获取信息、更新或配置 Muthur Command DNS 服务
  docker         Docker 后端，用于信息与 OCI 配置
  hardware       提供本机硬件信息
  help           任意命令的帮助
  host           控制运行 Muthur Command 的主机/系统
  info           提供 Muthur Command 的总体信息概览
  multicast      获取信息、更新或配置 Muthur Command Multicast
  network        网络：更新、信息与配置导入
  observer       获取信息、更新或配置 Muthur Command observer
  os             操作系统：更新、信息与配置导入
  resolution     Supervisor 问题中心：展示问题与建议方案
  backups        创建、恢复与删除备份
  supervisor     监视、控制与配置 Muthur Command Supervisor
```

## 安装

在 Muthur Command 系统上，CLI 由 **CLI 容器**提供；使用 **Muthur Command Operating System** 时，也可在设备终端中使用。

在上述系统上，CLI 会自动更新。

此外，应用商店中的 **SSH** 应用可使用本工具；部分社区应用也集成了它（例如 **Visual Studio Code** 应用）。

## 开发与贡献

### 前置条件

CLI 可通过 [开发者应用仓库](https://github.com/muthur-command/addons-development) 中的 **`remote_api`** 应用，与 Muthur Command Supervisor 远程交互。

安装并启动该应用后，在 **`remote_api`** 应用日志中会显示令牌，后续开发需要用到。

### 获取源码

在 [GitHub 上 fork](https://github.com/muthur-command/cli/fork) 本仓库，或克隆本仓库。

### 开发中使用

```shell
export SUPERVISOR_ENDPOINT=http://192.168.1.2
export SUPERVISOR_API_TOKEN=replace_this_with_remote_api_token
go run main.go info
```

**说明：** 将 `192.168.1.2` 替换为运行 **`remote_api`** 应用的 Muthur Command 实例 IP，并使用日志中提供的令牌。

### 构建

本仓库使用 Go modules；示例构建：

```bash
CGO_ENABLED=0 go build -ldflags="-s -w" -o "mc"
```

跨架构构建的更多说明见本仓库的 [构建工作流](https://github.com/muthur-command/cli/blob/master/.github/workflows/build.yml)。

### 提交变更

1. 在 fork/克隆的仓库上创建功能分支。
2. 提交你的修改。
3. 将本地修改在 **`master`** 分支上变基（rebase）。
4. 运行 `go test ./...` 并确认通过。
5. 运行 `gofmt -s` 保证代码格式正确。
6. 创建新的 Pull Request。

## 来源

- **上游：** [home-assistant/cli](https://github.com/home-assistant/cli) — Home Assistant Supervisor 命令行工具，本仓库由其移植而来。
- **本仓库：** **Muthur Command** 维护此 fork，供 **Muthur Command OS** 使用；行为可能随时间与上游产生差异。
- **许可：** 自上游继承的代码仍为 **Apache-2.0**；详见 [`LICENSE`](./LICENSE)。
