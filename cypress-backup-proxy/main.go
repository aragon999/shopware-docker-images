package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var ctx context.Context
var dClient *client.Client

func init() {
	ctx = context.Background()

	var err error

	dClient, err = client.NewEnvClient()
	if err != nil {
		panic(err)
	}
}

func restoreHandler(w http.ResponseWriter, r *http.Request) {
	host := strings.Split(r.Host, ".")[0]

	c, err := getMysqlContainer()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}

	log.Printf("Performing restoring for %s\n", host)

	execConfig := types.ExecConfig{
		AttachStderr: true,
		AttachStdin:  true,
		AttachStdout: true,
		Cmd:          []string{"bash", "-c", fmt.Sprintf("mysql -uroot -p%s < /tmp/%s.sql", os.Getenv("MYSQL_ROOT_PASSWORD"), host)},
		Tty:          true,
		Detach:       false,
	}

	create, err := dClient.ContainerExecCreate(ctx, c.ID, execConfig)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}

	execAttachConfig := types.ExecStartCheck{
		Detach: false,
		Tty:    true,
	}

	containerResponse, err := dClient.ContainerExecAttach(ctx, create.ID, execAttachConfig)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		fmt.Fprintln(w, err)
		return
	}

	defer containerResponse.Close()

	data, _ := ioutil.ReadAll(containerResponse.Reader)
	log.Printf("Output: %s\n", string(data))

	log.Printf("Completed restoring for %s\n", host)

	fmt.Fprint(w, "success")
}

func backupHandler(w http.ResponseWriter, r *http.Request) {
	host := strings.Split(r.Host, ".")[0]

	c, err := getMysqlContainer()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}

	log.Printf("Performing backup for %s\n", host)

	execConfig := types.ExecConfig{
		AttachStderr: true,
		AttachStdin:  true,
		AttachStdout: true,
		Cmd:          []string{"bash", "-c", fmt.Sprintf("mysqldump --add-drop-database -uroot -p%s --databases %s > /tmp/%s.sql", os.Getenv("MYSQL_ROOT_PASSWORD"), host, host)},
		Tty:          true,
		Detach:       false,
	}

	create, err := dClient.ContainerExecCreate(ctx, c.ID, execConfig)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}

	execAttachConfig := types.ExecStartCheck{
		Detach: false,
		Tty:    true,
	}

	containerResponse, err := dClient.ContainerExecAttach(ctx, create.ID, execAttachConfig)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}

	defer containerResponse.Close()

	data, _ := ioutil.ReadAll(containerResponse.Reader)
	log.Println(string(data))

	fmt.Fprint(w, "success")
}

func main() {
	http.HandleFunc("/backup", backupHandler)
	http.HandleFunc("/restore", restoreHandler)
	http.HandleFunc("/cleanup", restoreHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func getMysqlContainer() (*types.Container, error) {
	opts := types.ContainerListOptions{All: true}
	opts.Filters = filters.NewArgs()
	opts.Filters.Add("name", "shopware-docker-mysql-1")

	containers, err := dClient.ContainerList(context.Background(), opts)

	if err != nil {
		return nil, err
	}

	for _, container := range containers {
		return &container, nil
	}

	opts = types.ContainerListOptions{All: true}
	opts.Filters = filters.NewArgs()
	opts.Filters.Add("name", "shopware-docker_mysql_1")

	containers, err = dClient.ContainerList(context.Background(), opts)

	if err != nil {
		return nil, err
	}

	for _, container := range containers {
		return &container, nil
	}

	return nil, errors.New("cannot find mysql container")
}
