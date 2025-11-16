## 📦 Overview
This module handles user management with minimal manual wiring and minimal “magic.” It is designed to integrate seamlessly into your project.

# 🚀 Installation

```cmd
mycli install user
```

## 🔗 Wiring

Every module is designed to require the minimum ammount of manual wiring as possible, and the minimum of magic possible.

in main should call the Register method defined into "./users/routes/innit.go"

## 🛠 parameters

### Storage

Specify the storage backend to use with the module:
```cmd
mycli install user --storage mongo
```
- If omitted, the module will use the database already installed.
- If no database is installed, MongoDB will be used as the default.

> See more supported storage options [here]().