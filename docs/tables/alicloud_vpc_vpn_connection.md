# Table: alicloud_vpc_vpn_connection

An IPsec-VPN connection provides support to establish an encrypted communication tunnel between a VPN Gateway and a customer gateway.

## Examples

### Basic info

```sql
select
  name,
  vpn_connection_id,
  status,
  local_subnet,
  remote_subnet,
  vpn_gateway_id
from
  alicloud_vpc_vpn_connection;
```

### Get the vpn connections which are not healthy

```sql
select
  name,
  vpn_connection_id,
  vco_health_check ->> 'Status' as health_check_status,
  status
from
  alicloud_vpc_vpn_connection
where vco_health_check ->> 'Status' = 'failed';
```

### Get the BGP configuration information of vpn connections

```sql
select
  name,
  vpn_connection_id,
  vpn_bgp_config ->> 'EnableBgp' as enable_bgp,
  vpn_bgp_config ->> 'LocalAsn' as local_asn,
  vpn_bgp_config ->> 'LocalBgpIp' as local_bgp_ip,
  vpn_bgp_config ->> 'PeerAsn' as peer_asn,
  vpn_bgp_config ->> 'PeerBgpIp' as peer_bgp_ip,
  vpn_bgp_config ->> 'Status' as status,
  vpn_bgp_config ->> 'TunnelCidr' as tunnel_cidr
from
  alicloud_vpc_vpn_connection;
```


### Get the vpn connections where NAT traversal feature is enabled

```sql
select
  name,
  vpn_connection_id,
  enable_nat_traversal
from
  alicloud_vpc_vpn_connection
where enable_nat_traversal;
```
