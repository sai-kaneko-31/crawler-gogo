# crawler-gogo

## How To Run

### Build the Docker image and run the container

Execute the following command to create a directory with the current date and run the Docker container.

```
CUR_DATE=$(date +"%Y%m%d") && \
docker build . --tag  crawler-gogo:latest && \
docker run -v "$(pwd)/$CUR_DATE:/usr/src/app/$CUR_DATE" crawler-gogo:latest
```

> ![NOTE]
> Use `$(pwd)` to correctly obtain the current directory path.

> ![NOTE]
> You can modify the `/usr/src/app/` directory in the container based on your project's structure.