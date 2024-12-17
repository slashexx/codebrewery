# Contributing to CodeBrewery

Thank you for considering contributing to **CodeBrewery**! Whether it's reporting a bug, proposing new features, or submitting code changes, your input is highly valuable and appreciated.

This document provides guidelines on how to contribute effectively to the project.

---

## Table of Contents
- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
  - [Project Structure](#project-structure)
  - [Prerequisites](#prerequisites)
- [Setting Up the Project](#setting-up-the-project)
- [Contribution Workflow](#contribution-workflow)
  - [Reporting Bugs](#reporting-bugs)
  - [Suggesting Features](#suggesting-features)
  - [Code Contributions](#code-contributions)
    - [Style Guide](#style-guide)
    - [Testing](#testing)
- [Commit and PR Guidelines](#commit-and-pr-guidelines)
- [Code Review Process](#code-review-process)

---

## Code of Conduct

By participating in this project, you agree to adhere to the [Code of Conduct](CODE_OF_CONDUCT.md) to maintain a positive and inclusive environment for all contributors.

## Getting Started

### Project Structure
CodeBrewery follows a **microservices architecture**. Here's the high-level directory structure:

```
codebrewery/
  frontend/               # Frontend React app (TypeScript, Vite)
  backend/
    api-gateway/          # Central entry point for services
    user-service/         # User-related operations
    code-execution-service/  # Code execution engine
  k8s/                    # Kubernetes manifests
  docker/                 # Docker configuration files
```

### Prerequisites
Ensure you have the following tools installed:

- **Node.js (v18+)** and **npm**
- **Go (v1.21+)**
- **Docker** and **Docker Compose**
- **Kubernetes (Minikube or any cluster provider)**
- **Git**

> Install dependencies for the frontend/backend as specified in their respective `package.json` or `go.mod` files.

---

## Setting Up the Project

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/codebrewery.git
   cd codebrewery
   ```

2. **Setup Frontend:**
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

3. **Setup Backend Services:**
   Each service has its own `go.mod` file and `Dockerfile`. Start the services using Docker Compose:
   ```bash
   cd backend
   docker-compose up --build
   ```

4. **Verify setup:**
   Access the frontend at `http://localhost:5173` and confirm all backend APIs are running.

---

## Contribution Workflow

### Reporting Bugs

1. Check if the issue already exists in the [GitHub Issues](https://github.com/your-username/codebrewery/issues).
2. If not, open a new issue with:
   - Clear title
   - Steps to reproduce
   - Expected behavior
   - Screenshots/logs if applicable

### Suggesting Features

- Propose a feature by opening an issue under the "Feature Request" template.
- Clearly describe the use case, benefits, and possible implementation approach.

### Code Contributions

#### Style Guide
- **Backend:** Follow standard Go best practices.
- **Frontend:** Follow the Prettier and ESLint rules configured in the project.

Run formatters before submitting changes:
```bash
# For Go code
go fmt ./...

# For frontend
npm run lint --fix
```

#### Testing
- Write unit tests for new features or fixes.
- Run tests locally:
   ```bash
   # Frontend
   npm run test

   # Backend (Go tests)
   go test ./...
   ```

---

## Commit and PR Guidelines

1. **Commit Messages:**
   - Use concise, descriptive commits.
   - Follow conventional commit format:
     ```
     <type>: <subject>
     ```
     - `feat:` for new features
     - `fix:` for bug fixes
     - `docs:` for documentation changes
     - `chore:` for refactors, cleanups, etc.

     Example:
     ```
     feat: add user authentication in user-service
     ```

2. **Pull Requests:**
   - Fork the repository and work on a new branch.
     ```bash
     git checkout -b feature/your-feature-name
     ```
   - Ensure the code is tested and linted.
   - Provide a clear PR description:
     - **What:** Describe the changes.
     - **Why:** Reason for the change.
     - **How:** Implementation details.
   - Link related issues.

3. **Small, Focused Changes:**
   - PRs should be focused on solving one problem at a time.

---

## Code Review Process

- Once you submit a PR, the maintainers will review it.
- Address feedback promptly.
- Be respectful during code reviews — they are here to ensure quality and growth.

---

## Questions or Help?

Feel free to join the discussion on GitHub Issues or contact maintainers.

Let's build CodeBrewery together! 🚀

