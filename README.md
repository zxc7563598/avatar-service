# Avatar Service

<div align="center">
  <a href="./README.md">English</a>｜<a href="./README.zh-CN.md">简体中文</a>
  <hr width="50%"/>
</div>

A random avatar generation service written in Go, supporting multiple avatar styles (Identicon, Pixel).  
Ideal for personal projects, application registration avatars, fun tools, etc., making every default avatar unique.

**This project has been parsed by Zread. If you need a quick overview of the project, you can click here to view it：[Understand this project](https://zread.ai/zxc7563598/avatar-service)**

---

## Features

- **Multiple Avatar Styles**

  - **Identicon**: Symmetrical grid style (GitHub-like)
  - **Pixel**: Pixel block style

- **Reproducible**  
  The same seed always generates the same avatar
- **Lightweight**  
  No external dependencies, can be compiled into a single binary
- **Server Deployable**  
  HTTP API returns PNG images, easy to integrate with web services
- **High Performance**  
  Go’s high-performance rendering + PNG output + caching reduces repeated generation

---

## Project Structure

```
/avatar-service
├── main.go                 # Main file
├── internal/
│   ├── generator/
│   │   ├── identicon.go    # Identicon style generation logic
│   │   ├── pixel.go        # Pixel style generation logic
│   ├── cache/              # Simple in-memory cache logic
│   └── utils/              # Hash utility functions
└── README.md
```

---

## Installation & Running

1. Clone the repository:

```bash
git clone https://github.com/zxc7563598/avatar-service.git
cd avatar-service
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run directly:

```bash
go run main.go --port 8080
```

The service will start at `http://localhost:8080`​

4. Build and run as a binary:

```bash
go build -o app
./app --port=8080
```

---

## API Usage

### Get an Avatar

```
GET /avatar?seed=<seed>&style=<style>
```

- ​`seed`: String used to generate the avatar (same seed always produces the same avatar)
- ​`style`: Avatar style, optional values:

  - ​`identicon` -- Symmetrical grid style (GitHub-like)
  - ​`pixel` -- Pixel block style

- Example:

```
http://localhost:8080/avatar?seed=hejunjie&size=128&style=pixel
```

---

## Future Plans

I plan to add more generation types in the future to make avatars more diverse and fun:

- **Simple Face Combinations**: Different hair, eyes, accessories (glasses, etc.)
- **Color Randomization**: Random colors for each element
- **More Styles**: Beyond grids and pixel blocks, generate more personalized and cute default avatars

These enhancements aim to make every avatar more fun and unique.

---

## Contributing & Feedback

If you find this project interesting, feel free to:

- ⭐ Star the project
- 🐛 Submit issues or suggestions
- 📥 Fork and try using it in your own projects

Any feedback motivates me to improve this project, and contributions of new avatar styles and features are always welcome!
