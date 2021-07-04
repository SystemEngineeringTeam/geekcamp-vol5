
GeekCamp vol5 ドーパミン APP
===

This product is designed to help you release your dopamine🥳
Built with Golang,Docker,Nuxt.js,MySQL

[![GitHub issues](https://img.shields.io/github/issues/SystemEngineeringTeam/geekcamp-vol5?style=for-the-badge)](https://github.com/SystemEngineeringTeam/geekcamp-vol5/issues)

[![GitHub license](https://img.shields.io/github/license/SystemEngineeringTeam/geekcamp-vol5?style=for-the-badge)](https://github.com/SystemEngineeringTeam/geekcamp-vol5/blob/main/LICENSE)


## How to run this site

`$ cd frontend`

`$ yarn install`

### serve with hot reload at localhost:3000

`$ yarn dev`

### build for production and launch server

`$ yarn build`

`$ yarn start`

### generate static project

`$ yarn generate`

## How to use Docker-Compose and Makefile

```
$ make run/build
	$(COMPOSE) up -d --build
```
    
```
$ make run
	$(COMPOSE) up -d
```

```
$ make sh/database
	$(DOCKER) exec -it $(DB) bash
```

```
$ make down
	$(COMPOSE) down
```

```
$ make down/v
	$(COMPOSE) down -v
```

License
This repository is provided under the MIT License.


###### tags: `Documentation`
