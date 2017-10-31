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

var Help = `{
      "collapse": true,
      "height": 250,
      "panels": [
        {
          "content": "<font color=\"#ff9966\"><u><b>Node Graphs's variables</b></u> <br></font>\n<font color=\"#5cd6d6\"><b>Node:</b></font> shows a list of all nodes available in the environment. <br>\n<font color=\"#5cd6d6\"><b>NodeLabel:</b></font> dinamically discovers all node labels in the environment. When a label is chosen, it filters the 'Node' list, which shows only the labeled nodes <br><br>\n\n<font color=\"#ff9966\"><u><b>Openshift Graphs's variables</b></u> <br></font>\n<font color=\"#5cd6d6\"><b>showTotal:</b></font> if true, shows the total value for memory and CPU usages. <br>\n<font color=\"#5cd6d6\"><b>Namespace:</b></font> dinamically discovers all numerical namespaces in the environment. Choosing a namespace will restrict the graphs to that namespace.<br>\n<font color=\"#5cd6d6\"><b>Projects:</b></font> filters the shown projects for the variable 'Namespace'.  Choosing a project will restrict the graphs to that project.<br>\n<font color=\"#5cd6d6\"><b>Pods:</b></font> filters the shown pods for the variables 'Namespace'. Choosing a pod will restrict the graphs to that pod.<br>\n<font color=\"#5cd6d6\"><b>division:</b></font> selects which openshift module is represented in the graphics: pods, containers or projects.<br>\n<font color=\"#5cd6d6\"><b>referenceTo:</b></font> selects to which module the division variable is compared: for example, if 'division' is Pods and 'referenceTo' is Nodes, it will show in which node is that pod running .<br>",
          "description": "Help",
          "id": 40,
          "links": [],
          "mode": "html",
          "span": 12,
          "title": "Help",
          "transparent": false,
          "type": "text"
        }
      ],
      "repeat": null,
      "repeatIteration": null,
      "repeatRowId": null,
      "showTitle": false,
      "title": "Help",
      "titleSize": "h6"
    }`

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
              "expr": "sum (container_memory_usage_bytes{kubernetes_io_hostname=~\"$DedicatedNodes\",namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\",pod_name=~\"$Pods\"}) by ($division,$referenceTo)",
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "{{$division}} ({{$referenceTo}})",
              "metric": "container_memory_usage:sort_desc",
              "refId": "A",
              "step": 10
            },
            {
              "expr": "sum (container_memory_usage_bytes{showTotal=~\"$showTotal\",kubernetes_io_hostname=~\"$DedicatedNodes\",namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\",pod_name=~\"$Pods\"})",
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
          "title": "$division memory usage. Referenced to $referenceTo",
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
      "title": "Pods Memory usage",
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
              "expr": "sum (rate (container_cpu_usage_seconds_total{kubernetes_io_hostname=~\"$DedicatedNodes\",namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\",pod_name=~\"$Pods\"}[1m])) by ($division,$referenceTo)",
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "{{$division}} ({{$referenceTo}})",
              "metric": "container_cpu",
              "refId": "A",
              "step": 10
            },
            {
              "expr": "sum (rate (container_cpu_usage_seconds_total{showTotal=~\"$showTotal\",kubernetes_io_hostname=~\"$DedicatedNodes\",namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\",pod_name=~\"$Pods\"}[1m]))",
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
          "title": "$division CPU usage (1m avg). Referenced to $referenceTo",
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
              "expr": "sum (container_memory_usage_bytes{id=\"/\",kubernetes_io_hostname=~\"^$DedicatedNodes$\"}) / sum (machine_memory_bytes{kubernetes_io_hostname=~\"^$DedicatedNodes$\"}) * 100",
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
          "aliasColors": {},
          "bars": false,
          "datasource": "prometheus_source",
          "fill": 1,
          "id": 35,
          "legend": {
            "alignAsTable": true,
            "avg": false,
            "current": false,
            "max": false,
            "min": false,
            "rightSide": true,
            "show": true,
            "sideWidth": 500,
            "total": false,
            "values": false
          },
          "lines": false,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 4,
          "stack": false,
          "steppedLine": false,
          "targets": [
              {
                "expr": "sum(ALERTS{namespace=~\"$NAMESPACE-.*\",alertname=~\"PodHighMemoryConsumption\",alertname=~\"$alerts\"}) by (namespace,pod_name)",
                "intervalFactor": 1,
                "legendFormat": "project {{namespace}} - affected Pod:  {{pod_name}}",
                "refId": "A",
                "step": 1
              },
              {
                "expr": "sum(ALERTS{namespace=~\"$NAMESPACE-.*\",alertname=~\"NodeDown\",alertname=~\"$alerts\"}) by (namespace,pod_name,instance)",
                "intervalFactor": 1,
                "legendFormat": "Cannot reach {{instance}} ",
                "refId": "B",
                "step": 1
              },
              {
                "expr": "sum(ALERTS{namespace=~\"$NAMESPACE-.*\",alertname=~\"PodHighCPUConsumption\",alertname=~\"$alerts\"}) by (namespace,pod_name)",
                "intervalFactor": 1,
                "legendFormat": "project {{namespace}} - affected Pod:  {{pod_name}}",
                "refId": "C",
                "step": 1
              }
            ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Alert: $alerts",
          "tooltip": {
            "shared": true,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "time",
            "name": null,
            "show": false,
            "values": []
          },
          "yaxes": [
            {
              "format": "short",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": false
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
              "expr": "sum (container_memory_usage_bytes{id=\"/\",kubernetes_io_hostname=~\"^$DedicatedNodes$\"})",
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
      "title": "Nodes usage",
      "titleSize": "h6"
    }`

var ProjectResourceQuotas = `
{
      "collapse": false,
      "height": 250,
      "panels": [
        {
          "columns": [
            {
              "text": "Current",
              "value": "current"
            }
          ],
          "datasource": "prometheus_source",
          "fontSize": "90%",
          "id": 40,
          "links": [],
          "pageSize": null,
          "scroll": true,
          "showHeader": true,
          "sort": {
            "col": 0,
            "desc": true
          },
          "span": 2,
          "styles": [
            {
              "dateFormat": "YYYY-MM-DD HH:mm:ss",
              "pattern": "Time",
              "type": "date"
            },
            {
              "colorMode": null,
              "colors": [
                "rgba(245, 54, 54, 0.9)",
                "rgba(237, 129, 40, 0.89)",
                "rgba(50, 172, 45, 0.97)"
              ],
              "decimals": 2,
              "pattern": "/.*/",
              "thresholds": [],
              "type": "number",
              "unit": "short"
            }
          ],
          "targets": [
            {
              "expr": "resource_quota_hard_pods{namespace=\"$Projects\"}",
              "intervalFactor": 1,
              "legendFormat": "Limit Pods",
              "metric": "resource_quota_hard_pods",
              "refId": "A",
              "step": 1
            },
            {
              "expr": "resource_quota_hard_services{namespace=\"$Projects\"}",
              "intervalFactor": 1,
              "legendFormat": "Limit Services",
              "metric": "resource_quota_hard_services",
              "refId": "B",
              "step": 1
            },
            {
              "expr": "resource_quota_hard_replicationcontrollers{namespace=\"$Projects\"}",
              "intervalFactor": 1,
              "legendFormat": "Limit Replication controllers",
              "metric": "resource_quota_hard_replicationcontrollers",
              "refId": "C",
              "step": 1
            },
            {
              "expr": "resource_quota_used_pods{namespace=\"$Projects\"}",
              "intervalFactor": 1,
              "legendFormat": "Used Pods",
              "metric": "resource_quota_used_pods",
              "refId": "D",
              "step": 1
            },
            {
              "expr": "resource_quota_used_services{namespace=\"$Projects\"}",
              "intervalFactor": 1,
              "legendFormat": "Used Services",
              "metric": "resource_quota_used_services",
              "refId": "E",
              "step": 1
            },
            {
              "expr": "resource_quota_used_replicationcontrollers{namespace=\"$Projects\"}",
              "intervalFactor": 1,
              "legendFormat": "Used Replication controllers",
              "metric": "resource_quota_used_replicationcontrollers",
              "refId": "F",
              "step": 1
            }
          ],
          "title": "Project $Projects - Resource Quota Modules",
          "transform": "timeseries_aggregations",
          "type": "table"
        },
        {
          "aliasColors": {},
          "bars": true,
          "datasource": "prometheus_source",
          "fill": 1,
          "id": 41,
          "legend": {
            "alignAsTable": true,
            "avg": false,
            "current": true,
            "max": false,
            "min": false,
            "rightSide": true,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": false,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 5,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum (resource_quota_hard_memory{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\"}) by (namespace) ",
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "Reservation Limit",
              "refId": "A",
              "step": 10
            },
            {
              "expr": "sum (resource_quota_used_memory{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\"}) by (namespace) ",
              "hide": false,
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "Reserved",
              "refId": "B",
              "step": 10
            },
            {
              "expr": "sum (resource_quota_hard_limit_memory{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\"}) by (namespace) ",
              "hide": false,
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "Memory Limit",
              "refId": "C",
              "step": 10
            },
            {
              "expr": "sum (resource_quota_used_limit_memory{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\"}) by (namespace) ",
              "hide": false,
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "In use",
              "refId": "D",
              "step": 10
            },
            {
              "expr": "sum (container_memory_usage_bytes{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\"}) by (namespace) ",
              "hide": true,
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "In use (2)",
              "refId": "E",
              "step": 10
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Project $Projects - Resource Quota Memory",
          "tooltip": {
            "shared": false,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "series",
            "name": null,
            "show": true,
            "values": [
              "current"
            ]
          },
          "yaxes": [
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": false
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": true,
          "datasource": "prometheus_source",
          "fill": 1,
          "id": 42,
          "legend": {
            "alignAsTable": true,
            "avg": false,
            "current": true,
            "max": false,
            "min": false,
            "rightSide": true,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": false,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 5,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum (resource_quota_hard_cpu{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\"}/ 100000) by (namespace) ",
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "Reservation limit",
              "refId": "A",
              "step": 10
            },
            {
              "expr": "sum (resource_quota_used_cpu{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\"}/ 100000) by (namespace) ",
              "hide": false,
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "Reserved",
              "refId": "B",
              "step": 10
            },
            {
              "expr": "sum (resource_quota_hard_limit_cpu{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\"}/ 100000) by (namespace) ",
              "hide": false,
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "CPU limit",
              "refId": "C",
              "step": 10
            },
            {
              "expr": "sum (resource_quota_used_limit_cpu{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\"}/ 100000) by (namespace) ",
              "hide": false,
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "In use",
              "refId": "D",
              "step": 10
            },
            {
              "expr": "sum (rate (container_cpu_usage_seconds_total{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\"}[1m]))",
              "hide": true,
              "interval": "10s",
              "intervalFactor": 1,
              "legendFormat": "In use(2)",
              "refId": "E",
              "step": 10
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Project $Projects - Resource Quota CPU",
          "tooltip": {
            "shared": false,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "series",
            "name": null,
            "show": true,
            "values": [
              "current"
            ]
          },
          "yaxes": [
            {
              "format": "none",
              "label": "cores",
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "bytes",
              "label": "cores",
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
      "title": "Project Resource Quota",
      "titleSize": "h6"
    }
    `
var NodeLimits = `
{
      "collapse": false,
      "height": "250px",
      "panels": [
        {
          "aliasColors": {},
          "bars": true,
          "datasource": "prometheus_source",
          "fill": 1,
          "id": 55,
          "legend": {
            "alignAsTable": true,
            "avg": false,
            "current": true,
            "max": false,
            "min": false,
            "rightSide": true,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": false,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": true,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 12,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum (machine_memory_bytes{kubernetes_io_hostname=~\"^$DedicatedNodes$\"})",
              "intervalFactor": 1,
              "legendFormat": "Machine total memory",
              "refId": "B",
              "step": 1
            },
            {
              "expr": "sum(container_spec_memory_limit_bytes{container_name!=\"POD\",kubernetes_io_hostname=~\"$DedicatedNodes\",image!=\"\"}) ",
              "intervalFactor": 1,
              "legendFormat": "Node's total set limit",
              "refId": "D",
              "step": 1
            },
            {
              "expr": "sum(container_spec_memory_limit_bytes{kubernetes_io_hostname=~\"$DedicatedNodes\",image!=\"\"})  by ($division) != 0",
              "hide": false,
              "intervalFactor": 1,
              "legendFormat": "Limit - {{$division}}",
              "refId": "A",
              "step": 1
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Node $DedicatedNodes - Total memory limits set",
          "tooltip": {
            "shared": false,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "series",
            "name": null,
            "show": true,
            "values": [
              "current"
            ]
          },
          "yaxes": [
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "bytes",
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
      "title": "Nodes Limits Quotas",
      "titleSize": "h6"
    }
`
var PodsLimits = `{
      "collapse": false,
      "height": 250,
      "panels": [
        {
          "aliasColors": {},
          "bars": true,
          "datasource": "prometheus_source",
          "fill": 1,
          "id": 49,
          "legend": {
            "alignAsTable": true,
            "avg": false,
            "current": true,
            "max": false,
            "min": false,
            "rightSide": true,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": false,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": true,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "seriesOverrides": [],
          "span": 6,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(container_spec_memory_limit_bytes{container_name!=\"POD\",namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\",pod_name=~\"$Pods\"})  by (pod_name)",
              "intervalFactor": 1,
              "legendFormat": "Limit",
              "refId": "D",
              "step": 1
            },
            {
              "expr": "sum (container_memory_usage_bytes{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\",pod_name=~\"$Pods\"}) by (pod_name)",
              "intervalFactor": 1,
              "legendFormat": "Usage",
              "refId": "A",
              "step": 1
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Pod $Pods - Memory limit",
          "tooltip": {
            "shared": false,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "series",
            "name": null,
            "show": true,
            "values": [
              "current"
            ]
          },
          "yaxes": [
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "bytes",
              "label": null,
              "logBase": 1,
              "max": null,
              "min": null,
              "show": false
            }
          ]
        },
        {
          "aliasColors": {},
          "bars": true,
          "datasource": "prometheus_source",
          "fill": 1,
          "id": 50,
          "legend": {
            "alignAsTable": true,
            "avg": false,
            "current": true,
            "max": false,
            "min": false,
            "rightSide": true,
            "show": true,
            "total": false,
            "values": true
          },
          "lines": false,
          "linewidth": 1,
          "links": [],
          "nullPointMode": "null",
          "percentage": false,
          "pointradius": 5,
          "points": false,
          "renderer": "flot",
          "repeat": null,
          "seriesOverrides": [],
          "span": 6,
          "stack": false,
          "steppedLine": false,
          "targets": [
            {
              "expr": "sum(container_spec_cpu_quota{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\",pod_name=~\"$Pods\"} / 100000) by (pod_name)",
              "intervalFactor": 1,
              "legendFormat": "Limit",
              "refId": "A",
              "step": 1
            },
            {
              "expr": "sum (rate (container_cpu_usage_seconds_total{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\",pod_name=~\"$Pods\"}[1m])) by (pod_name)",
              "intervalFactor": 2,
              "legendFormat": "Usage",
              "refId": "B",
              "step": 2
            }
          ],
          "thresholds": [],
          "timeFrom": null,
          "timeShift": null,
          "title": "Pod $Pods - CPU limit",
          "tooltip": {
            "shared": false,
            "sort": 0,
            "value_type": "individual"
          },
          "type": "graph",
          "xaxis": {
            "mode": "series",
            "name": null,
            "show": true,
            "values": [
              "current"
            ]
          },
          "yaxes": [
            {
              "format": "none",
              "label": "cores",
              "logBase": 1,
              "max": null,
              "min": "0",
              "show": true
            },
            {
              "format": "bytes",
              "label": "cores",
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
      "title": "Pods Limits Quotas",
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

//TODO remove namespace from current value by the division selection
//TODO, change 1m averages to 5m or so, cause of the 120% error in gauges
//TODO overwrite should be a parameter only used when updating dashboard

func DashboardPanels(panelGauges bool, panelCpu bool, panelMemory bool, panelIOpressure bool, PanelResourcequotas bool) string {
	var DashboardPanels = dashboardHead //+ PanelIOpressure+ "," + GaugesDedicatedNodes + "," + PanelCPU + "," + PanelMemory

	//DashboardPanels = DashboardPanels + Help + ","
	DashboardPanels = DashboardPanels

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

	if PanelResourcequotas == true {
		DashboardPanels = DashboardPanels + ProjectResourceQuotas + ","

		DashboardPanels = DashboardPanels + NodeLimits + ","
		DashboardPanels = DashboardPanels + PodsLimits + ","
	}

	//The last comma has to be removed, otherwise the dashboard is invalid
	if strings.HasSuffix(DashboardPanels, ",") {
		DashboardPanels = DashboardPanels[:len(DashboardPanels)-len(",")]
	}

	return DashboardPanels
}

func Templating(varNode bool, varShowTotal bool, varDivision bool, varReferenceTo bool, varCustomer bool, Label string, varDedicatedNodes bool, varUseProdNodes bool, varNAMESPACE bool, Namespace string, varProjects bool, varPods bool, varAlerts bool) string {

	var Templating = templatingHead

	if varNode == true {
		Templating = Templating + templatingNode + ","
	}
	if varShowTotal == true {
		Templating = Templating + templatingShowTotal + ","
	}

	if varDivision == true {
		Templating = Templating + templatingDivision + ","
	}

	if varReferenceTo == true {
		Templating = Templating + templatingReferenceTo + ","
	}

	if varCustomer == true {
		Templating = Templating + templatingCustomer(Label) + ","
	}

	if varDedicatedNodes == true {
		Templating = Templating + templatingDedicatedNodes + ","
	}

	if varUseProdNodes == true {
		Templating = Templating + templatingUseProdNodes + ","
	}

	if varNAMESPACE == true {
		Templating = Templating + templatingNAMESPACE(Namespace) + ","
	}

	if varProjects == true {
		Templating = Templating + templatingProjects + ","
	}

	if varPods == true {
		Templating = Templating + templatingPods + ","
	}
	if varAlerts == true {
		Templating = Templating + templatingAlerts + ","
	}

	//The last comma has to be removed, otherwise the dashboard is invalid
	if strings.HasSuffix(Templating, ",") {
		Templating = Templating[:len(Templating)-len(",")]
	}

	return Templating
}

var templatingHead = `  
  ],
  "schemaVersion": 14,
  "style": "dark",
  "tags": [
    "kubernetes"
  ],
  "templating": {
    "list": [`

var templatingNode = `{
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
      }`
var templatingShowTotal = `{
        "allValue": ".*",
        "current": {
          "tags": [],
          "text": "false",
          "value": "false"
        },
        "hide": 0,
        "includeAll": true,
        "label": null,
        "multi": false,
        "name": "showTotal",
        "options": [
          {
            "selected": false,
            "text": "true",
            "value": "$__all"
          },
          {
            "selected": true,
            "text": "false",
            "value": "false"
          }
        ],
        "query": "false",
        "type": "custom"
      }`
var templatingDivision = `{
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
      }`
var templatingReferenceTo = `{
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
        "name": "referenceTo",
        "options": [
          {
            "selected": false,
            "text": "none",
            "value": "none"
          },
          {
            "selected": true,
            "text": "Pods",
            "value": "pod_name"
          },
          {
            "selected": false,
            "text": "Containers",
            "value": "container_name"
          },
          {
            "selected": false,
            "text": "Namespaces",
            "value": "namespace"
          }
        ],
        "query": "none,pod_name,container_name,namespace",
        "refresh": 0,
        "type": "custom"
      }`

func templatingCustomer(Label string) string {
	var templatingCustomer = `{
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
      }`

	return templatingCustomer
}

var templatingDedicatedNodes = `{
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
        "query": "up{customer=~\"$Customer\",prod=~\"$UseProdNodes\"}",
        "refresh": 1,
        "regex": "/.*instance=\"([^\"]*).*/",
        "sort": 1,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      }`
var templatingUseProdNodes = `{
        "allValue": ".*",
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
        "name": "UseProdNodes",
        "options": [],
        "query": "up{prod=~\".*\"}",
        "refresh": 1,
        "regex": "/.*prod=\"([^\"]*).*/",
        "sort": 1,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      }`

func templatingNAMESPACE(Namespace string) string {
	var templatingNAMESPACE = `{
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
	return templatingNAMESPACE
}

var templatingProjects = `{
        "allValue": ".*",
        "datasource": "prometheus_source",
        "hide": 0,
        "includeAll": true,
        "label": null,
        "multi": false,
        "name": "Projects",
        "options": [],
        "query": "container_memory_working_set_bytes{namespace=~\"$NAMESPACE-.*\",kubernetes_io_hostname=~\"$DedicatedNodes\"}",
        "refresh": 1,
        "regex": "/.*namespace=\"([^\"]*).*/",
        "sort": 1,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      }`
var templatingPods = `{
        "allValue": ".*",
        "current": {
          "text": "none",
          "value": "none"
        },
        "datasource": "prometheus_source",
        "hide": 0,
        "includeAll": true,
        "label": null,
        "multi": false,
        "name": "Pods",
        "options": [],
        "query": "container_memory_working_set_bytes{namespace=~\"$NAMESPACE-.*\",namespace=~\"$Projects\",kubernetes_io_hostname=~\"$DedicatedNodes\"}",
        "refresh": 1,
        "regex": "/.*pod_name=\"([^\"]*).*/",
        "sort": 1,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      }`
var templatingAlerts = `{
        "allValue": "",
        "current": {
        },
        "datasource": "prometheus_source",
        "hide": 0,
        "includeAll": false,
        "label": null,
        "multi": false,
        "name": "alerts",
        "options": [],
        "query": "ALERTS{alertname=~\".*\"}",
        "refresh": 1,
        "regex": "/.*alertname=\"([^\"]*).*/",
        "sort": 3,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      }`
