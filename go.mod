module github.com/bf2fc6cc711aee1a0c2a/cli

go 1.12

replace github.com/bf2fc6cc711aee1a0c2a/cli/client/mas => ./client/mas

require (
	github.com/antihax/optional v1.0.0
	github.com/bf2fc6cc711aee1a0c2a/cli/client/mas v0.0.0-00010101000000-000000000000
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/kataras/tablewriter v0.0.0-20180708051242-e063d29b7c23 // indirect
	github.com/landoop/tableprinter v0.0.0-20200805134727-ea32388e35c1
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
)
