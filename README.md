# Rubrik FileRetriever
Retrive files from Rubrik to container.

# Requirements
My Test Environment  

- golang = 1.14.1  
- Rubrik = 5.1.2  
- kubernetes = 1.17  

# How to use
You can use this as init container in k8s.  
[Dockerhub](https://hub.docker.com/r/fideltak/rubrik-file-retriever)  
In k8s manifest, describe enviroment values to retrive file from rubrik like below.  

```
spec:
  initContainers:
  - name: file-retriever01
    image: fideltak/rubrik-file-retriever:0.1
    imagePullPolicy: Always
    env:
      - name: rubrik_cdm_node_ip
        value: 172.16.14.220
      - name: rubrik_cdm_username
        value: tak
      - name: rubrik_cdm_password
        value: P@ssw0rd
      - name: RUBRICK_TARGET_HOST
        value: rubrik-sql01.hybrid-lab.local
      - name: RUBRICK_TARGET_FILESET
        value: sdk01
      - name: RUBRICK_TARGET_HOSTOS
        value: Linux
      - name: RUBRICK_TARGET_DATE
        value: 04-20-2020 10:52 AM
      - name: RUBRICK_TARGET_FILEPATH
        value: /rubrik/bu-tar01/test01.txt
      - name: DOWNLOAD_FILE_PATH
        value: /tmp/test01.txt
    volumeMounts:
      - mountPath: /tmp
        name: tmp
  containers:
  - name: target
    image: alpine:3.9
    command:
      - "cat"
      - "/tmp/test01.txt"
    volumeMounts:
      - mountPath: /tmp
        name: tmp
  volumes:
    - emptyDir: {}
      name: tmp
``` 

# Environment Value Definitions  
|Key|Value Example|Require|Description|
|:---|:---|:---|:---|
|rubrik\_cdm\_node\_ip|172.16.14.220|Yes|IP Address or Hostname of Rubrik cluster.|
|rubrik\_cdm\_username|tak|Yes|Username of Rubrik cluster.|
|rubrik\_cdm\_password|P@ssw0rd|Yes|Password of Rubrik cluster.|
|RUBRICK\_TARGET\_HOST|mysql01|Yes|Registered hostname which you want to restore files in Rubrik.|
|RUBRICK\_TARGET\_FILESET|fileset01|Yes|Fileset name.|
|RUBRICK\_TARGET\_HOSTOS|Linux|Yes|Registered host os type. Linux or Windows|
|RUBRICK\_TARGET\_DATE|04-20-2020 10:52 AM|Yes|Datetime of buckup you want to restore.|
|RUBRICK\_TARGET\_FILEPATH|/rubrik/test01.txt|Yes|File path od backup.|
|DOWNLOAD\_FILE\_PATH|/tmp/test01.txt|Yes|Download path in container.|
|POST\_SCRIPT\_PATH|/opt/mypost.sh|Option|Post shell script path. If you wanto to execure shell script after download files from Rubrik, you have copy scripts into container before that.|

# Use Post-script  
If you want use postscript, you have to copy it and build your docker image based on this image.  

```
# Dockerfile
FROM fideltak/rubrik-file-retriever

COPY mypost.sh /tmp/
``` 