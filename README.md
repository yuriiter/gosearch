# gosearch

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" alt="Go Language Badge">
  <img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" alt="License Badge">
</p>

`gosearch` is a lightweight command-line tool written in Go for quickly searching and viewing packages on the official Go package discovery site, [pkg.go.dev](https://pkg.go.dev). It provides a clean, colorized, and beautiful output format right in your terminal.

## ✨ Features

*   **Fast Search**: Query pkg.go.dev directly from the terminal.
*   **Detailed Results**: See the package path, synopsis, version, import count, and license at a glance.
*   **Colorized Output**: Easy-to-read, structured results using ANSI colors.
*   **Custom Limit**: Control the number of results shown with a CLI flag.
*   **URL Encoded**: Safely handles complex search queries.

## 🚀 Installation

### Using `go install`

The easiest way to install `gosearch` is using the `go install` command:

```bash
go install github.com/yuriiter/gosearch@latest
```

## 💡 Usage

```bash
gosearch [flags] <query>
```

### Examples

**Search for packages related to "http router"**

```bash
gosearch http router
```

**Search for "cobra" and limit the results to 3**

```bash
gosearch -limit 3 cobra
```

**Output Example:**

```
Searching pkg.go.dev for: cobra
Request URL: https://pkg.go.dev/search?limit=5&m=package&q=cobra

github.com/spf13/cobra (v1.10.1)
  Imports: 182,338 | License: Apache-2.0 | Updated: Sep  1, 2025
  Package cobra is a commander providing a simple interface to create powerful modern CLI interfaces.

github.com/cosmos/cosmos-sdk/server (v0.53.4)
  Imports: 4,509 | License: Apache-2.0 | Updated: Jul 25, 2025
  The commands from the SDK are defined with `cobra` and configured with the `viper` package.

github.com/muesli/mango-cobra (v1.3.0)
  Imports: 113 | License: MIT | Updated: Sep 10, 2025
  (No description available)
...
```

## ⚙️ Flags

| Flag | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `-limit` | `int` | `10` | Max number of search results to display. |


## 📄 License

This project is licensed under the MIT License.
