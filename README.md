## Golang Docker Remote API (daemon mode)
Still very simple but I hope it will continue to grow. join us and make this some code for fun.
Do not forget to give criticism and advices, clone and contribute! :)

### Build and run
1. Get
    ```
    go get -u -v github.com:orcinustools/godoremi
    ```
2. Build app
    ```
    go build
    ```
3. Install it
    ```
    make install
    ```
4. Then start it with systemd
    ```
    sudo systemctl start godoremi
    ```
    or
    ```
    sudo service godoremi status
    ```
5. And you can see status of the service with
    ```
    sudo systemctl status godoremi
    ```
6. Browse to [http://localhost:4125/](http://localhost:4125/)

### Available endpoint
#### Misc
- **GET** `/` See the system-wide information
- **GET** `/version` Show the docker version information
- **GET** `/networks` List networks
- **GET** `/networks/:id` Inspect network
- **GET** `/ping` Ping the docker server
- **GET** `/events` Monitor the Docker's events

#### Images
- **GET** `/images` List all of the images 
- **GET** `/images/search` Search images
- **GET** `/images/inspect/:name` Inspect an image
- **GET** `/images/history/:name` Display history an image

#### Containers
- **GET** `/containers` List all of the running containers
- **POST** `/containers/create` Create a container
- **POST** `/containers/start/:name` Start a container
- **POST** `/containers/restart/:name` Restart a container
- **POST** `/containers/rename/:name` Rename a container
- **POST** `/containers/stop/:name` Stop a container
- **POST** `/containers/kill/:name` Kill a container
- **GET** `/containers/inspect/:name` Inspect a container
- **GET** `/containers/top/:name` List processes running inside a container
- **GET** `/containers/changes/:name` Inspect changes on a container’s filesystem
- **GET** `/containers/stats/:name` Get container stats based on resource usage

#### Volumes
- **GET** `/volumes` List all of the volumes
- **GET** `/volumes/:name` Inspect an volume