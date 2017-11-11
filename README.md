## Interactive shell for controlling Honeywell's Evohome system

Copyright (c) 2017, Rick Hofstede

All rights reserved. This software is distributed under a BSD-style
license. For more information, see LICENSE.

### 1. Introduction

This library can be used for controlling Evohome systems from the
command line. The following features are currently supported:

- List all zones, including their current temperature and the target temperature.
- Set zone temperatures, either based on end time, or continuously.
- Cancel zone temperature overrides.

The following screenshots demonstrate some of this functionality.

#### List all zones

![evohome-cli show zones](https://github.com/rickhofstede/assets/raw/master/evohome-cli/evohome-cli_show_zones.png "evohome-cli: show zones")

#### Cancel zone temperature override

![evohome-cli cancel zone temperature override](https://github.com/rickhofstede/assets/raw/master/evohome-cli/evohome-cli_cancel_temp_override.png "evohome-cli: cancel zone temperature override")

### 2. How to build

Fetch all dependencies:
```
$ go get
```

Build from source:
```
$ go build                          # build for your own system/OS
$ GOOS=linux GOARCH=amd64 go build  # build for Linux/amd64
$ GOOS=darwin GOARCH=amd64 go build # build for Mac OS X
```

### 3. How to use

The Evohome CLI currently supports three main command keywords:

- `show`: display configuration/status.
- `set`: change configuration.
- `cancel`: stop a (temporary) configuration, such as a temperature override.

Here are some examples of currently supports commands per command keyword:

- `show`:
    - `show zones`: List all zones, including their current temperature and the target temperature.
    - `show zone Bathroom`: Show details of the "Bathroom" zone.
- `set`:
    - `set zone Bathroom temperature 19.5`: Set the temperature of the "Bathroom" zone to 19.5 degrees.
    - `set zone Bathroom temperature 19.5 until 2017/11/05 17:30`: Set the temperature of the "Bathroom" zone to 19.5 degrees until November 5, 2017, 17:30 (local time).
- `cancel`:
    - `cancel zone Bathroom temperature override`: Cancel the temperature override of the "Bathroom" zone.

### 4. Support

Please request support by creating an 'issue' [here](https://github.com/rickhofstede/evohome-cli/issues).
