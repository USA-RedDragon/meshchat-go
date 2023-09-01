package cmd

import (
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/USA-RedDragon/meshchat-go/internal/http"
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
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		return err
	}
	server := http.NewServer(port)
	if err := server.Start(); err != nil {
		return err
	}

	stop := func(sig os.Signal) {
		wg := new(sync.WaitGroup)

		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = server.Stop()
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
