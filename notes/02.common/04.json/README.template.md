{
"CreatedAt": "2023-12-31",
"UpdatedAt": "2023-12-31",
"Title": "JSON - Go notes"
}
---
# JSON

## Encoding
<<encoding.go>>
Output:
```JSON
{
 "IsRobot": false,
 "name": "John Doe",
 "age": 42,
 "weight": 75.5,
 "nicknames": [
  "Johnny",
  "J"
 ],
 "jobs": [
  {
   "company": "ACME",
   "position": "CEO",
   "started": "2015-01-01T00:00:00Z",
   "ended": null
  },
  {
   "company": "ACME",
   "position": "CTO",
   "started": "2011-01-01T00:00:00Z",
   "ended": "2015-01-01T00:00:00Z",
   "comments": "I was promoted"
  }
 ],
 "attributes": {
  "eyes": "blue",
  "hair": "brown"
 }
}
```

### OmitEmpty
The omitempty option in Go's JSON encoding works by omitting the field from the JSON output if the field has an empty value. An "empty" value means the zero value for the field's type, which includes:

- false for booleans,
 -0 for numeric types,
- "" (the empty string) for strings,
- nil for pointers, interfaces, maps, slices, and channels,
- the zero value for arrays, structs, and any other types.

However, it's important to note that for structs, omitempty will not omit the field if the struct is empty but not nil. This is because an empty struct still has a valid, non-nil zero value. If you want to omit a struct, the best solution would be to add it as a pointer in the struct.

## Decoding
As long as we know how the JSON data is structured, we can decode it into a struct. And we can use the parts of the data we need.

The struct must have exported fields, otherwise the decoder will not be able to set them.

The struct can have more fields than the JSON data, but the decoder will only set the fields that are present in the JSON data.

The struct must tag the fields with the name of the JSON field, otherwise the decoder will not be able to set them.
<<decoding.go>>
Output:
```bash
{Coord:{Lon:10.99 Lat:44.34} Weather:[{ID:501 Main:Rain Description:moderate rain Icon:10d}]}
```

## MarshalJSON interface
## UnmarshalJSON interface