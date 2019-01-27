package cmd

import (
	"errors"
	"skeletor"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var SkeletonUpdateCmd = &cobra.Command{
	Use:   "update-skeleton [skeleton_name]",
	Short: "Updates one skeleton",
	Long: `Update one skeleton

if you want to update one skeleton  directory you can run this command.
			`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("skeleton-update requires one arg(The skeleton name)")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		project, err := skeletor.NewProject("config.yml")

		if err != nil {
			panic(err.Error())
		}

		if skeleton, ok := project.GetSkeletons()[args[0]]; ok {
			log.Info("Creating skeleton ", skeleton.Name)

			if err := skeleton.Create(); err != nil {
				panic(err.Error())
			}
		} else {
			panic(errors.New("The skeleton " + args[0] + " doesn't exist."))
		}
	},
}

func init() {
	RootCmd.AddCommand(SkeletonUpdateCmd)
}
