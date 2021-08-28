package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/agent"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/cluster"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/container"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metricendpoint"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metriclabel"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metricname"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/metrictype"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/node"
	"github.com/willie-lin/kubeMonitor/pkg/nextserver/database/ent/proces"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Hello(client *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, "hello, welcome to kubeMonitor!!!")
	}
}

type Controller struct {
	client *ent.Client
}

// FindRemoteAgent find remote agent
func (controller *Controller) FindRemoteAgent() echo.HandlerFunc {
	return func(c echo.Context) error {
		ag := new(ent.Agent)
		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&ag); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return c.JSON(http.StatusBadRequest, "ParseInt: "+err.Error())

		}

		agents, err := controller.client.Agent.Query().Where(agent.MachineIdEQ(ag.MachineId)).First(context.Background())
		if err != nil {
			if ent.IsNotFound(err) {
				c.Logger().Error("Get: ", err)
				log.Fatal("Get: ", zap.Error(err))
				return c.String(http.StatusBadRequest, "Get: "+err.Error())
			}
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.JSON(http.StatusOK, &agents)
	}
}

// FindNodeByAgent find node agent
func (controller *Controller) FindNodeByAgent() echo.HandlerFunc {
	return func(c echo.Context) error {
		nd := new(ent.Node)

		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&nd); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return c.JSON(http.StatusBadRequest, "json: "+err.Error())
		}
		nodes, err := controller.client.Node.Query().Where(node.AgentIdEQ(nd.AgentId)).Only(context.Background())
		if err != nil {
			if !ent.IsNotFound(err) {
				return c.String(http.StatusBadRequest, "Get: "+err.Error())
			}
			return c.String(http.StatusNotFound, "Not Found")
		}
		return c.JSON(http.StatusOK, &nodes)

	}
}

// FindCluster Find Cluster
func (controller *Controller) FindCluster() echo.HandlerFunc {
	return func(c echo.Context) error {
		clu := new(ent.Cluster)

		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&clu); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return c.JSON(http.StatusBadRequest, "json:"+err.Error())
		}
		result, err := controller.client.Cluster.Query().Where(cluster.NameEQ(clu.Name)).Only(context.Background())
		if err != nil {
			return err
		}

		clusters, err := controller.client.Cluster.Create().
			SetID(result.ID).
			SetName(result.Name).
			SetDescription(result.Description).
			SetDisabled(result.Disabled).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(context.Background())

		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, &clusters)
	}
}

// FindMetricEndpoint findMetricEndpoint
func (controller *Controller) FindMetricEndpoint() echo.HandlerFunc {
	return func(c echo.Context) error {
		me := new(ent.MetricEndpoint)

		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&me); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return c.JSON(http.StatusBadRequest, "json:"+err.Error())
		}
		result, err := controller.client.MetricEndpoint.Query().Where(metricendpoint.PathEQ(me.Path)).First(context.Background())
		if err != nil {
			return fmt.Errorf("faild querying err: %w", err)
		}

		metricEndpoint, err := controller.client.MetricEndpoint.Create().SetPath(result.Path).Save(context.Background())
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, &metricEndpoint)
	}
}

// FindMetricType Find Metric Type
func (controller *Controller) FindMetricType() echo.HandlerFunc {
	return func(c echo.Context) error {
		mt := new(ent.MetricType)

		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&mt); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return c.JSON(http.StatusBadRequest, "json: "+err.Error())
		}
		result, err := controller.client.MetricType.Query().Where(metrictype.NameNEQ(mt.Name)).First(context.Background())
		if err != nil {
			return fmt.Errorf("faild querying err: %w", err)
		}
		metricType, err := controller.client.MetricType.Create().SetName(result.Name).Save(context.Background())
		return c.JSON(http.StatusOK, &metricType)
	}
}

