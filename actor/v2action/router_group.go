package v2action

import (
	"code.cloudfoundry.org/cli/actor/actionerror"

	"code.cloudfoundry.org/cli/api/router"
)

type RouterGroup router.RouterGroup

func (actor Actor) GetRouterGroupByName(routerGroupName string, client RouterClient) (RouterGroup, error) {
	routerGroups, err := client.GetRouterGroupsByName(routerGroupName)
	if err != nil {
		return RouterGroup{}, err
	}

	for _, routerGroup := range routerGroups {
		if routerGroup.Name == routerGroupName {
			return RouterGroup(routerGroup), nil
		}
	}
	return RouterGroup{}, actionerror.RouterGroupNotFoundError{Name: routerGroupName}
}
