package main

import "strings"

//DashboardNodesJSON blabla
//var DashboardJSON1A = dashboardHead + PanelIOpressure + "," + GaugesDedicatedNodes + "," + PanelCPU + "," + PanelMemory //+ LabelNodesA

//Between Panels a comma has to be written!!!

var dashboardHead = `{
  "dashboard": {
    "id": null,
    "title": "Tenant Dashboard",
    "tags": [ "templated" ],
    "timezone": "browser",
    "refresh": "10s",
  "rows": [`

var PanelIOpressure = `{
      "collapse": false,
      "height": "200px",
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "datasource": "prometheus_source",
          "decimals": 2,
          "editable": true,
          "error": false,
          "fill": 1,
          "grid": {},
          "height": "200px",
          "id": 32,
          "legend": {
            "alignAsTable": false,
            "avg": true,
            "current": true,
            "max": false,
            "min": false,
            "rightSide": false,
            "show": false,
            "sideWidth": 200,
            "sort": "current",
            "sortDesc": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 2,
          "links": [],
          "nullPointMode": "connected",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 12,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum (rate (container_network_receive_bytes_total{kubernetes_io_hostname=~\"^$DedicatedNodes$\"}[1m]))",
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "Received",
              "metric": "network",
              "refId": "A",
              "step": 10
            },
            {
              "expr": "- sum (rate (container_network_transmit_bytes_total{kubernetes_io_hostname=~\"^$DedicatedNodes$\"}[1m]))",
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "Sent",
              "metric": "network",
              "refId": "B",
              "step": 10
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Dedicated nodes: $DedicatedNodes Network I/O pressure",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 0,
            "value_type": "cumulative"
          },
          "transparent": false,
          "type": "graph",
          "xaxis": {
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "Bps",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "Bps",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": false
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": false,
      "title": "Network I/O pressure",
      "titleSize": "h6"
    }`

//Memory
var PanelMemory = `{
      "collapse": false,
      "height": "250px",
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "datasource": "prometheus_source",
          "decimals": 2,
          "editable": true,
          "error": false,
          "fill": 0,
          "grid": {},
          "id": 39,
          "legend": {
            "alignAsTable": true,
            "avg": true,
            "current": true,
            "max": false,
            "min": false,
            "rightSide": true,
            "show": true,
            "sideWidth": 200,
            "sort": "current",
            "sortDesc": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 2,
          "links": [],
          "nullPointMode": "connected",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 12,
          "stack": false,
          "steppedLine": true,
          "targets": [
            {
              "expr": "sum (container_memory_working_set_bytes{namespace=~\"$NAMESPACE-.*\"}) by ($division)",
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "{{$division}}",
              "metric": "container_memory_usage:sort_desc",
              "refId": "A",
              "step": 10
            },
            {
              "expr": "sum (container_memory_working_set_bytes{namespace=~\"$NAMESPACE-.*\"})",
              "hide": false,
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "Total",
              "metric": "container_memory_usage:sort_desc",
              "refId": "B",
              "step": 10
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "$division memory usage",
          "tooltip": {
            "msResolution": false,
            "shared": true,
            "sort": 2,
            "value_type": "cumulative"
          },
          "type": "graph",
          "xaxis": {
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": false
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": false,
      "title": "Pods CPU usage",
      "titleSize": "h6"
    }`

//CPU
var PanelCPU = `{
      "collapse": false,
      "height": "250px",
      "panels": [
        {
          "aliasColors": {},
          "bars": false,
          "datasource": "prometheus_source",
          "decimals": 3,
          "editable": true,
          "error": false,
          "fill": 0,
          "grid": {},
          "height": "",
          "id": 17,
          "legend": {
            "alignAsTable": true,
            "avg": true,
            "current": true,
            "max": false,
            "min": false,
            "rightSide": true,
            "show": true,
            "sort": "current",
            "sortDesc": true,
            "total": false,
            "values": true
          },
          "lines": true,
          "linewidth": 2,
          "links": [],
          "nullPointMode": "connected",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 12,
          "stack": false,
          "steppedLine": true,
          "targets": [
            {
              "expr": "sum (rate (container_cpu_usage_seconds_total{namespace=~\"$NAMESPACE-.*\"}[1m])) by ($division)",
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "{{$division}} ",
              "metric": "container_cpu",
              "refId": "A",
              "step": 10
            },
            {
              "expr": "sum (rate (container_cpu_usage_seconds_total{namespace=~\"$NAMESPACE-.*\"}[1m]))",
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "Total",
              "refId": "B",
              "step": 10
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "$division CPU usage (1m avg)",
          "tooltip": {
            "msResolution": true,
            "shared": true,
            "sort": 2,
            "value_type": "cumulative"
          },
          "transparent": false,
          "type": "graph",
          "xaxis": {
            "mode": "time",
            "name": null,
            "show": true,
            "values": []
          },
          "yaxes": [
            {
              "format": "none",
              "label": "cores",
              "logBase": 1,
              "max": null,
              "min": null,
              "show": true
            },
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": false
            }
          ]
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": false,
      "title": "Pods CPU usage",
      "titleSize": "h6"
    }`

