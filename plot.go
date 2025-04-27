package goplot

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

var gnuplotCmd string

func initialize() error {
	gnuplotExecutableName := "gnuplot"

	if runtime.GOOS == "windows" {
		gnuplotExecutableName = "gnuplot.exe"
	}

	var err error
	gnuplotCmd, err = exec.LookPath(gnuplotExecutableName)
	if err != nil {
		return fmt.Errorf("failed to find path to 'gnuplot':\n%v\n", err)
	}
	return nil
}

func compileData(data map[string]data) (string, error) {
	parts := make([]string, 0, len(data))
	for _, dat := range data {
		part, err := dat.compile()
		if err != nil {
			return "", fmt.Errorf("failed to compile data set: %w", err)
		}
		parts = append(parts, part)
	}
	return strings.Join(parts, ", "), nil
}

func (b *builder) DebugPlot() error {
	for _, command := range b.commands {
		fmt.Println(command.compile())
	}
	compiledData, err := compileData(b.data)
	if err != nil {
		return fmt.Errorf("failed to compile data: %w", err)
	}
	fmt.Println("plot " + compiledData)
	fmt.Println("exit")
	return nil
}

func (b *builder) Plot(ctx context.Context) error {
	err := sync.OnceValue(initialize)()
	if err != nil {
		return err
	}

	eg, egCtx := errgroup.WithContext(ctx)

	cmd := exec.CommandContext(egCtx, gnuplotCmd, "-persist")
	ready := make(chan struct{})

	eg.Go(func() error {
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return fmt.Errorf("failed to get stdin pipe: %w", err)
		}
		select {
		case <-ready:
		case <-ctx.Done():
		}
		cmd.Start()
		for _, command := range b.commands {
			_, err := io.WriteString(stdin, command.compile()+"\n")
			if err != nil {
				return fmt.Errorf("failed to set command: %w", err)
			}
		}
		compiledData, err := compileData(b.data)
		if err != nil {
			return fmt.Errorf("failed to compile plot command: %w", err)
		}
		_, err = io.WriteString(stdin, "plot "+compiledData+"\n")
		if err != nil {
			return fmt.Errorf("failed to set command: %w", err)
		}
		_, err = io.WriteString(stdin, "exit\n")
		if err != nil {
			return fmt.Errorf("failed to set command: %w", err)
		}
		cmd.Wait()
		return nil
	})

	eg.Go(func() error {
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return fmt.Errorf("failed to get stderr pipe: %w", err)
		}
		close(ready)
		stderrReader := bufio.NewReader(stderr)

		errmsg := ""
		for {
			line, _, err := stderrReader.ReadLine()
			if err != nil {
				break
			}
			errmsg += string(line) + "\n"
		}
		if len(errmsg) > 0 {
			return fmt.Errorf("gnuplot error: %w", errors.New(errmsg))
		}

		return nil
	})

	err = eg.Wait()
	if err != nil {
		return fmt.Errorf("failed to plot: %w", err)
	}
	return nil
}
