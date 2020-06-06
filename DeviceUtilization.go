package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type SnmpDeviceUtilization struct {
	Aggregations Aggregations `json:"aggregations"`
}
type MemorySize struct {
	Value float64 `json:"value"`
}
type Memory struct {
	Value float64 `json:"value"`
}
type Buckets struct {
	Key        string     `json:"key"`
	DocCount   int        `json:"doc_count"`
	MemorySize MemorySize `json:"MemorySize"`
	Memory     Memory     `json:"memory"`
}
type ByMemorySize struct {
	DocCountErrorUpperBound int       `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int       `json:"sum_other_doc_count"`
	Buckets                 []Buckets `json:"buckets"`
}
type DiskSize struct {
	Value float64 `json:"value"`
}
type Disk struct {
	Value float64 `json:"value"`
}
type Buckets struct {
	Key      string   `json:"key"`
	DocCount int      `json:"doc_count"`
	DiskSize DiskSize `json:"diskSize"`
	Disk     Disk     `json:"disk"`
}
type ByDiskStorage struct {
	DocCountErrorUpperBound int       `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int       `json:"sum_other_doc_count"`
	Buckets                 []Buckets `json:"buckets"`
}
type GetNestedStorage struct {
	DocCount      int           `json:"doc_count"`
	ByMemorySize  ByMemorySize  `json:"by_memory_size"`
	ByDiskStorage ByDiskStorage `json:"by_disk_storage"`
}
type GetCPUTolalLoadAverage struct {
	Value float64 `json:"value"`
}
type GetCPUAverage struct {
	Value float64 `json:"value"`
}
type Buckets struct {
	KeyAsString   string        `json:"key_as_string"`
	Key           int64         `json:"key"`
	DocCount      int           `json:"doc_count"`
	GetCPUAverage GetCPUAverage `json:"get_cpu_average"`
}
type GraphDataForCPU struct {
	Buckets []Buckets `json:"buckets"`
}
type DiskAverage struct {
	Value float64 `json:"value"`
}
type Buckets struct {
	KeyAsString time.Time   `json:"key_as_string"`
	Key         int64       `json:"key"`
	DocCount    int         `json:"doc_count"`
	DiskAverage DiskAverage `json:"disk_average"`
}
type GraphDataForDiskSpace struct {
	Buckets []Buckets `json:"buckets"`
}
type Buckets struct {
	Key                   string                `json:"key"`
	DocCount              int                   `json:"doc_count"`
	GraphDataForDiskSpace GraphDataForDiskSpace `json:"graph_data_for_disk_space"`
}
type ByDisk struct {
	DocCountErrorUpperBound int       `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int       `json:"sum_other_doc_count"`
	Buckets                 []Buckets `json:"buckets"`
}
type MemoryAverage struct {
	Value float64 `json:"value"`
}
type Buckets struct {
	KeyAsString   time.Time     `json:"key_as_string"`
	Key           int64         `json:"key"`
	DocCount      int           `json:"doc_count"`
	MemoryAverage MemoryAverage `json:"memory_average"`
}
type GraphDataForMemory struct {
	Buckets []Buckets `json:"buckets"`
}
type Buckets struct {
	Key                string             `json:"key"`
	DocCount           int                `json:"doc_count"`
	GraphDataForMemory GraphDataForMemory `json:"graph_data_for_memory"`
}
type ByMemory struct {
	DocCountErrorUpperBound int       `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int       `json:"sum_other_doc_count"`
	Buckets                 []Buckets `json:"buckets"`
}
type GetNestedStorageData struct {
	DocCount int      `json:"doc_count"`
	ByDisk   ByDisk   `json:"by_disk"`
	ByMemory ByMemory `json:"by_memory"`
}
type Aggregations struct {
	GetNestedStorage       GetNestedStorage       `json:"get_nested_storage"`
	GetCPUTolalLoadAverage GetCPUTolalLoadAverage `json:"get_cpu_tolal_load_average"`
	GraphDataForCPU        GraphDataForCPU        `json:"graph_data_for_cpu"`
	GetNestedStorageData   GetNestedStorageData   `json:"get_nested_storage_data"`
}

