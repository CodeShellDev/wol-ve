# Getting Started
First download all of the files __(except README.md and .gitignore)__.

Then you need to install python version 3.9,

> [!IMPORTANT]
> This program was only tested with version 3.9  
> so be careful when upgrading to the newest version of python

After that you need to initialize the program,

**Run**
```
sudo chmod +x ./init.sh
bash ./init.sh
```

To start the program,

**Run**
```
sudo chmod +x ./start.sh
bash ./start.sh
```

All done.

Now you can send JSON to the program on Port 9000 via http:
```json
{
    "id": "100",
    "ip": "192.168.1.1",
    "startupTime": "10" 
}
```
