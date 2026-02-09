# Git Push Script for Auth Service
# Usage: .\git-push.ps1 "your commit message"
# Or just: .\git-push.ps1 (will use default message)

param(
    [string]$CommitMessage = "Update auth-service"
)

$ErrorActionPreference = "Stop"

Write-Host "[START] Starting Git Push..." -ForegroundColor Cyan

# Check if git is initialized
if (-not (Test-Path ".git")) {
    Write-Host "[INIT] Initializing git repository..." -ForegroundColor Yellow
    git init
    git remote add origin https://github.com/luongtien872003/auth-service.git
}

# Check if remote exists
$remotes = git remote
if ($remotes -notcontains "origin") {
    Write-Host "[REMOTE] Adding remote origin..." -ForegroundColor Yellow
    git remote add origin https://github.com/luongtien872003/auth-service.git
}

# Add all files
Write-Host "[ADD] Adding files..." -ForegroundColor Yellow
git add .

# Commit
Write-Host "[COMMIT] $CommitMessage" -ForegroundColor Yellow
git commit -m "$CommitMessage"

# Set branch to main
git branch -M main

# Push
Write-Host "[PUSH] Pushing to GitHub..." -ForegroundColor Yellow
git push -u origin main

Write-Host "[DONE] Check: https://github.com/luongtien872003/auth-service" -ForegroundColor Green
