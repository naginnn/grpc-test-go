package main

import (
	"context"
	pb "github.com/naginnn/GoPythonGrpcTest/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pbtime "google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"time"
)

func main() {
	addr := "localhost:9999"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewOutliersClient(conn)
	//s := "cluster_config"

	req := &pb.ReportType{
		Report: pb.Reports_ClusterConfigMetric,
		Start:  time.Now().Unix(),
		//Param: wrapperspb.String("Bomba"),
		Param: "hdfs-name",
	}

	resp, err := client.Report(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("report: %v", len(resp.Report))

	//req := &pb.OutliersRequest{
	//	Metrics: dummyData(),
	//}
	//
	//resp, err := client.Detect(context.Background(), req)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//log.Printf("outliers at: %v", resp.Indices)
}

func dummyData() []*pb.Metric {
	const size = 1000
	out := make([]*pb.Metric, size)
	t := time.Date(2020, 5, 22, 14, 13, 11, 0, time.UTC)
	for i := 0; i < size; i++ {
		m := pb.Metric{
			Time: Timestamp(t),
			Name: "CPU",
			// normally we're below 40% CPU utilization
			Value: rand.Float64() * 40,
		}
		out[i] = &m
		t.Add(time.Second)
	}
	// Create some outliers
	out[7].Value = 97.3
	out[113].Value = 92.1
	out[835].Value = 93.2
	return out
}

// Timestamp converts time.Time to protobuf *Timestamp
func Timestamp(t time.Time) *pbtime.Timestamp {
	return &pbtime.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
