## v0.4.0 (July 12, 2023)

FEATURES:


* autodiscovery: Add support for PowerShell scripts / Windows agents ([#22](https://github.com/bastionzero/bastionzero-sdk-go/issues/22)).


* connections/rdp-sqlserver: Add support for RDP and SQL Server connections ([#23](https://github.com/bastionzero/bastionzero-sdk-go/issues/23)).


## v0.3.0 (June 07, 2023)

FEATURES:


* apikeys: Add support for api-keys API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* githubactions: Add support for github-actions API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* autodiscoveryscripts: Add support for getting ansible playbook ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* autodiscoveryscripts: Add support for getting container bash autodiscovery script ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* connections: Add support for connections API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* events: Add support for events API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* mfa: Add support for MFA API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* oktapublickeys: Add support for okta-public-keys API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* organization: Add support for GET organization ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* organization: Add support for GET BZCert validation info ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* organization: Add support for fetch groups API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* organization: Add support for DELETE IdP group credentials ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* organization: Add support for invalidating keycloak provider cache API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* organization: Add support for registration key settings and global registration key API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* organization: Add support for GET Slack integration details ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* organization: Add support for GET identity provider details ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* serviceaccounts: Add support for creating service account API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* serviceaccounts: Add support for PATCH service account ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* serviceaccounts: Add support for fetching service account info of current subject API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* serviceaccounts: Add support for invalidating Jwks URL cache API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* sessionrecordings: Add support for session-recordings API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* subjects: Add support for subjects API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* targets/bzero: Add support for DELETE Bzero target ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* targets/bzero: Add support for restart Bzero agent API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* targets/bzero: Add support for requesting Bzero agent logs API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* targets/database: Add support for listing DB targets with optional filter API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* targets/database: Add support for listing database types with SplitCert support API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* targets/kube: Add support for generate Kube agent YAML API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* targets/kube: Add support for DELETE Cluster target ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* targets/web: Add support for creating Web target API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* targets/web: Add support for DELETE Web target ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* users: Add support for me API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* users: Add support for DELETE user ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* users: Add support for PATCH user ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* users: Add support for listing users API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


* users: Add support for closing all connections for specific user API ([#20](https://github.com/bastionzero/bastionzero-sdk-go/issues/20)).


## v0.2.0 (May 12, 2023)

FEATURES:


* targets: Add support for POST (Create) and DELETE Database target ([#16](https://github.com/bastionzero/bastionzero-sdk-go/issues/16)).


## v0.1.0 (March 28, 2023)

FEATURES:


* autodiscoveryscripts/bzero: Add support for GET Bzero bash autodiscovery script ([#1](https://github.com/bastionzero/bastionzero-sdk-go/issues/1)).


* organization: Add support for GET list of IdP groups ([#1](https://github.com/bastionzero/bastionzero-sdk-go/issues/1)).


* serviceaccounts: Add support for GET service account by ID or list ([#1](https://github.com/bastionzero/bastionzero-sdk-go/issues/1)).


* users: Add support for GET user by ID/email or list ([#1](https://github.com/bastionzero/bastionzero-sdk-go/issues/1)).


* environments: Add support for all CRUD environment endpoints ([#1](https://github.com/bastionzero/bastionzero-sdk-go/issues/1)).


* policies: Add support for all CRUD policy endpoints, excluding the OrganizationControls policy type ([#1](https://github.com/bastionzero/bastionzero-sdk-go/issues/1)).


* targets/dynamicaccess: Add support for all CRUD dynamic access configuration (DAC) endpoints ([#1](https://github.com/bastionzero/bastionzero-sdk-go/issues/1)).


* targets: Add support for GET target by ID or list for all remaining target types ([#1](https://github.com/bastionzero/bastionzero-sdk-go/issues/1)).


* targets: Add support for PATCH target by ID for all remaining target types ([#1](https://github.com/bastionzero/bastionzero-sdk-go/issues/1)).


