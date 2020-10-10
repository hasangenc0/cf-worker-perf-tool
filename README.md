# Cf Worker Perf Cli Tool

## Build Project
```
make build
```

## Commands

#### Help
Help command shows usage information about the cli to
```
./cf-perf --help
```

#### Url
Url command prints the response from the specified url. (should be full url with starting http or https)
```
./cf-perf --url https://github.com
```

#### Profile
Profile Command sends request specified number of times. And generates report about responses. 
```
./cf-perf --url https://github.com --profile 30
```

Version shows package version:)
```
./cf-perf --version
```

## Examples
Cli tool shows help screen in default or you can specify --help argument. 
![help](./documents/help.png?raw=true")

### Using --url
--url command used to print my worker's homepage response.
![help](./documents/workerHomePageUrl.png?raw=true")

--url command used to print my workers links json response .
![help](./documents/urlLinks.png?raw=true")

--url command used to print google's home page.
![help](./documents/url.png?raw=true")

### Using --profile
--profile command used to report github.com.
![help](./documents/profileGithub.png?raw=true")

--profile command used to report the non existed page of github.com. 
![help](./documents/profileGithubNotDefinedPage.png?raw=true")

--profile command used to report my worker's homepage.
![help](./documents/profileWorkerHomePage.png?raw=true")

--profile command used to report my worker's /links endpoint.
![help](./documents/profileWorkerLinks.png?raw=true")

--profile command used to report my personal site which is a firebase FaaS. 
![help](./documents/profilePersonalSite.png?raw=true")

--profile command used to report facebook.com. Because of the Facebook redirects the request response body not showed. (Error Codes section actually is non 200 responses.)
![help](./documents/profileFacebook.png?raw=true")

## Author
[Hasan Genc](https://github.com/hasangenc0)
