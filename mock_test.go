package marathon

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"
)

const someDeployments = `[{
  "id": "97c136bf-5a28-4821-9d94-480d9fbb01c8",
  "version": "2015-09-30T09:09:17.614Z",
  "affectedApps": [ "/foo" ],
  "affectedPods": [ "/bla" ],
  "steps": [
      { "actions": [{ "action": "ScaleApplication", "app": "/foo" }]},
      { "actions": [{ "action": "ScalePod", "pod": "/bla" }]}
  ],
  "currentActions": [{
    "action": "ScaleApplication",
    "app": "/foo",
    "readinessCheckResults": [{
      "taskId": "foo.c9de6033",
      "lastResponse": { "body": "{}", "contentType": "application/json", "status": 500 },
      "name": "myReadyCheck",
      "ready": false
    }]}],
    "currentStep": 1,
    "totalSteps": 1
  }]`

var deployArray = `[]`

var appsArray = `{
 "apps": [{
   "id": "/infra/redis-1",
   "acceptedResourceRoles": [ "*" ],
   "backoffFactor": 1.15,
   "backoffSeconds": 1,
   "container": { "type": "DOCKER",
    "docker": { "image": "docker.io/redis-ha:5.0.5", "privileged": false, "forcePullImage": true },
	"volumes": [
     { "containerPath": "/data", "hostPath": "/var/lib/mesos/redis/1/data", "mode": "RW" },
	 { "containerPath": "/conf", "hostPath": "/data/redis-1/conf", "mode": "RW" },
	 { "containerPath": "/etc/localtime", "hostPath": "/etc/localtime", "mode": "RO" }
	],
	"portMappings": [
	 { "containerPort": 46379, "labels": { "VIP_0": "/redissrv:46379" }, "protocol": "tcp", "servicePort": 10013 }
	]
   },
   "cpus": 1,
   "env": { "REDISPORT": "46379", "REDISPRTY": "2" },
   "fetch": [
	{ "uri": "file:///data/registry-auth/docker.tar.gz", "extract": true, "executable": false, "cache": false }
   ],
   "healthChecks": [
	{ "gracePeriodSeconds": 60, "intervalSeconds": 5, "maxConsecutiveFailures": 3, "path": "", "portIndex": 0, "protocol": "TCP", "ipProtocol": "IPv4", "timeoutSeconds": 5, "delaySeconds": 15 }
   ],
   "instances": 1,
   "labels": { "ENVIRONMENT": "testing" },
   "maxLaunchDelaySeconds": 3600,
   "mem": 8192,
   "networks": [ { "mode": "container/bridge" } ],
   "upgradeStrategy": { "maximumOverCapacity": 0, "minimumHealthCapacity": 0 },
   "killSelection": "YOUNGEST_FIRST",
   "unreachableStrategy": { "inactiveAfterSeconds": 0, "expungeAfterSeconds": 0 },
   "role": "slave_public"
  },
  {
   "id": "/infra/broker-0",
   "acceptedResourceRoles": [ "*" ],
   "backoffFactor": 1.15,
   "backoffSeconds": 1,
   "container": { "type": "DOCKER",
    "docker": { "image": "docker.io/kafka-ha:2.0.1", "privileged": false, "forcePullImage": true },
    "volumes": [
	 { "containerPath": "/data", "hostPath": "/var/lib/mesos/broker-0/0", "mode": "RW" },
     { "containerPath": "/conf", "hostPath": "/var/lib/mesos/broker-0/1", "mode": "RW" },
	 { "containerPath": "/etc/localtime", "hostPath": "/etc/localtime", "mode": "RO" }
	],
	"portMappings": [ { "containerPort": 46379, "labels": { "VIP_0": "/broker-0:9092" }, "protocol": "tcp", "servicePort": 10014 } ]
   },
   "cpus": 1,
   "env": { "BROKERPORT": "9092" },
   "fetch": [ { "uri": "file:///data/registry-auth/docker.tar.gz", "extract": true, "executable": false, "cache": false } ],
   "healthChecks": [ { "gracePeriodSeconds": 60, "intervalSeconds": 5, "maxConsecutiveFailures": 3, "path": "", "portIndex": 0, "protocol": "TCP", "ipProtocol": "IPv4", "timeoutSeconds": 5, "delaySeconds": 15 } ],
   "instances": 1,
   "labels": { "ENVIRONMENT": "testing" },
   "maxLaunchDelaySeconds": 3600,
   "mem": 8192,
   "networks": [ { "mode": "container/bridge" } ],
   "upgradeStrategy": { "maximumOverCapacity": 0, "minimumHealthCapacity": 0 },
   "killSelection": "YOUNGEST_FIRST",
   "unreachableStrategy": { "inactiveAfterSeconds": 0, "expungeAfterSeconds": 0 },
   "role": "slave_public"
  }
 ]
}`

