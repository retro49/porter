package main

import "fmt"

const (
	USAGE_HELP    = "display this help message and exit"
        USAGE_VERSION = "print the version of porter"
	USAGE_HOST    = "host to specify host to scan"
	USAGE_NETWORK = "network  to specify network type"
	USAGE_START   = "to specify start port number"
	USAGE_END     = "to specify ending port number"
	USAGE_RANGE   = "to specify a port as range of values"
	USAGE_SKIP    = "to specify port(s) not to scan"
	USAGE_STEP    = "to specify a step for scanning ports"
	USAGE_PORT    = "to specify a single port to scan"
	USAGE_OUTPUT  = "to specify output file"
	USAGE_FORMAT  = "to specify output format"
        USAGE_THREADS = "to specify how  many threads to spawn"
        USAGE_TIMEOUT = "to specify the timeout to wait for  each port"
)

const (
	DEFAULT_HOST    string = "127.0.0.1"
	DEFAULT_NETWORK string = "tcp"
	DEFAULT_START   int    = 1
	DEFAULT_END     int    = (1 << 16) - 1
	DEFAULT_RANGE   string = ""
        DEFAULT_SKIP    int    = -1
        DEFAULT_OUTPUT  string = ""
        DEFAULT_FORMAT  string = "normal"
        DEFAULT_THREADS int    = 1
        DEFAULT_TIMEOUT int    = 5
)

// options are in order so take care.
var (
	ARG_MANUAL = fmt.Sprintf(`
Porter is a simple cli tool used for scanning ports on a host.
This tool uses a dialup method to scan a given port with the 
provided network type.

use -h | --help %s 

use -H | --host %s
if not provided localhost (127.0.0.1) will be used by default

use -n | --network %s 
By default tcp netowrk is used
available options for a network type is tcp and udp.

use -s | --start %s
start port must be 1, if not provided by default it is 1

use -e | --end %s
The maximum port number to scan is 65535, if
not provided the default will be 65535

use -r | --range %s
This format provided an easy representation of start and end option.
Inorder to use this option it is must that to provide a range syntax
using "..", Eg. 1..80, in this option 80 is exclusive
using "..=", Eg. 1..=80, specifies 80 is inclusive.

use -k | --skip %s
By using this option it is possible to skip a port to scan.
This argument can take a single or multiple ports.

use -S | --step %s
With the given start and end values this argument is used to
make a step or jump between ports. By defualt its 1 so that no
jump or step is made.

use -p | --port %s

use -o | --output %s

use -f | --format %s
Format option supports the result output as a human 
readale text or json result.

use -t | --threads %s
whether to scan in a single or concurrent thread
    `,

		USAGE_HELP,
		USAGE_HOST,
		USAGE_NETWORK,
		USAGE_START,
		USAGE_END,
		USAGE_RANGE,
		USAGE_SKIP,
		USAGE_STEP,
		USAGE_PORT,
		USAGE_OUTPUT,
		USAGE_FORMAT,
                USAGE_THREADS,
	)
)
