package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/ayupov-ayaz/grpc/pkg/api"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	ErrInvalidArgsCount    = errors.New("invalid args count")
	ErrParseArgToIntFailed = errors.New("parse arg to int failed")
)

func parseStrToUint32(str string) (int32, error) {
	x, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, ErrParseArgToIntFailed
	}

	return int32(x), nil
}

func getArgs() (x, y int32, err error) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	args := strings.Split(sc.Text(), " ")
	if len(args) != 2 {
		return 0, 0, ErrInvalidArgsCount
	}

	x, err = parseStrToUint32(args[0])
	if err != nil {
		return 0, 0, err
	}

	y, err = parseStrToUint32(args[1])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

func run(client api.AdderClient) error {
	req := &api.AddRequest{}

	for {
		x, y, err := getArgs()
		if err != nil {
			return fmt.Errorf("getArgs: %w", err)
		}

		req.X = x
		req.Y = y

		resp, err := client.Add(context.Background(), req)
		if err != nil {
			return err
		}

		fmt.Println("result: ", resp.Result)
	}
}

func main() {
	// grpc.WithInsecure - use only for test client
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	if err := run(api.NewAdderClient(conn)); err != nil {
		log.Fatal(err)
	}
}