var groupsArray = `{
  "id": "/infra",
  "apps": [
    {
      "id": "/infra/redis",
      "acceptedResourceRoles": [
        "*"
      ],
      "backoffFactor": 1.15,
      "backoffSeconds": 1,
      "container": {
        "type": "DOCKER",
        "docker": {
          "image": "docker.io/redis-ha:4.0.10",
          "privileged": true,
          "forcePullImage": true
        },
        "volumes": [
          {
            "containerPath": "/data",
            "hostPath": "/data/redis/data",
            "mode": "RW"
          },
          {
            "containerPath": "/conf",
            "hostPath": "/data/redis/conf",
            "mode": "RW"
          }
        ],
        "portMappings": [
          {
            "containerPort": 6379,
            "labels": {
              "VIP_0": "/redissrv:6379"
            },
            "protocol": "tcp",
            "servicePort": 10013
          }
        ]
      },
      "cpus": 4,
      "env": {
        "EMPTY": "true"
      },
      "fetch": [
        {
          "uri": "file:///data/registry-auth/docker.tar.gz",
          "extract": true,
          "executable": false,
          "cache": false
        }
      ],
      "healthChecks": [
        {
          "gracePeriodSeconds": 60,
          "intervalSeconds": 5,
          "maxConsecutiveFailures": 3,
          "path": "",
          "portIndex": 0,
          "protocol": "TCP",
          "ipProtocol": "IPv4",
          "timeoutSeconds": 5,
          "delaySeconds": 15
        }
      ],
      "instances": 0,
      "labels": {
        "ENVIRONMENT": "testing"
      },
      "maxLaunchDelaySeconds": 3600,
      "mem": 8192,
      "networks": [
        {
          "mode": "container/bridge"
        }
      ],
      "upgradeStrategy": {
        "maximumOverCapacity": 0,
        "minimumHealthCapacity": 0
      },
      "killSelection": "YOUNGEST_FIRST",
      "unreachableStrategy": {
        "inactiveAfterSeconds": 0,
        "expungeAfterSeconds": 0
      },
      "role": "slave_public"
    },
    {
      "id": "/infra/kong-v2",
      "acceptedResourceRoles": [
        "*"
      ],
      "backoffFactor": 1.15,
      "backoffSeconds": 1,
      "container": {
        "type": "DOCKER",
        "docker": {
          "image": "docker.io/kong:2.0.4-dev",
          "privileged": false,
          "forcePullImage": true
        },
        "portMappings": [
          {
            "containerPort": 8443,
            "hostPort": 8443,
            "labels": {
              "VIP_0": "/api:8443"
            },
            "protocol": "tcp",
            "servicePort": 10078
          },
          {
            "containerPort": 8001,
            "hostPort": 8001,
            "labels": {
              "VIP_0": "/api:8001"
            },
            "protocol": "tcp",
            "servicePort": 10079
          }
        ]
      },
      "cpus": 1,
      "fetch": [
        {
          "uri": "file:///data/registry-auth/docker.tar.gz",
          "extract": true,
          "executable": false,
          "cache": false
        }
      ],
      "healthChecks": [
        {
          "gracePeriodSeconds": 20,
          "intervalSeconds": 60,
          "maxConsecutiveFailures": 5,
          "path": "/",
          "portIndex": 1,
          "protocol": "MESOS_HTTP",
          "ipProtocol": "IPv4",
          "timeoutSeconds": 5,
          "delaySeconds": 15
        }
      ],
      "instances": 1,
      "labels": {
        "ENVIRONMENT": "testing"
      },
      "maxLaunchDelaySeconds": 300,
      "mem": 4096,
      "networks": [
        {
          "mode": "container/bridge"
        }
      ],
      "upgradeStrategy": {
        "maximumOverCapacity": 1,
        "minimumHealthCapacity": 0
      },
      "killSelection": "YOUNGEST_FIRST",
      "unreachableStrategy": {
        "inactiveAfterSeconds": 0,
        "expungeAfterSeconds": 0
      },
      "role": "slave_public"
    }
  ],
  "groups": [
    {
      "id": "/infra/kafka",
      "apps": [
        {
          "id": "/infra/kafka/broker-0",
          "acceptedResourceRoles": [
            "*"
          ],
          "backoffFactor": 1.15,
          "backoffSeconds": 1,
          "container": {
            "type": "DOCKER",
            "docker": {
              "image": "docker.io/kafka-ha:2.1.1",
              "privileged": false,
              "forcePullImage": true
            },
            "volumes": [
              {
                "containerPath": "/var/log/kafka",
                "hostPath": "/data/kafka/broker-0/0",
                "mode": "RW"
              },
              {
                "containerPath": "/var/lib/kafka",
                "hostPath": "/data/kafka/broker-0/1",
                "mode": "RW"
              }
            ],
            "portMappings": [
              {
                "containerPort": 9092,
                "labels": {
                  "VIP_0": "/broker-0:9092"
                },
                "protocol": "tcp",
                "servicePort": 10014
              }
            ]
          },
          "cpus": 1,
          "env": {
            "ADVERTISED_HOST": "broker-0.marathon.l4lb.thisdcos.directory",
            "ADVERTISED_PORT": "9092",
            "AUTO_CREATE_TOPICS": "true",
            "BROKER_ID": "0",
            "DEBUG": "false",
            "JAVA_OPTS": "-Duser.timezone=GMT-3 -Xms2g -Xmx2g",
            "ZK_SERVER": "zk.marathon.l4lb.thisdcos.directory:2181/kafka"
          },
          "fetch": [
            {
              "uri": "file:///data/registry-auth/docker.tar.gz",
              "extract": true,
              "executable": false,
              "cache": false
            }
          ],
          "healthChecks": [
            {
              "gracePeriodSeconds": 60,
              "intervalSeconds": 5,
              "maxConsecutiveFailures": 3,
              "path": "",
              "portIndex": 0,
              "protocol": "TCP",
              "ipProtocol": "IPv4",
              "timeoutSeconds": 5,
              "delaySeconds": 15
            }
          ],
          "instances": 1,
          "labels": {
            "ENVIRONMENT": "testing"
          },
          "maxLaunchDelaySeconds": 3600,
          "mem": 2048,
          "networks": [
            {
              "mode": "container/bridge"
            }
          ],
          "upgradeStrategy": {
            "maximumOverCapacity": 1,
            "minimumHealthCapacity": 1
          },
          "killSelection": "YOUNGEST_FIRST",
          "unreachableStrategy": {
            "inactiveAfterSeconds": 0,
            "expungeAfterSeconds": 0
          },
          "role": "slave_public"
        },
        {
          "id": "/infra/kafka/broker-1",
          "acceptedResourceRoles": [
            "*"
          ],
          "backoffFactor": 1.15,
          "backoffSeconds": 1,
          "container": {
            "type": "DOCKER",
            "docker": {
              "image": "docker.io/kafka-ha:2.1.1",
              "privileged": false,
              "forcePullImage": true
            },
            "volumes": [
              {
                "containerPath": "/var/log/kafka",
                "hostPath": "/data/kafka/broker-1/0",
                "mode": "RW"
              },
              {
                "containerPath": "/var/lib/kafka",
                "hostPath": "/data/kafka/broker-1/1",
                "mode": "RW"
              }
            ],
            "portMappings": [
              {
                "containerPort": 9092,
                "labels": {
                  "VIP_0": "/broker-1:9092"
                },
                "protocol": "tcp",
                "servicePort": 10014
              }
            ]
          },
          "cpus": 1,
          "env": {
            "ADVERTISED_HOST": "broker-1.marathon.l4lb.thisdcos.directory",
            "ADVERTISED_PORT": "9092",
            "AUTO_CREATE_TOPICS": "true",
            "BROKER_ID": "1",
            "DEBUG": "true",
            "JAVA_OPTS": "-Duser.timezone=GMT-3 -Xms2g -Xmx2g",
            "ZK_SERVER": "zk.marathon.l4lb.thisdcos.directory:2181/kafka"
          },
          "fetch": [
            {
              "uri": "file:///data/registry-auth/docker.tar.gz",
              "extract": true,
              "executable": false,
              "cache": false
            }
          ],
          "healthChecks": [
            {
              "gracePeriodSeconds": 60,
              "intervalSeconds": 5,
              "maxConsecutiveFailures": 3,
              "path": "",
              "portIndex": 0,
              "protocol": "TCP",
              "ipProtocol": "IPv4",
              "timeoutSeconds": 5,
              "delaySeconds": 15
            }
          ],
          "instances": 1,
          "labels": {
            "ENVIRONMENT": "testing"
          },
          "maxLaunchDelaySeconds": 3600,
          "mem": 2048,
          "networks": [
            {
              "mode": "container/bridge"
            }
          ],
          "upgradeStrategy": {
            "maximumOverCapacity": 1,
            "minimumHealthCapacity": 1
          },
          "killSelection": "YOUNGEST_FIRST",
          "unreachableStrategy": {
            "inactiveAfterSeconds": 0,
            "expungeAfterSeconds": 0
          },
          "role": "slave_public"
        },
        {
          "id": "/infra/kafka/broker-2",
          "acceptedResourceRoles": [
            "*"
          ],
          "backoffFactor": 1.15,
          "backoffSeconds": 1,
          "container": {
            "type": "DOCKER",
            "docker": {
              "image": "docker.io/kafka-ha:2.1.1",
              "privileged": false,
              "forcePullImage": true
            },
            "volumes": [
              {
                "containerPath": "/var/log/kafka",
                "hostPath": "/data/kafka/broker-2/0",
                "mode": "RW"
              },
              {
                "containerPath": "/var/lib/kafka",
                "hostPath": "/data/kafka/broker-2/1",
                "mode": "RW"
              }
            ],
            "portMappings": [
              {
                "containerPort": 9092,
                "labels": {
                  "VIP_0": "/broker-2:9092"
                },
                "protocol": "tcp",
                "servicePort": 10014
              }
            ]
          },
          "cpus": 1,
          "env": {
            "ADVERTISED_HOST": "broker-2.marathon.l4lb.thisdcos.directory",
            "ADVERTISED_PORT": "9092",
            "AUTO_CREATE_TOPICS": "true",
            "BROKER_ID": "2",
            "DEBUG": "true",
            "JAVA_OPTS": "-Duser.timezone=GMT-3 -Xms2g -Xmx2g",
            "ZK_SERVER": "zk.marathon.l4lb.thisdcos.directory:2181/kafka"
          },
          "fetch": [
            {
              "uri": "file:///data/registry-auth/docker.tar.gz",
              "extract": true,
              "executable": false,
              "cache": false
            }
          ],
          "healthChecks": [
            {
              "gracePeriodSeconds": 60,
              "intervalSeconds": 5,
              "maxConsecutiveFailures": 3,
              "path": "",
              "portIndex": 0,
              "protocol": "TCP",
              "ipProtocol": "IPv4",
              "timeoutSeconds": 5,
              "delaySeconds": 15
            }
          ],
          "instances": 1,
          "labels": {
            "ENVIRONMENT": "testing"
          },
          "maxLaunchDelaySeconds": 3600,
          "mem": 2048,
          "networks": [
            {
              "mode": "container/bridge"
            }
          ],
          "upgradeStrategy": {
            "maximumOverCapacity": 1,
            "minimumHealthCapacity": 1
          },
          "killSelection": "YOUNGEST_FIRST",
          "unreachableStrategy": {
            "inactiveAfterSeconds": 0,
            "expungeAfterSeconds": 0
          },
          "role": "slave_public"
        }
      ],
      "groups": [],
      "pods": [],
      "version": "2020-06-25T11:50:34.096Z",
      "versionInfo": {
        "lastScalingAt": "0001-01-01T00:00:00Z",
        "lastConfigChangeAt": "0001-01-01T00:00:00Z"
      }
    }
  ],
  "pods": [],
  "version": "2020-06-25T11:50:34.096Z",
  "versionInfo": {
    "lastScalingAt": "0001-01-01T00:00:00Z",
    "lastConfigChangeAt": "0001-01-01T00:00:00Z"
  }
}`

