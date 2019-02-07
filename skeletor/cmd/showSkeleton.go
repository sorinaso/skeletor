package cmd

import (
	"errors"
	"fmt"
	"skeletor"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

var SkeletonShowCmd = &cobra.Command{
	Use:   "show-skeleton [skeleton_name]",
	Short: "Shows one skeleton",
	Long: `Show one skeleton

if you want to show one skeleton information run this command
			`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("skeleton-show requires one arg(The skeleton name)")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		project, err := skeletor.NewProject("skeletor.yml")

		if err != nil {
			panic(err.Error())
		}

		if skeleton, ok := project.GetSkeletons()[args[0]]; ok {
			fmt.Println("Skeleton: ", skeleton.Name)
			fmt.Println("")
			spew.Dump(skeleton)
		} else {
			panic(errors.New("The skeleton " + args[0] + " doesn't exist."))
		}
	},
}

func init() {
	RootCmd.AddCommand(SkeletonShowCmd)
}
