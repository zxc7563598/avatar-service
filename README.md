# Avatar Service

一个用 Go 写的随机头像生成服务，支持多种风格的头像生成（Identicon、Pixel、Cartoon）。  
适合个人项目展示、应用注册头像生成、趣味工具等场景。

---

## 特性

- **多风格头像**
  - **Identicon**：对称格子风格（仿GitHub）
  - **Pixel**：像素方块风格
- **可复现**  
  相同 seed 永远生成相同头像
- **轻量**  
  无需外部依赖，可编译成单个二进制部署
- **可部署在服务器**  
  HTTP 接口返回 PNG 图片，方便集成到 Web 服务
- **高性能**
  Go 高性能绘图 + PNG 输出 + 缓存减少重复生成

---

## 目录结构

## 安装与运行

1. 克隆仓库：
```bash
git clone https://github.com/zxc7563598/avatar-service.git
cd avatar-service
```

2. 安装依赖：

```bash
go mod tidy
```

3. 编译并运行：

```bash
go run main.go --port 8080
```

默认服务启动在 `http://localhost:8080`​

3. 打包后运行：

```bash
go build -o app
./app --port=8080
```

## API 使用

### 获取头像

```
GET /avatar?seed=<seed>&style=<style>
```

- ​`seed`：字符串，用于生成头像（相同 seed 永远生成相同头像）
- ​`style`：头像风格，可选：

  - ​`identicon`​ -- 对称格子风格（仿GitHub）
  - ​`pixel`​ -- 像素方块风格
- 示例：

```
http://localhost:8080/avatar?seed=hejunjie&size=128&style=pixel
```