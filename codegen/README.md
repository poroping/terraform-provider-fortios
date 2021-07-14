Can use codegen to generate new resources/data_sources and documentation from API Schema.

To get schema from Fortigate GET a path with URL param action=schema ex. `https://fortigate:10443/api/v2/cmdb/firewall/access-proxy?action=schema` save complete output as PATH_NAME.json and ```go run template.go```

WIP

TODO:
- Generate tests
- Examples in docs
- Tidy/refactor template.go
- Tidy templates