# acme-dns-solver
It solves ACME dns-01 challenge for dns providers. (This is for AWS for now.)Thus, it provides ssl/tls certificates automatically according to the domain.

Create a new **.env** file in the root of the project.

```
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_REGION=

CA_DIR_URL=https://acme-staging-v02.api.letsencrypt.org/directory
EMAIL=
```

Give a domain string in **main.go** file.

```
//It'll get a domain parameter as a string
app.Start("")
```

Just run the following command. You will get "fullchain.pem" and "privkey.pem".

``go run .``