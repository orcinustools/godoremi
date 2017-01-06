## Golang Docker Remote API (daemon mode)
Still very simple but I hope it will continue to grow. join us and make this some code for fun.
Do not forget to give criticism and advices, clone and contribute! :)

### Run
1. Clone
    ```
    git clone git@github.com:imamdigmi/godoremi.git
    ```
2. Build app
    ```
    go build main.go
    ```
3. Install it
    ```
    sudo ./main install
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
6. Browse to [http://localhost:8080/](http://localhost:8080/)

### Available endpoint
- **GET** `/` See the information of the docker
- **GET** `/images` List all of the images
- **GET** `/containers` List all of the running containers 