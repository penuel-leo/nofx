# NOFX Web Dashboard

基于 Vite + React + TypeScript 的AI自动交易监控面板

## 技术栈

- **React 18** - UI框架
- **TypeScript** - 类型安全
- **Vite** - 构建工具
- **Tailwind CSS** - 样式框架
- **SWR** - 数据获取和缓存
- **Zustand** - 状态管理
- **Recharts** - 图表库

## 安装依赖

```bash
npm install
```

## 环境变量配置

首次运行前，请复制环境变量示例文件：

```bash
cp .env.example .env
```

可配置的环境变量：

| 变量名 | 说明 | 默认值 | 示例 |
|--------|------|--------|------|
| `VITE_API_BASE_URL` | 后端 API 基础地址 | `/api` | `http://localhost:8080/api` |

**配置说明：**
- **开发环境**：使用默认值 `/api`，通过 Vite 代理转发到 `http://localhost:8080`
- **生产环境**：可设置为完整的后端地址，如 `https://api.example.com/api`

## 运行开发服务器

```bash
npm run dev
```

访问 http://localhost:3000

## 构建生产版本

```bash
npm run build
```

构建产物输出到 `dist` 目录。

## Vercel 部署

> 📋 **完整部署检查清单**：请查看 [`DEPLOY_CHECKLIST.md`](./DEPLOY_CHECKLIST.md)

### 部署前验证（重要！）

在部署到 Vercel 之前，**必须先在本地验证**，避免反复部署浪费时间：

```bash
# 快速验证脚本（推荐）
./scripts/verify-build.sh

# 或者手动验证
npm run build    # 查看构建日志中的环境变量
npm run preview  # 访问 http://localhost:4173
```

**验证要点：**
1. ✅ 构建日志显示环境变量检查信息
2. ✅ 浏览器控制台显示 `🔧 API Configuration`
3. ✅ 页面可以正常访问（不是 404）
4. ✅ 刷新页面后仍然正常

### 1. 框架选择

在 Vercel 项目设置中选择 **`Vite`** 作为框架预设。

### 2. 环境变量配置

在 Vercel 项目设置的 Environment Variables 中添加：

| 变量名 | 值 | 说明 |
|--------|-----|------|
| `VITE_API_BASE_URL` | `https://your-backend.com/api` | 后端 API 地址 |

### 3. 部署配置

项目已包含 `vercel.json` 配置文件：
- **输出目录**：`dist`（Vite 默认输出）
- **构建命令**：`npm run build`
- **开发命令**：`npm run dev`

### 4. 验证部署

部署后，打开浏览器控制台，应该能看到类似输出：

```
🔧 API Configuration: {
  VITE_API_BASE_URL: "https://your-backend.com/api",
  API_BASE: "https://your-backend.com/api",
  mode: "production",
  isDev: false,
  isProd: true
}
```

### 常见问题

**Q: 构建失败提示找不到 build 目录？**  
A: 确保在 Vercel 中选择了 `Vite` 框架，或检查 `vercel.json` 中 `outputDirectory` 配置为 `dist`。

**Q: 环境变量未生效？**  
A: 确保在 Vercel 项目设置中添加了 `VITE_API_BASE_URL` 环境变量，并重新部署。

## 功能特性

### 实时监控
- **系统状态** - 运行状态、AI提供商、周期数
- **账户信息** - 净值、可用余额、总盈亏、保证金使用率
- **持仓列表** - 实时价格、盈亏、杠杆、强平价
- **决策日志** - 完整的AI思维链（可展开）、决策动作、执行结果

### AI思维链分析
每个决策记录都包含完整的AI思考过程：
- **第一步**：现有持仓分析（技术指标、盈亏评估）
- **第二步**：账户风险评估（保证金使用率、可用余额）
- **第三步**：新机会评估（候选币种筛选、技术形态分析）
- **第四步**：最终决策总结（平仓/开仓/持有决策）

点击 "💭 AI思维链分析" 即可展开查看完整分析过程！

### 自动刷新
- 系统状态、账户、持仓：每5秒刷新
- 决策日志、统计：每10秒刷新

### API集成

前端通过环境变量 `VITE_API_BASE_URL` 配置后端地址。

**配置方式：**
- 在 `.env` 文件中设置 `VITE_API_BASE_URL`
- 默认值为 `/api`，开发环境通过 Vite 代理转发到 `http://localhost:8080`
- 代码位置：`src/lib/api.ts` 第 13 行

**API端点：**
- `GET /api/status` - 系统状态
- `GET /api/account` - 账户信息
- `GET /api/positions` - 持仓列表
- `GET /api/decisions` - 决策日志（最近30条）
- `GET /api/decisions/latest` - 最新决策（最近5条）
- `GET /api/statistics` - 统计信息

## 项目结构

```
web/
├── src/
│   ├── components/      # React组件（待扩展）
│   ├── lib/
│   │   └── api.ts      # API调用函数
│   ├── store/          # Zustand状态管理（待扩展）
│   ├── types/
│   │   └── index.ts    # TypeScript类型定义
│   ├── App.tsx         # 主应用组件
│   ├── main.tsx        # 入口文件
│   └── index.css       # 全局样式
├── index.html          # HTML模板
├── vite.config.ts      # Vite配置
├── tailwind.config.js  # Tailwind配置
├── tsconfig.json       # TypeScript配置
└── package.json        # 依赖配置
```

## 注意事项

1. **确保后端API服务已启动**（默认端口8080）
2. **Node.js版本要求**：>= 18.0.0
3. **网络连接**：需要访问Binance API

## 开发计划

- [ ] 添加图表展示（账户净值走势、盈亏曲线）
- [ ] 添加决策详情页面（完整的CoT分析）
- [ ] 添加手动交易控制
- [ ] 添加参数配置页面
- [ ] 添加通知和告警系统
