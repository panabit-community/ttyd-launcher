# Panabit 网关插件系列之网页终端模拟器

## 摘要

[ttyd](https://github.com/tsl0922/ttyd) 是 @tsl0922 开发的一款网页终端模拟器。本项目提供了在 Panabit 智能应用网关平台上的一种插件封装，可以在浏览器管理页中直接交互式执行命令。

## 已知问题

当前项目仅为最小可用版本，下面列出一些亟待解决的问题。

+ 仅支持 linux/arm64 平台，例如 AX50C/AX40
+ 不支持自定义管理口 IP 地址（代码中硬编码了）
+ 不支持认证和其它参数，目前默认启动一个监听 7681 端口的终端，使用一次后自动退出