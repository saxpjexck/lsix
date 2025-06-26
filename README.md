<img src="image.gif">

## **Overview**:
A single-file tool that activates JetBrains IDEs with one click — no manual activation code input required.

### ✅ Features:
- Supports activation of **paid plugins**, such as **Rainbow Brackets**
- Automatically **backs up and restores** the original state before patching
- **Cross-platform**: compatible with macOS, Linux, and Windows

### 🔧 v3.1 Fixes & Improvements:
- Improved compatibility with **residual configurations** left by other activation scripts (e.g., environment variables and incorrect permission handling)
- Added **parallel plugin request support** to speed up startup
- Introduced **plugin caching**, allowing usage even when offline or under poor network conditions
- Changed the **file extraction path** to avoid polluting the current working directory
- Added support for the **`plugin-privacy`** plugin
- Supports activating plugins like **Rainbow Brackets** (some plugins have built-in time-based license checks — to avoid being flagged as abnormal, it's recommended to set the expiration date to **2 years from today** rather than an excessively long period)

## How 2 Use

**Only once click, automatically activated**

- Click the ide card to automatically inject ja-netfilter and activation key, and restart the IDE to see that it has been activated.
- Click the plugin or ide cards to copy the activation key to the clipboard, and you can manually enter the key for activation
- It also supports one-click removal of activation configuration
- It is single bin file, after execution, the relevant files for activation will be released in the current directory. After activation, the ide will reference the directory and do not delete it.
- warning: Some plugins have built-in time detection mechanisms. Setting an expiration time too long may cause the license to be marked as an exception. Consider adjusting the expiration date of these plugins to two years from today

## dev

install go-bindata

```bash
go install github.com/go-bindata/go-bindata/v3/go-bindata@latest
export PATH=$PATH:$(go env GOPATH)/bin
go-bindata --version
go-bindata -o internal/util/access.go -pkg util static/... templates/... cache/...
go run cmd/main.go
```

## run it !

mac linux windows

```
make run
```

## make it !

mac or linux ：

```bash
make install-bindata
export PATH=$PATH:$(go env GOPATH)/bin
make build
make run
```

windows use powershell run:

```powershell
.\build.ps1
```

## Star History

[![Stargazers over time](https://starchart.cc/saxpjexck/lsix.svg?variant=adaptive)](https://starchart.cc/saxpjexck/lsix)

# JetBrains Activation Code Generator

A userscript that adds a button on the JetBrains plugin homepage to generate activation codes.

## Installation

### Prerequisites
- A userscript manager extension installed in your browser:
  - [Tampermonkey](https://www.tampermonkey.net/) (Recommended)
  - [Violentmonkey](https://violentmonkey.github.io/)
  - [Greasemonkey](https://www.greasespot.net/)

### Installation Steps
1. Install one of the userscript managers mentioned above
2. Click on the userscript manager icon in your browser
3. Click "Create a new script"
4. Copy and paste the entire contents of `JetBrains Activation Code Generator.js` into the editor
5. Save the script (Ctrl+S or File > Save)

## Usage

1. Visit any JetBrains plugin page (e.g., https://plugins.jetbrains.com/plugin/...)
2. You will see two new buttons in the plugin header:
   - "Get permanent activation code"
   - "Get a 2-year activation code"
3. Click either button to generate the corresponding activation code
4. The code will be automatically copied to your clipboard
5. Use the copied code to activate your JetBrains plugin

## Updating Keys

The script includes two methods to update the keys in `JetBrains Activation Code Generator.js`:

### Prerequisites

1. Place your original key file at `bin/.jetbra-free/jetbra.key`
2. Place your certificate file at `bin/.jetbra-free/jetbra.pem`

### Using Bash Script (Linux/macOS/WSL)

1. Make the script executable:
```bash
chmod +x update_keys.sh
```

2. Run the script:
```bash
./update_keys.sh
```

### Using PowerShell Script (Windows)

1. Open PowerShell and navigate to the script directory

2. If you haven't already, set the execution policy to allow running scripts:
```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

3. Run the script:
```powershell
.\update_keys.ps1
```

### What the Scripts Do

Both scripts will:
1. Convert the key to PKCS8 format using OpenSSL
2. Read the new key and certificate files
3. Update the JavaScript file with the new values
4. Clean up temporary files
5. Display a success message when complete

### Troubleshooting

- Make sure OpenSSL is installed and available in your system PATH
- Ensure you have read/write permissions for all files
- If you encounter any issues, try running the scripts with administrator/sudo privileges
- Always backup your files before running the update scripts

## Browser Compatibility

The userscript has been tested and works with:
- Google Chrome
- Mozilla Firefox
- Microsoft Edge
- Opera

## Notes

- The script only works on JetBrains plugin pages
- Make sure to update the keys if you want to use your own certificates
- The generated activation codes are for personal use only
