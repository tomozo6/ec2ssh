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
	grepword     string
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
		grepword:   viper.GetString("grep"),
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
		// 名前が空じゃないインスタンスだけ処理する(空なインスタンスはEC2)
		if v.Name != nil {
			// Online状態のインスタンスのみ処理する
			if v.PingStatus == "Online" {
				var e instanceInfo
				e.Name = *v.Name
				e.InstanceId = *v.InstanceId
				e.Ip = *v.IPAddress

				// インスタンス名かインスタンスIDにgrepwordが含まれていたらappend
				if strings.Contains(e.Name, d.grepword) {
					d.SSMInstances = append(d.SSMInstances, e)
				} else if strings.Contains(e.InstanceId, d.grepword) {
					d.SSMInstances = append(d.SSMInstances, e)
				}

			}
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
			// Running状態のEC2のみ処理する
			if i.State.Code == 16 {
				var e instanceInfo
				e.InstanceId = *i.InstanceId
				e.Ip = *i.PrivateIpAddress

				// インスタンス名をネームタグから取得する
				for _, t := range i.Tags {
					if *t.Key == "Name" {
						e.Name = *t.Value
					}
				}

				// インスタンス名かインスタンスIDにgrepwordが含まれていたらappend
				if strings.Contains(e.Name, d.grepword) {
					d.EC2Instances = append(d.EC2Instances, e)
				} else if strings.Contains(e.InstanceId, d.grepword) {
					d.EC2Instances = append(d.EC2Instances, e)
				}
			}
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
		Label:    `{{ " ? " | green | bold }}{{ . | yellow | bold }}`,
		Active:   fmt.Sprintf(`{{ ">" | blue | bold }} {{ .Name | cyan | bold }} ({{ .%s | red | bold }})`, t),
		Inactive: fmt.Sprintf(`  {{ .Name }} ({{ .%s }})`, t),
		Selected: `{{ ">" | blue| bold }} {{ .Name | cyan | bold }}`,
		Details: `
{{ "--------- Instance ----------" | faint }}
{{ "Name:" | faint }}       {{ .Name | faint }}
{{ "InstanceId:" | faint }} {{ .InstanceId | faint }}
{{ "Ip:" | faint }}         {{ .Ip | faint }}`,
	}

	searcher := func(input string, index int) bool {
		v := e[index]
		name := strings.Replace(strings.ToLower(v.Name + v.InstanceId), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Choose Instance:",
		Items:     e,
		Templates: templates,
		Size:      15,
		Searcher:  searcher,
	}

	// 実行処理、Prompt同様、特定の操作でエラーが返ってくるので、同様のエラー処理が必要になります。
	// 戻り値は、第1は選択したデータのindex、第2は選択した行の文字列が返ってきます。

	// Ctrl + c もしくは Ctrl + dの場合はエラーが返ってくるので、型チェックして真の場合は終了
	// それ以外のエラーは出力してから異常終了
	i, _, err := prompt.Run()

	if err == promptui.ErrEOF {
		fmt.Println("Exit ec2ssh, because ctrl+D has been entered.")
		os.Exit(0)
	}

	if err == promptui.ErrInterrupt {
		fmt.Println("Exit ec2ssh, because ctrl+C has been entered.")
		os.Exit(-1)
	}

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
