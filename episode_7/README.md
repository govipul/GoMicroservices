# Problem faced in episode 7
Documenting the REST API with Swagger

### first which i noiced that nic has created a beautiful env package and is using that package to manage the server address. ###
So for that please follow the following steps:
1. ´´´go get github.com/nicholasjackson/env´´´
2. Add the above package in your import files
3. create a variable env with following parameters : ´´´var bindAddress = env.String("BIND_ADDRESS", false, ":9000", "Bind address for the server")´´´
4. parse the env using ´´´env.Parse()´´´
5. use the above vairibale in code
´´´
s := http.Server{
        Addr: *bindAddress,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
´´´
### Challanges with Make and swagger ###
1. Make is very specific to Linux so its very difficult to use the make file to work on windows, unless we tryout with Cygwin, which i have not tried.
2. So to use the make file work i have finally have to move to WSL2 with windows version.
3. Installing WSL2 is very simple and easy to use. Specially if we are using VS code.
4. Even though in the make file we are using the installation of swagger, but as per some of stackoverflow comments we have to install swagger manually.
´´´bash
echo "deb https://dl.bintray.com/go-swagger/goswagger-debian ubuntu main" | sudo tee /etc/apt/sources.list.d/goswagger.list
apt install swagger
´´´ and to validate the version we can use following command 
´´´root:episode_7# swagger version
version: v0.26.1
commit: (unknown, mod sum: "h1:1XUWLnH6hKxHzeKjJfA2gHkSqcT1Zgi4q/PZp2hDdN8=")
´´´
5. After this i faced a issue when running command ´´´´make swagger´´´ was generating the empty YAML file.
6. After debugging even though i was using 1 tab in the handler class for documentation, but still getting the empty swagger file. So have to 2 tabs instead and it worked :-)
