package gluecmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-glue/authenticator"
	"github.com/Appkube-awsx/awsx-glue/client"
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

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
	
		if authFlag {
			jobName, _ := cmd.Flags().GetString("jobName")
			if jobName != "" {
				getGlueDetails(region, crossAccountRoleArn, acKey, secKey, jobName, externalId)
			} else {
				log.Fatalln("jobName not provided. Program exit")
			}
		}
	},
}

func getGlueDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string,  externalId string,jobName string) *glue.GetJobOutput {
	log.Println("Getting aws job Name data")
	listClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	input := &glue.GetJobInput{
		JobName: aws.String(jobName),
		}
	glueDetailsResponse, err := listClient.GetJob(input)
	log.Println(glueDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return glueDetailsResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("jobName", "t", "", "job name")

	if err := GetConfigDataCmd.MarkFlagRequired("jobName"); err != nil {
		fmt.Println(err)
	}
	
}
