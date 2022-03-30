package server

import (
	"context"
	"errors"
	"time"

	pb "github.com/mragungsetiaji/grpc-example/protocol"
	"google.golang.org/grpc"
)

// startJobOnWorker translates the http start request to grpc
// request on the workers.
// Returns:
// 		- string: job id
// 		- error: nil if no error
func StartJobOnWorker(req APIStartJobReq) (string, error) {
	workersMutex.Lock()
	defer workersMutex.Unlock()

	worker, ok := workers[req.WorkerID]
	if !ok {
		return "", errors.New("worker not found")
	}

	conn, err := grpc.Dial(worker.addr, grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()
	c := pb.NewWorkerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	startJobReq := pb.StartJobReq{
		Command: req.Command,
		Path:    req.Path,
	}

	r, err := c.StartJob(ctx, &startJobReq)
	if err != nil {
		return "", err
	}

	return r.JobID, nil
}