func main()  {
	res := ` "aggregations" : {
    "get_nested_storage" : {
      "doc_count" : 10,
      "by_disk" : {
        "doc_count_error_upper_bound" : 0,
        "sum_other_doc_count" : 0,
        "buckets" : [
          {
            "key" : "system disk",
            "doc_count" : 5,
            "diskSize" : {
              "value" : 131072.0
            },
            "disk" : {
              "value" : 20552.0
            }
          }
        ]
      },
      "by_memory" : {
        "doc_count_error_upper_bound" : 0,
        "sum_other_doc_count" : 0,
        "buckets" : [
          {
            "key" : "main memory",
            "doc_count" : 5,
            "MemorySize" : {
              "value" : 131072.0
            },
            "memory" : {
              "value" : 28212.0
            }
          }
        ]
      }
    },
    "get_cpu_tolal_load_average" : {
      "value" : 1.0
    },
    "graph_data_for_cpu" : {
      "buckets" : [
        {
          "key_as_string" : "2020-5-7 06:00",
          "key" : 1588831200000,
          "doc_count" : 1,
          "get_cpu_average" : {
            "value" : 1.0
          }
        },
        {
          "key_as_string" : "2020-5-7 06:01",
          "key" : 1588831260000,
          "doc_count" : 1,
          "get_cpu_average" : {
            "value" : 1.0
          }
        },
        {
          "key_as_string" : "2020-5-7 06:02",
          "key" : 1588831320000,
          "doc_count" : 1,
          "get_cpu_average" : {
            "value" : 1.0
          }
        },
        {
          "key_as_string" : "2020-5-7 06:03",
          "key" : 1588831380000,
          "doc_count" : 1,
          "get_cpu_average" : {
            "value" : 1.0
          }
        },
        {
          "key_as_string" : "2020-5-7 06:04",
          "key" : 1588831440000,
          "doc_count" : 1,
          "get_cpu_average" : {
            "value" : 1.0
          }
        }
      ]
    },
    "get_nested_storage_data" : {
      "doc_count" : 10,
      "by_disk" : {
        "doc_count_error_upper_bound" : 0,
        "sum_other_doc_count" : 0,
        "buckets" : [
          {
            "key" : "system disk",
            "doc_count" : 5,
            "graph_data_for_disk_space" : {
              "buckets" : [
                {
                  "key_as_string" : "2020-05-07T06:00:00.000Z",
                  "key" : 1588831200000,
                  "doc_count" : 1,
                  "disk_average" : {
                    "value" : 20552.0
                  }
                },
                {
                  "key_as_string" : "2020-05-07T06:01:00.000Z",
                  "key" : 1588831260000,
                  "doc_count" : 1,
                  "disk_average" : {
                    "value" : 20552.0
                  }
                },
                {
                  "key_as_string" : "2020-05-07T06:02:00.000Z",
                  "key" : 1588831320000,
                  "doc_count" : 1,
                  "disk_average" : {
                    "value" : 20552.0
                  }
                },
                {
                  "key_as_string" : "2020-05-07T06:03:00.000Z",
                  "key" : 1588831380000,
                  "doc_count" : 1,
                  "disk_average" : {
                    "value" : 20552.0
                  }
                },
                {
                  "key_as_string" : "2020-05-07T06:04:00.000Z",
                  "key" : 1588831440000,
                  "doc_count" : 1,
                  "disk_average" : {
                    "value" : 20552.0
                  }
                }
              ]
            }
          }
        ]
      },
      "by_memory" : {
        "doc_count_error_upper_bound" : 0,
        "sum_other_doc_count" : 0,
        "buckets" : [
          {
            "key" : "main memory",
            "doc_count" : 5,
            "graph_data_for_memory" : {
              "buckets" : [
                {
                  "key_as_string" : "2020-05-07T06:00:00.000Z",
                  "key" : 1588831200000,
                  "doc_count" : 1,
                  "memory_average" : {
                    "value" : 28092.0
                  }
                },
                {
                  "key_as_string" : "2020-05-07T06:01:00.000Z",
                  "key" : 1588831260000,
                  "doc_count" : 1,
                  "memory_average" : {
                    "value" : 28116.0
                  }
                },
                {
                  "key_as_string" : "2020-05-07T06:02:00.000Z",
                  "key" : 1588831320000,
                  "doc_count" : 1,
                  "memory_average" : {
                    "value" : 28440.0
                  }
                },
                {
                  "key_as_string" : "2020-05-07T06:03:00.000Z",
                  "key" : 1588831380000,
                  "doc_count" : 1,
                  "memory_average" : {
                    "value" : 28192.0
                  }
                },
                {
                  "key_as_string" : "2020-05-07T06:04:00.000Z",
                  "key" : 1588831440000,
                  "doc_count" : 1,
                  "memory_average" : {
                    "value" : 28220.0
                  }
                }
              ]
            }
          }
        ]
      }
    }
  }`
	var snmpResponse SnmpDeviceUtilization
	err := json.NewDecoder(strings.NewReader(res)).Decode(&snmpResponse)
	if err != nil {
		fmt.Println("Some errror while decoding:: ", err.Error())
	}
	fmt.Println(snmpResponse)
}


