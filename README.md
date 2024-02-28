### Create
#### localhost:3000/business (method:POST)
```sh
{
    "Location" : "jakarta",
    "Latitude" : 12345,
    "Longitude" : 54321,
    "Term" : "testterm",
    "radius" : 333,
    "categories" : "testcat",
    "locale" : "testlocale",
    "price" : 200000,
    "open_now" : true,
    "open_at" : 5431123,
    "attributes" : "afasvda",
    "sort_by" : "asvdav",
    "device_platform" : "asvadva",
    "reservation_date" : "2024-09-08",
    "reservation_time" : "20:04",
    "reservation_covers" : 2,
    "matches_party_size_param" : true,
    "limit" : 20,
    "offset" : 1
}
```

### Update
#### localhost:3000/business/id (method:PUT)

```sh
{
    "Location" : "tangerang",
    "Latitude" : 12349,
    "Longitude" : 54329,
    "Term" : "testterm2",
    "radius" : 339,
    "categories" : "cattest",
    "locale" : "localtest",
    "price" : 200009,
    "open_now" : false,
    "open_at" : 5431129,
    "attributes" : "afasvd9",
    "sort_by" : "asvda9",
    "device_platform" : "asvadv9",
    "reservation_date" : "2024-09-01",
    "reservation_time" : "20:11",
    "reservation_covers" : 1,
    "matches_party_size_param" : false,
    "limit" : 29,
    "offset" : 2
}
```

### Delete
#### localhost:3000/business/id (method:DEL)

### Search
#### localhost:3000/business/search/test/jakarta/123/321/333 (method:GET)

#### Another Example of Search without Location Parameter
#### localhost:3000/business/search/test//123/321/333 (method:GET)

### ** For delete and search no request body needed, will return response 200 if succeed