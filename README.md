# aws-ipranges

By default all results are in json format.

# Services
```
AMAZON
AMAZON_APPFLOW
AMAZON_CONNECT
API_GATEWAY
CHIME_MEETINGS
CHIME_VOICECONNECTOR
CLOUD9
CLOUDFRONT
CODEBUILD
DYNAMODB
EC2
EC2_INSTANCE_CONNECT
GLOBALACCELERATOR
KINESIS_VIDEO_STREAMS
ROUTE53
ROUTE53_HEALTHCHECKS
ROUTE53_HEALTHCHECKS_PUBLISHING
S3
WORKSPACES_GATEWAYS
```

# ex1
```
$ ./aws-ipranges -region sa-east-1 -service API_GATEWAY
```
```
[{"ip_prefix":"15.228.72.64/26","region":"sa-east-1","service":"API_GATEWAY"},{"ip_prefix":"15.228.97.0/24","region":"sa-east-1","service":"API_GATEWAY"},{"ip_prefix":"18.229.100.0/26","region":"sa-east-1","service":"API_GATEWAY"},{"ip_prefix":"18.229.99.0/24","region":"sa-east-1","service":"API_GATEWAY"},{"ip_prefix":"18.230.54.0/23","region":"sa-east-1","service":"API_GATEWAY"}]
```

# ex2
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
