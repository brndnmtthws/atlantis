# Atlantis

[Atlantis](https://www.runatlantis.io/) is a tool for safe collaboration on [Terraform](https://www.terraform.io/) repositories.

## Introduction

This chart creates a single pod in a StatefulSet running Atlantis.  Atlantis persists Terraform [plan files](https://www.terraform.io/docs/commands/plan.html) and [lock files](https://www.terraform.io/docs/state/locking.html) to disk for the duration of a Pull/Merge Request.  These files are stored in a PersistentVolumeClaim to survive Pod failures.

## Prerequisites

-   Kubernetes 1.9+
-   PersistentVolume support

## Required Configuration

In order for Atlantis to start and run successfully, all of the following must be true:

1.  At least one of the following sets of credentials must be defined:

    -   `github`
    -   `gitlab`
    -   `bitbucket`

    Refer to [values.yaml](values.yaml) for detailed examples.

2.  Supply a value for `orgWhitelist`, e.g. `github.org/my_company/*`.

## Customization

The following options are supported.  See [values.yaml](values.yaml) for more detailed documentation and examples:

| Parameter                                   | Description                                                                                                                                                                                                                                                                                                                                                | Default                |
| ------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------- |
| `atlantis_data_storage`                     | The amount of storage available for Atlantis' data directory (mostly used to check out git repositories).                                                                                                                                                                                                                                                  | `5Gi`                  |
| `orgWhiteList`                              | A whitelist of repositories from which Atlantis will accept webhooks. **This value must be changed for Atlantis to function correctly.**  Accepts wildcard characters (`*`).  Multiple values may be comma-separated.                                                                                                                                      | `github.com/yourorg/*` |
| `github.user`                               | The name of the Atlantis GitHub user.,This value does not be defined if Atlantis is not working against GitHub repositories.                                                                                                                                                                                                                               | n/a                    |
| `github.token`                              | The personal access token for the Atlantis GitHub user.,This value should not be defined if Atlantis is not integrated with GitHub repositories.                                                                                                                                                                                                           | n/a                    |
| `github.secret`                             | The repository or organization-wide secret for the Atlantis GitHub integration.,All repositories in GitHub that are to be integrated with Atlantis must share the same value.,For this reason, the Atlantis maintainers recommend an organization-scoped webhook.,This value should not be defined if Atlantis is not integrated with GitHub repositories. | n/a                    |
| `github.hostname`                           | GitHub Enterprise only: The hostname of your GitHub Enterprise installation.                                                                                                                                                                                                                                                                               | n/a                    |
| `gitlab.user`                               | The repository or organization-wide secret for the Atlantis GitLab,integration.,All repositories in GitHub that are to be integrated with,Atlantis must share the same value.,For this reason, the Atlantis,maintainers recommend an organization-scoped webhook.,This value should,not be defined if Atlantis is not integrated with GitLab repositories. | n/a                    |
| `gitlab.token`                              | The personal access token for the Atlantis GitHub user.,This value should not be defined if Atlantis is not integrated with GitHub repositories.                                                                                                                                                                                                           | n/a                    |
| `gitlab.secret`                             | The repository secret for the Atlantis GitLab integration.,All repositories in GitLab that are to be integrated with,Atlantis must share the same value.,(Unlike GitHub, GitLab does not support organization-wide integrations.)                                                                                                                          | n/a                    |
| `gitlab.hostname`                           | GitLab Enterprise only: The hostname of your GitLab Enterprise installation.                                                                                                                                                                                                                                                                               | n/a                    |
| `bitbucket.user`                            | The name of the Atlantis Bitbucket user.,This value does not be defined if Atlantis is not working against Bitbucket repositories.                                                                                                                                                                                                                         | n/a                    |
| `bitbucket.token`                           | The personal access token for the Atlantis Bitbucket user.,This value should not be defined if Atlantis is not integrated with Bitbucket repositories.                                                                                                                                                                                                     | n/a                    |
| `bitbucket.secret`                          | Bitbucket Server only: The webhook secret for Bitbucket repositories.                                                                                                                                                                                                                                                                                      | n/a                    |
| `bitbucket.base_url`                        | Bitbucket Server only: The hostname of your Bitbucket Server installation.                                                                                                                                                                                                                                                                                 | n/a                    |
| `gitconfig`                                 | Contents of a file to be mounted to `~atlantis/.gitconfig`.  Use to allow redirection for Terraform modules in private git repositories.                                                                                                                                                                                                                   | n/a                    |
| `aws.credentials`                           | Contents of a file to be mounted to `~atlantis/.aws/credentials`.                                                                                                                                                                                                                                                                                          | n/a                    |
| `aws.config`                                | Contents of a file to be mounted to `~atlantis/.aws/config`.                                                                                                                                                                                                                                                                                               | n/a                    |
| `serviceAccountSecrets.credentials`         | JSON object representing secrets for a Google Cloud Platform production service account.  Only applicable if hosting Atlantis on GKE.                                                                                                                                                                                                                      | n/a                    |
| `serviceAccountSecrets.credentials-staging` | JSON object representing secrets for a Google Cloud Platform staging,service account.,Only applicable if hosting Atlantis on GKE.                                                                                                                                                                                                                          | n/a                    |

## Testing the Deployment

To perform a smoke test of the deployment (i.e. ensure that the Atlantis UI is up and running):

1.  Install the chart.  Supply your own values file or use `test-values.yaml`, which has a minimal set of values required in order for Atlantis to start.

    ```bash
    helm install -f test-values.yaml --name my-atlantis . --debug
    ```

2.  Run the tests:
    ```bash
    helm test my-atlantis
    ```
