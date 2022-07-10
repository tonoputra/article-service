# Quick start guide

## Project Structure

## Initial project setup
1. Create a file `.env` based on file `.env.example`
2. Run in terminal `docker volume create --name=tono-modules` to create a volume **tono_go-modules**

## Run Project
- This project is part of **Docker Compose** or **Kubernetes**. So it must be ran from `docker-compose` directory
- Using Docker compose: `docker-compose up` or `docker-compose up <service_name>`
- To stop container run `docker-compose down`

# References
- [Compiler Daemon](https://github.com/githubnemo/CompileDaemon)
- [Projecy Layout](https://github.com/golang-standards/project-layout)

## Gitflow Workflow
- [Gitflow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow#:~:text=The%20overall%20flow%20of%20Gitflow,branch%20is%20created%20from%20develop&text=When%20a%20feature%20is%20complete%20it%20is%20merged%20into%20the,merged%20into%20develop%20and%20main)
- [git-flow-cheatshee](https://danielkummer.github.io/git-flow-cheatsheet/)

## Docker Image Build, CI/CD
- [Golang Build Image](https://docs.docker.com/language/golang/build-images/)
- [Gitlab CI/CD](https://rizkimufrizal.github.io/belajar-gitlab-continuous-integration-dan-continuous-deployment)

## Connection MongoDB
- [MongoDB](https://www.mongodb.com/blog/post/quick-start-golang-mongodb-starting-and-setup)
