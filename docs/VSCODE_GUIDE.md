# VS Code Quick Start Guide

## 🚀 Getting Started

### Opening the Project

```bash
code .
```

### First-Time Setup

When you open VS Code for the first time, you'll be prompted to install recommended extensions:

1. **Go** (`golang.go`) - Essential for Go development
2. **Docker** (`ms-azuretools.vscode-docker`) - Docker support
3. **REST Client** (`humao.rest-client`) - Test API endpoints

Click "Install All" or install individually.

---

## 🎯 Quick Actions

### Run Application

**Method 1: Debug (F5)**

- Press `F5` to run with debugger attached
- MongoDB starts automatically
- Set breakpoints by clicking left of line numbers
- Use Debug Console to inspect variables

**Method 2: Tasks Menu**

- `Ctrl+Shift+P` → Type "Tasks: Run Task"
- Select "Run Application"
- Background task (check Terminal panel)

**Method 3: Hot Reload**

- `Ctrl+Shift+P` → "Tasks: Run Task" → "Run with Hot Reload"
- Code changes trigger automatic restart
- Faster development iteration

### Run Tests

**Method 1: Command Palette**

- `Ctrl+Shift+P` → "Go: Test Package"
- Tests current package

**Method 2: Tasks**

- `Ctrl+Shift+P` → "Tasks: Run Task" → "Run Tests"
- Runs all tests with race detector

**Method 3: Code Lens**

- Look for "run test" | "debug test" above test functions
- Click to run individual tests
- Great for TDD workflow

### Debug Tests

**Method 1: Current Test**

1. Open test file
2. Place cursor on test function name
3. `F5` or Debug sidebar → "Debug Current Test"

**Method 2: Code Lens**

- Click "debug test" above test function
- Breakpoints work in tests too!

---

## ⚡ Keyboard Shortcuts

### Running & Debugging

| Action | Shortcut |
|--------|----------|
| Start Debugging | `F5` |
| Run without Debugging | `Ctrl+F5` |
| Stop Debugging | `Shift+F5` |
| Restart Debugging | `Ctrl+Shift+F5` |
| Toggle Breakpoint | `F9` |
| Step Over | `F10` |
| Step Into | `F11` |
| Step Out | `Shift+F11` |
| Continue | `F5` |

### Code Navigation

| Action | Shortcut |
|--------|----------|
| Go to Definition | `F12` |
| Peek Definition | `Alt+F12` |
| Go to References | `Shift+F12` |
| Go to Symbol | `Ctrl+Shift+O` |
| Go to File | `Ctrl+P` |
| Command Palette | `Ctrl+Shift+P` |

### Editing

| Action | Shortcut |
|--------|----------|
| Format Document | `Shift+Alt+F` |
| Organize Imports | `Ctrl+Shift+O` (auto on save) |
| Rename Symbol | `F2` |
| Quick Fix | `Ctrl+.` |
| Show Hover | `Ctrl+K Ctrl+I` |

### Testing

| Action | Shortcut |
|--------|----------|
| Run Tests in File | `Ctrl+Shift+P` → "Go: Test File" |
| Run Test at Cursor | `Ctrl+Shift+P` → "Go: Test Function At Cursor" |
| Toggle Test Coverage | `Ctrl+Shift+P` → "Go: Toggle Test Coverage" |

---

## 📂 Workspace Layout

### Recommended Panel Layout

**Left Sidebar:**

- Explorer (files)
- Search
- Source Control (Git)
- Debug
- Extensions

**Bottom Panel:**

- Terminal (for command output)
- Problems (errors/warnings)
- Output (task output)
- Debug Console (during debugging)

**Tip:** Use `Ctrl+B` to toggle sidebar, `Ctrl+J` to toggle panel.

---

## 🔧 Available Tasks

### Build Tasks

```
Build Application          - Compile to bin/cadastro-pessoas
Clean Build               - Remove old binaries and rebuild
Build Docker Image        - Create Docker image
```

### Run Tasks

