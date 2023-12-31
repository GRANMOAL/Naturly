package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/GRANMOAL/Naturly/awsgo"
	"github.com/GRANMOAL/Naturly/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(SecretName string) (models.Secret, error) {
	var datosSecret models.Secret
	fmt.Println("> Pido Secreto" + SecretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(SecretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Lectura de Secret OK " + SecretName)
	return datosSecret, nil
}
