package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {

			wantedHost := c.Args().Get(0)
			var wantedIndex int64
			if len(c.Args().Get(1)) == 0 {
				wantedIndex = 0
			} else {
				wantedIndex, _ = strconv.ParseInt(c.Args().Get(1), 10, 64)
			}
			sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("eu-west-1")}))
			ec2Client := ec2.New(sess)

			var input ec2.DescribeInstancesInput

			if strings.HasPrefix(wantedHost, "i-") {
				// this is probably an instance ID
				input = ec2.DescribeInstancesInput{
					InstanceIds: []*string{
						aws.String(wantedHost),
					},
				}
			} else {
				input = ec2.DescribeInstancesInput{
					Filters: []*ec2.Filter{
						&ec2.Filter{
							Name:   aws.String("tag:Name"),
							Values: []*string{aws.String(wantedHost)},
						},
					},
				}
			}

			instances, _ := ec2Client.DescribeInstances(&input)
			if len(instances.Reservations[0].Instances) < int(wantedIndex+1) {
				os.Exit(1)
			}
			instance := instances.Reservations[0].Instances[wantedIndex]
			fmt.Println(*instance.PrivateIpAddress)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
