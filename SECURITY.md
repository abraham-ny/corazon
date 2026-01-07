# Security Policy

## Supported Versions

We release patches for security vulnerabilities in the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take the security of Corazon seriously. If you believe you have found a security vulnerability, please report it to us responsibly.

### Please Do Not
- Open a public GitHub issue for security vulnerabilities
- Disclose the vulnerability publicly before it has been addressed

### How to Report

**Email**: security@yourdomain.com (if you have one)

**GitHub Security Advisory**: Use GitHub's [private vulnerability reporting](https://github.com/yourusername/corazon/security/advisories/new)

### What to Include

Please include as much of the following information as possible:

- Type of vulnerability (e.g., file inclusion, path traversal, etc.)
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if applicable)
- Impact of the vulnerability
- Suggested fix (if you have one)

### What to Expect

1. **Acknowledgment**: We'll acknowledge receipt of your report within 48 hours
2. **Assessment**: We'll assess the vulnerability and determine its severity
3. **Fix Development**: We'll work on a fix for confirmed vulnerabilities
4. **Disclosure**: We'll coordinate disclosure timing with you
5. **Credit**: We'll credit you in the security advisory (unless you prefer to remain anonymous)

### Response Timeline

- **Critical vulnerabilities**: Patch within 7 days
- **High severity**: Patch within 30 days
- **Medium/Low severity**: Patch in next release cycle

## Security Best Practices

When using Corazon:

### 1. Verify Downloads
Always verify the SHA256 checksum of downloaded binaries:

```bash
# Linux/macOS
sha256sum crx-linux-amd64.tar.gz
cat crx-linux-amd64.tar.gz.sha256

# Windows
certutil -hashfile crx-windows-amd64.zip SHA256
type crx-windows-amd64.zip.sha256
```

### 2. Use Ignore Files
Always use `.gitignore` or `.zignore` to exclude sensitive files:

```gitignore
# Sensitive files
.env
.env.*
*.key
*.pem
config/secrets.json
credentials/
```

### 3. Review Package Contents
Before distributing extensions, verify what's included:

```bash
# Extract and inspect
unzip -l extension.zip
# or
tar -tzf extension.tar.gz
```

### 4. Use -a Flag Carefully
The `-a` flag bypasses all ignore rules. Only use it when you specifically need to package everything:

```bash
# This will include EVERYTHING, including sensitive files
crx -a ./extension extension.zip
```

### 5. Keep Corazon Updated
Always use the latest version to ensure you have the latest security patches:

```bash
crx -v  # Check your version
```

## Known Security Considerations

### Path Traversal
Corazon validates paths to prevent directory traversal attacks. Files outside the source directory cannot be included in packages.

### Symbolic Links
Corazon follows symbolic links but validates that they point to files within the source directory tree.

### File Permissions
Corazon preserves file permissions in the archive. Ensure sensitive files have appropriate permissions before packaging.

## Security Updates

Security updates are announced via:
- GitHub Security Advisories
- Release notes with `[SECURITY]` tag
- GitHub releases page

Subscribe to the repository to receive notifications.

## Scope

This security policy applies to:
- The Corazon command-line tool
- Official release binaries
- Source code in this repository

Out of scope:
- Third-party forks
- Modified versions
- Extensions packaged by Corazon (these are user responsibility)

## Contact

For general security questions: security@yourdomain.com
For vulnerability reports: Use GitHub Security Advisory

Thank you for helping keep Corazon and its users safe!