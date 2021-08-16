Can use codegen to generate new resources/data_sources and documentation from API Schema.

To get schema from Fortigate GET a path with URL param action=schema ex. `https://fortigate:10443/api/v2/cmdb/firewall/access-proxy?action=schema` save complete output in ./apischema/PATH_NAME.json

**Optional:** 
Save example HCL as ./examples/PATH_NAME.d.txt or ./examples/PATH_NAME.r.txt for data source and resource docs respectively.

run ```go run template.go```



TODO:
- Generate tests
- Tidy/refactor template.go
- Tidy templates
- Deal with secrets with reversible encryption (plain-text-password=1)