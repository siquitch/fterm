# Flutterterm: A Flutter CLI Wrapper

**Flutterterm** is a cli tool that manages flutter run configurations

## Inspiration

The flutter cli is annoying to use when you need to manage things like 
environment variables, or test your app on multiple platforms

flutter-tools.nvim has a nice way to make pre-set configurations, and vs-code
has launch.json, but these are limited to running inside their respective tools,
which I have found to be inconvenient. 

This tool aims to add functionality directly on the command line by creating
a more interactive way to use it. If you've ever needed to manage different 
run configurations, then this tool can help. 

## Usage

Run ```flutterterm init``` in the root of your flutter project.

This will create a `.fterm_config.json` file

### Example `.fterm_config.json` File

```json
{
  "version": "0.0.3",
  "fvm": false,
  "default_config": "default",
  "configs": [
    {
      "name": "Debug",
      "description": "Run app in debug mode",
      "mode": "debug",
      "flavor": "",
      "target": "lib/main.dart",
      "dart_define_from_file": "",
      "additional_args": null,
      "favorite": false
    },
    {
      "name": "Release",
      "description": "Run app in release mode",
      "mode": "release",
      "flavor": "",
      "target": "lib/main.dart",
      "dart_define_from_file": "",
      "additional_args": null,
      "favorite": false
    }
  ],
  "last": ""
}
```

### **Configuration Fields**

- **`name`**: The name of the configuration.
- **`description`**: A description of the configuration.
- **`mode`**: The mode to run your app in (`debug`, `release`, or `profile`).
- **`flavor`**: The flavor to use (equivalent to `--flavor`).
- **`target`**: The target file to run (`-t` or `--target`). Defaults to `main.dart` if not specified.
- **`dart_define_from_file`**: Defines environment variables from a file (equivalent to `--dart-define-from-file`).
- **`additional_args`**: Pass in additional args for the command
- **`favorite`**: Favorite this config to sort to the top (not yet implemented)


### **Commands**

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
#### Run the last run config
```bash
flutterterm run -l
```
