package cmd

import (
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/ztrue/shutdown"
)

//nolint:golint,gochecknoglobals
var (
	serverCmd = &cobra.Command{
		Use:               "server",
		Short:             "Start the daemon server",
		RunE:              runServer,
		SilenceErrors:     true,
		DisableAutoGenTag: true,
	}
)

//nolint:golint,gochecknoinits
func init() {
	serverCmd.Flags().IntP("port", "p", 3333, "port to listen on")
	RootCmd.AddCommand(serverCmd)
}

func runServer(cmd *cobra.Command, _ []string) error {
	stop := func(sig os.Signal) {
		wg := new(sync.WaitGroup)

		wg.Add(1)
		go func() {
			defer wg.Done()
		}()

		const timeout = 10 * time.Second
		c := make(chan struct{})
		go func() {
			defer close(c)
			wg.Wait()
		}()

		select {
		case <-c:
			os.Exit(0)
		case <-time.After(timeout):
			os.Exit(1)
		}
	}
	shutdown.AddWithParam(stop)
	shutdown.Listen(syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT)

	return nil
}
