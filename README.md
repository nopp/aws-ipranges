# aws-ipranges

By default all results are in json format.

# Ex1
```
$ ./aws-ipranges -region sa-east-1 -service API_GATEWAY
```
```
[{"ip_prefix":"15.228.72.64/26","region":"sa-east-1","service":"API_GATEWAY"},{"ip_prefix":"15.228.97.0/24","region":"sa-east-1","service":"API_GATEWAY"},{"ip_prefix":"18.229.100.0/26","region":"sa-east-1","service":"API_GATEWAY"},{"ip_prefix":"18.229.99.0/24","region":"sa-east-1","service":"API_GATEWAY"},{"ip_prefix":"18.230.54.0/23","region":"sa-east-1","service":"API_GATEWAY"}]
```

# Ex2
```
$ ./aws-ipranges -region sa-east-1 -service API_GATEWAY -json=false
```
```
sa-east-1 API_GATEWAY 15.228.72.64/26
sa-east-1 API_GATEWAY 15.228.97.0/24
sa-east-1 API_GATEWAY 18.229.100.0/26
sa-east-1 API_GATEWAY 18.229.99.0/24
sa-east-1 API_GATEWAY 18.230.54.0/23
```