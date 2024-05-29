// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package controller

import (
	"dataverse/consts"

	ametrics "github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/utils/wrappers"
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	createAsset prometheus.Counter
	mintAsset   prometheus.Counter
	burnAsset   prometheus.Counter

	transfer prometheus.Counter

	createOrder prometheus.Counter
	fillOrder   prometheus.Counter
	closeOrder  prometheus.Counter

	importAsset prometheus.Counter
	exportAsset prometheus.Counter

	createProject   prometheus.Counter
	createUpdate    prometheus.Counter
	registerMachine prometheus.Counter
}

func newMetrics(gatherer ametrics.MultiGatherer) (*metrics, error) {
	m := &metrics{
		createAsset: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "actions",
			Name:      "create_asset",
			Help:      "number of create asset actions",
		}),
		mintAsset: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "actions",
			Name:      "mint_asset",
			Help:      "number of mint asset actions",
		}),
		burnAsset: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "actions",
			Name:      "burn_asset",
			Help:      "number of burn asset actions",
		}),
		transfer: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "actions",
			Name:      "transfer",
			Help:      "number of transfer actions",
		}),
		createOrder: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "actions",
			Name:      "create_order",
			Help:      "number of create order actions",
		}),
		fillOrder: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "actions",
			Name:      "fill_order",
			Help:      "number of fill order actions",
		}),
		closeOrder: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "actions",
			Name:      "close_order",
			Help:      "number of close order actions",
		}),
		importAsset: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "actions",
			Name:      "import_asset",
			Help:      "number of import asset actions",
		}),
		exportAsset: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "actions",
			Name:      "export_asset",
			Help:      "number of export asset actions",
		}),
		createProject: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "actions",
			Name:      "project",
			Help:      "no of projects created",
		}),
		createUpdate: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "actions",
			Name:      "update",
			Help:      "no of updates created",
		}),
		registerMachine: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "actions",
			Name:      "register_machine",
			Help:      "no of machines cid's",
		}),
	}
	r := prometheus.NewRegistry()
	errs := wrappers.Errs{}
	errs.Add(
		r.Register(m.createAsset),
		r.Register(m.mintAsset),
		r.Register(m.burnAsset),

		r.Register(m.transfer),

		r.Register(m.createOrder),
		r.Register(m.fillOrder),
		r.Register(m.closeOrder),

		r.Register(m.importAsset),
		r.Register(m.exportAsset),
		r.Register(m.createProject),
		r.Register(m.createUpdate),
		r.Register(m.registerMachine),
		gatherer.Register(consts.Name, r),
	)
	return m, errs.Err
}
