# Snippy

A simple way to store and manage snippets of code.

## Client Usage

To initialise snippy:

```bash
snippy init --url "https://snippy.my.domain" --auth "my-auth-token"
```

To upload a snippet:

```bash
snippy upload example-id example-file.go
```

To upload a private snippet:

```bash
snippy upload -p example-id example-file.go
```

To delete a snippet:

```bash
snippy delete example-id
```

To download a snippet:

```bash
snippy download example-id
```

To download a snippet to your clipboard:

```bash
snippy download -c example-id
```

## Server Usage

Snippy uses Docker and docker-compose to run its server. You will need both installed to run it.

To run the server:

```bash
docker-compose up -d
```

You will need a `.env` file in the root of the project with the following variables:

```bash
SNIPPY_URL=https://snippy.my.domain
SNIPPY_AUTH=my-auth-token
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
