![image](https://hub.steampipe.io/images/plugins/turbot/alicloud-social-graphic.png)

# Alibaba Cloud Plugin for Steampipe

Use SQL to query infrastructure including servers, networks, identity and more from Alibaba Cloud.

* **[Get started →](https://hub.steampipe.io/plugins/turbot/alicloud)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/alicloud/tables)
* Community: [Slack Channel](https://steampipe.io/community/join)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-alicloud/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):
```shell
steampipe plugin install alicloud
```

Run a query:
```sql
select display_name, create_date from alicloud_ram_user;
```

## Developing

Prerequisites:
- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-alicloud.git
cd steampipe-plugin-alicloud
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:
```
make
```

Configure the plugin:
```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/alicloud.spc
```

Try it!
```
steampipe query
> .inspect alicloud
```

Further reading:
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-alicloud/blob/main/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Alibaba Cloud Plugin](https://github.com/turbot/steampipe-plugin-alicloud/labels/help%20wanted)
