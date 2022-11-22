# cx-keycloak-utils

## installing go on mac

```shell
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```
```shell
brew install go
```
## programs
### add-to-innovators-group

#### install
```shell
go install github.com/hokiegeek/cx-keycloak-utils/cmd/cx-add-to-innovators-group
```

#### run
```shell
cx-add-to-innovators-group -config CONFIG_FILE.yaml -users USERS_CSV.csv
```