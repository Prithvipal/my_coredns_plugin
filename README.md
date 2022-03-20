# my_coredns_plugin

## Steps

1. Clone coredns repo
https://github.com/coredns/coredns
2. Add following line in `plugin.cfg` file of coredns repo:
```
myplugin:github.com/Prithvipal/my_coredns_plugin
```
3. Run make command in coredns repo
```
make
```
4. add `myplugin` in Corefile
```Corefile
. {
   myplugin
   log
}

```
6. Run coredns by following command in coredns repo. Note: set valid path of -conf flag. In my case it I have added relative path
```
coredns -conf ../../coredns/Corefile
```
7. run dig command to verify that myplugin is added sucessfully or not
```
dig @localhost google.com
```
When you see the coredns logs, you will see following logs which is added as fmt.Println in myplugin.
```
=====Yoooo Serving DNS
```
