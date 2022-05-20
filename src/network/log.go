package network

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/signal"
	"syscall"
)

var reader io.ReadCloser

func GetDockerClient() (*client.Client, error) {
	opts, err := client.NewClientWithOpts()
	if err != nil {
		return nil, err
	}
	return opts, nil
}

func retrieveLogStream(containerName, tail string, follow bool) (io.ReadCloser, error) {
	docClient, err := GetDockerClient()
	if err != nil {
		return nil, err
	}
	return docClient.ContainerLogs(context.Background(), containerName, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: false,
		Follow:     follow,
		Tail:       tail,
	})
}

func getReadCloser(containerName, tail string, follow bool) io.ReadCloser {
	if reader == nil {
		tempReader, err := retrieveLogStream(containerName, tail, follow)
		if err != nil {
			cobra.CompErrorln(err.Error())
			return nil
		}
		reader = tempReader
	}
	return reader
}

func ReadLog(containerName, tail string, follow bool) {
	stream := getReadCloser(containerName, tail, follow)
	if stream == nil {
		fmt.Printf("can't read log for %s. Check if the container is running or not\n", containerName)
		return
	}
	streamReader := make([]byte, 20)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-c
		stream.Close()
		os.Exit(0)
	}()

	for {
		n, err := stream.Read(streamReader)
		if err != nil {
			if err == io.EOF {
				fmt.Print(string(streamReader[:n]))
			}
			stream.Close()
			return
		}
		fmt.Print(string(streamReader[:n]))
	}
}
