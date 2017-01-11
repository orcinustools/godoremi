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

#### Volumes
- **GET** `/volumes` List all of the volumes
- **GET** `/volumes/:name` Inspect an volume