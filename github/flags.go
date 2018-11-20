package github

import "gopkg.in/urfave/cli.v2"

func Flags() []cli.Flag {
	return []cli.Flag{
		TagNameFlag(),
		RefFlag(),
	}
}

func TagNameFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:  "tag-name",
		Usage: "a tag name of a GitHub release",
	}
}

func RefFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:  "ref",
		Usage: "a branch name or hash to be operated",
	}
}
