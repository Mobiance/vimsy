# Vimsy 🌀

> A simple CLI tool to manage and switch between multiple Neovim configurations.

Vimsy helps developers manage different Neovim setup profiles with ease. Whether you're testing a plugin-heavy config, starting fresh, or working on your own Vim distribution—`vimsy` lets you switch between them with a single command.

---

## ✨ Features

- 🔧 Add new Neovim configuration profiles
- 🔀 Switch between existing configurations
- 📁 Supports custom config directories
- 🧹 Clean separation of profiles without cluttering your `.config/nvim`

---

## 📦 Installation

```bash
git clone https://github.com/mobiance/vimsy.git
cd vimsy
go install .
```

Make sure `$GOPATH/bin` is in your `PATH`.

---

## 🛠 Usage

```bash
vimsy [command]
```

###  Vimsy CLI Commands

| Command                        | Description                         |
|--------------------------------|-------------------------------------|
|`vimsy create <name>`          | Create a new config profile         |
|`vimsy edit [name]`            | Edit a profile (or current, if none)|
|`vimsy current`                | Show current active config          |
|`vimsy delete <name>`          | Delete a config profile             |
|`vimsy list`                   | List all config profiles            |
|`vimsy rename <old> <new>`     | Rename a config profile             |
|`vimsy use <name>`          | Switch to selected config           |

---

## 📂 Config Structure

All Vimsy-managed configs are stored in:

```
~/.config/vimsy/configs/
├── my-config/
├── minimal/
└── lunarvim-clone/
```

Your `~/.config/nvim` becomes a symlink pointing to one of these configs.
To add a new config, simply create a new configuration directory under `~/.config/vimsy/configs/`.

---

## 🧪 TODO

- [ ] Backup and restore support
- [ ] Plugin profiles

---

## 👨‍💻 Author

Made with 💙 by [Shubham Sharma](https://github.com/mobiance)  
Inspired by dotfile and dev setup struggles.

---

## 🪪 License

MIT © 2025 Shubham Sharma
