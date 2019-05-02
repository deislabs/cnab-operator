package swarm

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	swarmtypes "github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"gotest.tools/poll"
)

// NoTasksForService verifies that there are no more tasks for the given service
func NoTasksForService(ctx context.Context, client client.ServiceAPIClient, serviceID string) func(log poll.LogT) poll.Result {
	return func(log poll.LogT) poll.Result {
		tasks, err := client.TaskList(ctx, types.TaskListOptions{
			Filters: filters.NewArgs(
				filters.Arg("service", serviceID),
			),
		})
		if err == nil {
			if len(tasks) == 0 {
				return poll.Success()
			}
			if len(tasks) > 0 {
				return poll.Continue("task count for service %s at %d waiting for 0", serviceID, len(tasks))
			}
			return poll.Continue("waiting for tasks for service %s to be deleted", serviceID)
		}
		// TODO we should not use an error as indication that the tasks are gone. There may be other reasons for an error to occur.
		return poll.Success()
	}
}

// NoTasks verifies that all tasks are gone
func NoTasks(ctx context.Context, client client.ServiceAPIClient) func(log poll.LogT) poll.Result {
	return func(log poll.LogT) poll.Result {
		tasks, err := client.TaskList(ctx, types.TaskListOptions{})
		switch {
		case err != nil:
			return poll.Error(err)
		case len(tasks) == 0:
			return poll.Success()
		default:
			return poll.Continue("waiting for all tasks to be removed: task count at %d", len(tasks))
		}
	}
}

// RunningTasksCount verifies there are `instances` tasks running for `serviceID`
func RunningTasksCount(client client.ServiceAPIClient, serviceID string, instances uint64) func(log poll.LogT) poll.Result {
	return func(log poll.LogT) poll.Result {
		filter := filters.NewArgs()
		filter.Add("service", serviceID)
		tasks, err := client.TaskList(context.Background(), types.TaskListOptions{
			Filters: filter,
		})
		switch {
		case err != nil:
			return poll.Error(err)
		case len(tasks) == int(instances):
			for _, task := range tasks {
				if task.Status.State != swarmtypes.TaskStateRunning {
					return poll.Continue("waiting for tasks to enter run state")
				}
			}
			return poll.Success()
		default:
			return poll.Continue("task count at %d waiting for %d", len(tasks), instances)
		}
	}
}
