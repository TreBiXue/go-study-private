package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

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

var rootDir string

// LOCAL
func init() {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("os.Getwd failed:%v", err)
	}
	cfg, err := ini.Load(filepath.Join(rootDir, "asset", "config.ini"))
	if err != nil {
		fmt.Printf("error %v", err)
	}
	spannerConfig = SpannerDBConfigList{
		ProjectID:  cfg.Section("spanner_setting").Key("SPANNER_PROJECT_ID").String(),
		InstanceID: cfg.Section("spanner_setting").Key("SPANNER_INSTANCE_ID").String(),
		DBName:     cfg.Section("spanner_setting").Key("SPANNER_DATABASE_ID").String(),
	}
}

// TODO :DEV
// func init() {
// 	spannerConfig = SpannerDBConfigList{
// 		ProjectID:  os.Getenv("SPANNER_PROJECT_ID"),
// 		InstanceID: os.Getenv("SPANNER_INSTANCE_ID"),
// 		DBName:     os.Getenv("SPANNER_DATABASE_ID"),
// 	}
// }

func InitSpannerDB() *spanner.Client {
	ctx := context.Background()
	dbPath := fmt.Sprintf("projects/%s/instances/%s/databases/%s",
		spannerConfig.ProjectID, spannerConfig.InstanceID, spannerConfig.DBName)
	fmt.Println(dbPath)
	//TODO :DEV
	// client, err := spanner.NewClient(ctx, dbPath)

	//LOCAL
	// client, err := spanner.NewClient(ctx, dbPath, option.WithCredentialsFile("../asset/admin_api_serviceaccount.json"))
	client, err := spanner.NewClient(ctx, dbPath, option.WithCredentialsFile(filepath.Join(rootDir, "asset", "admin_api_serviceaccount.json")))
	if err != nil {
		fmt.Printf("error %v", err)
	}
	return client

}
