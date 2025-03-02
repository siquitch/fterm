# Flutterterm: A Flutter CLI Wrapper

**Flutterterm** is a cli tool that manages flutter run configurations

## Inspiration

The flutter cli is annoying to use when you need to manage things like 
environment variables, or test your app on multiple platforms

flutter-tools.nvim has a nice way to make pre-set configurations, and vs-code
has launch.json, but these are limited to running inside the respective tools,
which I have found to be inconvenient. 

This tool aims to add functionality directly on the command line by creating
a more interactive way to use it. If you've ever needed to manage different 
run configurations, then this tool can help. 

## Usage

Run ```flutterterm config init``` in the root of your flutter project.

This will create a `.fterm_config.json` file

### Example `.fterm_config.json` File

```json
{
  "version": "0.0.3",
  "default_config": "default",
  "configs": [
    {
      "name": "default",
      "description": "The default run configuration",
      "mode": "debug",
      "flavor": "",
      "target": "lib/main.dart",
      "dart_define_from_file": "",
      "additional_args": null
    }
  ],
  "favorite_configs": [],
  "devices": {
    "favorite_devices": [],
    "device_configs": []
  }
}
```

### **Configuration Fields**

- **`name`**: The name of the configuration.
- **`mode`**: The mode to run your app in (`debug`, `release`, or `profile`).
- **`flavor`**: The flavor to use (equivalent to `--flavor`).
- **`target`**: The target file to run (`-t` or `--target`). Defaults to `main.dart` if not specified.
- **`dart_define_from_file`**: Defines environment variables from a file (equivalent to `--dart-define-from-file`).
- **`additional_args`**: Pass in additional args for the command


### **Running Commands**

#### Get emulators to launch
```bash
flutterterm emulators
```
#### Run with config options listed
```bash
flutterterm run
```
#### Run a specific config
```bash
flutterterm run [config]
```
#### Run with the default config
```bash
flutterterm run -d
```

