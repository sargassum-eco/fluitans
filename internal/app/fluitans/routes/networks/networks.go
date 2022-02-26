package networks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sargassum-eco/fluitans/internal/app/fluitans/auth"
	ztc "github.com/sargassum-eco/fluitans/internal/clients/zerotier"
	"github.com/sargassum-eco/fluitans/internal/clients/ztcontrollers"
	"github.com/sargassum-eco/fluitans/pkg/zerotier"
)

type NetworksData struct {
	Controller ztcontrollers.Controller
	Networks   map[string]zerotier.ControllerNetwork
}

func getNetworksData(
	ctx context.Context, c *ztc.Client, cc *ztcontrollers.Client,
) ([]NetworksData, error) {
	controllers, err := cc.GetControllers()
	if err != nil {
		return nil, err
	}

	networkIDs, err := c.GetAllNetworkIDs(ctx, controllers, cc)
	if err != nil {
		return nil, err
	}

	networks, err := c.GetAllNetworks(ctx, controllers, networkIDs)
	if err != nil {
		return nil, err
	}

	networksData := make([]NetworksData, len(controllers))
	for i, controller := range controllers {
		networksData[i].Controller = controller
		networksData[i].Networks = networks[i]
	}
	return networksData, nil
}

func (s *Service) getNetworks() echo.HandlerFunc {
	t := "networks/networks.page.tmpl"
	s.r.MustHave(t)
	return func(c echo.Context) error {
		// Check authentication & authorization
		a, _, err := auth.GetWithSession(c, s.sc)
		if err != nil {
			return err
		}

		// Run queries
		networksData, err := getNetworksData(c.Request().Context(), s.ztc, s.ztcc)
		if err != nil {
			return err
		}

		// Produce output
		return s.r.CacheablePage(c.Response(), c.Request(), t, networksData, a)
	}
}

func (s *Service) postNetworks() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check authentication & authorization
		if err := auth.RequireAuthorized(c, s.sc); err != nil {
			return err
		}

		// Parse params
		name := c.FormValue("controller")
		if name == "" {
			return echo.NewHTTPError(
				http.StatusBadRequest, "zerotier controller name not specified",
			)
		}

		// Run queries
		controller, err := s.ztcc.FindController(name)
		if err != nil {
			return err
		}
		if controller == nil {
			return echo.NewHTTPError(
				http.StatusNotFound, fmt.Sprintf("zerotier controller %s not found", name),
			)
		}

		createdNetwork, err := s.ztc.CreateNetwork(c.Request().Context(), *controller, s.ztcc)
		if err != nil {
			return err
		}
		created := createdNetwork.Id
		if created == nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "network status unknown")
		}

		// Redirect user
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/networks/%s", *created))
	}
}
