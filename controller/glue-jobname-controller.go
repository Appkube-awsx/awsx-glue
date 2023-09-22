package controller

import (
	"github.com/Appkube-awsx/awsx-glue/cmd"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/glue"
	"log"
)

func GetGlueByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) (*glue.ListJobsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetGlueByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetGlueByUserCreds(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) (*glue.ListJobsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
	return GetGlueByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetGlueByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) (*glue.ListJobsOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := cmd.GetListGlue(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetGlue(clientAuth *client.Auth) (*glue.ListJobsOutput, error) {
	response, err := cmd.GetListGlue(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
