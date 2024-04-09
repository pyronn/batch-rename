
# batch-rename

这是一个用Go编写的命令行工具，用于批量重命名文件。它支持完全重命名、前缀重命名和后缀重命名三种模式，并提供了灵活的文件选择机制，包括全选目录下的所有文件、通过正则表达式筛选文件，以及用户自定义选择文件。

## 功能特点

- **完全重命名**：将选中的所有文件重命名为指定的名称，如果有多个文件，则在名称后自动加上序号。
- **前缀重命名**：在选中文件的原有名称前添加指定的前缀。
- **后缀重命名**：在选中文件的原有名称后添加指定的后缀。
- **灵活的文件选择**：支持全选目录下的所有文件、通过正则表达式筛选文件，以及用户自定义选择文件。

## 安装

确保你已经安装了Go环境（版本1.13及以上）。然后执行以下命令：

```bash
go install github.com/pyronn/filerenamer@latest
```

这将编译并安装文件重命名工具到你的`$GOPATH/bin`目录下。

## 使用方法

在命令行中，你可以通过以下方式使用文件重命名工具：

```bash
filerenamer [options]
```

### 选项

- `-dir`：指定包含要重命名文件的目录。
- `-type`：选择重命名模式（`full`、`prefix`、`suffix`）。
- `-name`：在`full`模式下指定新名称。
- `-prefix`：在`prefix`模式下指定前缀。
- `-suffix`：在`suffix`模式下指定后缀。
- `-regex`：启用正则表达式筛选文件。
- `-pattern`：指定用于筛选文件的正则表达式模式。
- `-ext`：是否同时重命名文件的扩展名。

## 示例

完全重命名目录`/path/to/files`下的所有文件为`newname`，并自动添加序号：

```bash
filerenamer -dir /path/to/files -type full -name newname
```

为目录`/path/to/files`下的所有文件添加前缀`prefix_`：

```bash
filerenamer -dir /path/to/files -type prefix -prefix prefix_
```

## 许可证

本项目使用MIT许可证。详情见[LICENSE](LICENSE)文件。

