if ! command -v conventional-changelog &> /dev/null
then
  echo "Please install conventional-changelog-cli [https://www.npmjs.com/package/conventional-changelog-cli]"
  exit 1
fi

conventional-changelog -p angular -i CHANGELOG.md -s -r 0