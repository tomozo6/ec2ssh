package pkg

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

type App struct {
	user         string
	ssmEnabled   bool
	cfg          aws.Config
	SSMInstances []instanceInfo
	EC2Instances []instanceInfo
}

type instanceInfo struct {
	Name       string
	InstanceId string
	Ip         string
}

func NewApp() (*App, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	d := &App{
		cfg:        cfg,
		user:       viper.GetString("ssh-user"),
		ssmEnabled: viper.GetBool("session-manager"),
	}
	return d, nil
}

func (d *App) GetSSMinstancesInfo() error {
	client := ssm.NewFromConfig(d.cfg)
	input := &ssm.DescribeInstanceInformationInput{}

	result, err := getSSMInstances(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		fmt.Println(err)
		return err
	}

	for _, v := range result.InstanceInformationList {
		if v.Name != nil {
			var e instanceInfo
			e.Name = *v.Name
			e.InstanceId = *v.InstanceId
			e.Ip = *v.IPAddress
			d.SSMInstances = append(d.SSMInstances, e)
		}
	}
	return nil
}

func (d *App) GetEC2instancesInfo() error {
	client := ec2.NewFromConfig(d.cfg)
	input := &ec2.DescribeInstancesInput{}

	result, err := getInstances(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		fmt.Println(err)
		return err
	}

	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			var e instanceInfo
			e.InstanceId = *i.InstanceId
			e.Ip = *i.PrivateIpAddress

			for _, t := range i.Tags {
				if *t.Key == "Name" {
					e.Name = *t.Value
				}
			}
			d.EC2Instances = append(d.EC2Instances, e)
		}
	}
	return nil
}

func (d *App) Ssh() {
	var e []instanceInfo
	var t string
	if d.ssmEnabled {
		e = append(d.EC2Instances, d.SSMInstances...)
		t = "InstanceId"
	} else {
		e = d.EC2Instances
		t = "Ip"
	}

	templates := &promptui.SelectTemplates{
		Label: "{{ . }}?",
		// Active:   "\U0001F336 {{ .Name | cyan }} ({{ .InstanceId | red }})",
		Active:   fmt.Sprintf("\U0001F449 {{ .Name | cyan }} ({{ .%s | red }})", t),
		Inactive: fmt.Sprintf("{{ .Name | cyan }} ({{ .%s | red }})", t),
		Selected: "\U0001F449 {{ .Name | green | cyan }}",
		Details: `
	--------- Instance ----------
	{{ "Name:" | faint }}   {{ .Name }}
	{{ "InstanceId:" | faint }}  {{ .InstanceId }}
	{{ "Ip:" | faint }}    {{ .Ip }}`,
	}

	searcher := func(input string, index int) bool {
		v := e[index]
		name := strings.Replace(strings.ToLower(v.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Which instance",
		Items:     e,
		Templates: templates,
		Size:      15,
		Searcher:  searcher,
	}

	// 実行処理、Prompt同様、特定の操作でエラーが返ってくるので、同様のエラー処理が必要になります。
	// 戻り値は、第1は選択したデータのindex、第2は選択した行の文字列が返ってきます。
	i, _, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
	}

	var arg string
	if d.ssmEnabled {
		if d.user == "" {
			arg = e[i].InstanceId
		} else {
			arg = d.user + "@" + e[i].InstanceId
		}
	} else {
		if d.user == "" {
			arg = e[i].Ip
		} else {
			arg = d.user + "@" + e[i].Ip
		}
	}

	fmt.Println("ssh " + arg)

	cmd := exec.Command("ssh", arg)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}