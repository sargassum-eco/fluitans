package ztcontrollers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// All Controllers

func (c *Client) GetControllers() ([]Controller, error) {
	// TODO: make these methods of a controllers client object
	// TODO: look up the controllers from a database, if one is specified!
	controllers := make([]Controller, 0)
	envController := c.Config.Controller

	if envController != nil {
		controllers = append(controllers, *envController)
	}
	return controllers, nil
}

func (c *Client) ScanControllers(ctx context.Context, controllers []Controller) ([]string, error) {
	eg, ctx := errgroup.WithContext(ctx)
	addresses := make([]string, len(controllers))
	for i, controller := range controllers {
		eg.Go(func(i int) func() error {
			return func() error {
				client, cerr := controller.NewClient()
				if cerr != nil {
					return nil
				}

				res, err := client.GetStatusWithResponse(ctx)
				if err != nil {
					return err
				}
				addresses[i] = *res.JSON200.Address
				return nil
			}
		}(i))
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}

	for i, v := range controllers {
		err := c.Cache.SetControllerByAddress(addresses[i], v)
		if err != nil {
			return nil, err
		}
	}

	return addresses, nil
}

// Individual Controller

func (c *Client) FindController(name string) (*Controller, bool, error) {
	found := false
	controllers, err := c.GetControllers()
	if err != nil {
		return nil, false, err
	}

	for _, v := range controllers {
		if v.Name == name {
			return &v, true, nil
		}
	}
	return nil, found, nil
}

func (c *Client) checkCachedController(ctx context.Context, address string) (*Controller, error) {
	controller, cacheHit, err := c.Cache.GetControllerByAddress(address)
	if err != nil {
		return nil, err
	}

	if !cacheHit {
		return nil, nil
	}

	client, cerr := controller.NewClient()
	if cerr != nil {
		return nil, cerr
	}

	res, err := client.GetStatusWithResponse(ctx)
	if err != nil {
		// evict the cached controller if its authtoken is stale or it no longer responds
		c.Cache.UnsetControllerByAddress(address)
		return nil, err
	}

	if address != *res.JSON200.Address { // cached controller's address is stale
		c.Logger.Warnf(
			"zerotier controller %s's address has changed from %s to %s",
			controller.Server, address, *res.JSON200.Address,
		)
		err = c.Cache.SetControllerByAddress(*res.JSON200.Address, *controller)
		if err != nil {
			return nil, err
		}

		c.Cache.UnsetControllerByAddress(address)
		return nil, nil
	}

	return controller, nil
}

func (c *Client) FindControllerByAddress(ctx context.Context, address string) (*Controller, error) {
	controller, err := c.checkCachedController(ctx, address)
	if err != nil {
		// Log the error and proceed to manually query all controllers
		c.Logger.Error(err, errors.Wrap(err, fmt.Sprintf(
			"couldn't handle the cache entry for the zerotier controller with address %s", address,
		)))
	} else if controller != nil {
		return controller, nil
	}

	// Query the list of all known controllers
	controllers, err := c.GetControllers()
	if err != nil {
		return nil, err
	}

	c.Logger.Warnf(
		"rescanning zerotier controllers due to a stale/missing controller for %s in cache", address,
	)
	addresses, err := c.ScanControllers(ctx, controllers)
	if err != nil {
		return nil, err
	}

	for i, v := range controllers {
		if addresses[i] == address {
			return &v, nil
		}
	}

	return nil, echo.NewHTTPError(
		http.StatusNotFound, fmt.Sprintf("zerotier controller not found with address %s", address),
	)
}
