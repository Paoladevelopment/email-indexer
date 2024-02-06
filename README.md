# Email indexer from the Enron database
## First steps
- Download the latest version of ZincSearch from its [repository](https://github.com/zincsearch/zincsearch/releases)
- Download the Enron emails database [here](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz)
This repository contains two aplications: `enron-indexer` and `search-api`.

## enron-indexer
This app parses the raw files from the Enron Database and upload emails to a ZincZearch index.

### Environment variables
Don't forget to include your `.env` file, which should contain the following environment variables:

```bash
ZINC_FIRST_ADMIN_USER="admin"
ZINC_FIRST_ADMIN_PASSWORD="Complexpass#123"
ZINC_SERVER_URL = "http://localhost:4080/api/"
```

