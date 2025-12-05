# 🚀 Glowllama Quick Start

## Release Without Code Signing (Easiest)

### No GitHub Secrets Needed! ✅

**1. Push your changes:**
```bash
git add .
git commit -m "🚀 release: v1.0.0-alpha"
git push origin main
```

**2. Create and push tag:**
```bash
git tag v1.0.0-alpha
git push origin v1.0.0-alpha
```

**3. Wait for build (~30-60 min)**
- Go to **Actions** tab on GitHub
- Watch the `release` workflow complete

**4. Publish the release:**
- Go to **Releases** tab
- Find the draft release
- Click **Edit**
- Click **Publish release**

Done! 🎉

---

## What You Get

### Binaries for All Platforms:
- ✅ macOS (Intel + Apple Silicon)
- ✅ Windows (CPU + CUDA 12 + CUDA 13)
- ✅ Linux (x64, ARM64, ROCm)
- ✅ SHA256 checksums

### Total: ~15+ build artifacts

---

## Optional: Code Signing

### For macOS Signed Binaries

**Add these GitHub Secrets:**
1. `APPLE_IDENTITY` - Apple Team ID
2. `APPLE_PASSWORD` - App-specific password
3. `MACOS_SIGNING_KEY` - Base64 P12 certificate
4. `MACOS_SIGNING_KEY_PASSWORD` - P12 password

**Add these GitHub Variables:**
1. `APPLE_TEAM_ID` - Apple Team ID
2. `APPLE_ID` - Your Apple ID email

### For Windows Signed Binaries

**Add these GitHub Secrets:**
1. `WINDOWS_SIGNING_KEY` - Base64 PFX certificate
2. `WINDOWS_SIGNING_KEY_PASSWORD` - PFX password

---

## First-Time Setup

### 1. Create `release` Environment
```
Settings → Environments → New environment → "release"
```

### 2. Enable Actions
```
Settings → Actions → General → Allow all actions
```

### 3. Configure Permissions
```
Settings → Actions → General → Workflow permissions
→ ✅ Read and write permissions
```

---

## Test Locally First

```bash
# Build for your platform
go build -o glowllama .

# Test it works
./glowllama --version

# Test rendering
echo "# Test\nThis is **bold**" | ./glowllama run llama3.3
```

---

## Release Checklist

```
[ ] Code committed and pushed
[ ] Version correct in version/version.go (1.0.0-alpha)
[ ] Binary tested locally
[ ] Tag created: v1.0.0-alpha
[ ] Tag pushed to GitHub
[ ] Workflow completed (check Actions tab)
[ ] Release published (check Releases tab)
[ ] Announce on socials 📢
```

---

## Common Commands

```bash
# List existing tags
git tag -l

# Delete a tag (if you made a mistake)
git tag -d v1.0.0-alpha
git push origin :refs/tags/v1.0.0-alpha

# Create a new tag
git tag v1.0.0-alpha
git push origin v1.0.0-alpha

# View release workflow logs
gh run list --workflow=release.yaml
gh run view <run-id> --log
```

---

## Need Help?

📖 **See full guide:** [RELEASE_GUIDE.md](RELEASE_GUIDE.md)
🐛 **Issues?** Check the troubleshooting section in RELEASE_GUIDE.md
💬 **Questions?** Open a GitHub issue

---

**That's it!** You're ready to ship Glowllama to the world 🦙✨
