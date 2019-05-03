package commands

import (
	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
)

func formatJSON(pb proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		Indent:       "  ",
		EmitDefaults: true,
	}
	return marshaler.MarshalToString(pb)
}

func Add(cmd *cobra.Command) {
	//AddDPOS(cmd) //No one should be using the old DPOS version
	AddGeneralCommands(cmd)
}
