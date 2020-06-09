package githubInfo

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
)

func getDetails(username string) UserInfo{
	client := github.NewClient(nil)
	ctx := context.Background()
	//list all organizations for user GabielFemi
	orgs, _, err := client.Organizations.List(ctx, username, nil)
	if err != nil {
		logrus.Fatalln(err)
	}
	var details UserInfo
	details.Username = username
	details.NumberOfJoinedOrganizations = len(orgs)
	logrus.Println(username, " has joined ", len(orgs), " GitHub organizations ")
	return details
}