//Memory
var GaugesDedicatedNodes = `{
      "collapse": false,
      "height": "250px",
      "panels": [
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": true,
          "colors": [
            "rgba(50, 172, 45, 0.97)",
            "rgba(237, 129, 40, 0.89)",
            "rgba(245, 54, 54, 0.9)"
          ],
          "datasource": "prometheus_source",
          "editable": true,
          "error": false,
          "format": "percent",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": true,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "height": "250px",
          "id": 4,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": "",
          "postfixFontSize": "50%",
          "prefix": "",
          "prefixFontSize": "50%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "span": 4,
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": false,
            "lineColor": "rgb(31, 120, 193)",
            "show": false
          },
          "targets": [
            {
              "expr": "sum (container_memory_working_set_bytes{id=\"/\",kubernetes_io_hostname=~\"^$DedicatedNodes$\"}) / sum (machine_memory_bytes{kubernetes_io_hostname=~\"^$DedicatedNodes$\"}) * 100",
              "interval": "10s",
              "intervalFactor": 1,
              "refId": "A",
              "step": 10
            }
          ],
          "thresholds": "65, 90",
          "title": "Dedicated nodes: $DedicatedNodes memory usage",
          "transparent": false,
          "type": "singlestat",
          "valueFontSize": "80%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "current"
        },
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": true,
          "colors": [
            "rgba(50, 172, 45, 0.97)",
            "rgba(237, 129, 40, 0.89)",
            "rgba(245, 54, 54, 0.9)"
          ],
          "datasource": "prometheus_source",
          "editable": true,
          "error": false,
          "format": "percent",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": true,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "height": "250px",
          "id": 36,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": "",
          "postfixFontSize": "50%",
          "prefix": "",
          "prefixFontSize": "50%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "span": 4,
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": false,
            "lineColor": "rgb(31, 120, 193)",
            "show": false
          },
          "targets": [
            {
              "expr": "sum (rate (container_cpu_usage_seconds_total{id=\"/\",kubernetes_io_hostname=~\"^$DedicatedNodes$\"}[1m])) / sum (machine_cpu_cores{kubernetes_io_hostname=~\"^$DedicatedNodes$\"}) * 100",
              "interval": "10s",
              "intervalFactor": 1,
              "refId": "A",
              "step": 10
            }
          ],
          "thresholds": "65, 90",
          "title": "Dedicated nodes: $DedicatedNodes CPU usage (1m avg)",
          "transparent": false,
          "type": "singlestat",
          "valueFontSize": "80%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "current"
        },
        {
          "content": "",
          "id": 35,
          "links": [],
          "mode": "markdown",
          "span": 4,
          "title": "",
          "transparent": true,
          "type": "text"
        },
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": false,
          "colors": [
            "rgba(50, 172, 45, 0.97)",
            "rgba(237, 129, 40, 0.89)",
            "rgba(245, 54, 54, 0.9)"
          ],
          "datasource": "prometheus_source",
          "decimals": 2,
          "editable": true,
          "error": false,
          "format": "bytes",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": false,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "height": "1px",
          "id": 9,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": "",
          "postfixFontSize": "20%",
          "prefix": "",
          "prefixFontSize": "20%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "span": 2,
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": false,
            "lineColor": "rgb(31, 120, 193)",
            "show": false
          },
          "targets": [
            {
              "expr": "sum (container_memory_working_set_bytes{id=\"/\",kubernetes_io_hostname=~\"^$DedicatedNodes$\"})",
              "interval": "10s",
              "intervalFactor": 1,
              "refId": "A",
              "step": 10
            }
          ],
          "thresholds": "",
          "title": "Used",
          "type": "singlestat",
          "valueFontSize": "50%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "current"
        },
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": false,
          "colors": [
            "rgba(50, 172, 45, 0.97)",
            "rgba(237, 129, 40, 0.89)",
            "rgba(245, 54, 54, 0.9)"
          ],
          "datasource": "prometheus_source",
          "decimals": 2,
          "editable": true,
          "error": false,
          "format": "bytes",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": false,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "height": "1px",
          "id": 10,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": "",
          "postfixFontSize": "50%",
          "prefix": "",
          "prefixFontSize": "50%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "span": 2,
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": false,
            "lineColor": "rgb(31, 120, 193)",
            "show": false
          },
          "targets": [
            {
              "expr": "sum (machine_memory_bytes{kubernetes_io_hostname=~\"^$DedicatedNodes$\"})",
              "interval": "10s",
              "intervalFactor": 1,
              "refId": "A",
              "step": 10
            }
          ],
          "thresholds": "",
          "title": "Total",
          "type": "singlestat",
          "valueFontSize": "50%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "current"
        },
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": false,
          "colors": [
            "rgba(50, 172, 45, 0.97)",
            "rgba(237, 129, 40, 0.89)",
            "rgba(245, 54, 54, 0.9)"
          ],
          "datasource": "prometheus_source",
          "decimals": 2,
          "editable": true,
          "error": false,
          "format": "none",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": false,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "height": "1px",
          "id": 37,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": " cores",
          "postfixFontSize": "30%",
          "prefix": "",
          "prefixFontSize": "50%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "span": 2,
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": false,
            "lineColor": "rgb(31, 120, 193)",
            "show": false
          },
          "targets": [
            {
              "expr": "sum (rate (container_cpu_usage_seconds_total{id=\"/\",kubernetes_io_hostname=~\"^$DedicatedNodes$\"}[1m]))",
              "interval": "10s",
              "intervalFactor": 1,
              "refId": "A",
              "step": 10
            }
          ],
          "thresholds": "",
          "title": "Used",
          "type": "singlestat",
          "valueFontSize": "50%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "current"
        },
        {
          "cacheTimeout": null,
          "colorBackground": false,
          "colorValue": false,
          "colors": [
            "rgba(50, 172, 45, 0.97)",
            "rgba(237, 129, 40, 0.89)",
            "rgba(245, 54, 54, 0.9)"
          ],
          "datasource": "prometheus_source",
          "decimals": 2,
          "editable": true,
          "error": false,
          "format": "none",
          "gauge": {
            "maxValue": 100,
            "minValue": 0,
            "show": false,
            "thresholdLabels": false,
            "thresholdMarkers": true
          },
          "height": "1px",
          "id": 38,
          "interval": null,
          "links": [],
          "mappingType": 1,
          "mappingTypes": [
            {
              "name": "value to text",
              "value": 1
            },
            {
              "name": "range to text",
              "value": 2
            }
          ],
          "maxDataPoints": 100,
          "nullPointMode": "connected",
          "nullText": null,
          "postfix": " cores",
          "postfixFontSize": "30%",
          "prefix": "",
          "prefixFontSize": "50%",
          "rangeMaps": [
            {
              "from": "null",
              "text": "N/A",
              "to": "null"
            }
          ],
          "span": 2,
          "sparkline": {
            "fillColor": "rgba(31, 118, 189, 0.18)",
            "full": false,
            "lineColor": "rgb(31, 120, 193)",
            "show": false
          },
          "targets": [
            {
              "expr": "sum (machine_cpu_cores{kubernetes_io_hostname=~\"^$DedicatedNodes$\"})",
              "interval": "10s",
              "intervalFactor": 1,
              "refId": "A",
              "step": 10
            }
          ],
          "thresholds": "",
          "title": "Total",
          "type": "singlestat",
          "valueFontSize": "50%",
          "valueMaps": [
            {
              "op": "=",
              "text": "N/A",
              "value": "null"
            }
          ],
          "valueName": "current"
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": false,
      "title": "Total usage",
      "titleSize": "h6"
    }`

