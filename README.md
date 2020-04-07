# My Go Web
Golang web application using echo and golang template

## Owner
[Tommy Nurwantoro](https://github.com/tommynurwantoro)

## Onboarding and Development Guide

### Prerequisite
- Git
- Go 1.13 or later
- [Air](https://github.com/cosmtrek/air) (Optional)

### Installation

- Install Git  
  See [Git Installation](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

- Install Go (Golang)  
  See [Golang Installation](https://golang.org/doc/install)

- Clone this repo in your local

  ```sh
  git clone git@github.com:tommynurwantoro/mygoweb.git
  ```

- Go to `mygoweb` directory, then sync the module

  ```sh
  cd mygoweb
  make mod
  ```

- Copy env.sample and static_config.yml.sample if necessary, modify the env and config value(s)

  ```sh
  cp env.sample .env
  cp static_config.yml.sample static_config.yml
  ```

- Run Web if you use Air (Development only)

  ```sh
  air
  ```

- Run Web

  ```sh
  make run
  ```
  
### Deployment
  
- Copy mygoweb.service to /lib/systemd/system/

  ```sh
  cp mygoweb.service /lib/systemd/system/mygoweb.service
  ```

- Run web

  ```sh
  service mygoweb start
  ```
  
- Stop web (If needed)

  ```sh
  service mygoweb stop
  ```
