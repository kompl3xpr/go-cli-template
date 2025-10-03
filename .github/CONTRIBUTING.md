# Contribution Guidelines

Thank you for your interest in Go CLI Template! We welcome all forms of contributions, including but not limited to bug reports, feature suggestions, documentation improvements, and code submissions.

## New to Open Source?

If you're new to contributing to open source projects, don't worry! This guide will help you successfully complete your first contribution. If you encounter any issues during the process, feel free to ask questions in the issue tracker, and we'll be happy to assist.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How to Contribute](#how-to-contribute)
  - [Reporting Bugs](#reporting-bugs)
  - [Making Suggestions](#making-suggestions)
  - [Submitting Code](#submitting-code)
- [Development Environment Setup](#development-environment-setup)
- [Code Style](#code-style)
- [Commit & Pull Request Conventions](#commit--pull-request-conventions)
- [Pull Request Process](#pull-request-process)

## Code of Conduct

Please read and adhere to our Code of Conduct to ensure a friendly and inclusive community environment.

## How to Contribute

### Reporting Bugs

If you discover a bug, please first check if it has already been reported in the [issue list](issues).

**Important: Please use the issue template!**

When creating an issue, GitHub will automatically display the template. This template contains essential information that helps us understand and resolve issues more efficiently. Please complete all sections of the template.

### Making Suggestions

For new feature suggestions, please also use the corresponding issue template. Describe your requirements, use cases, and expected behavior in detail.

### Submitting Code

<details>
<summary>Expand: Detailed Process</summary>

1. Fork this repository

2. [Set up the development environment](#development-environment-setup)

3. Create your feature branch (`git checkout -b feat/amazing-feature`)
    - Branch names must follow [this format](#branch-naming-convention)

4. Make your code changes

5. Stage your changes (`git add .`)

6. Commit your changes (`git commit -m 'feat: add some amazing feature'`)
    - Commit messages must follow [this format](#commit-message-convention)

7. Push to the branch (`git push origin feat/amazing-feature`)

8. Open a Pull Request on GitHub

</details>

## Development Environment Setup

### 0. Prerequisites

#### Tools

Ensure you have `git` (version 2.9 or higher), Go toolchain, `golangci-lint`, and `make` installed.

<details>
<summary>Expand: Tool Installation on Windows</summary>

On Windows, we recommend using chocolatey (a software management solution) for installation. Steps:

1. Right-click the Start menu and open PowerShell or Terminal **as Administrator**.

2. Run the following command to install chocolatey (Warning: command may be outdated, please refer to the [official website](https://chocolatey.org/install)):
    ```shell
    Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
    ```

3. Wait until you see `Chocolatey CLI (choco.exe) is now ready.`

4. Execute the following command (remove parameters for tools you **already** have installed):
    ```shell
    choco install git golang golangci-lint make
    ```

5. Wait until the terminal indicates successful installation.

</details>

#### Initial Git Setup

*Skip this step if already configured.*

<details>
<summary>Expand: Setup Steps</summary>

1. Set username and email

    ```bash
    git config --global user.name YourGitHubUsername
    git config --global user.email YourGitHubEmail
    ```

2. Optional: [Generate SSH key and upload public key to GitHub](https://docs.github.com/en/authentication/connecting-to-github-with-ssh)

3. Optional: [Generate GPG signature and upload public key to GitHub](https://docs.github.com/en/authentication/managing-commit-signature-verification/generating-a-new-gpg-key)

</details>

### 1. Clone your forked project locally:
```bash
git clone https://github.com/YourUsername/go-cli-template.git
cd go-cli-template
```

### 2. Add upstream remote repository:
```bash
# Execute in project root directory
git remote add upstream https://github.com/kompl3xpr/go-cli-template.git
```

### 3. Git Hooks Setup

**Important: Git Hooks setup is mandatory!**

To ensure code quality, we use Git Hooks to automatically run code checks and tests during commits.
```bash
# Execute in project root directory
git config core.hooksPath .githooks
```
For Linux/Unix systems, also run the following command to ensure hook scripts have execute permissions:
```bash
# Execute in project root directory
chmod -R +x .githooks/
```

After setup, the following will automatically run on each commit:
- Code style checks, static analysis, compilation attempts, and test execution
- Automatic generation of header comments for source code
- Verification of commit message and branch naming conventions

If hooks fail, please fix the relevant issues before committing.

### 4. Build and test the project
```bash
# Execute in project root directory
make
```
Or equivalently:
```bash
# Execute in project root directory
make clear
make build
make test
```

If build or tests fail, check if your environment configuration is incorrect. If the failure is not due to local configuration issues, please submit a relevant Issue.

## Code Style

Please follow the project's code style conventions:
- Use `make fmt` to format code
- Write clear comments
- Add appropriate test cases

## Commit & Pull Request Conventions

### Commit Message Convention
```bash
# Note the space after the colon
# Commit format can be
git commit -m "type: description"
# or
git commit -m "type(module/subdirectory): description"

# Examples:
git commit -m "feat: add `--reload` command for CLI"
git commit -m "fix(internal/api): status code of GET v1/users"
git commit -m "docs: fix typo in README.md"
```

#### Commit Types

Must be one of: `feat` `fix` `docs` `style` `refactor` `test` `chore` `ci`

| Type | Description |
| :--- | :--- |
| `feat` | **New feature**, involving addition or modification of source code. |
| `fix` | **Bug fix**, involving source code modifications to resolve issues. |
| `docs` | **Documentation updates**, such as README, code comments, etc. |
| `style` | **Code style** adjustments, formatting changes that don't affect code meaning (spaces, indentation). |
| `refactor` | **Code refactoring**, structural code changes that neither fix bugs nor add features. |
| `test` | **Test case** additions or modifications, including unit and integration tests. |
| `chore` | **Build/toolchain** changes, such as dependency updates, Makefile or script modifications. |
| `ci` | **CI configuration** updates, modifying continuous integration processes or scripts (e.g., `.github/workflows`). |

#### Module/Subdirectory

1. Optional but recommended
2. Describes the area/section affected by the commit, e.g., `(cmd)` indicates command-line related changes

#### Commit Description

1. English imperative mood, lowercase first letter, space-separated words, e.g., `remove heading emojis in README.md`
2. Minimum 6 characters, maximum 80 characters

## Branch Naming Convention
```bash
# Use hyphens instead of underscores
# Use forward slash "/" not backslash "\"
git checkout -b type/description

# Examples:
git checkout -b feat/add-reload-command-for-cli
git checkout -b fix/api-get-user-404
git checkout -b docs/readme-typo
```

### Branch Types
See [Commit Types](#commit-types)

### Branch Description
1. English imperative mood, lowercase letters, numbers, and hyphens only, hyphen-separated words, e.g., `remove-heading-emojis-in-readme`
2. Minimum 6 characters, maximum 80 characters

## Pull Request Process

1. Ensure your fork is synchronized with the upstream repository
2. Create a new branch for your feature
3. Ensure all tests pass
4. Submit clear commit messages
5. Open a Pull Request and describe your changes

Our maintainers will review your PR and may suggest modifications. This is normal in open-source collaboration - don't be discouraged!

Thank you again for your contributions! 🎉