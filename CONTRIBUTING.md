# Contributing to Corazon

Thank you for considering contributing to Corazon! This document outlines the process for contributing to the project.

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for everyone.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the [issue tracker](https://github.com/yourusername/corazon/issues) to avoid duplicates.

When creating a bug report, include:

- **Clear title and description**
- **Steps to reproduce** the issue
- **Expected behavior** vs actual behavior
- **System information** (OS, Go version)
- **Extension structure** that caused the issue (if applicable)
- **Error messages** or logs

Example:
```markdown
**Description**: Corazon fails to package extensions with symlinks

**Steps to Reproduce**:
1. Create extension with symlinked file: `ln -s target.js link.js`
2. Run `crx ./extension output.zip`
3. Error occurs

**Expected**: Should follow symlinks and package target file
**Actual**: Error message: "failed to add link.js: ..."

**System**: macOS 14.2, Go 1.21.5
```

### Suggesting Features

Feature suggestions are welcome! Please:

1. Check if the feature already exists or has been requested
2. Explain the problem your feature would solve
3. Describe the proposed solution
4. Consider alternative solutions
5. Include example use cases

### Pull Requests

1. **Fork the repository** and create a branch from `main`
2. **Make your changes** following the coding standards
3. **Add tests** if you're adding functionality
4. **Update documentation** if you're changing behavior
5. **Run tests** to ensure everything passes
6. **Submit a pull request** with a clear description

#### Pull Request Guidelines

- Use clear, descriptive commit messages
- Keep changes focused (one feature/fix per PR)
- Update README.md if adding user-facing features
- Add comments for complex logic
- Follow Go conventions and idioms

#### Commit Message Format

```
<type>: <description>

[optional body]

[optional footer]
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Examples:
```
feat: add support for .npmignore files

fix: handle symlinks in extension directories

docs: update installation instructions for Windows
```

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git

### Setup Steps

```bash
# Clone your fork
git clone https://github.com/yourusername/corazon.git
cd corazon

# Add upstream remote
git remote add upstream https://github.com/originaluser/corazon.git

# Create a branch
git checkout -b feature/my-feature

# Make changes and commit
git add .
git commit -m "feat: add my feature"

# Push to your fork
git push origin feature/my-feature
```

### Building

```bash
# Build the binary
go build -o crx .

# Run the binary
./crx --help
```

### Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detection
go test -race ./...
```

### Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` to format code
- Use `go vet` to check for common issues
- Keep functions focused and concise
- Add comments for exported functions and complex logic

Example:
```go
// PackageExtension creates a ZIP archive of the browser extension.
// It respects ignore rules unless packageAll is true.
func PackageExtension(sourceDir, outputPath string, packageAll bool) error {
    // Implementation
}
```

## Project Structure

```
corazon/
├── main.go              # Main application logic
├── README.md            # Project documentation
├── LICENSE              # MIT License
├── CONTRIBUTING.md      # This file
├── .github/
│   └── workflows/
│       └── release.yml  # Release automation
└── go.mod               # Go module file
```

## Testing Checklist

Before submitting a PR, ensure:

- [ ] Code builds without errors
- [ ] All tests pass
- [ ] New features have tests
- [ ] Documentation is updated
- [ ] Code follows project style
- [ ] Commit messages are clear
- [ ] No sensitive information is included

## Areas for Contribution

We welcome contributions in these areas:

- 🐛 **Bug fixes** - Fix reported issues
- ✨ **Features** - Add new ignore patterns or options
- 📚 **Documentation** - Improve guides and examples
- 🧪 **Tests** - Add test coverage
- 🚀 **Performance** - Optimize packaging speed
- 🌍 **Localization** - Add translations (if applicable)

## Questions?

If you have questions:

1. Check the [README](README.md)
2. Search [existing issues](https://github.com/yourusername/corazon/issues)
3. Open a new [discussion](https://github.com/yourusername/corazon/discussions)

## Recognition

Contributors will be:
- Listed in release notes
- Mentioned in the README (for significant contributions)
- Credited in commit history

Thank you for contributing to Corazon! 🫀