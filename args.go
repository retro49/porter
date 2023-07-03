package main

import "fmt"

const (
	ARG_PORTER_USAGE_HELP    = "-h | --help to display this and exit"
	ARG_PORTER_USAGE_HOST    = "-H | --host to specify host to scan"
	ARG_PORTER_USAGE_NETWORK = "-n | --network  to specify network type"
	ARG_PORTER_USAGE_START   = "-s | --start to specify start port address"
	ARG_PORTER_USAGE_END     = "-e | --end to specify ending port address"
	ARG_PORTER_USAGE_RANGE   = "-r | --range to specify a port as range of values"
	ARG_PORTER_USAGE_SKIP    = "-k | --skip to specify port(s) not to scan"
	ARG_PORTER_USAGE_STEP    = "-S | --step to specify a step for scanning ports"
	ARG_PORTER_USAGE_PORT    = "-p | --port to specify a single port to scan"
	ARG_PORTER_USAGE_OUTPUT  = "-o | --output to specify output file"
	ARG_PORTER_USAGE_FORMAT  = "-f | --format to specify output format"
        ARG_PORTER_USAGE_MODE    = "-m | --mode to specify scan mode"
)

const (
	ARG_PORTER_DEFAULT_HOST    string = "127.0.0.1"
	ARG_PORTER_DEFAULT_NETWORK string = "tcp"
	ARG_PORTER_DEFAULT_START   int    = 1
	ARG_PORTER_DEFAULT_END     int    = (1 << 16) - 1
	ARG_PORTER_DEFAULT_RANGE   string = "1"
)

// options are in order so take care.
var (
	ARG_MANUAL = fmt.Sprintf(`
+-+-+-+-+-+-+
|p|o|r|t|e|r|
+-+-+-+-+-+-+

Porter is a simple cli tool used for scanning ports on a host.
This tool uses a dialup method to scan a given port with the 
provided network type.

use %s 
use %s
if not provided localhost (127.0.0.1) will be used by default
use %s 
By default tcp netowrk is used
available options for a network type is tcp, udp and ip
use %s
start port must be 1, if not provided by default it is 1
use %s
end or the maximum port number to scan. The maximum port number to
scan is 65535, if not provided the default will be 65535
use %s
This format provided an easy representation of start and end option.
Inorder to use this option it is must that to provide a range syntax
using "..", Eg. 1..80, in this option 80 is exclusive
using "..=", Eg. 1..=80, specifies 80 is inclusive.
use %s
By using this option it is possible to skip a port to scan.
This argument can take a single or multiple ports.
use %s
With the given start and end values this argument is used to
make a step or jump between ports. By defualt its 1 so that no
jump or step is made.
use %s
use %s
use %s
Format option supports the result output as a standart text or json result.
use %s
whether to scan in a single or concurrent thread
    `,

		ARG_PORTER_USAGE_HELP,
		ARG_PORTER_USAGE_HOST,
		ARG_PORTER_USAGE_NETWORK,
		ARG_PORTER_USAGE_START,
		ARG_PORTER_USAGE_END,
		ARG_PORTER_USAGE_RANGE,
		ARG_PORTER_USAGE_SKIP,
		ARG_PORTER_USAGE_STEP,
		ARG_PORTER_USAGE_PORT,
		ARG_PORTER_USAGE_OUTPUT,
		ARG_PORTER_USAGE_FORMAT,
                ARG_PORTER_USAGE_MODE,
	)
)