```
Run Application           - Start with MongoDB
Run with Hot Reload       - Auto-restart on changes
Run Docker Container      - Run in Docker
```

### Test Tasks

```
Run Tests                 - All tests with race detector
Run Tests with Coverage   - Generate coverage report
Run Integration Tests     - Full integration suite
Run Linter                - golangci-lint checks
```

### Utility Tasks

```
Format Code               - gofmt + goimports
Download Dependencies     - go mod download & tidy
```

### Docker Tasks

```
docker:start-mongodb      - Start MongoDB container
docker:stop-mongodb       - Stop all containers
docker:logs-mongodb       - View MongoDB logs
```

---

## 🐛 Debug Configurations

### 1. Launch Application

- **Use when:** Starting app for first time
- **Features:** Full debugger, MongoDB auto-start
- **Breakpoints:** ✓
- **Hot Reload:** ✗

### 2. Launch Application (No Docker)

- **Use when:** MongoDB already running
- **Features:** Faster startup
- **Breakpoints:** ✓
- **Hot Reload:** ✗

### 3. Debug Current Test

- **Use when:** Debugging single test
- **Features:** Isolated test execution
- **Breakpoints:** ✓
- **Tip:** Place cursor on test function name first

### 4. Debug All Tests

- **Use when:** Debugging test suite
- **Features:** All tests with debugger
- **Breakpoints:** ✓

### 5. Debug Integration Tests

- **Use when:** Testing with real MongoDB
- **Features:** MongoDB auto-start, test environment
- **Breakpoints:** ✓

### 6. Attach to Process

- **Use when:** App already running
- **Features:** Attach debugger to live process
- **Breakpoints:** ✓ (in running code)

---

## 🔍 IntelliSense & Code Assistance

### Auto-Complete

Type and press `Ctrl+Space` to trigger:

- Function signatures
- Struct fields
- Import suggestions
- Code snippets

### Code Actions (Quick Fixes)

Place cursor on error/warning:

- Press `Ctrl+.` for quick fixes
- Extract function
- Generate interface implementation
- Add missing imports

### Hover Information

Hover over any symbol or press `Ctrl+K Ctrl+I`:

- Function documentation
- Type information
- Parameter hints

---

## 📝 Snippets

Type these prefixes and press `Tab`:

```go
// main - main function
func main() {
    $0
}

// func - function declaration
func $1($2) $3 {
    $0
}

// if - if statement
if $1 {
    $0
}

// for - for loop
for $1 := 0; $1 < $2; $1++ {
    $0
}

// test - test function
func Test$1(t *testing.T) {
    $0
}
```

---

## 🧪 Testing Workflow (TDD)

### Red-Green-Refactor Cycle

1. **Write Test** (RED)

   ```
   - Create test file: *_test.go
   - Write failing test
   - Click "run test" → should FAIL
   ```

2. **Write Code** (GREEN)

   ```
   - Implement minimal code
   - Click "run test" → should PASS
   ```

3. **Refactor** (REFACTOR)

   ```
   - Improve code quality
   - Click "run test" → still PASS
   ```

### Test Coverage View

1. `Ctrl+Shift+P` → "Go: Toggle Test Coverage"
2. Green = covered
3. Red = not covered
4. Aim for 80%+ coverage

---

## 🐳 Docker Integration

### Start MongoDB

**Method 1: Automatic**

- Press `F5` - MongoDB starts automatically

**Method 2: Manual**

- `Ctrl+Shift+P` → "Tasks: Run Task" → "docker:start-mongodb"

### View Logs

- `Ctrl+Shift+P` → "Tasks: Run Task" → "docker:logs-mongodb"
- Or use Docker extension sidebar

### Stop Containers

- `Ctrl+Shift+P` → "Tasks: Run Task" → "docker:stop-mongodb"

---

## 🎨 Customization

### Settings

Open settings: `Ctrl+,`

**Popular Go Settings:**

