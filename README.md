# PoGoPass - Secure Password Manager

PoGoPass is a lightweight, open-source password manager designed to help you securely store and manage your sensitive information. With a simple yet effective design, PoGoPass ensures that your passwords are encrypted and easily accessible only with your master password.

---

## 🚀 Getting Started

To get started, you need to **clone the repository** or download the latest release from our [GitHub Releases](https://github.com/arwindpianist/PoGoPass/releases). Follow the instructions below to set up PoGoPass on your system.

### **Prerequisites**
- **Go 1.17+** (for building from source)
- A **SQLite database** for storing passwords (handled automatically)
- Your **master password** to securely encrypt and decrypt the data

---

## 🎨 Welcome Screen

When you run the app for the first time, you'll be welcomed with the following message:

```
 _______             ______             _______                              
/       \           /      \           /       \                             
$$$$$$$  | ______  /$$$$$$  |  ______  $$$$$$$  | ______    _______  _______ 
$$ |__$$ |/      \ $$ | _$$/  /      \ $$ |__$$ |/      \  /       |/       |
$$    $$//$$$$$$  |$$ |/    |/$$$$$$  |$$    $$/ $$$$$$  |/$$$$$$$//$$$$$$$/ 
$$$$$$$/ $$ |  $$ |$$ |$$$$ |$$ |  $$ |$$$$$$$/  /    $$ |$$      \$$      \ 
$$ |     $$ \__$$ |$$ \__$$ |$$ \__$$ |$$ |     /$$$$$$$ | $$$$$$  |$$$$$$  |
$$ |     $$    $$/ $$    $$/ $$    $$/ $$ |     $$    $$ |/     $$//     $$/ 
$$/       $$$$$$/   $$$$$$/   $$$$$$/  $$/       $$$$$$$/ $$$$$$$/ $$$$$$$/  
```

---

## 🔑 Features

- **AES-256 Encryption**: Your passwords are stored securely using AES-256 encryption.
- **Cross-platform**: Build for Linux, Windows, and macOS.
- **Lightweight**: A simple and minimal design to keep the focus on security.
- **Command-line Interface (CLI)**: Add, list, and manage your passwords with ease.
- **Offline**: No need for an internet connection to store or retrieve your passwords.

---

## 🛠 Installation

### From Source

1. Clone the repository:

```bash
git clone https://github.com/arwindpianist/PoGoPass.git
cd PoGoPass
```

2. Build the app:

```bash
make build
```

This will build executables for Linux, macOS, and Windows.

3. Run the application:

```bash
./PoGoPass-linux  # For Linux
./PoGoPass-darwin # For macOS
./PoGoPass-windows # For Windows
```

---

## 📝 Commands

Here’s a list of available commands you can use within the PoGoPass application:

- `add`: Add a new password entry.
- `list`: List all stored passwords.
- `delete`: Remove a password entry.
- `show`: Show a specific password entry.
- `reindex`: Rebuild the index for optimized access.
- `help`: Display help information.
- `quit`: Exit the application.

---

## 💡 Usage Example

Once you've run the app, you will be prompted to enter your **master password**. After verifying your master key, you can start managing your passwords. Here's an example session:

```bash
Enter your master password: ********
Master key verified successfully
Enter command (add, list, delete, show, reindex, help, quit): add
Enter service name (e.g. Facebook): twitter
Enter username: arwindpianist
Enter password: *********
Auto-generated password: Fr1Y@4n-TM2)J0#I
Password added successfully!
```

---

## ⚙️ Configuration

You can configure the app by setting up a `.env` file to store sensitive information securely. Example:

```
PEPPER=PoGoPassPepper!
SALT=PoGoPassSalt1234
HASH_FILE=masterkey.hash
```

---

## 📝 License

PoGoPass is open-source and released under the MIT License. See [LICENSE](LICENSE) for more information.

---

## 🤝 Contributing

We welcome contributions! If you'd like to contribute to the project, feel free to fork the repository, make your changes, and submit a pull request.

---

## 🌟 Credits

PoGoPass is developed and maintained by **[arwindpianist](https://github.com/arwindpianist)**.

---

## 📦 Releases

You can download the latest releases directly from [GitHub Releases](https://github.com/arwindpianist/PoGoPass/releases).

---

## 🛠 Tools & Technologies

- **Go**: Programming language used for development.
- **SQLite**: Database used for storing passwords.
- **AES-256**: Encryption algorithm used for securing passwords.
- **GoReleaser**: Tool used for building and releasing the app.

---