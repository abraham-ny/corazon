# Corazon 🫀

A fast, cross-platform browser extension packaging tool written in Go.

[![Release](https://img.shields.io/github/v/release/abraham-ny/corazon)](https://github.com/abraham-ny/corazon/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/abraham-ny/corazon)](go.mod)

## Features

- **Fast and lightweight** - Single binary with no dependencies
- **Smart ignore patterns** - Respects `.gitignore` and custom `.zignore` files
- **Security-first** - Automatically excludes sensitive files (`.env`, `*.key`, etc.)
- **Cross-platform** - Works on Linux, macOS, Windows, and FreeBSD
- **Simple CLI** - Intuitive command-line interface
- **Flexible** - Package everything or use intelligent filtering

## Installation

### Download Pre-built Binaries

Download the latest release for your platform from the [releases page](https://github.com/yourusername/corazon/releases).

#### Linux/macOS
```bash
# Download and extract
tar -xzf crx-*.tar.gz

# Move to PATH
sudo mv crx /usr/local/bin/

# Verify installation
crx -v
```

#### Windows
1. Download the `.zip` file for your architecture
2. Extract `crx.exe`
3. Add to your PATH or move to a directory in your PATH
4. Verify: `crx -v`

### Build from Source

```bash
# Clone repository
git clone https://github.com/yourusername/corazon.git
cd corazon

# Build
go build -o crx .

# Install (optional)
go install
```

## Usage

### Basic Usage

```bash
# Package extension with ignore rules
crx <source_directory> <output_file>

# Examples
crx ./my-extension extension.zip
crx ./my-extension extension.crx
```

### Package Everything (Bypass Ignore Rules)

```bash
# Include all files, even those in .gitignore/.zignore
crx -a ./my-extension extension-full.zip
```

### Options

```
-a          Package all files, ignore .gitignore and .zignore
-v, -version Show version information
```

## Ignore Rules

Corazon respects multiple ignore file formats to help you exclude unnecessary or sensitive files:

### Default Ignores

The following files and patterns are ignored by default:
- `.git/` - Git repository data
- `.gitignore` - Git ignore file
- `.zignore` - Corazon ignore file
- `node_modules/` - Node.js dependencies
- `.DS_Store` - macOS metadata
- `Thumbs.db` - Windows thumbnails
- `*.log` - Log files
- `.env`, `.env.local` - Environment variables
- `*.pem`, `*.key` - Private keys

### Custom Ignore Rules

#### `.gitignore`
Corazon automatically reads `.gitignore` files in your extension directory.

#### `.zignore`
Create a `.zignore` file for Corazon-specific ignore rules. Uses the same syntax as `.gitignore`:

```gitignore
# Ignore source files
src/
*.ts

# Ignore documentation
docs/
README.md

# But include the built output
!dist/

# Ignore test files
**/*test.js
*.spec.js
```

### Ignore Pattern Syntax

- `filename` - Ignore specific file
- `directory/` - Ignore entire directory
- `*.ext` - Wildcard for file extensions
- `**/pattern` - Match in any directory
- `!pattern` - Negate (don't ignore)
- `#comment` - Comments

## Examples

### Basic Chrome Extension

```bash
# Project structure
my-extension/
├── manifest.json
├── popup.html
├── popup.js
├── icon.png
├── src/           # TypeScript sources
└── node_modules/  # Dependencies

# Package (excludes src/ and node_modules/)
crx ./my-extension extension.zip
```

### Include Everything

```bash
# Package including all sources and dependencies
crx -a ./my-extension extension-dev.zip
```

### Custom Ignore Rules

Create `.zignore`:
```gitignore
# Ignore all markdown except README
*.md
!README.md

# Ignore screenshots
screenshots/
```

Then package:
```bash
crx ./my-extension extension.zip
```

## Output Formats

Both `.zip` and `.crx` extensions produce valid ZIP archives that work as browser extensions:

```bash
crx ./extension extension.zip  # Standard ZIP
crx ./extension extension.crx  # Chrome extension format
```

## Requirements

- **Source directory must contain `manifest.json`**
- Go 1.21+ (for building from source)

## Platform Support

| Platform | Architecture | Status |
|----------|-------------|---------|
| Linux    | amd64       | ✅      |
| Linux    | arm64       | ✅      |
| Linux    | 386         | ✅      |
| macOS    | amd64       | ✅      |
| macOS    | arm64       | ✅      |
| Windows  | amd64       | ✅      |
| Windows  | 386         | ✅      |
| Windows  | arm64       | ✅      |
| FreeBSD  | amd64       | ✅      |

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

Corazon is released under the [MIT License](LICENSE).

## Support

- **Bug reports**: [Issue tracker](https://github.com/abraham-ny/corazon/issues)
- **Feature requests**: [Issue tracker](https://github.com/abraham-ny/corazon/issues)
- **Documentation**: [Wiki](https://github.com/abraham-ny/corazon/wiki)

## Acknowledgments

Built using Go's standard library.

---