var appRedis = `{
  "app":   {
   "id": "/infra/redis-1",
   "acceptedResourceRoles": [ "*" ],
   "backoffFactor": 1.15,
   "backoffSeconds": 1,
   "container": { "type": "DOCKER",
    "docker": { "image": "docker.io/redis-ha:5.0.5", "privileged": false, "forcePullImage": true, "parameters": [{"key": "add-host", "value": "10.128.64.32"}]},
	"volumes": [
     { "containerPath": "/data", "hostPath": "/var/lib/mesos/redis/1/data", "mode": "RW" },
	 { "containerPath": "/conf", "hostPath": "/data/redis-1/conf", "mode": "RW" },
	 { "containerPath": "/etc/localtime", "hostPath": "/etc/localtime", "mode": "RO" }
	],
	"portMappings": [
	 { "containerPort": 46379, "labels": { "VIP_0": "/redissrv:46379" }, "protocol": "tcp", "servicePort": 10013 }
	]
   },
   "cpus": 1,
   "env": { "REDISPORT": "46379", "REDISPRTY": "2" },
   "fetch": [
	{ "uri": "file:///data/registry-auth/docker.tar.gz", "extract": true, "executable": false, "cache": false }
   ],
   "healthChecks": [
	{ "gracePeriodSeconds": 60, "intervalSeconds": 5, "maxConsecutiveFailures": 3, "path": "", "portIndex": 0, "protocol": "TCP", "ipProtocol": "IPv4", "timeoutSeconds": 5, "delaySeconds": 15 }
   ],
   "instances": 1,
   "labels": { "ENVIRONMENT": "testing" },
   "maxLaunchDelaySeconds": 3600,
   "mem": 8192,
   "networks": [ { "mode": "container/bridge" } ],
   "upgradeStrategy": { "maximumOverCapacity": 0, "minimumHealthCapacity": 0 },
   "killSelection": "YOUNGEST_FIRST",
   "unreachableStrategy": { "inactiveAfterSeconds": 0, "expungeAfterSeconds": 0 },
   "role": "slave_public"
  }
}`