// FindMetricName findMetricName
func (controller *Controller) FindMetricName() echo.HandlerFunc {
	return func(c echo.Context) error {
		mn := new(ent.MetricName)
		metricType := new(ent.MetricType)

		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&mn); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return c.JSON(http.StatusBadRequest, "json: "+err.Error())
		}

		result, err := controller.client.MetricName.Query().Where(metricname.NameEQ(mn.Name)).First(context.Background())
		if err != nil {
			return fmt.Errorf("faild querying err: %w", err)
		}
		metricName, err := controller.client.MetricName.Create().
			SetName(result.Name).
			SetTypeId(metricType.ID).
			Save(context.Background())
		return c.JSON(http.StatusOK, &metricName)
	}
}

// FindMetricLabel findMetricLabel
func (controller *Controller) FindMetricLabel() echo.HandlerFunc {
	return func(c echo.Context) error {
		ml := new(ent.MetricLabel)

		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&ml); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return c.JSON(http.StatusBadRequest, "json: "+err.Error())
		}

		result, err := controller.client.MetricLabel.Query().Where(metriclabel.LabelEQ(ml.Label)).Only(context.Background())
		if err != nil {
			return fmt.Errorf("faild querying err: %w", err)
		}

		metricLabel, err := controller.client.MetricLabel.Create().
			SetLabel(result.Label).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).Save(context.Background())

		return c.JSON(http.StatusOK, &metricLabel)
	}
}

// FindNode findNode
func (controller *Controller) FindNode() echo.HandlerFunc {
	return func(c echo.Context) error {
		nd := new(ent.Node)
		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&nd); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return c.JSON(http.StatusBadRequest, "json: "+err.Error())
		}

		node, err := controller.client.Node.Query().Where(node.HostEQ(nd.Host)).Where(node.ClusterId(nd.ClusterId)).Only(context.Background())
		if err != nil {
			return fmt.Errorf("faild querying err: %w", err)
		}
		return c.JSON(http.StatusOK, &node)
	}
}

// FindNodeById FindNodeById
func (controller *Controller) FindNodeById() echo.HandlerFunc {
	return func(c echo.Context) error {
		nd := new(ent.Node)
		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&nd); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return c.JSON(http.StatusBadRequest, "json: "+err.Error())
		}

		node, err := controller.client.Node.Query().Where(node.IDEQ(nd.ID)).Where(node.ClusterId(nd.ClusterId)).Only(context.Background())
		if err != nil {
			return fmt.Errorf("faild querying err: %w", err)
		}
		return c.JSON(http.StatusOK, &node)
	}
}

// FindContainer find Container
func (controller *Controller) FindContainer() echo.HandlerFunc {
	return func(c echo.Context) error {
		cot := new(ent.Container)
		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&cot); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return c.JSON(http.StatusBadRequest, "json: "+err.Error())
		}

		container, err := controller.client.Container.Query().Where(container.NameEQ(cot.Name)).Where(container.ClusterIdEQ(cot.ClusterId)).Where(container.NodeIdEQ(cot.NodeId)).Only(context.Background())
		if err != nil {
			return fmt.Errorf("faild querying err: %w", err)
		}
		return c.JSON(http.StatusOK, &container)
	}
}

// FindProcess find Process
func (controller *Controller) FindProcess() echo.HandlerFunc {
	return func(c echo.Context) error {
		ps := new(ent.Proces)
		log, _ := zap.NewDevelopment()
		if err := json.NewDecoder(c.Request().Body).Decode(&ps); err != nil {
			log.Fatal("json decode error", zap.Error(err))
			return c.JSON(http.StatusBadRequest, "json: "+err.Error())
		}

		processes, err := controller.client.Proces.Query().Where(proces.NameEQ(ps.Name)).Where(proces.PIdEQ(ps.PId)).
			Where(proces.NodeIdEQ(ps.NodeId)).Where(proces.ClusterIdEQ(ps.ClusterId)).Only(context.Background())
		if err != nil {
			return fmt.Errorf("faild querying err: %w", err)
		}
		return c.JSON(http.StatusOK, &processes)
	}
}
