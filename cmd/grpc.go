package cmd

import (
	"os"

	"github.com/tianopo/pix/application/grpc"
	"github.com/tianopo/pix/infrastructure/db"

	"github.com/spf13/cobra"
)

var portNumber int

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p",50051, "gRPC Server port")
}