var appBroker = `{
 "app": {
   "id": "/infra/broker-0",
   "acceptedResourceRoles": [ "*" ],
   "backoffFactor": 1.15,
   "backoffSeconds": 1,
   "container": { "type": "DOCKER",
    "docker": { "image": "docker.io/kafka-ha:2.0.1", "privileged": false, "forcePullImage": true },
    "volumes": [
	 { "containerPath": "/data", "hostPath": "/var/lib/mesos/broker-0/0", "mode": "RW" },
     { "containerPath": "/conf", "hostPath": "/var/lib/mesos/broker-0/1", "mode": "RW" },
	 { "containerPath": "/etc/localtime", "hostPath": "/etc/localtime", "mode": "RO" }
	],
	"portMappings": [ { "containerPort": 46379, "labels": { "VIP_0": "/broker-0:9092" }, "protocol": "tcp", "servicePort": 10014 } ]
   },
   "cpus": 1,
   "env": { "BROKERPORT": "9092" },
   "fetch": [ { "uri": "file:///data/registry-auth/docker.tar.gz", "extract": true, "executable": false, "cache": false } ],
   "healthChecks": [ { "gracePeriodSeconds": 60, "intervalSeconds": 5, "maxConsecutiveFailures": 3, "path": "", "portIndex": 0, "protocol": "TCP", "ipProtocol": "IPv4", "timeoutSeconds": 5, "delaySeconds": 15 } ],
   "instances": 1,
   "labels": { "ENVIRONMENT": "testing" },
   "maxLaunchDelaySeconds": 3600,
   "mem": 8192,
   "networks": [ { "mode": "container/bridge" } ],
   "upgradeStrategy": { "maximumOverCapacity": 0, "minimumHealthCapacity": 0 },
   "killSelection": "YOUNGEST_FIRST",
   "unreachableStrategy": { "inactiveAfterSeconds": 0, "expungeAfterSeconds": 0 },
   "role": "slave_public"
  }
}`

