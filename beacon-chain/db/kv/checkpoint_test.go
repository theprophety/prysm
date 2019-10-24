package kv

import (
	"context"
	"testing"

	"github.com/gogo/protobuf/proto"
	ethpb "github.com/prysmaticlabs/prysm/proto/eth/v1alpha1"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	"github.com/prysmaticlabs/prysm/shared/bytesutil"
)

func TestStore_JustifiedCheckpoint_CanSaveRetrieve(t *testing.T) {
	db := setupDB(t)
	defer teardownDB(t, db)
	ctx := context.Background()
	cp := &ethpb.Checkpoint{
		Epoch: 10,
		Root:  []byte{'A'},
	}

	if err := db.SaveJustifiedCheckpoint(ctx, cp); err != nil {
		t.Fatal(err)
	}

	retrieved, err := db.JustifiedCheckpoint(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(cp, retrieved) {
		t.Errorf("Wanted %v, received %v", cp, retrieved)
	}
}

func TestStore_FinalizedCheckpoint_CanSaveRetrieve(t *testing.T) {
	db := setupDB(t)
	defer teardownDB(t, db)
	ctx := context.Background()
	root := bytesutil.ToBytes32([]byte{'B'})
	cp := &ethpb.Checkpoint{
		Epoch: 5,
		Root:  root[:],
	}

	// a state is required to save checkpoint
	if err := db.SaveState(ctx, &pb.BeaconState{}, root); err != nil {
		t.Fatal(err)
	}

	if err := db.SaveFinalizedCheckpoint(ctx, cp); err != nil {
		t.Fatal(err)
	}

	retrieved, err := db.FinalizedCheckpoint(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(cp, retrieved) {
		t.Errorf("Wanted %v, received %v", cp, retrieved)
	}
}

func TestStore_JustifiedCheckpoint_DefaultCantBeNil(t *testing.T) {
	db := setupDB(t)
	defer teardownDB(t, db)
	ctx := context.Background()

	genesisRoot := [32]byte{'A'}
	if err := db.SaveGenesisBlockRoot(ctx, genesisRoot); err != nil {
		t.Fatal(err)
	}

	cp := &ethpb.Checkpoint{Root: genesisRoot[:]}
	retrieved, err := db.JustifiedCheckpoint(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(cp, retrieved) {
		t.Errorf("Wanted %v, received %v", cp, retrieved)
	}
}

func TestStore_FinalizedCheckpoint_DefaultCantBeNil(t *testing.T) {
	db := setupDB(t)
	defer teardownDB(t, db)
	ctx := context.Background()

	genesisRoot := [32]byte{'B'}
	if err := db.SaveGenesisBlockRoot(ctx, genesisRoot); err != nil {
		t.Fatal(err)
	}

	cp := &ethpb.Checkpoint{Root: genesisRoot[:]}
	retrieved, err := db.FinalizedCheckpoint(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if !proto.Equal(cp, retrieved) {
		t.Errorf("Wanted %v, received %v", cp, retrieved)
	}
}

func TestStore_FinalizedCheckpoint_StateMustExist(t *testing.T) {
	db := setupDB(t)
	defer teardownDB(t, db)
	ctx := context.Background()
	cp := &ethpb.Checkpoint{
		Epoch: 5,
		Root:  []byte{'B'},
	}

	if err := db.SaveFinalizedCheckpoint(ctx, cp);  err != errMissingStateForFinalizedCheckpoint {
		t.Fatalf("wanted err %v, got %v", errMissingStateForFinalizedCheckpoint, err)
	}
}
