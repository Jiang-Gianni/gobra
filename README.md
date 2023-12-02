# gobra
Learning Cobra-CLI tool and Viper

This uses `GOBRA_HOSTS_FILE`

```bash
GOBRA_HOSTS_FILE=newFile.hosts ./gobra hosts add host01 host02
```


This uses the value of `hosts-file` set in `config.yaml`

```bash
./gobra hosts list --config config.yaml
```


This uses `GOBRA_HOSTS_FILE`

```bash
GOBRA_HOSTS_FILE=newFile.hosts ./gobra hosts list --config config.yaml
```



This uses the `hosts-file` flag

```bash
GOBRA_HOSTS_FILE=newFile.hosts ./gobra hosts list --config config.yaml --hosts-file=pScan.hosts
```