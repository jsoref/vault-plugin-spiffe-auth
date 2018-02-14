# Imeplementation Design
This document will contain the initial implementation design.


## Example Trust Domains
The below tables show some example SPIFFE trust domains and how they may map to a Vault cluster and auth point, the current examples would validate based on and individual trust domains CA, there is currently no concept for heirachical trust domains and validation based on trust chain.

### Simple example where an organisation has a single trust domain per environment
| Trust Domain            | Vault Cluster | Auth Mount Point |
| ----------------------- | ------------- | ---------------- |
| spiffe://prod.acme.net/ | Production    | /v1/auth/spiffe  |
| spiffe://stg.acme.net/  | Staging       | /v1/auth/spiffe  |
| spiffe://dev.acme.net/  | Development   | /v1/auth/spiffe  |
  
  
### Example showing `spiffe://prod.acme.net/` as a global identity and a trust domain per geo location.
| Trust Domain               | Vault Cluster | Auth Mount Point   | Comments |
| -------------------------- | ------------- | ------------------ | ------------------------------------------ |
| spiffe://us.prod.acme.net/ | Production    | /v1/auth/spiffe/us | Vault premium with performance replication | 
| spiffe://eu.prod.acme.net/ | Production    | /v1/auth/spiffe/eu | "" |
| spiffe://ap.prod.acme.net/ | Production    | /v1/auth/spiffe/ap | "" |


### Trust domain per department or application boundary
In the below example, the two trust domains `insurance` and `consumer` would most probably share the same cluster in an enterprise.  However, Support may or may not use it's own cluster, ideally support would require access to secrets such a Database Users, AWS credentials, etc, therefore it would make sense to allow access to the main Vault cluster instead of having to replicate and maintain secrets in two clusters.  Policy in Vault would allow for privelidge to be restricted to the right levels ensuring any sensitive infomation which support are not allowed to access remains secret.  The organisation would most likely leverage Vault premium's capability to run in more than one datacenter. If the organisation was particularly security adverse then they may use their own infra for support application secrets and allow support personnele to auth the production Vault cluster to obtain the secrets required to solve problems.

| Trust Domain                                  | Vault Cluster  | Auth Mount Point         | Comments  |
| --------------------------------------------- | -------------- | ------------------------ | --------- |
| spiffe://insurance.bigbank.com/               | Production     | v1/auth/spiffe/insurance |           |
| spiffe://consumer.bigbank.com/                | Production     | v1/auth/spiffe/consumer  |           |
| spiffe://support-site.prod.consumer.acme.net/ | Production?    | v1/auth/spiffe/support   | support-site has dedicated infra |