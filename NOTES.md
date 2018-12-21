# Notes

## Packages

Where should a `Transaction` type live? A transaction is quite a core piece of information for Bud. It can be used in a lot of different places in the `moneysvc`.

* `money`: Fairly broad, everything to do with Bud is related to money in some way. This is probably going to be the "core" public package for the `moneysvc` though. We'd have a `money.Transaction` and a `money.TransactionGateway` I suppose?
* `transactions`: Would produce a stutter. So, something inbetween these two in terms of specificity. Unless it was called `transactions.Entry` or something similar. Perhaps this one does work? Maybe under `money`? Or is that just pointless? Gateway would be `transactions.Gateway`. This one does lead to shorter type names too overall by keeping the domain a bit more specific. Would we risk encountering cyclic dependencies?
* `bud`: Is this too broad? Or does it just show how re-usable this is? At least `money` is more closely limited to the `moneysvc` application which would be used to interact with transactions. This is probably too broad really.
