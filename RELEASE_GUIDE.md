# Glowllama Release Guide

## Overview
This guide covers how to publish Glowllama binaries to GitHub Releases using the existing GitHub Actions workflow.

## Prerequisites

### 1. GitHub Repository Setup
- Repository must be public or you have a paid GitHub plan for private releases
- Repository must have GitHub Actions enabled
- You need admin access to configure secrets and variables

### 2. Required Accounts & Services
- **GitHub Account** - For releases
- **Apple Developer Account** - For macOS code signing (optional but recommended)
- **Windows Code Signing Certificate** - For Windows binaries (optional)

## GitHub Secrets & Variables Setup

### Navigate to Repository Settings
```
Your Repo → Settings → Secrets and variables → Actions
```

### Required Secrets (Minimal Setup)

For a **basic release without code signing**, you only need:

#### ✅ No secrets required for basic unsigned releases!

The workflow uses `${{ github.token }}` which is automatically provided by GitHub Actions.

### Optional Secrets (For Code Signing)

#### macOS Code Signing (Optional)
Add these under **Secrets** tab:

1. **`APPLE_IDENTITY`** - Your Apple Developer Team ID
   - Example: `ABCD123456`
   - Find at: https://developer.apple.com/account

2. **`APPLE_PASSWORD`** - App-specific password
   - Generate at: https://appleid.apple.com/account/manage
   - Create under "App-Specific Passwords"

3. **`MACOS_SIGNING_KEY`** - Base64-encoded P12 certificate
   ```bash
   # Export from Keychain as .p12, then:
   base64 -i certificate.p12 | pbcopy
   # Paste the output as the secret value
   ```

4. **`MACOS_SIGNING_KEY_PASSWORD`** - Password for the P12 file

#### macOS Code Signing Variables
Add these under **Variables** tab:

1. **`APPLE_TEAM_ID`** - Your Apple Developer Team ID (same as APPLE_IDENTITY)
2. **`APPLE_ID`** - Your Apple ID email

#### Windows Code Signing (Optional)

1. **`WINDOWS_SIGNING_KEY`** - Base64-encoded PFX certificate
2. **`WINDOWS_SIGNING_KEY_PASSWORD`** - Password for the PFX file

### Environment Setup

The workflow uses a `release` environment. Create it:

1. Go to: **Settings → Environments**
2. Click **New environment**
3. Name it: `release`
4. (Optional) Add protection rules:
   - Required reviewers
   - Deployment branches (only allow tags matching `v*`)

## How to Create a Release

### Step 1: Update Version (if needed)
```bash
# Already set to 1.0.0-alpha in version/version.go
# For new releases, update:
vim version/version.go
```

### Step 2: Commit and Push Changes
```bash
git add .
git commit -m "🚀 release: prepare v1.0.0-alpha"
git push origin main
```

### Step 3: Create and Push a Tag
```bash
# Create a tag (must start with 'v')
git tag v1.0.0-alpha

# Push the tag to trigger the release workflow
git push origin v1.0.0-alpha
```

### Step 4: Monitor the Release

1. Go to: **Actions** tab in your repository
2. You'll see the `release` workflow running
3. Monitor the progress of each job:
   - `darwin-build` - macOS binaries
   - `windows-depends` & `windows-app` - Windows binaries
   - `linux-build` - Linux binaries
   - `release` - Creates GitHub Release

### Step 5: Publish the Release

After the workflow completes:

1. Go to: **Releases** tab
2. Find the draft release for `v1.0.0-alpha`
3. Review the auto-generated release notes
4. Edit the description if needed
5. Uncheck "Set as a pre-release" if it's a stable release
6. Click **Publish release**

## What Gets Built

The release workflow builds:

### macOS (darwin)
- `glowllama-darwin-amd64.tgz` - Intel Macs
- `glowllama-darwin-arm64.tgz` - Apple Silicon
- `Glowllama-darwin.dmg` - Installer (if code signed)
- `Glowllama-darwin.zip` - Zip archive (if code signed)

