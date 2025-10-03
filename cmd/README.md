# Application Commands

## Overview

This directory contains application commands generated and managed using **Cobra CLI**. All command-line interface (CLI) related code is organized under this directory.

## Command Structure

```
cmd/
├── root.go              # Root command configuration and initialization
└── ...                  # Other subcommands
```

## Development Standards
*Note: All command-related development should be managed through the Cobra CLI tool to maintain structural consistency and maintainability.*

Files in this directory should be created exclusively through Cobra CLI, as follows:
```bash
# Example: Create a 'serve' command, generating corresponding source code files
cobra-cli add serve

# Example: Create a 'create' subcommand under the 'config' command, generating corresponding source code files
cobra-cli add create -p 'configCmd' # Note: Default parent command variable name is parent command name + 'Cmd'
```

Cobra CLI can be installed using this command:
```bash
go install github.com/spf13/cobra-cli@latest
```