var kongApp = `{
	"app": {
		"id": "/infra/kong-v2",
		"acceptedResourceRoles": [
			"*"
		],
		"backoffFactor": 1.15,
		"backoffSeconds": 1,
		"container": {
			"type": "DOCKER",
			"docker": {
				"forcePullImage": true,
				"image": "kong:2.0.2",
				"parameters": [],
				"privileged": false
			},
			"volumes": [],
			"portMappings": [
				{
					"containerPort": 8443,
					"hostPort": 11567,
					"labels": {
						"VIP_0": "/api:8443"
					},
					"name": "proxy-v2",
					"protocol": "tcp",
					"servicePort": 0
				},
				{
					"containerPort": 8001,
					"hostPort": 11568,
					"labels": {
						"VIP_0": "/api:8001"
					},
					"name": "admin-v2",
					"protocol": "tcp",
					"servicePort": 0
				}
			]
		},
		"cpus": 1,
		"disk": 0,
		"executor": "",
		"fetch": [
			{
				"uri": "file:///data/registry-auth/docker.tar.gz",
				"extract": true,
				"executable": false,
				"cache": false
			}
		],
		"healthChecks": [
			{
				"gracePeriodSeconds": 20,
				"intervalSeconds": 60,
				"maxConsecutiveFailures": 5,
				"path": "/",
				"portIndex": 1,
				"protocol": "MESOS_HTTP",
				"ipProtocol": "IPv4",
				"timeoutSeconds": 5,
				"delaySeconds": 15
			}
		],
		"instances": 1,
		"labels": {
			"ENVIRONMENT": "testing"
		},
		"maxLaunchDelaySeconds": 300,
		"mem": 4096,
		"gpus": 0,
		"networks": [
			{
				"mode": "container/bridge"
			}
		],
		"requirePorts": false,
		"upgradeStrategy": {
			"maximumOverCapacity": 1,
			"minimumHealthCapacity": 0
		},
		"killSelection": "YOUNGEST_FIRST",
		"unreachableStrategy": {
			"inactiveAfterSeconds": 0,
			"expungeAfterSeconds": 0
		},
		"role": "slave_public"
	}
}`

