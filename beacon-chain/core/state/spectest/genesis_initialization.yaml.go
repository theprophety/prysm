// Code generated by yaml_to_go. DO NOT EDIT.
// source: genesis_initialization_minimal.yaml

package spectest

import pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"

type GenesisInitializationTest struct {
	Title         string   `yaml:"title"`
	Summary       string   `yaml:"summary"`
	ForksTimeline string   `yaml:"forks_timeline"`
	Forks         []string `yaml:"forks"`
	Config        string   `yaml:"config"`
	Runner        string   `yaml:"runner"`
	Handler       string   `yaml:"handler"`
	TestCases     []struct {
		Description   string          `yaml:"description"`
		Eth1BlockHash []byte          `yaml:"eth1_block_hash"`
		Eth1Timestamp uint64          `yaml:"eth1_timestamp"`
		Deposits      []*pb.Deposit   `yaml:"deposits"`
		State         *pb.BeaconState `yaml:"state"`
	} `yaml:"test_cases"`
}