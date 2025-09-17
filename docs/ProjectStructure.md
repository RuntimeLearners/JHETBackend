# 项目目录结构规范

此文档规定了这个后端项目的目录结构，可能随着开发进程而发生更改。

```
PROJECT
├─configs #配置文件，不应包含代码
├─docs #项目文档，仅供开发参考
├─internal #私有业务代码
│  └─global #全局(结构体等)定义(也称config)
└─packages #外部库
    └─utils #外部现成单元
```

注：通过shell的 `tree` 命令来生成目录结构图