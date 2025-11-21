# moon-phase-display
A little go program to print out the current moon phase as a visually appealing string to be used in bars like polybar. Will be expanded in the future

## Usage
General:
```bash
moon-phase-display <mode> [latitude] [longitude]
```
Valid modes are: `moon`
latitude and longitude are optional and also useless for now

Building the project:
```bash
# build the binary and put it in a desired directory
go build -o /usr/local/bin/moon-phase-display
```

Polybar module:
```toml
[module/moonphase]
type = custom/script
exec = moon-phase-display moon
interval = 43200 # refresh every 12 hours
```
