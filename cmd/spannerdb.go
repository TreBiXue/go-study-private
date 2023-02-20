package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"google.golang.org/api/option"
	"gopkg.in/ini.v1"
)

type SpannerDBConfigList struct {
	ProjectID  string
	InstanceID string
	DBName     string
}

var spannerConfig SpannerDBConfigList

func init() {
	cfg, err := ini.Load("../assets/config.ini")
	if err != nil {
		fmt.Printf("error %v", err)
	}
	spannerConfig = SpannerDBConfigList{
		ProjectID:  cfg.Section("spanner_setting").Key("SPANNER_PROJECT_ID").String(),
		InstanceID: cfg.Section("spanner_setting").Key("SPANNER_INSTANCE_ID").String(),
		DBName:     cfg.Section("spanner_setting").Key("SPANNER_DATABASE_ID").String(),
	}
}

func InitSpannerDB() *spanner.Client {
	ctx := context.Background()
	dbPath := fmt.Sprintf("projects/%s/instances/%s/databases/%s",
		spannerConfig.ProjectID, spannerConfig.InstanceID, spannerConfig.DBName)
	fmt.Println(dbPath)
	client, err := spanner.NewClient(ctx, dbPath, option.WithCredentialsFile("../assets/admin_api_serviceaccount.json"))
	if err != nil {
		fmt.Printf("error %v", err)
	}
	return client

}
