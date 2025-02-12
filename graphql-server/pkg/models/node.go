package models

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	graphql "github.com/99designs/gqlgen/graphql"
	gqlerrcode "github.com/99designs/gqlgen/graphql/errcode"
	servererrors "github.com/oursky/likedao/pkg/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type NodeID struct {
	EntityType string
	ID         string
}

func (nodeID NodeID) String() string {
	return fmt.Sprintf("%s_%s", nodeID.EntityType, nodeID.ID)
}

func ParseNodeID(nodeID string) (NodeID, error) {
	parts := strings.Split(nodeID, "_")
	if len(parts) != 2 {
		return NodeID{}, servererrors.ErrInvalidNodeID
	}
	return NodeID{
		EntityType: parts[0],
		ID:         strings.ToLower(parts[1]),
	}, nil
}

func ExtractObjectIDs(nodeIDs []NodeID) []string {
	objectIDs := make([]string, 0, len(nodeIDs))
	for _, nodeID := range nodeIDs {
		objectIDs = append(objectIDs, nodeID.ID)
	}
	return objectIDs
}

// MarshalNodeID is used by gqlgen to write ID type to output.
func MarshalNodeID(nodeID NodeID) graphql.Marshaler {
	empty := NodeID{}
	if nodeID == empty {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		_, err := io.WriteString(w, strconv.Quote(nodeID.String()))
		if err != nil {
			panic(err)
		}
	})
}

// UnmarshalNodeID is used by gqlgen to read ID from input.
func UnmarshalNodeID(v interface{}) (NodeID, error) {
	if tmpStr, ok := v.(string); ok {
		nid, err := ParseNodeID(tmpStr)
		if err != nil {
			gqlerr := gqlerror.Errorf(`invalid node ID "%s"`, tmpStr)
			gqlerrcode.Set(gqlerr, gqlerrcode.ValidationFailed)
			return NodeID{}, gqlerr
		}

		return nid, nil
	}
	gqlerr := gqlerror.Errorf("unable to decode Node ID")
	gqlerrcode.Set(gqlerr, gqlerrcode.ValidationFailed)
	return NodeID{}, gqlerr
}
