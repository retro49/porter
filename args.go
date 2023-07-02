package main

import "fmt"

const (
	ARG_HELP_PORTER_HELP    = "-h | --help to display this and exit"
	ARG_HELP_PORTER_HOST    = "-H | --host to specify host to scan"
	ARG_HELP_PORTER_NETWORK = "-n | --network  to specify network type"
	ARG_HELP_PORTER_START   = "-S | --start to specify start port address"
	ARG_HELP_PORTER_END     = "-E | --end to specify ending port address"
	ARG_HELP_PORTER_RANGE   = "-r | --range to specify a port as range of values"
	ARG_HELP_PORTER_SKIP    = "-k | --skip to specify port not to scan"
	ARG_HELP_PORTER_STEP    = "-e | --step to specify a step for scanning ports"
	ARG_HELP_PORTER_PORT    = "-p | --port to specify a single port to scan"
	ARG_HELP_PORTER_OUTPUT  = "-o | --output to specify output file"
	ARG_HELP_PORTER_FORMAT  = "-f | --format to specify output format"
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
use %s 
available options for a network type is tcp, udp and ip
use %s
start port must be 1, if not provided by default it is 1
use %s
end or the maximum port number to scan. The maximum port number to
scan is 65535, if not provided the default will be 65535
use %s
This format provided an easy representation of start and end option.
Inorder to use this option it is must that to provide a range syntax
using "-", Eg. 1-80, in this option 80 is exclusive
using "-=", Eg. 1-=80, specifies 80 is inclusive.
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
    `,

		ARG_HELP_PORTER_HELP,
		ARG_HELP_PORTER_HOST,
		ARG_HELP_PORTER_NETWORK,
		ARG_HELP_PORTER_START,
		ARG_HELP_PORTER_END,
		ARG_HELP_PORTER_RANGE,
		ARG_HELP_PORTER_SKIP,
		ARG_HELP_PORTER_STEP,
		ARG_HELP_PORTER_PORT,
		ARG_HELP_PORTER_OUTPUT,
		ARG_HELP_PORTER_FORMAT,
	)
)
