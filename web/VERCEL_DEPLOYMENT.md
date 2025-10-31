# Vercel 部署指南

## 项目配置

此项目已按照 Vercel 最新标准进行优化配置。

## Vercel Dashboard 配置

### 1. 项目设置 (Project Settings)

进入 Vercel Dashboard → 您的项目 → Settings → General

#### Root Directory
```
web
```

#### Build & Development Settings

- **Framework Preset**: `Vite`
- **Build Command**: `npm run build`（或留空让 Vercel 自动检测）
- **Output Directory**: `dist`（或留空让 Vercel 自动检测）
- **Install Command**: `npm install`（或留空让 Vercel 自动检测）
- **Development Command**: `vite`（或留空让 Vercel 自动检测）

### 2. 推荐配置（Override 设置）

建议**不启用 Override**，让 Vercel 自动检测：
- ⬜ Build Command (关闭 Override)
- ⬜ Output Directory (关闭 Override)
- ⬜ Install Command (关闭 Override)
- ⬜ Development Command (关闭 Override)

### 3. Node.js 版本
```
20.x 或 22.x
```

### 4. Deployment Protection

如果启用了 Standard Protection，请注意：
- 预览部署需要登录才能访问
- 生产部署（main 分支）应该是公开的

## 文件说明

### vercel.json
```json
{
  "rewrites": [{ "source": "/(.*)", "destination": "/" }]
}
```

最简化配置，只包含 SPA 路由重写规则。

### vite.config.ts
包含必要的构建配置：
- `publicDir`: 指定 public 目录
- `build.outDir`: 输出目录为 dist
- `build.assetsDir`: 静态资源目录为 assets

## 部署流程

1. **提交代码**
   ```bash
   git add .
   git commit -m "optimize: vercel deployment config"
   git push origin main
   ```

2. **等待 Vercel 自动部署**（2-3 分钟）

3. **检查部署**
   - 进入 Vercel Dashboard → Deployments
   - 查看最新部署状态
   - 点击 "Visit" 访问部署的网站

## 故障排查

### 如果部署失败

1. **检查构建日志**
   - 在 Deployment 详情页查看完整日志
   - 确认 `npm run build` 执行成功

2. **检查 Source/Output**
   - 点击部署详情 → Source → Output
   - 确认以下文件存在：
     - index.html
     - assets/index-*.js
     - assets/index-*.css
     - vite.svg

3. **本地测试构建**
   ```bash
   cd web
   npm run build
   npm run preview
   ```

### 如果无法访问

1. **清除浏览器缓存**
2. **使用无痕模式访问**
3. **检查 Deployment Protection 设置**
4. **确认使用正确的 URL**（生产 URL 而非预览 URL）

## 项目结构

```
web/
├── public/          # 静态资源目录
│   └── vite.svg    # 网站图标
├── src/            # 源代码
├── dist/           # 构建输出（Git 已忽略）
├── vercel.json     # Vercel 配置
├── vite.config.ts  # Vite 配置
└── package.json    # 项目依赖
```

## 注意事项

1. **不要修改 .gitignore 中的 `web/dist/`** - 构建产物不应提交到 Git
2. **vercel.json 保持最简配置** - 让 Vercel 自动处理大部分配置
3. **确保 public 目录存在** - Vite 会将其内容复制到 dist

## 环境变量

如需配置环境变量（如 API 地址）：

1. Vercel Dashboard → Settings → Environment Variables
2. 添加变量，例如：
   - `VITE_API_URL`: API 服务器地址
3. 重新部署以应用变量

## 优化建议

当前构建警告提示 JavaScript 包过大（> 500KB），建议：

1. **代码分割**：使用动态导入
   ```typescript
   const CompetitionPage = lazy(() => import('./components/CompetitionPage'))
   ```

2. **手动分块**：配置 `vite.config.ts` 中的 `build.rollupOptions.output.manualChunks`

3. **依赖优化**：审查并移除未使用的依赖

