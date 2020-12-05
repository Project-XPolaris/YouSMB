# YouSMB

For samba in linux,provide a convenient interface to maintain Samba on Linux
,easy to use

## API

The program will automatically apply the new configuration and restart the SMB service.

### Get config
`/config` [GET] get samba config 

#### response
```json
{
    "sections": [
        {
            "name": "global",
            "fields": {
                "log file": "/var/log/samba/log.%m",
                "logging": "file",
                ...
            }
        },
        {
            "name": "sambashare",
            "fields": {
                "browsable": "yes",
                ...
            }
        },
    ]
}
```
### Add Folder
`/folders/add` [POST] add share folder
#### body
```json
{
    "name":"share",
    "properties":{
        "available": "yes",
        ...
    }
}
```
### Update Folder Config
`/folders/update` [POST] update share folder

#### body
```json
{
    "name":"share",
    "properties":{
        "available": "yes",
        ...
    }
}
```

### Delete Folder
`/folders/delete?name={folder_name}` [DELETE] remove share folder

note: only remove folder on config,not delete real folder

## Config

The configuration file needs to be placed with the program, named `config.json`

example (NAS in ubuntu 20.04)
```json
{
  "addr": ":8200",
  "smb_config_path": "/etc/samba/smb.conf",
  "start_script": "sudo service smbd start",
  "restart_script": "sudo service smbd restart",
  "stop_script": "sudo service smbd stop"
}
```
`addr` : service address

`smb_config_path`: smb config file position

`start_script`: start smb service command

`restart_script`: restart smb service command

`stop_script`: stop smb service command

TODO:
- [ ] Web GUI