var times int = 0

func MockMarathonServer() *httptest.Server {
	// Mock Marathon server
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fakeDeploy := &Response{
			ID:      "d4b75430-8ee6-47e9-95f2-6cf297aaac00",
			Version: time.Now(),
		}
		buffer, _ := json.Marshal(fakeDeploy)

		switch r.URL.Path {

		case "/ping":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`pong`))

		case "/v2/info":
			fakeInfo := &Info{
				Name:           "mock_marathon",
				Version:        "v1.0.0",
				Buildref:       "2020.01.01",
				Elected:        false,
				Leader:         "127.0.0.10:8080",
				FrameworkID:    "97c136bf-5a28-4821-9d94-480d9fbb01c8",
				MarathonConfig: Config{},
				ZookeeperConfig: ZkConfig{
					Zk:                     "127.0.0.10:2181",
					ZkCompression:          false,
					ZkCompressionThreshold: 0,
					ZkConnectionTimeout:    0,
					ZkMaxNodeSize:          0,
					ZkMaxVersions:          0,
					ZkSessionTimeout:       0,
					ZkTimeout:              0,
				},
				HTTPConf: HTTPConfig{},
			}
			infoBuffer, _ := json.Marshal(fakeInfo)

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(infoBuffer))

		case "/v2/deployments/":

			switch r.Method {

			case http.MethodGet:
				// Define a "break condition" for test eventually finish of a deployment
				times++

				if times > 12 {
					deployArray = `[]`
				}

				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write([]byte(deployArray))

			case http.MethodDelete:
				w.WriteHeader(http.StatusAccepted)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			}

		case "/v2/deployments/97c136bf-5a28-4821-9d94-480d9fbb01c8":

			switch r.Method {

			case http.MethodDelete:
				w.WriteHeader(http.StatusAccepted)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			}

		case "/v2/apps/":

			switch r.Method {

			case http.MethodGet:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write([]byte(appsArray))

			case http.MethodPost:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			}

		case "/v2/groups/infra":

			switch r.Method {

			case http.MethodGet:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write([]byte(groupsArray))

			case http.MethodPost:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			case http.MethodDelete:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			}

		case "/v2/apps/infra/redis-1":

			switch r.Method {

			case http.MethodGet:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write([]byte(appRedis))

			case http.MethodPatch:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			case http.MethodPut:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			case http.MethodPost:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			case http.MethodDelete:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			}

		case "/v2/apps/infra/broker-0":

			switch r.Method {

			case http.MethodGet:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write([]byte(appBroker))

			case http.MethodPatch:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			case http.MethodPut:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			case http.MethodPost:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			case http.MethodDelete:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			}

		case "/v2/apps/infra/redis-1/restart":

			switch r.Method {

			case http.MethodPost:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)
			}

		case "/v2/apps/infra/broker-0/restart":

			switch r.Method {

			case http.MethodPost:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)
			}

		case "/v2/apps/infra/kong-v2":

			switch r.Method {

			case http.MethodGet:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write([]byte(kongApp))

			case http.MethodDelete:
				w.WriteHeader(http.StatusOK)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write(buffer)

			}
		}
	}),
	)
}
