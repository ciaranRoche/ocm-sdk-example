package main

import (
	"fmt"
	"log"

	"github.com/openshift-online/ocm-sdk-go"
	v1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
)

func main() {
	logger, err := sdk.NewGoLoggerBuilder().
		Debug(true).
		Build()
	if err != nil {
		log.Fatalf("can't build logger: %v\n", err)
	}

	// https://clusters-service.apps-crc.testing
	// https://api.stage.openshift.com
	// https://api.integration.openshift.com
	// Create the connection, and remember to close it:
	connection, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens("<<OFFLINE_ACCESS_TOKEN>>").
		URL("https://api.stage.openshift.com").
		Insecure(true).
		Build()
	if err != nil {
		log.Fatalf("Can't build connection: %v", err)
	}
	defer connection.Close()

	// example using client id and secret to create sdk connection
	//connection, err := sdk.NewConnectionBuilder().
	//	Logger(logger).
	//	Client("xxxxxxxxx", "xxxxxx-xxxxxx-xxxxxxx-xxxxxxx-xxxx").
	//	URL("https://api.integration.openshift.com").
	//	Insecure(true).
	//	Build()
	//if err != nil {
	//	log.Fatalf("Can't build connection: %v", err)
	//}
	//defer connection.Close()

	clusterName := "my-test-cluster"
	cluster, err := v1.NewCluster().
		Name(clusterName).
		Region(v1.NewCloudRegion().ID("eu-west-1")).
		Properties(map[string]string{
			"fake_cluster": "true",
		}).
		Build()
	if err != nil {
		log.Fatalf("could not build cluster : %v", err)
	}

	clusterResp, err := connection.ClustersMgmt().V1().Clusters().Add().Body(cluster).Send()
	if err != nil {
		log.Fatalf("could not create cluster : %v", err)
	}

	clusterID := clusterResp.Body().ID()
	fmt.Println("cluster created - ", clusterID)

	clusterGetResp, err := connection.ClustersMgmt().V1().Clusters().Cluster(clusterID).Get().Send()
	if err != nil {
		log.Fatalf("could not get cluster : %v", err)
	}

	fmt.Println("got cluster - ", clusterGetResp.Body().ID())

	clusterDeleteResp, err := connection.ClustersMgmt().V1().Clusters().Cluster(clusterID).Delete().Send()
	if err != nil {
		log.Fatalf("could not delete cluster : %v", err)
	}

	fmt.Println("deleted cluster - ", clusterDeleteResp.Status())

}
