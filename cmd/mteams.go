package cmd

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/hugo-lorenzo-mato/go-notify/mteams"
	"github.com/hugo-lorenzo-mato/go-notify/util"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"time"
)

var Cmds []cli.Command

var hookUrl string
var cardTitle string
var cardText string
var uriAction string
var planKey string

func init() {
	hookurlflag := cli.StringFlag{
		Name:        "hookUrl",
		Destination: &hookUrl,
		Usage:       "Channel webhook URL",
	}
	cardTitleFlag := cli.StringFlag{
		Name:        "cardTitle",
		Destination: &cardTitle,
		Usage:       "Card title",
	}
	cardTextflag := cli.StringFlag{
		Name:        "cardText",
		Destination: &cardText,
		Usage:       "Card text",
	}
	planKeyflag := cli.StringFlag{
		Name:        "planKey",
		Destination: &planKey,
		Usage:       "Plan key",
	}
	uriActionflag := cli.StringFlag{
		Name:        "uriAction",
		Destination: &uriAction,
		Usage:       "URI action",
	}
	messageCardFlags := []cli.Flag{
		hookurlflag,
		cardTitleFlag,
		cardTextflag,
		uriActionflag,
		planKeyflag,
	}
	messageCard := cli.Command{
		Name:   "messageCard",
		Usage:  "Report any type of information to Microsoft Teams Channel using messageCard card type",
		Action: messageCardToTeams,
		Flags:  messageCardFlags,
	}
	AdaptiveCard := cli.Command{
		Name:   "adaptiveCard",
		Usage:  "Report any type of information to Microsoft Teams Channel using adaptiveCard card type",
		Action: adaptiveCardToTeams,
	}
	cmd := cli.Command{
		Name:        "microsoftTeams",
		Usage:       "Microsoft Teams Channel report options",
		Flags:       []cli.Flag{},
		Subcommands: []cli.Command{messageCard, AdaptiveCard},
	}
	Cmds = append(Cmds, cmd)
}

func sendMessageCard() error {
	log.Debugf("    sendTheMessage - Started")
	var webhookUrl string
	mClient, err := mteams.NewClient()
	if err != nil {
		return errors.Wrap(err, "Error in mteams.NewClient fucntion - sendMessageCard section")
	}
	if mteams.IsValidWebhookURL(hookUrl) {
		webhookUrl = hookUrl
	} else {
		return errors.Errorf("Not valid hookurl, please, check it - %s", hookUrl)
	}
	card, err := mteams.CreateNewMessageCard()
	if err != nil {
		return errors.Wrap(err, "Error in mteams.CreateNewMessageCard fucntion - sendMessageCard section")
	}
	card.Title = cardTitle
	currenttTime := time.Now()
	for i, section := range card.Sections {
		if section.ActivitySubtitle == "DATE" {
			card.Sections[i].ActivitySubtitle = currenttTime.Format("02-01-2006 15:04:05 Monday")
			for j, fact := range card.Sections[i].Facts {
				if fact.Value == "PLAN" {
					fmt.Println("Entramos y...", planKey)
					card.Sections[i].Facts[j].Value = planKey
				}
			}
			card.Sections[i].Text = cardText
		}
	}
	log.Debug(card)
	card.PotentialAction[0].Targets[0].URI = uriAction
	webhookMessage, _ := mteams.MarshalMessageAlert(card)
	log.Debugf("    sendTheMessage - Finished")
	return mClient.Send(webhookUrl, webhookMessage)
}

func messageCardToTeams(ctx *cli.Context) error {
	log.Debugf("messageCardToTeams - Started")
	err := sendMessageCard()
	util.CheckErr(err)
	log.Debugf("messageCardToTeams - Finished")
	return nil
}

func sendAdaptiveCard() error {
	fmt.Println("Method not yet implemented")
	return nil
}

func adaptiveCardToTeams(ctx *cli.Context) error {
	log.Debugf("adaptiveCardToTeams - Started")
	err := sendAdaptiveCard()
	util.CheckErr(err)
	log.Debugf("adaptiveCardToTeams - Finished")
	return nil
}
