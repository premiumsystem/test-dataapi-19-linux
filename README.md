![Go](https://github.com/premiumsystem/test-dataapi-19-linux/workflows/Go/badge.svg)
# test-dataapi-19-linux
Test of Data API on FMS 19

## How to install and use
Download the program
```
wget https://github.com/premiumsystem/test-dataapi-19-linux/releases/download/1.0.1/test-dataapi-19-linux_v1.0.1_linux_amd64.tar.gz
```
and if you like to verify it after download
```
wget https://github.com/premiumsystem/test-dataapi-19-linux/releases/download/1.0.1/test-dataapi-19-linux_v1.0.1_linux_amd64.sha512
```
to verify the download after, run
```
sha512sum -c test-dataapi-19-linux_v1.0.1_linux_amd64.sha512
```

To extract the program, run
```
tar -zxpf test-dataapi-19-linux_v1.0.1_linux_amd64.tar.gz
```

Make sure it can run with
```
chmod +x test-dataapi-19-linux
```

Now create a settings.json file in the same folder as the program with the settings that we need, it should contain something like this. If you are feeling lazy just run the program once and it will on first run create the settings.json file and you can after edit it.
```json
{
	"host":"https://THE_FILEMAKER_HOST",
	"filename":"FILEMAKE_FILENAME",
	"layout": "LAYOUT_TO_GET_DATA_FROM",
	"user":"DATA_API_USERNAME",
	"pass":"DATA_API_PASSWORD",
	"show_done": false,
	"no_of_request":2500,
	"no_of_conccurent":50
}
```
Where the `no_of_request` is the total number of requests that will be executed, they will all each login to get the data from the layout and then logout.

The `no_of_conccurent` are how many request that will run concurrent while doing all the requests.

The `show_done` if true will display when a worker are done with a single request. If we get an error you will see the output anyway.

Now that we are all set up, run the program
```
./test-dataapi-19-linux
```


