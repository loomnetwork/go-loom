package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/loomnetwork/go-loom/builtin/types/address_mapper"
	"github.com/loomnetwork/go-loom/cli"
)

const AddressMapperContractName = "addressmapper"

func GetMapping() *cobra.Command {
	return &cobra.Command{
		Use:   "get-mapping",
		Short: "Get mapping address",
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp address_mapper.AddressMapperGetMappingResponse
			from, err := cli.ParseAddress(args[0])
			if err != nil {
				return err
			}
			err = cli.StaticCallContract(AddressMapperContractName, "GetMapping", &address_mapper.AddressMapperGetMappingRequest{
				From: from.MarshalPB(),
			}, &resp)
			if err != nil {
				return err
			}
			out, err := formatJSON(&resp)
			if err != nil {
				return err
			}
			fmt.Println(out)
			return nil
		},
	}
}

func ListMapping() *cobra.Command {
	return &cobra.Command{
		Use:   "list-mapping",
		Short: "List mapping address",
		RunE: func(cmd *cobra.Command, args []string) error {
			var resp address_mapper.AddressMapperListMappingResponse
			err := cli.StaticCallContract(AddressMapperContractName, "ListMapping", &address_mapper.AddressMapperListMappingRequest{}, &resp)
			if err != nil {
				return err
			}
			out, err := formatJSON(&resp)
			if err != nil {
				return err
			}
			fmt.Println(out)
			return nil
		},
	}
}

func AddAddressMapper(root *cobra.Command) {
	root.AddCommand(
		GetMapping(),
		ListMapping(),
	)
}
