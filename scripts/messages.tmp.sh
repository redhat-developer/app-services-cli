
cat **/*.go | grep MustLocalize > messages.toml

## Grep
## From .*MustLocalize\(" to [
## From .*MustLocalizeError\(" to [
## From ".* to ] \n one = '' \n 