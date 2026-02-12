# ANALISIS & PLANNING: Dashboard Monitoring API

**Tanggal:** 12 Februari 2026  
**Project:** Winbu.NET API Scraper  
**Status:** BELUM ADA DASHBOARD

---

## 📊 ANALISIS KONDISI SAAT INI

### ✅ Yang Sudah Ada:
1. **Swagger UI** - Dokumentasi API interaktif di `/swagger/`
2. **Health Check** - Endpoint `/health` untuk monitoring basic
3. **CORS Support** - Sudah ada middleware CORS
4. **Dynamic Config** - Config sudah menggunakan environment variables
5. **Cache System** - Ada cache manager di `utils/cache.go`

### ❌ Yang Belum Ada:
1. **Dashboard UI** - Tidak ada interface monitoring visual
2. **Admin Panel** - Tidak ada panel untuk update config
3. **Metrics/Analytics** - Tidak ada tracking request, response time, dll
4. **Database** - Config masih hardcoded di code, tidak ada persistence
5. **Authentication** - Tidak ada auth untuk admin
6. **Logging Dashboard** - Logs hanya di console
7. **Alert System** - Tidak ada notifikasi jika API error

---

## 🎯 KEBUTUHAN DASHBOARD

### 1. **Configuration Management**
**Problem:** Base URL hardcoded di `config/config.go`
```go
BaseURL: getEnv("BASE_URL", "https://winbu.net"),
```

**Kebutuhan:**
- ✅ Update Base URL tanpa restart server
- ✅ Update scraping settings (timeout, rate limit, retry)
- ✅ Enable/disable cache
- ✅ History perubahan config

### 2. **API Monitoring**
**Kebutuhan:**
- Real-time API status (up/down)
- Request count per endpoint
- Response time average
- Error rate tracking
- Success rate per scraper

### 3. **Scraper Health Check**
**Kebutuhan:**
- Test scraper ke domain target
- Validate CSS selectors masih valid
- Alert jika struktur HTML berubah
- Test semua endpoint secara otomatis

### 4. **Data Analytics**
**Kebutuhan:**
- Total anime scraped
- Most requested endpoints
- Peak usage hours
- Cache hit/miss ratio
- Bandwidth usage

---

## 🏗️ ARSITEKTUR DASHBOARD

### Option A: **Embedded Dashboard (Recommended)**
**Stack:**
- Backend: Go (existing)
- Frontend: HTML + Vanilla JS + TailwindCSS
- Database: SQLite (lightweight, no external dependency)
- Auth: JWT token simple

**Pros:**
- ✅ Single binary deployment
- ✅ No external dependencies
- ✅ Lightweight
- ✅ Easy to maintain

**Cons:**
- ⚠️ Limited real-time features
- ⚠️ Simple UI capabilities

### Option B: **Separate Dashboard with Modern Stack**
**Stack:**
- Backend: Go API (existing) + Admin API
- Frontend: React/Vue/Svelte
- Database: PostgreSQL/MySQL
- Auth: OAuth2 or JWT

**Pros:**
- ✅ Modern UI/UX
- ✅ Rich features
- ✅ Scalable

**Cons:**
- ❌ Complex deployment
- ❌ More dependencies
- ❌ Higher resource usage

### Option C: **Hybrid - Go + HTMX**
**Stack:**
- Backend: Go (existing)
- Frontend: HTMX + AlpineJS + TailwindCSS
- Database: SQLite
- Auth: Session-based

**Pros:**
- ✅ Modern UX with minimal JS
- ✅ Server-side rendering
- ✅ Easy to maintain
- ✅ Lightweight

**Cons:**
- ⚠️ Learning curve for HTMX

---

## 📋 FITUR DASHBOARD - DETAIL SPEC