var DashboardJSON1B = `
    ]
  },
  "time": {
    "from": "now-5m",
    "to": "now"
  },
  "timepicker": {
    "refresh_intervals": [
      "5s",
      "10s",
      "30s",
      "1m",
      "5m",
      "15m",
      "30m",
      "1h",
      "2h",
      "1d"
    ],
    "time_options": [
      "5m",
      "15m",
      "1h",
      "6h",
      "12h",
      "24h",
      "2d",
      "7d",
      "30d"
    ]
  },
    "schemaVersion": 6,
    "version": 0
  },
  "overwrite": true 
}

`

func DashboardTemplating(Label string, Namespace string) string { //TODO : hide the Customer and the Nodes

	var template = `  
  ],
  "schemaVersion": 14,
  "style": "dark",
  "tags": [
    "kubernetes"
  ],
  "templating": {
    "list": [
      {
        "allValue": ".*",
        "current": {
          "text": "All",
          "value": "$__all"
        },
        "datasource": "prometheus_source",
        "hide": 2,
        "includeAll": true,
        "label": null,
        "multi": false,
        "name": "Node",
        "options": [],
        "query": "label_values(kubernetes_io_hostname)",
        "refresh": 1,
        "regex": "",
        "sort": 0,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      },
            {
        "allValue": null,
        "current": {
          "tags": [],
          "text": "Pods",
          "value": "pod_name"
        },
        "hide": 0,
        "includeAll": false,
        "label": null,
        "multi": false,
        "name": "division",
        "options": [
          {
            "selected": false,
            "text": "Pods",
            "value": "pod_name"
          },
           {
            "selected": false,
            "text": "Containers",
            "value": "container_name"
          },
          {
            "selected": true,
            "text": "Projects",
            "value": "namespace"
          }
        ],
        "query": "pod_name,container_name,namespace",
        "type": "custom"
      },
      {
        "current": {
          "tags": [],
          "text": "` + Label + `",
          "value": "` + Label + `"
        },
        "hide": 2,
        "label": null,
        "name": "Customer",
        "options": [
          {
            "selected": true,
            "text": "` + Label + `",
            "value": "` + Label + `"
          }
        ],
        "query": "` + Label + `",
        "type": "constant"
      },
      {
        "allValue": null,
        "current": {
          "tags": [],
          "text": "All",
          "value": "$__all"
        },
        "datasource": "prometheus_source",
        "hide": 0,  
        "includeAll": true,
        "label": null,
        "multi": false,
        "name": "DedicatedNodes",
        "options": [
          {
            "selected": true,
            "text": "All",
            "value": "$__all"
          }
        ],
        "query": "up{customer=\"$Customer\"}",
        "refresh": 1,
        "regex": "/.*instance=\"([^\"]*).*/",
        "sort": 0,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      },
	  {
        "current": {
          "text": "` + Namespace + `",
          "value": "` + Namespace + `"
        },
        "hide": 2,
        "label": null,
        "name": "NAMESPACE",
        "options": [
          {
            "selected": true,
            "text": "` + Namespace + `",
            "value": "` + Namespace + `"
          }
        ],
        "query": "` + Namespace + `",
        "type": "constant"
      }`

	return template
}

