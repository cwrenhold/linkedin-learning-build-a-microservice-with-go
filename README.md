This is based on following through the LinkedIn Learning course found here:

<https://www.linkedin.com/learning/build-a-microservice-with-go/>

This project has been setup with a dev container configuration, after cloning the repo, connect to the database with psql in the dev container and run the schema and data SQL files in the `dat` folder.

You can connect to the postgres database from the dev container using the following commany:
```bash
psql -h db -U postgres
```

The password will be `postgres` by default, but this can be changed using the `.env` file in the `.devcontainer` folder.