### 1. **Dashboard Home**
```
┌─────────────────────────────────────────┐
│ Winbu.NET API Dashboard                 │
├─────────────────────────────────────────┤
│ ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐   │
│ │ API  │ │Total │ │Error │ │Uptime│   │
│ │Status│ │Req   │ │Rate  │ │99.9% │   │
│ │  🟢  │ │12.5K │ │ 0.1% │ │      │   │
│ └──────┘ └──────┘ └──────┘ └──────┘   │
├─────────────────────────────────────────┤
│ Endpoint Performance (Last 24h)         │
│ ┌───────────────────────────────────┐   │
│ │ /anime-terbaru    [████░] 85ms   │   │
│ │ /movie            [███░░] 92ms   │   │
│ │ /home             [██░░░] 120ms  │   │
│ │ /search           [█████] 65ms   │   │
│ └───────────────────────────────────┘   │
├─────────────────────────────────────────┤
│ Recent Errors (Last 1h)                 │
│ • None - All systems operational! ✓     │
└─────────────────────────────────────────┘
```

### 2. **Configuration Panel**
```
┌─────────────────────────────────────────┐
│ Configuration Management                │
├─────────────────────────────────────────┤
│ Target Website Settings                 │
│ ┌─────────────────────────────────────┐ │
│ │ Base URL:                           │ │
│ │ [https://winbu.net             ] 💾 │ │
│ │ Last Updated: 2 hours ago           │ │
│ │ Status: ✅ Reachable (45ms)         │ │
│ └─────────────────────────────────────┘ │
├─────────────────────────────────────────┤
│ Scraping Settings                       │
│ ┌─────────────────────────────────────┐ │
│ │ Timeout:      [30s ▼]              │ │
│ │ Rate Limit:   [1s  ▼]              │ │
│ │ Max Retries:  [3   ▼]              │ │
│ │ User Agent:   [Mozilla/5.0...    ] │ │
│ └─────────────────────────────────────┘ │
├─────────────────────────────────────────┤
│ Cache Settings                          │
│ ┌─────────────────────────────────────┐ │
│ │ Enable Cache: [✓] ON                │ │
│ │ Cache TTL:    [5m ▼]                │ │
│ │ Cache Size:   2.3 MB / 100 MB       │ │
│ │ [🗑️ Clear Cache]                    │ │
│ └─────────────────────────────────────┘ │
└─────────────────────────────────────────┘
```

### 3. **Health Check Panel**
```
┌─────────────────────────────────────────┐
│ Scraper Health Check                    │
├─────────────────────────────────────────┤
│ [🔄 Run All Tests]                      │
├─────────────────────────────────────────┤
│ ✅ Domain Accessibility                 │
│    https://winbu.net - 200 OK (45ms)    │
├─────────────────────────────────────────┤
│ ✅ Homepage Scraper                     │
│    Top 10: 10 items                     │
│    New Episodes: 20 items               │
│    Confidence: 1.00                     │
├─────────────────────────────────────────┤
│ ✅ Anime Terbaru Scraper                │
│    Page 1: 20 items                     │
│    Confidence: 1.00                     │
├─────────────────────────────────────────┤
│ ✅ Movie Scraper                        │
│    Page 1: 30 items                     │
│    Confidence: 1.00                     │
├─────────────────────────────────────────┤
│ ⚠️ Detail Scraper                       │
│    Warning: Slow response (2.5s)        │
│    Action: [View Details]               │
└─────────────────────────────────────────┘
```

### 4. **Analytics Dashboard**
```
┌─────────────────────────────────────────┐
│ Analytics & Insights                    │
├─────────────────────────────────────────┤
│ Request Trend (Last 7 days)             │
│ ┌─────────────────────────────────────┐ │
│ │  ▄                                  │ │
│ │ ▄█▄  ▄▄                             │ │
│ │ ███▄▄██▄                            │ │
│ │ ████████▄                           │ │
│ │ Mon Tue Wed Thu Fri Sat Sun         │ │
│ └─────────────────────────────────────┘ │
├─────────────────────────────────────────┤
│ Most Popular Endpoints                  │
│ 1. /anime-terbaru    (45%)              │
│ 2. /home             (25%)              │
│ 3. /movie            (18%)              │
│ 4. /search           (12%)              │
└─────────────────────────────────────────┘
```

