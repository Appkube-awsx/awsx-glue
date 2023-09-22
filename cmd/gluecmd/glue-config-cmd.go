package gluecmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)

		if err != nil {
			cmd.Help()
			return
		}
	
		if authFlag {
			jobName, _ := cmd.Flags().GetString("jobName")

			if jobName != "" {
				GetGlueDetails(jobName, *clientAuth)
			} else {
				log.Fatalln("jobName not provided. Program exit")
			}
		}
	},
}

func GetGlueDetails(jobName string, auth client.Auth) (*glue.GetJobOutput, error) {

	log.Println("Getting aws job Name data")

	glueClient := client.GetClient(auth, client.GLUE_CLIENT).(*glue.Glue)

	input := &glue.GetJobInput{
		JobName: aws.String(jobName),
		}

	glueDetailsResponse, err := glueClient.GetJob(input)

	log.Println(glueDetailsResponse.String())

	if err != nil {
		log.Fatalln("Error:", err)
	}
	return glueDetailsResponse,err
}

func init() {
	GetConfigDataCmd.Flags().StringP("jobName", "t", "", "job name")

	if err := GetConfigDataCmd.MarkFlagRequired("jobName"); err != nil {
		fmt.Println(err)
	}
	
}
