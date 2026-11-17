# [Wails-JSON-Editor](https://github.com/William-Weng?tab=repositories&q=wails)

![Go](https://img.shields.io/badge/Go-1.27.1-00ADD8?logo=go&logoColor=white)
![Node.js](https://img.shields.io/badge/Node.js-20.19.2-339933?logo=nodedotjs&logoColor=white)
![Wails](https://img.shields.io/badge/Wails-v3.0.0-DF0000?logo=wails&logoColor=white)
![LICENSE](https://img.shields.io/github/license/William-Weng/Wails-JSON-Editor?style=flat&label=LICENSE&color=yellow)
![Tag](https://img.shields.io/github/v/tag/William-Weng/Wails-JSON-Editor?style=flat&label=Tag)
![Stars](https://img.shields.io/github/stars/William-Weng/Wails-JSON-Editor?style=flat&label=Stars)

使用 Wails 3、Svelte 5 與 TypeScript 打造的桌面 JSON 編輯器。

提供純文字與樹狀檢視，支援 JSON 格式化、壓縮、語法驗證，以及獨立調整兩側編輯器字級。

https://github.com/user-attachments/assets/846d5028-9936-45bd-ba7f-eb7274dba1db

## 功能特色

- 雙編輯器介面
  - CODE：純文字 JSON 編輯，附行號與捲動同步。
  - JSON：使用樹狀結構瀏覽與編輯資料。
- 三種檢視模式
  - Text：只顯示文字編輯器。
  - JSON：只顯示樹狀編輯器。
  - Text JSON：並排顯示兩個編輯器。
- JSON 格式化：使用 2 個空格縮排。
- JSON 壓縮：移除非必要空白與換行。
- JSON 解析錯誤提示。
- 狀態列顯示文字行數。
- 兩側編輯器可分別調整字級，並支援重設。
- 使用可複用的 Svelte 元件組織介面。
- 使用原生 CSS 與 CSS 變數管理樣式，不依賴 LESS。

## 技術組成

| 技術 | 用途 |
| --- | --- |
| Wails 3 | 桌面應用程式框架 |
| Svelte 5 | 元件化介面與響應式狀態 |
| TypeScript | 前端型別檢查 |
| svelte-jsoneditor | JSON 樹狀編輯器 |
| Vite | 前端開發與建置 |
| CSS | 元件樣式與共用主題變數 |

## 使用方式

### 編輯 JSON

在 CODE 面板輸入或貼上 JSON。

當內容可以成功解析時，JSON 面板會同步更新；若內容尚未完成或格式有誤，狀態列會顯示解析錯誤，JSON 面板則保留最近一次有效的內容。

也可以直接在 JSON 樹狀面板修改資料，變更會同步回 CODE 面板。

### 切換檢視

使用底部狀態列的切換按鈕：

| 模式 | 說明 |
| --- | --- |
| Text | 只顯示 CODE 面板 |
| JSON | 只顯示 JSON 面板 |
| Text JSON | 同時顯示兩個面板 |

### 格式化與壓縮

- Format：重新整理 JSON，使用 2 個空格縮排。
- Minify：將 JSON 轉換為不含非必要空白與換行的精簡內容。

如果文字內容不是有效的 JSON，操作會顯示錯誤訊息，而不覆寫原本的文字。

### 調整字級

兩個面板都提供獨立的字級控制：

- `-`：縮小字級。
- `+`：放大字級。
- 重設：恢復預設字級。

預設字級為 `14px`，可調整範圍為 `10px` 至 `32px`。

## 本機開發

### 環境需求

請先準備：

- Go
- Node.js 與 npm
- Wails 3 CLI
- 作業系統所需的 Wails 開發依賴

環境安裝方式請參閱 [Wails 官方網站](https://wails.io/) 的 Wails 3 文件。

### 安裝前端依賴

在專案根目錄執行：

```bash
cd frontend
npm install
cd ..
```

### 啟動開發模式

在專案根目錄執行：

```bash
wails3 dev
```

## 前端結構

```text
frontend/
└── src/
    ├── lib/
    │   ├── components/
    │   │   ├── CodePane.svelte
    │   │   ├── JsonPane.svelte
    │   │   ├── PaneHeader.svelte
    │   │   ├── FontControls.svelte
    │   │   ├── StatusBar.svelte
    │   │   ├── ViewSwitch.svelte
    │   │   └── TransformActions.svelte
    │   └── styles/
    │       └── theme.css
    └── App.svelte
```

### 元件職責

| 元件 | 職責 |
| --- | --- |
| `App.svelte` | 管理 JSON 內容、錯誤訊息、檢視模式與字級狀態 |
| `CodePane.svelte` | 純文字編輯、行號顯示與行號捲動同步 |
| `JsonPane.svelte` | 包裝 JSON 樹狀編輯器 |
| `PaneHeader.svelte` | 共用面板標題與控制區 |
| `FontControls.svelte` | 字級調整與重設按鈕 |
| `StatusBar.svelte` | 顯示行數、驗證訊息與操作區 |
| `ViewSwitch.svelte` | 切換文字、樹狀與並排檢視 |
| `TransformActions.svelte` | 顯示格式化與壓縮操作按鈕 |
| `theme.css` | 定義共用色彩與全域基礎樣式 |

## 狀態與資料同步

應用程式以 `App.svelte` 作為主要狀態容器：

```text
CODE 文字輸入
    │
    ▼
更新來源文字
    │
    ▼
嘗試解析 JSON
    │
    ├── 成功 → 更新 JSON 樹狀面板，清除錯誤
    │
    └── 失敗 → 保留輸入文字，顯示錯誤
```

由 JSON 面板編輯時，則將編輯結果轉換為文字，再同步更新 CODE 面板。

文字輸入與 JSON 解析分開處理，讓使用者在輸入尚未完成的 JSON 時，仍可保留目前的編輯內容。

## 樣式設計

每個 Svelte 元件使用自己的 `<style>` 管理局部樣式，共用色彩則定義於 `theme.css`。

例如：

```css
:root {
  --editor-bg: #132e38;
  --editor-text: #d5e4eb;
  --status-bg: #0b222c;
  --status-text: #8ba1ac;
}
```

這種方式讓元件樣式保持集中，也方便統一調整介面配色。

## 注意事項

- 格式化與壓縮操作會重新序列化 JSON，不會保留原本的空白與排版。
- 由樹狀面板修改資料後，文字面板會重新產生格式化的 JSON。
- JSON 解析與序列化使用 JavaScript 的 `JSON.parse()` 與 `JSON.stringify()`；處理超過 JavaScript 安全整數範圍的數值時，請留意精度問題。
- 目前編輯內容由前端記憶體中的狀態管理，不應視為已儲存至檔案。

## [建置指令整理](https://v3.wails.io/zh-tw/guides/build/building/)

| 目的 | 指令 |
| --- | --- |
| 建立新專案 | `wails3 init -n <專案名稱> -t <前端框架>` |
| 產生 bindings | `wails3 generate bindings` |
| 更新 Windows 與 macOS 專用圖示檔案 | `wails3 generate icons -input build/appicon.png -windowsfilename build/windows/icon.ico -macfilename build/darwin/icons.icns` |
| 更新建置資源 | `wails3 update build-assets -config build/config.yml -dir build` |
| 拉取（下載）用於跨平台交叉編譯的 Docker 映像檔 | `wails3 task setup:docker` |
| 打包 macOS arm64 | `wails3 package GOOS=darwin GOARCH=arm64` |
| 建置 Windows x64 | `wails3 build GOOS=windows GOARCH=amd64` |
| 打包 Linux arm64 | `wails3 build GOOS=linux GOARCH=arm64` |

## 致謝

- [Wails](https://wails.io/)
- [Svelte](https://svelte.dev/)
- [svelte-jsoneditor](https://github.com/josdejong/svelte-jsoneditor)
- [Vite](https://vite.dev/)