### Windows
- `GlowllamaSetup.exe` - Installer
- `glowllama-windows-amd64.zip` - CPU-only
- `glowllama-windows-amd64-cuda12.zip` - CUDA 12 support
- `glowllama-windows-amd64-cuda13.zip` - CUDA 13 support

### Linux
- `glowllama-linux-amd64.tgz` - CPU-only
- `glowllama-linux-amd64-cuda12.tgz` - CUDA 12
- `glowllama-linux-amd64-cuda13.tgz` - CUDA 13
- `glowllama-linux-arm64.tgz` - ARM64 (Raspberry Pi, etc.)
- `glowllama-linux-amd64-rocm.tgz` - AMD ROCm support

### Checksums
- `sha256sum.txt` - SHA256 checksums for all files

## Quick Start (Without Code Signing)

**Minimum steps to get a release out:**

```bash
# 1. Make sure your changes are committed
git add .
git commit -m "🚀 release: v1.0.0-alpha"
git push origin main

# 2. Create and push tag
git tag v1.0.0-alpha
git push origin v1.0.0-alpha

# 3. Wait for GitHub Actions to complete (~30-60 min)

# 4. Go to GitHub Releases and publish the draft
```

No secrets needed! The workflow will build unsigned binaries.

## Common Issues & Solutions

### Issue: "Resource not accessible by integration"
**Solution**: Check that the `release` environment exists and workflow has `contents: write` permission (already configured).

### Issue: macOS binaries won't run (unsigned)
**Solution**: Users need to run:
```bash
xattr -cr /path/to/glowllama
```
Or add code signing secrets.

### Issue: Windows "Unknown Publisher" warning
**Solution**: Add Windows code signing certificate or users click "More info" → "Run anyway".

### Issue: Workflow doesn't trigger
**Solution**: 
- Tag must start with `v` (e.g., `v1.0.0`, not `1.0.0`)
- Check GitHub Actions are enabled
- Check workflow file has no syntax errors

## Testing Before Release

### Test build locally (macOS):
```bash
./scripts/build_darwin.sh
ls -l dist/
```

### Test build locally (Linux):
```bash
./scripts/build_linux.sh
ls -l dist/
```

### Test the binary:
```bash
./glowllama --version
./glowllama run --help
```

## Release Checklist

- [ ] All changes committed and pushed
- [ ] Version updated in `version/version.go`
- [ ] CHANGELOG updated (if you maintain one)
- [ ] Tests passing locally
- [ ] Tag created: `v1.0.0-alpha`
- [ ] Tag pushed to GitHub
- [ ] GitHub Actions workflow started
- [ ] Workflow completed successfully
- [ ] Draft release created
- [ ] Release notes reviewed
- [ ] Release published

## Advanced: Custom Workflows

### Release to Multiple Registries

You can extend the workflow to also push to:
- **Docker Hub** - Add Docker build/push step
- **Homebrew** - Create a tap repository
- **AUR** (Arch Linux) - Create PKGBUILD

### Automated Releases

Add to workflow to auto-publish:
```yaml
- name: Publish release
  run: gh release edit ${GITHUB_REF_NAME} --draft=false
```

## Security Best Practices

1. **Never commit secrets** - Use GitHub Secrets
2. **Use environment protection** - Require manual approval
3. **Rotate certificates** - Update annually
4. **Minimize permissions** - Only grant what's needed
5. **Review dependencies** - Check for vulnerabilities with `go mod tidy`

## Summary

**Minimal Setup (Unsigned Binaries):**
- ✅ No secrets required
- ✅ No special configuration
- ✅ Just push a tag starting with `v`
- ✅ GitHub Actions builds everything
- ✅ Publish from Releases tab

**Full Setup (Signed Binaries):**
- Add Apple Developer secrets (macOS signing)
- Add Windows certificate secrets (Windows signing)
- Configure `release` environment
- Same release process

---

**Ready to release?** Just run:
```bash
git tag v1.0.0-alpha && git push origin v1.0.0-alpha
```

Then go to the **Releases** tab after ~30-60 minutes and publish! 🚀
