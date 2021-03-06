## sentrytool group list

list groups

### Synopsis


List groups specified in the command line. If no groups are specified, list all groups.
Without -v flag, just prints group names. If '-v' flag is provided, provided group to roles mapping.


```
sentrytool group list
```

### Examples

```

  sentrytool group list -v
  admin_group = admin
  finance_group = admin, customer
  user_group = customer
```

### Options inherited from parent commands

```
  -C, --component string   sentry client component
      --config string      config file (default is $HOME/.sentrytool.yaml)
  -H, --host string        hostname for Sentry server (default "localhost")
  -P, --port string        port for Sentry server (default "8038")
  -r, --role string        role name
  -U, --username string    user name (default "akolb")
  -v, --verbose            verbose mode
```

### SEE ALSO
* [sentrytool group](sentrytool_group.md)	 - list, add or remove groups

###### Auto generated by spf13/cobra on 7-Dec-2016
