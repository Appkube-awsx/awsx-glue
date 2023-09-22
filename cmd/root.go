/*
	Copyright © 2023 Afreen khan <afreen.khan@synectiks.com>
*/
package cmd

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-glue/cmd/gluecmd"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/spf13/cobra"
)
// AwsxGlueCmd represents the base command when called without any subcommands
var AwsxGlueCmd = &cobra.Command{
	Use:   "getGlueDetails",
	Short: "getGlueDetails command gets resource counts",
	Long:  `getGlueDetails command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command get glue details started")

		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			GetListGlue(*clientAuth)
		} else {
			cmd.Help()
			return
		}
	},
}

// json.Unmarshal
func GetListGlue(auth client.Auth) (*glue.ListJobsOutput, error) {

	log.Println("getting glue job  list summary")

	listClient := client.GetClient(auth, client.GLUE_CLIENT).(*glue.Glue)

	listRequest := &glue.ListJobsInput{}

	listResponse, err := listClient.ListJobs(listRequest)

	if err != nil {
		log.Fatalln("Error:in getting glue list", err)
	}
	log.Println(listResponse)

	return listResponse, err
}


func Execute() {
	err := AwsxGlueCmd.Execute()

	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	AwsxGlueCmd.AddCommand(gluecmd.GetConfigDataCmd)
	
	AwsxGlueCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxGlueCmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxGlueCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxGlueCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxGlueCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxGlueCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxGlueCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	AwsxGlueCmd.PersistentFlags().String("externalId", "", "aws external id auth")
	
}
