# getip

Returns a single ip address that can be used for other CLI applications

Usage: `getip i-00d5339e6b098faad` or `getip EC2InstanceName`

If you have more than one EC2 instance with the same name you can pass a second numeric input, starting at zero. Eg. `getip JenkinsWorker 0`