# jiratray
system tray that displays your jira issues


### Usage
```git clone https://github.com/Hansuuuuuuuuuu/jiratray.git
cd jiratray
go get
go build
./jiratray <args>
Usage of ./jiratray:
  -h, --host string
    	Jira Host Address (without http://)
  -i, --interval int
    	Update Interval (in seconds) (default 30)
  -p, --password string
    	Jira Password
  -x, --port int
    	Jira Host Port (default 443)
  -u, --user string
    	Jira Username
```

### Sample Commands
```./jiratray -h jira.com -x 443 -i 30 -u hansuuuuuuuuuu -p password123
./jiratray -h jira.com -u hansuuuuuuuuuu -p password123
```
