// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"io/ioutil"

	"github.com/ar4mirez/learning/todo"
	"github.com/micro/protobuf/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := grpc.Dial(":8888", grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("could not connect to backend: %v", err)
		}
		client := todo.NewTaskClient(conn)
		return list(context.Background(), client)
	},
}

func list() error {
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		return fmt.Errorf("could not read %s: %v", dbPath, err)
	}

	for {
		if len(b) == 0 {
			return nil
		} else if len(b) < 4 {
			return fmt.Errorf("remining odd %d bytes, what to do?", len(b))
		}

		var length int64
		if err := gob.NewDecoder(bytes.NewReader(b[:4])).Decode(&length); err != nil {
			return fmt.Errorf("could not decode message length: %v", err)
		}
		b = b[4:]

		var task todo.Task
		if err := proto.Unmarshal(b[:length], &task); err != nil {
			return fmt.Errorf("could not read task: %v", err)
		}
		b = b[length:]

		if task.Done {
			fmt.Printf("✅")
		} else {
			fmt.Printf("😱")
		}

		fmt.Printf(" %s\n", task.Text)
	}
}