---

## 💾 DATABASE SCHEMA

### Table: `config`
```sql
CREATE TABLE config (
    id INTEGER PRIMARY KEY,
    key VARCHAR(100) UNIQUE NOT NULL,
    value TEXT NOT NULL,
    description TEXT,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(100)
);

-- Indexes
CREATE INDEX idx_config_key ON config(key);
```

### Table: `api_metrics`
```sql
CREATE TABLE api_metrics (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    endpoint VARCHAR(255) NOT NULL,
    method VARCHAR(10) NOT NULL,
    status_code INTEGER NOT NULL,
    response_time_ms INTEGER NOT NULL,
    error_message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_metrics_endpoint ON api_metrics(endpoint, created_at);
CREATE INDEX idx_metrics_created ON api_metrics(created_at);
```

### Table: `health_checks`
```sql
CREATE TABLE health_checks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    scraper_name VARCHAR(100) NOT NULL,
    status VARCHAR(20) NOT NULL, -- success, warning, error
    items_found INTEGER,
    confidence_score REAL,
    response_time_ms INTEGER,
    error_message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_health_scraper ON health_checks(scraper_name, created_at);
```

### Table: `users` (Optional - untuk auth)
```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) DEFAULT 'admin',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## 🔐 DYNAMIC CONFIG IMPLEMENTATION

### Current Flow (Hardcoded):
```
main.go → config.Load() → getEnv() → Default: "https://winbu.net"
```

### New Flow (Dynamic):
```
main.go → config.LoadDynamic() 
         ↓
    Check Database
         ↓
    If exists: Use DB value
    If not: Use env/default → Save to DB
         ↓
    Return config
```

### API Endpoints untuk Config:
```
GET    /admin/config           - List all config
GET    /admin/config/:key      - Get specific config
PUT    /admin/config/:key      - Update config
POST   /admin/config/reload    - Reload config without restart
```

### Implementation:
```go
// config/dynamic_config.go
type DynamicConfig struct {
    db *sql.DB
}

func (dc *DynamicConfig) Get(key string) (string, error) {
    var value string
    err := dc.db.QueryRow(
        "SELECT value FROM config WHERE key = ?", 
        key,
    ).Scan(&value)
    return value, err
}

