# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 0.x     | :white_check_mark: |

## Reporting a Vulnerability

If you discover a security vulnerability in safeconv, please report it responsibly.

### How to Report

1. **Do not** open a public GitHub issue for security vulnerabilities
2. Email the maintainers directly (see repository owner's profile for contact info)
3. Or use GitHub's [private vulnerability reporting](https://github.com/heisenbergs-uncertainty/safeconv/security/advisories/new)

### What to Include

- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

### Response Timeline

- **Initial response**: Within 48 hours
- **Status update**: Within 7 days
- **Fix timeline**: Depends on severity, typically within 30 days

### After Reporting

- We will acknowledge receipt of your report
- We will investigate and determine the impact
- We will work on a fix and coordinate disclosure
- We will credit you in the security advisory (unless you prefer anonymity)

## Security Considerations

safeconv is designed to prevent integer overflow vulnerabilities (gosec G115, G109). However:

- **Performance-critical paths**: If you've validated inputs externally, you may use Go's built-in conversions directly in hot paths
- **Must variants**: `Must*` functions panic on error - use them only for initialization or when failure indicates a programming error, not for untrusted input

## Scope

This security policy covers:

- The safeconv Go library
- Security vulnerabilities in conversion logic
- Incorrect error handling that could lead to silent data corruption

Out of scope:

- Vulnerabilities in dependencies (we have none)
- Issues in applications using safeconv
