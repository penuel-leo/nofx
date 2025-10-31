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

### 部署前本地验证

```bash
# 1. 构建
npm run build

# 2. 预览（可选）
npm run preview
npm run dev
# 访问 http://localhost:4173
```

### Vercel 配置步骤

#### 1. 框架选择

在 Vercel 项目设置中选择 **`Vite`** 作为框架预设。

#### 2. 环境变量配置

在 Vercel 项目设置 → Environment Variables 中添加：

| 变量名 | 值（示例） |
|--------|-----------|
| `VITE_API_BASE_URL` | `https://nofx-backend.onrender.com/api` |

**应用到所有环境：** Production、Preview、Development

#### 3. 构建配置（自动检测，无需修改）

- Build Command: `npm run build`
- Output Directory: `dist`
- Install Command: `npm install`

#### 4. 部署验证

部署成功后：

1. **访问 Vercel 提供的 URL**
2. **打开浏览器控制台**（F12 → Console）
3. **查看日志输出**：

```javascript
🔧 API Configuration: {
  VITE_API_BASE_URL: "https://nofx-backend.onrender.com/api",
  API_BASE: "https://nofx-backend.onrender.com/api",
  mode: "production",
  isDev: false,
  isProd: true
}
```

4. **刷新页面**，确认仍然可以正常访问（SPA 路由测试）

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
