# Artifact References

## Usage Example:

First artifact:

```
echo '{
    "$schema": "https://json-schema.org/draft/2019-09/schema",
    "$id": "https://example.com/schema1.json",
    "type": "object",
    "$defs": {
        "name": {
            "type": "string"
        }
    }
}' > artifact1.json
```

Second artifact:

```
echo '{
    "$schema": "https://json-schema.org/draft/2019-09/schema",
    "$id": "https://example.com/schema2.json",
    "type": "object",
    "properties": {
        "external": {
            "$ref": "schema1.json#/$defs/name"
        }
    },
    "$defs": {
        "personal": {
            "email": {
                "type": "string",
                "format": "email"
            },
            "birthday": {
                "type": "string",
                "format": "date"
            }
        }
    }
}' > artifact2.json
```

Create first artifact:

```
cat artifact1.json | ./rhoas service-registry artifact create
```

Create second artifact (Note: The ID of the first artifact will be different from the one in the following example):

```
cat artifact2.json | ./rhoas service-registry artifact create -r 'https://example.com/schema1.json#/$defs/name=:0430a24f-ae35-485b-ada8-61376f1286d7:'
```

List references of an artifact:

```
./rhoas service-registry artifact get --artifact-id 6b07bbc9-2fae-4768-929c-27f03895451d --references
```