//TODO remove namespace from current value by the division selection
//TODO, change 1m averages to 5m or so, cause of the 120% error in gauges
//TODO overwrite should be a parameter only used when updating dashboard

func DashboardPanels(panelGauges bool, panelCpu bool, panelMemory bool, panelIOpressure bool) string {
	var DashboardPanels = dashboardHead //+ PanelIOpressure+ "," + GaugesDedicatedNodes + "," + PanelCPU + "," + PanelMemory

	if panelGauges == true {
		DashboardPanels = DashboardPanels + GaugesDedicatedNodes + ","
	}
	if panelCpu == true {
		DashboardPanels = DashboardPanels + PanelCPU + ","
	}
	if panelMemory == true {
		DashboardPanels = DashboardPanels + PanelMemory + ","
	}
	if panelIOpressure == true {
		DashboardPanels = DashboardPanels + PanelIOpressure + ","
	}

	//The last comma has to be removed, otherwise the dashboard is invalid
	if strings.HasSuffix(DashboardPanels, ",") {
		DashboardPanels = DashboardPanels[:len(DashboardPanels)-len(",")]
	}

	/*
	     	if panelGauges == true {
	   		DashboardPanels = DashboardPanels + "," + GaugesDedicatedNodes
	   	}
	   	if panelCpu == true {
	   		DashboardPanels = DashboardPanels + "," + PanelCPU
	   	}
	   	if panelMemory == true {
	   		DashboardPanels = DashboardPanels + "," + PanelMemory
	   	}
	   	if panelIOpressure == true {
	   		DashboardPanels = DashboardPanels + "," + PanelIOpressure
	   	}
	*/

	return DashboardPanels
}
