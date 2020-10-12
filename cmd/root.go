package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd) //go run main.go word --str=kevintest --mode=1
	rootCmd.AddCommand(timeCmd) //go run main.go time calc -c="2009-09-15 12:02:03" -d=5h
	rootCmd.AddCommand(sqlCmd)  //go run main.go sql struct --username root --password 123456 --dbType mysql --db tour --table "blog_tag"
}
