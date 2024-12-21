<p align="center">
  <img src="./frontend/static/favicon.png" alt="Gunnel Logo" width="200">
</p>

<h1 align="center">Gunnel</h1>

<p align="center">
  <strong>Simple and Intuitive SSH Tunnel Manager</strong>
</p>

# Overview

Gunnel is a cross-platform desktop application that makes managing SSH tunnels easy. Built with Go and Svelte, it provides a modern user interface to create, manage and monitor SSH port forwards without dealing with complex command line arguments.

# Features

- 🚀 Easy-to-use graphical interface for SSH tunnels
- 🔒 Secure SSH key authentication support
- 🔄 Multiple concurrent tunnel support
- 💾 Save and restore tunnel configurations
- 🖥️ Cross-platform (Windows, macOS, Linux)

# Installation

## Download

Download the latest release from the [Releases](https://github.com/Padi2312/gunnel/releases) page.

> [!IMPORTANT] 
> The application is not signed yet, so you may need to allow it to run on your system.

> [!NOTE]
> In case you don't trust the binary, you can build it from source.

## Build from Source

### Prerequisites

- Go 1.21 or higher
- Node.js 18+ and pnpm

### Building

To create production builds:

```bash
wails build
```

The built binaries will be available in the  `bin` directory.

# Usage

1. Launch Gunnel
2. Create a new tunnel configuration with:
   - Remote host details
   - Local and remote ports
   - Username
3. Click "Connect" to establish the tunnel

# Contributing

Contributions are welcome! Please feel free to submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Development Setup

1. Clone the repository:
```bash
git clone https://github.com/yourusername/gunnel.git
cd gunnel
```

2. Install Wails CLI:
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

3. Start development server:
```bash
wails dev
```

# License

This project is licensed under the MIT License - see the LICENSE file for details.

# Acknowledgments

- Built with [Wails](https://wails.io/)
- UI powered by [Svelte](https://svelte.dev/)
