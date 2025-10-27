# Avatar Service

<div align="center">
  <a href="./README.md">English</a>｜<a href="./README.zh-CN.md">简体中文</a>
  <hr width="50%"/>
</div>

一个用 Go 写的随机头像生成服务，支持多种风格的头像生成（Identicon、Pixel）。  
适合个人项目展示、应用注册头像生成、趣味工具等场景，让每个默认头像都独一无二。

**本项目已经经由 Zread 解析完成，如果需要快速了解项目，可以点击此处进行查看：[了解本项目](https://zread.ai/zxc7563598/avatar-service)**

---

## 特性

- **多风格头像**

  - **Identicon**：对称格子风格（仿 GitHub）
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

```
/avatar-service
├── main.go                 # 主文件
├── internal/
│   ├── generator/
│   │   ├── identicon.go    # identicon 风格生成逻辑
│   │   ├── pixel.go        # pixel 风格生成逻辑
│   ├── cache/              # 简易内存缓存逻辑
│   └── utils/              # hash 工具函数
└── README.md
```

---

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

4. 打包后运行：

```bash
go build -o app
./app --port=8080
```

---

## API 使用

### 获取头像

```
GET /avatar?seed=<seed>&style=<style>
```

- ​`seed`：字符串，用于生成头像（相同 seed 永远生成相同头像）
- ​`style`：头像风格，可选：

  - ​`identicon` -- 对称格子风格（仿 GitHub）
  - ​`pixel` -- 像素方块风格

- 示例：

```
http://localhost:8080/avatar?seed=hejunjie&size=128&style=pixel
```

---

## 未来扩展

我计划在未来增加更多生成类型，让头像更加丰富和可玩：

- **简笔人脸组合**：不同头发、眼睛、饰品（眼镜等）
- **颜色随机化**：每个元素可随机搭配颜色
- **更多风格**：不仅限于格子和像素风格，可以生成更个性化、可爱的默认头像

希望通过这些扩展，让每个头像都更有趣，也更有个性化的体验。

---

## 参与与反馈

如果你觉得这个项目有趣，欢迎：

- ⭐ 给项目点 Star
- 🐛 提交 Issue 或建议
- 📥 Fork 并尝试在自己的项目中使用

任何反馈都会让我更有动力去完善这个项目，也欢迎贡献新的头像风格和功能！
