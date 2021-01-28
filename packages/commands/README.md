# UCM Commands

## Command: `list`
Displays list of named config sets.
```
> ucm list
Universal Configs Manager:
  - personal
  - work     <= current
```

## Command: `show`
Displays configuration values of a name config set or from a requested file.

### Parameters
- `-n`, `--name` &mdash; name of a known configurations set.
- `-f`: `-file` &mdash; path or an URL to a file with config values

### Examples of usage:
```
> ucm show --n personal
> ucm show --name work
> ucm show --file ~/.my-config-value.yaml
> ucm show --f https://example.com/some-config.yaml
```

## Command: `use`
Switches configuration values using a named configs set or values from a file.

### Parameters
- `-n`, `--name` &mdash; name of a known configurations set.
- `-f`: `-file` &mdash; path or an URL to a file with config values

### Examples of usage:
```
> ucm use --n personal
> ucm use --name work
> ucm use --file ~/.my-config-value.yaml
> ucm use --f https://example.com/some-config.yaml
```