func (dc *DynamicConfig) Set(key, value string) error {
    _, err := dc.db.Exec(`
        INSERT INTO config (key, value, updated_at) 
        VALUES (?, ?, CURRENT_TIMESTAMP)
        ON CONFLICT(key) DO UPDATE SET 
            value = excluded.value,
            updated_at = CURRENT_TIMESTAMP
    `, key, value)
    return err
}
```

---

## 📦 IMPLEMENTATION PHASES

### **Phase 1: Foundation (Week 1)**
- [ ] Setup SQLite database
- [ ] Create database schema
- [ ] Implement dynamic config loading
- [ ] Add config API endpoints
- [ ] Basic authentication (username/password)

### **Phase 2: Dashboard UI (Week 2)**
- [ ] Create dashboard layout (HTML + TailwindCSS)
- [ ] Configuration management page
- [ ] Simple monitoring page
- [ ] Login page

### **Phase 3: Monitoring (Week 3)**
- [ ] Add metrics middleware
- [ ] Implement health check scheduler
- [ ] Create analytics endpoints
- [ ] Dashboard charts (Chart.js)

### **Phase 4: Polish (Week 4)**
- [ ] Add real-time updates (SSE/WebSocket)
- [ ] Alert system (email/webhook)
- [ ] Export reports (PDF/CSV)
- [ ] Documentation

---

## 🎨 TECHNOLOGY STACK RECOMMENDATION

### **Recommended: Option C (Hybrid)**

**Backend:**
- Go (existing codebase)
- SQLite (database/go-sqlite3)
- JWT for auth (golang-jwt/jwt)

**Frontend:**
- HTMX (hypermedia)
- AlpineJS (interactivity)
- TailwindCSS (styling)
- Chart.js (analytics charts)

**Why this stack:**
- ✅ No build step needed
- ✅ Single binary deployment
- ✅ Modern UX
- ✅ Easy maintenance
- ✅ Low resource usage

---

## 📂 NEW PROJECT STRUCTURE

```
winbutv/
├── main.go
├── config/
│   ├── config.go              # Existing
│   ├── dynamic_config.go      # NEW - Dynamic config loader
│   └── migrations.go          # NEW - DB migrations
├── dashboard/
│   ├── handlers.go            # NEW - Dashboard handlers
│   ├── middleware.go          # NEW - Auth middleware
│   ├── templates/             # NEW - HTML templates
│   │   ├── layout.html
│   │   ├── dashboard.html
│   │   ├── config.html
│   │   ├── health.html
│   │   └── login.html
│   └── static/                # NEW - Static assets
│       ├── css/
│       │   └── tailwind.min.css
│       └── js/
│           ├── htmx.min.js
│           ├── alpine.min.js
│           └── chart.min.js
├── database/
│   ├── db.go                  # NEW - Database connection
│   ├── queries.go             # NEW - SQL queries
│   └── schema.sql             # NEW - Database schema
├── metrics/
│   ├── collector.go           # NEW - Metrics collector
│   └── middleware.go          # NEW - Metrics middleware
└── winbu.db                   # NEW - SQLite database file
```

---

## 💰 ESTIMASI EFFORT

### Development Time:
- **Phase 1:** 3-5 days
- **Phase 2:** 4-6 days
- **Phase 3:** 3-5 days
- **Phase 4:** 2-3 days

**Total:** ~3-4 weeks for full dashboard

### Quick MVP (Minimum Viable Product):
Focus on Phase 1 + Basic UI from Phase 2
**Time:** ~1 week
**Features:**
- Dynamic config management
- Basic dashboard UI
- Config update without restart

---

## 🚀 QUICK START - MINIMAL DASHBOARD

Jika ingin cepat, bisa mulai dengan:

1. **Environment-based config** (Already supported!)
   ```bash
   # Update via environment variable
   export BASE_URL=https://new-domain.com
   # Restart app
   ```

2. **Simple Web UI for config** (3-4 hours work)
   - Single HTML page
   - Form to update env file
   - Restart app via API

3. **Docker with env file** (1-2 hours)
   - docker-compose.yml with env_file
   - Update .env without rebuild
   - Auto-restart on config change

---

## ✅ KESIMPULAN & REKOMENDASI

### Jawaban Pertanyaan Anda:

**Q1: Apakah sudah ada dashboard?**
**A:** ❌ Belum ada. Hanya ada Swagger UI untuk dokumentasi API.

**Q2: Bisa update URL tanpa hardcode?**
**A:** ✅ **BISA!** Ada 3 cara:

**Cara 1: Environment Variable (Sudah ada!)**
```bash
export BASE_URL=https://winbu.net
go run main.go
```

**Cara 2: Docker Compose + .env file**
```yaml
# docker-compose.yml
environment:
  - BASE_URL=${BASE_URL}
```

**Cara 3: Dashboard with Database (Perlu development)**
- Full dashboard dengan UI
- Update tanpa restart (hot reload)
- History tracking
- **Effort:** 1-4 minggu

### Rekomendasi Action:

**IMMEDIATE (Hari ini):**
1. ✅ Gunakan environment variable (sudah support!)
2. ✅ Buat file `.env` untuk config
3. ✅ Deploy dengan docker-compose

**SHORT TERM (1-2 minggu):**
1. 🔨 Build minimal dashboard (Phase 1 + 2)
2. 🔨 Dynamic config with SQLite
3. 🔨 Basic monitoring UI

**LONG TERM (1 bulan):**
1. 🎯 Full dashboard dengan analytics
2. 🎯 Alert system
3. 🎯 Auto health check

---

**Mau mulai dari mana?** 
1. Pakai env variable dulu (instant)?
2. Build minimal dashboard (1 minggu)?
3. Full-featured dashboard (3-4 minggu)?
