package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	replayBufferCmd = &cobra.Command{
		Use:   "replaybuffer",
		Short: "manage replay buffer",
		Long:  `The replaybuffer command manages the replay buffer`,
		RunE:  nil,
	}

	startReplayBufferCmd = &cobra.Command{
		Use:   "start",
		Short: "Starts replay buffer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return startReplayBuffer()
		},
	}

	stopReplayBufferCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stops replay buffer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return stopReplayBuffer()
		},
	}

	saveReplayBufferCmd = &cobra.Command{
		Use:   "save",
		Short: "Saves replay buffer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return saveReplayBuffer()
		},
	}

	replayBufferStatusCmd = &cobra.Command{
		Use:   "status",
		Short: "Reports replay buffer status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return replayBufferStatus()
		},
	}
)

func startReplayBuffer() error {
	_, err := client.ReplayBuffer.StartReplayBuffer()
	return err
}

func stopReplayBuffer() error {
	_, err := client.ReplayBuffer.StopReplayBuffer()
	return err
}

func saveReplayBuffer() error {
	_, err := client.ReplayBuffer.SaveReplayBuffer()
	return err
}

func replayBufferStatus() error {
	r, err := client.ReplayBuffer.GetReplayBufferStatus()
	if err != nil {
		return err
	}

	fmt.Printf("Replay Buffer active: %s\n", strconv.FormatBool(*r.IsReplayBufferActive))
	return nil
}

func init() {
	replayBufferCmd.AddCommand(startReplayBufferCmd)
	replayBufferCmd.AddCommand(stopReplayBufferCmd)
	replayBufferCmd.AddCommand(saveReplayBufferCmd)
	replayBufferCmd.AddCommand(replayBufferStatusCmd)

	rootCmd.AddCommand(replayBufferCmd)
}
