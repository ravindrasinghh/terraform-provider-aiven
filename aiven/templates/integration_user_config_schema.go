// Code generated by go generate; DO NOT EDIT.
// This file was generated at Wed  1 Jul 2020 10:11:07 CEST

package templates

import "encoding/json"

var (
  // inline JSON was taken from a file:
  // integrations_user_config_schema.json
  integrationJSON = `{
  "dashboard": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "datadog": {
    "additionalProperties": false,
    "properties": {
      "kafka_custom_metrics": {
        "items": {
          "example": "kafka.log.log_size",
          "maxLength": 1024,
          "title": "Metric name",
          "type": "string"
        },
        "maxItems": 1024,
        "title": "List of custom metrics",
        "type": "array"
      }
    },
    "type": "object"
  },
  "external_elasticsearch_logs": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "jolokia": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "kafka_connect": {
    "additionalProperties": false,
    "properties": {
      "kafka_connect": {
        "additionalProperties": false,
        "properties": {
          "config_storage_topic": {
            "example": "__connect_configs",
            "maxLength": 249,
            "title": "The name of the topic where connector and task configuration data are stored.This must be the same for all workers with the same group_id.",
            "type": "string"
          },
          "group_id": {
            "example": "connect",
            "maxLength": 249,
            "title": "A unique string that identifies the Connect cluster group this worker belongs to.",
            "type": "string"
          },
          "offset_storage_topic": {
            "example": "__connect_offsets",
            "maxLength": 249,
            "title": "The name of the topic where connector and task configuration offsets are stored.This must be the same for all workers with the same group_id.",
            "type": "string"
          },
          "status_storage_topic": {
            "example": "__connect_status",
            "maxLength": 249,
            "title": "The name of the topic where connector and task configuration status updates are stored.This must be the same for all workers with the same group_id.",
            "type": "string"
          }
        },
        "title": "Kafka Connect service configuration values",
        "type": "object"
      }
    },
    "title": "Integration user config",
    "type": "object"
  },
  "kafka_mirrormaker": {
    "additionalProperties": false,
    "properties": {
      "cluster_alias": {
        "description": "The alias under which the Kafka cluster is known to MirrorMaker. Can contain the following symbols: ASCII alphanumerics, '.', '_', and '-'.",
        "example": "kafka-abc",
        "maxLength": 128,
        "pattern": "^[a-zA-Z0-9_.-]+$",
        "title": "Kafka cluster alias",
        "type": "string",
        "user_error": "Must consist of alpha-numeric ASCII characters, dashes, underscores, and dots."
      }
    },
    "title": "Integration user config",
    "type": "object"
  },
  "logs": {
    "additionalProperties": false,
    "properties": {
      "elasticsearch_index_days_max": {
        "default": 3,
        "example": 5,
        "maximum": 10000,
        "minimum": 1,
        "title": "Elasticsearch index retention limit",
        "type": "integer"
      },
      "elasticsearch_index_prefix": {
        "default": "logs",
        "example": "logs",
        "maxLength": 1024,
        "minLength": 1,
        "title": "Elasticsearch index prefix",
        "type": "string"
      }
    },
    "type": "object"
  },
  "metrics": {
    "additionalProperties": false,
    "properties": {
      "database": {
        "example": "metrics",
        "maxLength": 40,
        "pattern": "^[_A-Za-z0-9][-_A-Za-z0-9]{0,39}$",
        "title": "Name of the database where to store metric datapoints. Only affects PostgreSQL destinations. Defaults to 'metrics'. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
        "type": "string",
        "user_error": "Must consist of alpha-numeric characters, underscores or dashes, may not start with dash, max 40 characters"
      },
      "retention_days": {
        "example": 30,
        "maximum": 10000,
        "minimum": 0,
        "title": "Number of days to keep old metrics. Only affects PostgreSQL destinations. Set to 0 for no automatic cleanup. Defaults to 30 days.",
        "type": "integer"
      },
      "ro_username": {
        "example": "metrics_reader",
        "maxLength": 40,
        "pattern": "^[_A-Za-z0-9][-._A-Za-z0-9]{0,39}$",
        "title": "Name of a user that can be used to read metrics. This will be used for Grafana integration (if enabled) to prevent Grafana users from making undesired changes. Only affects PostgreSQL destinations. Defaults to 'metrics_reader'. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
        "type": "string",
        "user_error": "Must consist of alpha-numeric characters, dots, underscores or dashes, may not start with dash or dot, max 40 characters"
      },
      "source_mysql": {
        "additionalProperties": false,
        "properties": {
          "telegraf": {
            "additionalProperties": false,
            "properties": {
              "gather_event_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.EVENT_WAITS",
                "type": "boolean"
              },
              "gather_file_events_stats": {
                "example": false,
                "title": "gather metrics from PERFORMANCE_SCHEMA.FILE_SUMMARY_BY_EVENT_NAME",
                "type": "boolean"
              },
              "gather_index_io_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_INDEX_USAGE",
                "type": "boolean"
              },
              "gather_info_schema_auto_inc": {
                "example": false,
                "title": "Gather auto_increment columns and max values from information schema",
                "type": "boolean"
              },
              "gather_innodb_metrics": {
                "example": true,
                "title": "Gather metrics from INFORMATION_SCHEMA.INNODB_METRICS",
                "type": "boolean"
              },
              "gather_perf_events_statements": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.EVENTS_STATEMENTS_SUMMARY_BY_DIGEST",
                "type": "boolean"
              },
              "gather_process_list": {
                "example": true,
                "title": "Gather thread state counts from INFORMATION_SCHEMA.PROCESSLIST",
                "type": "boolean"
              },
              "gather_slave_status": {
                "example": true,
                "title": "Gather metrics from SHOW SLAVE STATUS command output",
                "type": "boolean"
              },
              "gather_table_io_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_TABLE",
                "type": "boolean"
              },
              "gather_table_lock_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.TABLE_LOCK_WAITS",
                "type": "boolean"
              },
              "gather_table_schema": {
                "example": true,
                "title": "Gather metrics from INFORMATION_SCHEMA.TABLES",
                "type": "boolean"
              },
              "perf_events_statements_digest_text_limit": {
                "example": 120,
                "maximum": 2048,
                "minimum": 1,
                "title": "Truncates digest text from perf_events_statements into this many characters",
                "type": "integer"
              },
              "perf_events_statements_limit": {
                "example": 250,
                "maximum": 4000,
                "minimum": 1,
                "title": "Limits metrics from perf_events_statements",
                "type": "integer"
              },
              "perf_events_statements_time_limit": {
                "example": 86400,
                "maximum": 2592000,
                "minimum": 1,
                "title": "Only include perf_events_statements whose last seen is less than this many seconds",
                "type": "integer"
              }
            },
            "title": "Configuration options for Telegraf MySQL input plugin",
            "type": "object"
          }
        },
        "title": "Configuration options for metrics where source service is MySQL",
        "type": "object"
      },
      "username": {
        "example": "metrics_writer",
        "maxLength": 40,
        "pattern": "^[_A-Za-z0-9][-._A-Za-z0-9]{0,39}$",
        "title": "Name of the user used to write metrics. Only affects PostgreSQL destinations. Defaults to 'metrics_writer'. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
        "type": "string",
        "user_error": "Must consist of alpha-numeric characters, dots, underscores or dashes, may not start with dash or dot, max 40 characters"
      }
    },
    "title": "Integration user config",
    "type": "object"
  },
  "mirrormaker": {
    "additionalProperties": false,
    "properties": {
      "mirrormaker_whitelist": {
        "default": ".*",
        "example": ".*",
        "maxLength": 1000,
        "minLength": 1,
        "title": "Mirrormaker topic whitelist",
        "type": "string"
      }
    },
    "type": "object"
  },
  "prometheus": {
    "additionalProperties": false,
    "properties": {
      "source_mysql": {
        "additionalProperties": false,
        "properties": {
          "telegraf": {
            "additionalProperties": false,
            "properties": {
              "gather_event_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.EVENT_WAITS",
                "type": "boolean"
              },
              "gather_file_events_stats": {
                "example": false,
                "title": "gather metrics from PERFORMANCE_SCHEMA.FILE_SUMMARY_BY_EVENT_NAME",
                "type": "boolean"
              },
              "gather_index_io_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_INDEX_USAGE",
                "type": "boolean"
              },
              "gather_info_schema_auto_inc": {
                "example": false,
                "title": "Gather auto_increment columns and max values from information schema",
                "type": "boolean"
              },
              "gather_innodb_metrics": {
                "example": true,
                "title": "Gather metrics from INFORMATION_SCHEMA.INNODB_METRICS",
                "type": "boolean"
              },
              "gather_perf_events_statements": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.EVENTS_STATEMENTS_SUMMARY_BY_DIGEST",
                "type": "boolean"
              },
              "gather_process_list": {
                "example": true,
                "title": "Gather thread state counts from INFORMATION_SCHEMA.PROCESSLIST",
                "type": "boolean"
              },
              "gather_slave_status": {
                "example": true,
                "title": "Gather metrics from SHOW SLAVE STATUS command output",
                "type": "boolean"
              },
              "gather_table_io_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_TABLE",
                "type": "boolean"
              },
              "gather_table_lock_waits": {
                "example": false,
                "title": "Gather metrics from PERFORMANCE_SCHEMA.TABLE_LOCK_WAITS",
                "type": "boolean"
              },
              "gather_table_schema": {
                "example": true,
                "title": "Gather metrics from INFORMATION_SCHEMA.TABLES",
                "type": "boolean"
              },
              "perf_events_statements_digest_text_limit": {
                "example": 120,
                "maximum": 2048,
                "minimum": 1,
                "title": "Truncates digest text from perf_events_statements into this many characters",
                "type": "integer"
              },
              "perf_events_statements_limit": {
                "example": 250,
                "maximum": 4000,
                "minimum": 1,
                "title": "Limits metrics from perf_events_statements",
                "type": "integer"
              },
              "perf_events_statements_time_limit": {
                "example": 86400,
                "maximum": 2592000,
                "minimum": 1,
                "title": "Only include perf_events_statements whose last seen is less than this many seconds",
                "type": "integer"
              }
            },
            "title": "Configuration options for Telegraf MySQL input plugin",
            "type": "object"
          }
        },
        "title": "Configuration options for metrics where source service is MySQL",
        "type": "object"
      }
    },
    "title": "Integration user config",
    "type": "object"
  },
  "read_replica": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "rsyslog": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  },
  "signalfx": {
    "additionalProperties": false,
    "title": "Integration user config",
    "type": "object"
  }
}`
)

func init() {
  var integrationSchema map[string]interface{}
  if err := json.Unmarshal([]byte(integrationJSON), &integrationSchema); err != nil {
    panic("cannot unmarshal user configuration options integration JSON', error :" + err.Error())
  }
  userConfigSchemas["integration"] = integrationSchema
}