```json
{
    "go.formatTool": "goimports",      // Auto-fix imports
    "go.lintTool": "golangci-lint",    // Use golangci-lint
    "go.testFlags": ["-v", "-race"],   // Race detector
    "editor.formatOnSave": true,       // Auto-format
    "go.coverOnSave": false            // Coverage on save
}
```

### Themes

- `Ctrl+K Ctrl+T` - Change color theme
- Popular: Dark+ (default), Monokai, Solarized

### Extensions

Recommended extensions already configured in `.vscode/extensions.json`:

- Install from Extensions sidebar (`Ctrl+Shift+X`)

---

## 🚨 Troubleshooting

### "Go tools are missing"

**Solution:**

```bash
# VS Code will prompt - click "Install All"
# Or manually:
Ctrl+Shift+P → "Go: Install/Update Tools"
```

### MongoDB connection failed

**Check:**

```bash
# Is MongoDB running?
docker ps | grep mongo

# Start manually:
docker-compose up -d mongodb
```

### Debugger won't start

**Check:**

1. `go.mod` exists?
2. No compilation errors? (check Problems panel)
3. Port 8080 free? (`lsof -i :8080`)

### Hot reload not working

**Check:**

1. Air installed? (`which air`)
2. `.air.toml` exists?
3. Correct directory? (workspace root)

### Tests not discovered

**Solution:**

```bash
# Reload Go extension
Ctrl+Shift+P → "Go: Restart Language Server"
```

---

## 📚 Resources

### Official Docs

- [Go in VS Code](https://code.visualstudio.com/docs/languages/go)
- [Debugging in VS Code](https://code.visualstudio.com/docs/editor/debugging)
- [Tasks in VS Code](https://code.visualstudio.com/docs/editor/tasks)

### Project Docs

- `README.md` - Project overview
- `CONTRIBUTING.md` - Contribution guidelines
- `AGENTS.md` - Development guide
- `.specify/memory/constitution.md` - Code principles

### Extension Docs

- [Go Extension](https://github.com/golang/vscode-go)
- [Docker Extension](https://github.com/microsoft/vscode-docker)

---

## 💡 Tips & Tricks

### Multi-Cursor Editing

- `Alt+Click` - Add cursor
- `Ctrl+Alt+Down/Up` - Add cursor below/above
- `Ctrl+D` - Select next occurrence

### Split Editors

- `Ctrl+\` - Split editor
- `Ctrl+1/2/3` - Focus editor group
- Drag tabs to split

### Terminal

- `Ctrl+`` - Toggle terminal
- `Ctrl+Shift+`` - New terminal
- Multiple terminals for different tasks

### Zen Mode

- `Ctrl+K Z` - Distraction-free coding
- `Esc Esc` - Exit Zen mode

### Git Integration

- `Ctrl+Shift+G` - Source Control panel
- Click files to see diff
- Stage with `+` icon
- Commit with `Ctrl+Enter`

---

## ⚡ Power User Workflow

### Morning Startup

```
1. Open VS Code: code .
2. Pull latest: Ctrl+Shift+G → Pull
3. Start MongoDB: Ctrl+Shift+P → docker:start-mongodb
4. Run tests: Ctrl+Shift+P → Run Tests
5. Start coding: F5 (debug mode)
```

### Feature Development

```
1. Create branch: Ctrl+Shift+G → Create Branch
2. Write test (RED): test_*.go
3. Run test: Click "run test" → FAIL
4. Write code (GREEN): *.go
5. Run test: Click "run test" → PASS
6. Refactor: Format with Shift+Alt+F
7. Run test: Still PASS
8. Commit: Ctrl+Shift+G → Commit
```

### Before Push

```
1. Run all tests: Ctrl+Shift+P → Run Tests
2. Run linter: Ctrl+Shift+P → Run Linter
3. Check coverage: Ctrl+Shift+P → Run Tests with Coverage
4. Review changes: Ctrl+Shift+G
5. Push: Ctrl+Shift+G → Push
```

---

**Happy Coding! 🎉**

For questions, check `AGENTS.md` or team documentation.
