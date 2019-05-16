package cmd

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"os"
	"strings"

	"github.com/replicatedhq/libyaml"

	"github.com/awalterschulze/gographviz"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var options = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ",
	"BA", "BB", "BC", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BK", "BL", "BM", "BN", "BO", "BP", "BQ", "BR", "BS", "BT", "BU", "BV", "BW", "BX", "BY", "BZ",
	"CA", "CB", "CC", "CD", "CE", "CF", "CG", "CH", "CI", "CJ", "CK", "CL", "CM", "CN", "CO", "CP", "CQ", "CR", "CS", "CT", "CU", "CV", "CW", "CX", "CY", "CZ"}

var NODE_LABEL = 0
var GRAPH_LABEL = 1

func intToRGB(i int) string {
	i &= 0x00FFFFFF
	hex := strings.ToUpper(fmt.Sprintf("%x", i))

	for len(hex) < 6 {
		hex = "0" + hex
	}
	return `"#` + hex + `"`
}

var EventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Goviz event orchestration graph",
	Long:  `Generate a events graph for an Replicated app YAML`,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("yaml").Value.String()
		graphName := "events"

		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Print("Error: path is not valid\n")
			os.Exit(1)
		}

		var root libyaml.RootConfig
		if err = yaml.Unmarshal(bytes, &root); err != nil {
			fmt.Printf("failed to parse app YAML: %s\n", err)
			os.Exit(1)
		}

		graph := gographviz.NewGraph()
		graph.SetName(graphName)
		graph.SetDir(true)

		//var m map[string]string
		keys := make(map[string]string)
		graphKeys := make(map[*libyaml.Component]string)

		//// Add the clusters
		//for _, component := range root.Components {
		//	clusterName := "cluster_"+ strconv.Itoa(GRAPH_LABEL)
		//	graphKeys[component] = clusterName
		//
		//	clusterAttributes := make(map[string]string)
		//	clusterAttributes["label"] = "\"" + component.Name + "\""
		//	graph.AddSubGraph(graphName, "cluster_"+ strconv.Itoa(GRAPH_LABEL), clusterAttributes)
		//	GRAPH_LABEL = GRAPH_LABEL + 1
		//}

		// Add the nodes in the each cluster
		for _, component := range root.Components {
			for _, container := range component.Containers {
				from := fmt.Sprintf("\"%s:%s\"", component.Name, container.ImageName)
				label := findKey(keys, from)
				attrs := make(map[string]string)
				attrs["label"] = fmt.Sprintf("\"%s:%s\"", component.Name, container.ImageName)

				//f := fnv.New32()
				//f.Write([]byte(attrs["label"]))
				//attrs["color"] = intToRGB(int(f.Sum32()))

				graph.AddNode(graphKeys[component], label, attrs)
			}
		}

		// Add the edges
		for _, component := range root.Components {
			for _, container := range component.Containers {
				from := fmt.Sprintf("\"%s:%s\"", component.Name, container.ImageName)
				fromLabel := findKey(keys, from)

				for _, event := range container.PublishEvents {
					for _, subscription := range event.Subscriptions {
						to := fmt.Sprintf("\"%s:%s\"", subscription.ComponentName, subscription.ContainerName)
						toLabel := findKey(keys, to)
						attributes := make(map[string]string)
						attributes["label"] = "\" " + eventText(event) + " \""
						attributes["minlen"] = "2"

						f := fnv.New32()
						f.Write([]byte(attributes["label"]))
						attributes["color"] = intToRGB(int(f.Sum32()))

						//attributes["fontname"] = "\"times:italic\""
						graph.AddEdge(fromLabel, toLabel, true, attributes)
					}
				}
			}
		}

		output := graph.String()
		fmt.Printf("%s", output)
	},
}

func eventText(event *libyaml.ContainerEvent) string {
	return event.Trigger
}

func findKey(m map[string]string, label string) string {
	if m[label] == "" {
		m[label] = options[NODE_LABEL]
		NODE_LABEL = NODE_LABEL + 1
		return m[label]
	}

	return m[label]
}

func init() {
	EventsCmd.PersistentFlags().StringP("yaml", "y", "", "path to app YAML")
}
