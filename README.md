# NocoDB Reminder Emails

![quality workflow](https://github.com/MasterEvarior/nocodb-reminder-emails/actions/workflows/quality.yaml/badge.svg) ![release workflow](https://github.com/MasterEvarior/nocodb-reminder-emails/actions/workflows/publish.yaml/badge.svg)

This is just a ultra small utility software, which sends me reminders from my Kanban board in [NocoDB](https://nocodb.com/).

## Build

To build the container yourself, simply clone the repository and then build the container with the provided docker file. You can the run it as described in the section below.

```shell
docker build --tag nocodb-reminder-emails .
```

Alternatively you can build the binary directly with Go.

```shell
go build -o ./nocodb-reminder-emails
```

## Run

The easiest way to run this, is to use the provided container.

```shell
docker run -d \
  -e NDBRE_BASE_URL="https://nocodb.mydomain.com" \
  -e NDBRE_API_TOKEN="xyz" \
  -e NDBRE_EMAIL_FROM="reminders@mydomain.com" \
  -e NDBRE_SMTP_SERVER="smtp.mydomain.com:25" \
  -e NDBRE_EMAIL_TO="reminders@mydomain.com" \
  -p 8080:8080 \
  ghcr.io/masterevarior/nocodb-reminder-emails:latest
```

You should now see the UI at http://localhost:8080

### Environment Variables

| Name                  | Description                             | Example                       | Mandatory  |
|-----------------------|-----------------------------------------|-------------------------------|------------|
| NDBRE_BASE_URL        | URL to your NocoDB instance             | `https://nocodb.mydomain.com` | ✅         |
| NDBRE_API_TOKEN       | API token for your NocoDB instance      | `xyz`                         | ✅         |
| NDBRE_EMAIL_FROM      | Whiche email should send the reminder   | `reminders@mydomain.com`      | ✅         |
| NDBRE_SMTP_SERVER     | URL of your SMTP server                 | `smtp.mydomain.com:25`        | ✅         |
| NDBRE_EMAIL_TO        | Which email should receive the reminder | `reminders@mydomain.com`      | ✅         |

## Development

### Linting

Linting is done with [golangci-lint](https://golangci-lint.run/), which can be run like so:

```shell
golangci-lint run
```

Run all other linters with the treefmt command. Note that the command does not install the required formatters.

```shell
treefmt
```

### Git Hooks

There are some hooks for formatting and the like. To use those, execute the following command:

```shell
git config --local core.hooksPath .githooks/
```

### Nix

If you are using [NixOS or the Nix package manager](https://nixos.org/), there is a dev shell available for your convenience. This will install Go, everything needed for formatting, set the Git hooks and some default environment variables. Start it with this command:

```shell
nix develop
```

If you happen to use [nix-direnv](https://github.com/nix-community/nix-direnv), this is also supported.

## Improvements, issues and more

Pull requests, improvements and issues are always welcome.
