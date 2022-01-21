---
subcategory: ""
layout: "fortios"
page_title: "Debugging Provider"
description: |-
  "Debugging Provider"
---

## Collecting debug information

This information will be crucial to quickly resolving any issues.

SSH to your FortiGate. Close any open GUI windows as this will generate spam during the debug. Run the following commands:

    ```
    diagnose debug enable
    diagnose debug application httpsd -1
    diagnose debug cli 8
    ```

Run terraform apply with debugging enabled:

    ```sh
    TF_LOG=debug terraform apply
    ```

Raise an issue on [GitHub](https://github.com/poroping/terraform-provider-fortios/issues)