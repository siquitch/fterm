# Flutterterm: A Flutter CLI Wrapper

**Flutterterm** is a convenient CLI wrapper for running Flutter commands.

## Usage

A configuration file is **not required**, but it is recommended to get the most out of this tool.  

To set up a configuration file, create a `.fterm_config.json` file in the root directory of your Flutter project. Flutterterm will look for this file when you run commands.

### Without a Configuration File

If you choose not to use a configuration file, Flutterterm will act as if it is executing:
```flutter run```

### Example `.fterm_config.json` File

```json
[
    {
        "name": "Dev",
        "mode": "debug",
        "flavor": "dev",
        "target": "main.dart",
        "dart_define_from_file": "file.json"
    },
    {
        "name": "Prod",
        "mode": "release",
        "flavor": "prod",
        "target": "main.dart"
    }
]
```

### **Configuration Fields**

- **`name`**: The name of the configuration.
- **`mode`**: The mode to run your app in (`debug`, `release`, or `profile`).
- **`flavor`**: The flavor to use (equivalent to `--flavor`).
- **`target`**: The target file to run (`-t` or `--target`). Defaults to `main.dart` if not specified.
- **`dart_define_from_file`**: Defines environment variables from a file (equivalent to `--dart-define-from-file`).


### **Running Commands**

#### **flutterterm emulators**
This command allows you to start an existing emulator or create a new one
```bash
// Start an existing emulator
flutterterm emulators

// Create a new emulator
flutterterm emulators create
```

####  **flutterterm run**
This command wraps the ```flutter run``` command
```bash
flutterterm run
```
