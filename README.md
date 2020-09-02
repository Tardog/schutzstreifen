你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Schutzstreifen
A web application that displays a map with user-contributed hazards for cyclists.
Built with the [Buffalo framework](https://gobuffalo.io).

## What's the meaning of the project name?
"Schutzstreifen" is German for "protective strip". In legal terminology, it describes the kind of cheap bicycle path that gets painted on the road, often leading to conflicts between drivers and cyclists due to cars parking on or too close next to the strip, close passes on the road, among others. Many of these contraptions are inherently unsafe for cyclists, which has earned them the nickname of "Gefährdungsstreifen" (literally "hazard strip").

## ⛔️ Attention: Work in progress ⛔️
This application is not feature ready yet (as in: pre-alpha status):
- Several routes are only implemented in the most basic way possible
- Validation and error handling require work to improve security and user experience
- The frontend is pretty bare-bones right now and needs a lot of love and polish to be presentable
- Every single user has access to all POST endpoints, which is not desirable - there should be an admin user e.g. for managing available hazard types, removing users etc.

When the first alpha is ready, the app will be made available online. You can always download the code and run it on your local machine, of course.

## Installation
If you are setting up the application for the first time, use the following make target:
```
make setup
```

## Run locally
For development, having Go and Go dev tools installed on your machine is extremely helpful.

Instead of also installing Buffalo and PostgresSQL, you can simply run the app with Docker. Otherwise, please refer to the Buffalo documentation on how to run the framework natively.

Use the following make commands to start/stop the containers:
```
make start
make stop
```
When all containers are up, the application is available at http://localhost:22080/

To see debug output (logs) during development, use the following start command instead:
```
make start-dev
```

### Display container log output
With the default ```start``` target, containers are started in detached mode. If you prefer running them in the foreground, use the ```make start-dev``` target instead.

## Run tests
```
make test
```

## Buffalo CLI
Use this shortcut to open a shell inside the dev container. You can then use the full Buffalo CLI, including Pop for setting up the database and running migrations:
```
make cli
```

## Force rebuilding containers
This will skip the Docker build cache.
```
make build-containers
```

## Compile CSS
Create CSS files from Sass sources:
```
make css
```

## Run database migrations
Execute all migrations (only works if the database already exists).
```
make migrate
```
