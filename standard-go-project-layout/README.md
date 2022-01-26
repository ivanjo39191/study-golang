# Standard Go Project Layout

## Go 目錄結構

### `/cmd`

採用 `/cmd/[appname]/main.go` 進行組織

cmd 本專案的主要應用程式

cmd 目錄下的文件**不該有太多程式碼，不該包含業務邏輯**

通常主要應用程式只會有一個小小的 `main` 函式，`main` 函式主要做的事情就是負責應用程式的生命週期，然後大部分的程式都是從 `/internal` 和 `/pkg` 匯入呼叫並執行，除此之外應該什麼都沒有！

### `/internal`

**私有應用程式**和**函式庫**的程式碼，是你不希望其他人在其應用程式或函式庫中匯入的程式碼。請注意：這個目錄結構是由 Go 編譯器本身所要求的。

一般會把私有應用程式放到 internal 中，例如 `/internal/app`

可被其他應用程式匯入的套件會放到 `/pkg` 目錄

只在應用程式內部共享的套件會放到 `/internal/pkg`

你在專案目錄下的任何子目錄都可以包含 `internal` 目錄。

### `/pkg`

函式庫的程式碼當然可以讓外部應用程式來使用 (例如：`/pkg/mypubliclib`)，其他專案會匯入這些函式庫，並且期待它們能正常運作。

注意：使用 `internal` 目錄可以確保私有套件不會被匯入到其他專案使用，因為它是由 Go 的編譯器強制執行的，所以是比較好的解決方案。

### `/vendor`

應用程式的相依套件可透過手動管理，或使用你喜歡的相依性套件管理工具。

注意：從 `[Go 1.13](https://golang.org/doc/go1.13#modules)` 開始，Go 預設啟用了**模組的代理伺服器** (module proxy) 功能 (預設使用 `[https://proxy.golang.org](https://proxy.golang.org/)` 作為模組的代理伺服器)。如果你可以使用 module proxy 的話，那麼你根本不需要 `vendor` 目錄

## Web 應用程式目錄**(Web Application Directories)**

Web 應用程式相關的元件：靜態 Web 檔案、伺服器端範本與 SPAs 相關檔案。

## **服務應用程式目錄 (Service Application Directories)**

### **`/api`**

API 定義的目錄。

如果我們採用的是 grpc 那這裡面一般放的就是 proto 文件，也可能是 OpenAPI/Swagger 規格、JSON schema 檔案、各種協定定義檔。

```go
.
└── api
    └── product_name // 產品名稱
        └── app_name // 應用名稱
            └── v1   // 版本號
                └── v1.proto
```

## **通用應用程式目錄 (Common Application Directories)**

### **`/configs`**

組態設定的檔案範本或預設設定。

將你的 `confd` 或 `consul-template` 範本檔案放在這裡。

### **`/init`**

放置 System init (`systemd`, `upstart`, `sysv`) 與 Process manager/supervisor (`runit`, `supervisor`) 相關設定。

### **`/scripts`**

放置要執行各種建置、安裝、分析等操作的命令腳本

這些腳本可以讓你在根目錄的 `Makefile` 更小、更簡單。

### **`/build`**

封裝套件與持續整合(CI)。

將你的雲端 (AMI)、容器 (Docker)、OS (deb, rpm, pkg) 套件的組態設定與腳本放在 `/build/package` 目錄下。

將你的 CI (Travis CI, CircleCI, Drone CI) 的組態設定與腳本放在 `/build/ci` 目錄中。請注意：有些 CI 工具 (例如 Travis CI 等)，它們對這些組態設定檔案的位置非常挑剔。如果可能的話，請嘗試將檔案放在 `/build/ci` 目錄中，並**連結** (linking) 這些檔案到 CI 工具期望它們出現的位置。

### **`/deployments`**

IaaS、PaaS、系統和容器編配部署的組態設定與範本 (docker-compose、kubernetes/helm、mesos、terraform、bosh)。注意：在某些儲存庫中（特別是那些部署在 kubernetes 的應用程式），這個目錄會被命名為 `/deploy`。

### **`/test`**

額外的外部測試應用程式和測試資料。你可以自在的調整你在 `/test` 目錄中的結構。對於較大的專案來說，通常會有一個 `data` 資料夾也是蠻正常的。例如：如果你需要 Go 忽略這些目錄下的檔案，你可以使用 `/test/data` 或 `/test/testdata` 當作你的目錄名稱。請注意：Go 還會忽略以 `.` 或 `_` 開頭的目錄或檔案，所以你在測試資料的目錄命名上，將擁有更大的彈性。

## **其他目錄**

### **`/docs`**

設計和使用者文件 (除了 `godoc` 自動產生的文件之外)。

### **`/tools`**

這個專案的支援工具。注意：這些工具可以從 `/pkg` 和 `/internal` 目錄匯入程式碼。

### **`/examples`**

放置關於你的應用程式或公用函式庫的使用範例。

### **`/third_party`**

外部輔助工具、Forked 程式碼，以及其他第三方工具 (例如：Swagger UI)。

### **`/githooks`**

Git hooks。

### **`/assets`**

其他要一併放入儲存庫的相關檔案 (例如圖片、Logo、... 等等)。

### **`/website`**

如果你不使用 GitHub Pages 的話，這裡可以放置專案的網站相關資料。

## 參考資料：

[https://github.com/golang-standards/project-layout/blob/master/README_zh.md](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)

[https://lailin.xyz/post/go-training-week4-project-layout.html](https://lailin.xyz/post/go-training-week4-project-layout.html)