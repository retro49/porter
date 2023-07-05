module github.com/retro49/porter

go 1.20

replace github.com/retro49/porter/plogger => ./plogger

require (
	github.com/akamensky/argparse v1.4.0
	github.com/retro49/porter/plogger v0.0.0-00010101000000-000000000000
	github.com/retro49/porter/scanner v0.0.0-00010101000000-000000000000
)

replace github.com/retro49/porter/scanner => ./scanner
