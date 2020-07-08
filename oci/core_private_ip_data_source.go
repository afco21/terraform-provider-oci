// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"
)

func init() {
	RegisterDatasource("oci_core_private_ip", CorePrivateIpDataSource())
}

func CorePrivateIpDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["private_ip_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(CorePrivateIpResource(), fieldMap, readSingularCorePrivateIp)
}

func readSingularCorePrivateIp(d *schema.ResourceData, m interface{}) error {
	sync := &CorePrivateIpDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient()

	return ReadResource(sync)
}

type CorePrivateIpDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.GetPrivateIpResponse
}

func (s *CorePrivateIpDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CorePrivateIpDataSourceCrud) Get() error {
	request := oci_core.GetPrivateIpRequest{}

	if privateIpId, ok := s.D.GetOkExists("private_ip_id"); ok {
		tmp := privateIpId.(string)
		request.PrivateIpId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.GetPrivateIp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CorePrivateIpDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HostnameLabel != nil {
		s.D.Set("hostname_label", *s.Res.HostnameLabel)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.IsPrimary != nil {
		s.D.Set("is_primary", *s.Res.IsPrimary)
	}

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VlanId != nil {
		s.D.Set("vlan_id", *s.Res.VlanId)
	}

	if s.Res.VnicId != nil {
		s.D.Set("vnic_id", *s.Res.VnicId)
	}

	return nil
}
